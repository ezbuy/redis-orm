package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/ezbuy/redis-orm/trace/wrapper"
	wpsql "github.com/ezbuy/redis-orm/trace/wrapper/sql"
	"github.com/golang-sql/sqlexp"
)

// Operation defines database common operations
type Operation interface {
	QueryContext(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, sql string, args ...interface{}) (sql.Result, error)
}

func NewCustmizedTracer(op Operation) Operation {
	return op
}

func NewDefaultTracer(db sqlexp.Querier, enableRawQuery bool, instance string, user string) Operation {
	return &DefaultTracer{
		db:               db,
		isRawQueryEnable: enableRawQuery,
		wrappers: []wrapper.Wrapper{
			wpsql.NewMySQLTracer(instance, user),
		},
	}
}

type DefaultTracer struct {
	db               sqlexp.Querier
	isRawQueryEnable bool
	wrappers         []wrapper.Wrapper
}

func (t *DefaultTracer) QueryContext(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error) {
	for _, w := range t.wrappers {
		w.Do(ctx, sql)
		defer w.Close()
	}
	rows, err := t.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (t *DefaultTracer) ExecContext(ctx context.Context, sql string, args ...interface{}) (sql.Result, error) {
	for _, w := range t.wrappers {
		w.Do(ctx, sql)
		defer w.Close()
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
