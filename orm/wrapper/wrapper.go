package wrapper

import (
	"context"
	"database/sql"
)

type QueryContextFunc func(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)

type ExecContextFunc func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

// Wrapper defines database common operations
type Wrapper interface {
	WrapQueryContext(ctx context.Context, fn QueryContextFunc, sql string, args ...interface{}) QueryContextFunc
	WrapQueryExecContext(ctx context.Context, fn ExecContextFunc, sql string, args ...interface{}) ExecContextFunc
	Close()
}
