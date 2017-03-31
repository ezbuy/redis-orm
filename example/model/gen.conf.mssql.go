package model

import (
	"github.com/ezbuy/redis-orm/orm"
	"time"
)

var (
	_mssql_store *orm.DBStore
)

type MsSQLConfig struct {
	Host            string
	Port            int
	UserName        string
	Password        string
	Database        string
	PoolSize        int
	ConnMaxLifeTime time.Duration
}

func MsSQLSetup(cf *MsSQLConfig) {
	store, err := orm.NewDBStore("mssql", cf.Host, cf.Port, cf.Database, cf.UserName, cf.Password)
	if err != nil {
		panic(err)
	}

	store.SetConnMaxLifetime(time.Hour)
	if cf.ConnMaxLifeTime > 0 {
		store.SetConnMaxLifetime(cf.ConnMaxLifeTime)
	}
	store.SetMaxIdleConns(cf.PoolSize)
	store.SetMaxOpenConns(cf.PoolSize)
	_mssql_store = store
}

func MsSQL() *orm.DBStore {
	return _mssql_store
}
