package orm

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/ezbuy/redis-orm/trace/database/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type DB interface {
	Query(sql string, args ...interface{}) (*sql.Rows, error)
	Exec(sql string, args ...interface{}) (sql.Result, error)
	SetError(err error)
	ContextDB
}

type ContextDB interface {
	QueryContext(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, sql string, args ...interface{}) (sql.Result, error)
}

type DBStore struct {
	*sql.DB
	debug   bool
	slowlog time.Duration
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

func NewDBDSNStore(driver, dsn string) (*DBStore, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	return &DBStore{db, false, time.Duration(0)}, nil
}

func NewDBStoreCharset(driver, host string, port int, database, username, password, charset string) (*DBStore, error) {
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
			database,
			charset)
	case "mssql":
		dsn = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
			host, username, password, port, database)
	default:
		return nil, fmt.Errorf("unsupport db driver: %s", driver)
	}
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	return &DBStore{db, false, time.Duration(0)}, nil
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

func (store *DBStore) QueryContext(ctx context.Context, sql string,
	args ...interface{}) (*sql.Rows, error) {
	t := mysql.NewDefaultTracer(store.DB, false)
	return t.QueryContext(ctx, sql, args...)
}

func (store *DBStore) ExecContext(ctx context.Context, sql string,
	args ...interface{}) (sql.Result, error) {
	t := mysql.NewDefaultTracer(store.DB, false)
	return t.ExecContext(ctx, sql, args...)
}

type DBTx struct {
	tx           *sql.Tx
	debug        bool
	slowlog      time.Duration
	err          error
	rowsAffected int64
}

func (store *DBStore) BeginTx() (*DBTx, error) {
	tx, err := store.Begin()
	if err != nil {
		return nil, err
	}

	return &DBTx{
		tx:      tx,
		debug:   store.debug,
		slowlog: store.slowlog,
	}, nil
}

func (tx *DBTx) Close() error {
	if tx.err != nil {
		return tx.tx.Rollback()
	}
	return tx.tx.Commit()
}

func (tx *DBTx) Query(sql string, args ...interface{}) (*sql.Rows, error) {
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
	result, err := tx.tx.Query(sql, args...)
	if err != nil {
		tx.err = err
	}
	return result, tx.err
}

func (tx *DBTx) Exec(sql string, args ...interface{}) (sql.Result, error) {
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

func (tx *DBTx) QueryContext(ctx context.Context, sql string,
	args ...interface{}) (*sql.Rows, error) {
	t := mysql.NewDefaultTracer(tx.tx, false)
	return t.QueryContext(ctx, sql, args...)
}

func (tx *DBTx) ExecContext(ctx context.Context, sql string,
	args ...interface{}) (sql.Result, error) {
	t := mysql.NewDefaultTracer(tx.tx, false)
	return t.ExecContext(ctx, sql, args...)
}

func (tx *DBTx) SetError(err error) {
	tx.err = err
}

func TransactFunc(ctx context.Context, db *DBStore, txFunc func(context.Context, *DBTx) error) (err error) {
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

	err = txFunc(ctx, tx)
	return err
}

type Transactor interface {
	Transact(ctx context.Context, tx *DBTx) error
}

func Transact(ctx context.Context, db *DBStore, t Transactor) error {
	return TransactFunc(ctx, db, t.Transact)
}
