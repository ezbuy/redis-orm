package orm

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb" // register driver for go-mssqldb
	"github.com/ezbuy/wrapper/database"
	_ "github.com/go-sql-driver/mysql" // register driver for mysql
)

type DB interface {
	Query(sql string, args ...interface{}) (*sql.Rows, error)
	Exec(sql string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, sql string, args ...interface{}) (sql.Result, error)
	SetError(err error)
}

type DBStore struct {
	*sql.DB
	debug    bool
	slowlog  time.Duration
	wrappers []database.Wrapper
}

func NewDBStore(driver, host string, port int, database, username, password string) (*DBStore, error) {
	var dsn string
	switch strings.ToLower(driver) {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&autocommit=true&parseTime=True",
			username,
			password,
			host,
			port,
			database)
	case "mssql":
		dsn = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
			host, username, password, port, database)
	default:
		return nil, fmt.Errorf("unsupport db driver: %s", driver)
	}
	return NewDBDSNStore(driver, dsn)
}

func NewDBStoreWithRawDB(db *sql.DB) *DBStore {
	wps := []database.Wrapper{
		// insert common wrappers here...
	}
	return &DBStore{db, false, time.Duration(0), wps}
}

func NewDBDSNStore(driver, dsn string) (*DBStore, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	wps := []database.Wrapper{
		// insert common wrappers here...
	}
	return &DBStore{db, false, time.Duration(0), wps}, nil
}

func NewDBStoreCharset(driver, host string, port int, databaseName, username, password, charset string) (*DBStore, error) {
	var dsn string
	switch strings.ToLower(driver) {
	case "mysql":
		if charset == "" {
			charset = "utf8"
		}
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&autocommit=true&parseTime=True",
			username,
			password,
			host,
			port,
			databaseName,
			charset)
	case "mssql":
		dsn = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
			host, username, password, port, databaseName)
	default:
		return nil, fmt.Errorf("unsupport db driver: %s", driver)
	}
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	wps := []database.Wrapper{
		// insert common wrappers here...
	}
	return &DBStore{db, false, time.Duration(0), wps}, nil
}

func (store *DBStore) Debug(b bool) {
	store.debug = b
}

func (store *DBStore) SlowLog(duration time.Duration) {
	store.slowlog = duration
}

func (store *DBStore) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	t1 := time.Now()
	if store.slowlog > 0 {
		defer func(t time.Time) {
			span := time.Now().Sub(t1)
			if span > store.slowlog {
				log.Println("SLOW: ", span.String(), sql, args)
			}
		}(t1)
	}
	if store.debug {
		log.Println("DEBUG: ", sql, args)
	}
	return store.DB.Query(sql, args...)
}

func (store *DBStore) Exec(sql string, args ...interface{}) (sql.Result, error) {
	t1 := time.Now()
	if store.slowlog > 0 {
		defer func(t time.Time) {
			span := time.Now().Sub(t1)
			if span > store.slowlog {
				log.Println("SLOW: ", span.String(), sql, args)
			}
		}(t1)
	}
	if store.debug {
		log.Println("DEBUG: ", sql, args)
	}
	return store.DB.Exec(sql, args...)
}

func (store *DBStore) SetError(err error) {}

func (store *DBStore) Close() error {
	if err := store.DB.Close(); err != nil {
		return err
	}
	store.DB = nil
	return nil
}

func (store *DBStore) AddWrappers(wp ...database.Wrapper) {
	store.wrappers = append(store.wrappers, wp...)
}

func (store *DBStore) QueryContext(ctx context.Context, query string,
	args ...interface{}) (*sql.Rows, error) {
	fn := func(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
		return store.DB.QueryContext(ctx, query, args...)
	}
	for _, wp := range store.wrappers {
		fn = wp.WrapQueryContext(fn, query, args...)
	}
	return fn(ctx, query, args...)
}

func (store *DBStore) ExecContext(ctx context.Context, query string,
	args ...interface{}) (sql.Result, error) {
	fn := func(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
		return store.DB.ExecContext(ctx, query, args...)
	}
	for _, wp := range store.wrappers {
		fn = wp.WrapExecContext(fn, query, args...)
	}
	return fn(ctx, query, args...)
}

type DBTx struct {
	tx           *sql.Tx
	debug        bool
	slowlog      time.Duration
	err          error
	rowsAffected int64
	wrappers     []database.Wrapper
	afterCommit  func(err error)
}

