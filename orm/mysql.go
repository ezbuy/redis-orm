package orm

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLStore struct {
	*sql.DB
	debug   bool
	slowlog time.Duration
}

func NewMySQLStore(host string, port int, database, username, password string) (*MySQLStore, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&autocommit=true&parseTime=True",
		username,
		password,
		host,
		port,
		database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &MySQLStore{db, false, time.Duration(0)}, nil
}

func (store *MySQLStore) Debug(b bool) {
	store.debug = b
}

func (store *MySQLStore) SlowLog(duration time.Duration) {
	store.slowlog = duration
}

func (store *MySQLStore) Query(sql string, args ...interface{}) (*sql.Rows, error) {
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

func (store *MySQLStore) Exec(sql string, args ...interface{}) (sql.Result, error) {
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

func (store *MySQLStore) Close() {
	store.DB = nil
}

type MySQLTx struct {
	*sql.Tx
	debug   bool
	slowlog time.Duration
}

func (store *MySQLStore) BeginTx() (*MySQLTx, error) {
	tx, err := store.Begin()
	if err != nil {
		return nil, err
	}
	return &MySQLTx{tx, store.debug, store.slowlog}, nil
}

func (tx *MySQLTx) Query(sql string, args ...interface{}) (*sql.Rows, error) {
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
	return tx.Tx.Query(sql, args...)
}

func (tx *MySQLTx) Exec(sql string, args ...interface{}) (sql.Result, error) {
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
	return tx.Tx.Exec(sql, args...)
}
