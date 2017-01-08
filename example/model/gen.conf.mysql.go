package model

//! conf.mysql

import (
	"github.com/ezbuy/redis-orm/orm"
	"time"
)

var (
	_mysql_store *orm.MySQLStore
)

type MySQLConfig struct {
	Host            string
	Port            int
	UserName        string
	Password        string
	Database        string
	PoolSize        int
	ConnMaxLifeTime time.Duration
}

func MySQLSetup(cf *MySQLConfig) {
	store, err := orm.NewMySQLStore(cf.Host, cf.Port, cf.Database, cf.UserName, cf.Password)
	if err != nil {
		panic(err)
	}

	store.SetConnMaxLifetime(time.Hour)
	if cf.ConnMaxLifeTime > 0 {
		store.SetConnMaxLifetime(cf.ConnMaxLifeTime)
	}
	store.SetMaxIdleConns(cf.PoolSize)
	store.SetMaxOpenConns(cf.PoolSize)
	_mysql_store = store
}
