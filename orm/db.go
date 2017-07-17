package orm

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

type DB interface {
	Query(sql string, args ...interface{}) (*sql.Rows, error)
	Exec(sql string, args ...interface{}) (sql.Result, error)
	SetError(err error)
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
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&autocommit=true&parseTime=True",
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
	tx.err = err
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
	tx.err = err
	return result, tx.err
}

func (tx *DBTx) SetError(err error) {
	tx.err = err
}