func (store *DBStore) BeginTx() (*DBTx, error) {
	return store.BeginTxContext(context.Background())
}

func (store *DBStore) BeginTxContext(ctx context.Context) (*DBTx, error) {
	tx, err := store.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &DBTx{
		tx:       tx,
		debug:    store.debug,
		slowlog:  store.slowlog,
		wrappers: store.wrappers,
	}, nil
}

func (tx *DBTx) Close() error {
	if tx.err != nil {
		return tx.tx.Rollback()
	}
	err := tx.tx.Commit()
	tx.afterCommit(err)
	return err
}

func (tx *DBTx) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	return tx.QueryContext(context.Background(), sql, args...)
}

func (tx *DBTx) QueryContext(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error) {
	t1 := time.Now()
	if tx.slowlog > 0 {
		defer func(t time.Time) {
			span := time.Now().Sub(t1)
			if span > tx.slowlog {
				log.Println("SLOW: ", span.String(), sql, args)
			}
		}(t1)
	}
	if tx.debug {
		log.Println("DEBUG: ", sql, args)
	}
	result, err := tx.tx.QueryContext(ctx, sql, args...)
	if err != nil {
		tx.err = err
	}
	return result, tx.err
}

func (tx *DBTx) Exec(sql string, args ...interface{}) (sql.Result, error) {
	return tx.ExecContext(context.Background(), sql, args...)
}

func (tx *DBTx) ExecContext(ctx context.Context, sql string, args ...interface{}) (sql.Result, error) {
	t1 := time.Now()
	if tx.slowlog > 0 {
		defer func(t time.Time) {
			span := time.Now().Sub(t1)
			if span > tx.slowlog {
				log.Println("SLOW: ", span.String(), sql, args)
			}
		}(t1)
	}
	if tx.debug {
		log.Println("DEBUG: ", sql, args)
	}
	result, err := tx.tx.Exec(sql, args...)
	if err != nil {
		tx.err = err
	}
	return result, tx.err
}

func (tx *DBTx) QueryContext(ctx context.Context, query string,
	args ...interface{}) (*sql.Rows, error) {
	fn := func(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
		return tx.tx.QueryContext(ctx, query, args...)
	}
	for _, wp := range tx.wrappers {
		fn = wp.WrapQueryContext(fn, query, args...)
	}
	return fn(ctx, query, args...)
}

func (tx *DBTx) ExecContext(ctx context.Context, query string,
	args ...interface{}) (sql.Result, error) {
	fn := func(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
		return tx.tx.ExecContext(ctx, query, args...)
	}
	for _, wp := range tx.wrappers {
		fn = wp.WrapExecContext(fn, query, args...)
	}
	return fn(ctx, query, args...)
}

func (tx *DBTx) SetError(err error) {
	tx.err = err
}

func (tx *DBTx) AfterCommit(afterCommit func(err error)) {
	tx.afterCommit = afterCommit
}

func TransactFunc(db *DBStore, txFunc func(*DBTx) error) (err error) {
	tx, err := db.BeginTx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.SetError(fmt.Errorf("panic: %v", p))
			tx.Close()
			panic(p)
		} else if err != nil {
			tx.SetError(err)
			tx.Close()
		} else {
			err = tx.Close()
		}
	}()

	err = txFunc(tx)
	return err
}

func TransactFuncContext(ctx context.Context, db *DBStore, txFunc func(ctx context.Context, tx *DBTx) error) (err error) {
	tx, err := db.BeginTxContext(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.SetError(fmt.Errorf("panic: %v", p))
			tx.Close()
			panic(p)
		} else if err != nil {
			tx.SetError(err)
			tx.Close()
		} else {
			err = tx.Close()
		}
	}()

	err = txFunc(ctx, tx)
	return err
}

type Transactor interface {
	Transact(tx *DBTx) error
}

func Transact(db *DBStore, t Transactor) error {
	return TransactFunc(db, t.Transact)
}

type TransactorWithContext interface {
	TransactContext(ctx context.Context, tx *DBTx) error
}

func TransactContext(ctx context.Context, db *DBStore, t TransactorWithContext) error {
	return TransactFuncContext(ctx, db, t.TransactContext)
}

func BeginTx(ctx context.Context, db *sql.DB) (*DBTx, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &DBTx{
		tx: tx,
	}, nil
}
