package orm

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLStore struct {
	*sql.DB
	debug bool
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
	return &MySQLStore{db, false}, nil
}

func (store *MySQLStore) Debug(b bool) {
	store.debug = b
}

func (store *MySQLStore) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	if store.debug {
		fmt.Println("DEBUG: ", sql, args)
	}
	return store.DB.Query(sql, args...)
}

func (store *MySQLStore) Exec(sql string, args ...interface{}) (sql.Result, error) {
	if store.debug {
		fmt.Println("DEBUG: ", sql, args)
	}
	return store.DB.Exec(sql, args...)
}

type MySQLTx struct {
	*sql.Tx
	debug bool
}

func NewMySQLTx(tx *sql.Tx) *MySQLTx {
	return &MySQLTx{tx, false}
}

func (tx *MySQLTx) Debug(b bool) {
	tx.debug = b
}

func (tx *MySQLTx) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	if tx.debug {
		fmt.Println("DEBUG: ", sql, args)
	}
	return tx.Tx.Query(sql, args...)
}

func (tx *MySQLTx) Exec(sql string, args ...interface{}) (sql.Result, error) {
	if tx.debug {
		fmt.Println("DEBUG: ", sql, args)
	}
	return tx.Tx.Exec(sql, args...)
}
