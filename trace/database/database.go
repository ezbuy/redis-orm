package database

import (
	"github.com/ezbuy/redis-orm/trace/database/mysql"
)

var db *DB

type DBMySQL struct {
	mysql.Operation
}

func NewTracer() *DB {
	db = new(DB)
	return db
}

type DB struct {
	mysql mysql.Operation
}

func MySQL() *DBMySQL {
	return &DBMySQL{
		db.mysql,
	}
}

func (db *DB) AddMySQL(op mysql.Operation) {
	db.mysql = op
}
