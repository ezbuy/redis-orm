package mysql

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/ezbuy/redis-orm/orm/wrapper"
	"github.com/golang-sql/sqlexp"
	"github.com/opentracing/opentracing-go"
	tags "github.com/opentracing/opentracing-go/ext"
)

type Tracer struct {
	instance  string
	statement string
	dbtype    string
	user      string
	span      opentracing.Span
}

func NewMySQLTracer(ins string, user string) *Tracer {
	return &Tracer{
		dbtype:   "mysql",
		instance: ins,
		user:     user,
	}
}

func (t *Tracer) Do(ctx context.Context, statement string) {
	tracer := opentracing.GlobalTracer()
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return
	}
	span = tracer.StartSpan("SQL Operation", opentracing.ChildOf(span.Context()))
	tags.DBInstance.Set(span, t.instance)
	tags.DBStatement.Set(span, statement)
	tags.DBType.Set(span, "mysql")
	tags.DBUser.Set(span, t.user)
	ctx = opentracing.ContextWithSpan(ctx, span)
	t.span = span
}

func (t *Tracer) Close() {
	t.span.Finish()
}

func NewCustmizedTracer(tracer wrapper.Wrapper) wrapper.Wrapper {
	return tracer
}

func NewDefaultTracerWrapper(db sqlexp.Querier, enableRawQuery bool, instance string, user string) wrapper.Wrapper {
	return &DefaultTracer{
		db:                    db,
		isRawQueryEnable:      enableRawQuery,
		isIgnoreSeleteColumns: true,
		tracer:                NewMySQLTracer(instance, user),
	}
}

type DefaultTracer struct {
	db                    sqlexp.Querier
	isRawQueryEnable      bool
	isIgnoreSeleteColumns bool
	tracer                *Tracer
}

func (t *DefaultTracer) WrapQueryContext(ctx context.Context, fn wrapper.QueryContextFunc,
	query string, args ...interface{}) wrapper.QueryContextFunc {
	t.tracer.Do(ctx, t.hackQueryBuilder(query, args...))
	return fn
}

func (t *DefaultTracer) WrapQueryExecContext(ctx context.Context, fn wrapper.ExecContextFunc,
	query string, args ...interface{}) wrapper.ExecContextFunc {
	t.tracer.Do(ctx, t.hackQueryBuilder(query, args...))
	return fn
}

func (t *DefaultTracer) Close() {
	t.tracer.Close()
}

func (t *DefaultTracer) hackQueryBuilder(query string, args ...interface{}) string {
	if t.isRawQueryEnable {
		q := strings.Replace(query, "?", "%v", -1)
		query = fmt.Sprintf(q, args...)
	}
	if t.isIgnoreSeleteColumns {
		query = strings.Replace(query, "select", "SELECT", -1)
		query = strings.Replace(query, "from", "FROM", -1)
		r := regexp.MustCompile("SELECT (.*) FROM")
		query = r.ReplaceAllString(query, "SELECT ... FROM")
	}
	return query
}
