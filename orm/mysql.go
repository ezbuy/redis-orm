package orm

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLStore struct {
	*sql.DB
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
	return &MySQLStore{db}, nil
}
