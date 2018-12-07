package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/golang-sql/sqlexp"
	"github.com/opentracing/opentracing-go"
	tags "github.com/opentracing/opentracing-go/ext"
)

const (
	dbTracerServiceMySQL = "MySQL"
)

// Operation defines database common operations
type Operation interface {
	QueryContext(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, sql string, args ...interface{}) (sql.Result, error)
}

func NewCustmizedTracer(op Operation) Operation {
	return op
}

func NewDefaultTracer(db sqlexp.Querier, enableRawQuery bool) Operation {
	return &DefaultTracer{
		tracer:           opentracing.GlobalTracer(),
		db:               db,
		isRawQueryEnable: enableRawQuery,
	}
}

type DefaultTracer struct {
	tracer           opentracing.Tracer
	db               sqlexp.Querier
	isRawQueryEnable bool
}

func (t *DefaultTracer) QueryContext(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error) {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := t.tracer.StartSpan(dbTracerServiceMySQL, opentracing.ChildOf(span.Context()))
		tags.SpanKindRPCClient.Set(span)
		tags.PeerService.Set(span, dbTracerServiceMySQL)
		span.SetTag("sql.query", t.hackQueryBuilder(sql, args...))
		ctx = opentracing.ContextWithSpan(ctx, span)
		defer span.Finish()
	}
	rows, err := t.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (t *DefaultTracer) ExecContext(ctx context.Context, sql string, args ...interface{}) (sql.Result, error) {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := t.tracer.StartSpan(dbTracerServiceMySQL, opentracing.ChildOf(span.Context()))
		tags.SpanKindRPCClient.Set(span)
		tags.PeerService.Set(span, dbTracerServiceMySQL)
		span.SetTag("sql.query", t.hackQueryBuilder(sql, args...))
		ctx = opentracing.ContextWithSpan(ctx, span)
		defer span.Finish()
	}
	res, err := t.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *DefaultTracer) hackQueryBuilder(query string, args ...interface{}) string {
	if t.isRawQueryEnable {
		q := strings.Replace(query, "?", "%v", -1)
		return fmt.Sprintf(q, args...)
	}
	return query
}
