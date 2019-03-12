package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ezbuy/redis-orm/orm"
	"strings"
	"time"
)

var (
	_ sql.DB
	_ time.Time
	_ fmt.Formatter
	_ strings.Reader
	_ orm.VSet
)

type UserInfo struct {
	Id       int32  `db:"id"`
	Name     string `db:"name"`
	Mailbox  string `db:"mailbox"`
	Password string `db:"password"`
	Sex      bool   `db:"sex"`
}

type _UserInfoMgr struct {
}

var UserInfoMgr *_UserInfoMgr

func (m *_UserInfoMgr) NewUserInfo() *UserInfo {
	return &UserInfo{}
}

type _UserInfoDBMgr struct {
	db orm.DB
}

func (m *_UserInfoMgr) DB(db orm.DB) *_UserInfoDBMgr {
	return UserInfoDBMgr(db)
}

func UserInfoDBMgr(db orm.DB) *_UserInfoDBMgr {
	if db == nil {
		panic(fmt.Errorf("UserInfoDBMgr init need db"))
	}
	return &_UserInfoDBMgr{db: db}
}

func (m *_UserInfoDBMgr) QueryBySQL(q string, args ...interface{}) (results []*UserInfo, err error) {
	rows, err := m.db.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("UserInfo fetch error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var result UserInfo
		err = rows.Scan(&(result.Id), &(result.Name), &(result.Mailbox), &(result.Password), &(result.Sex))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("UserInfo fetch result error: %v", err)
	}
	return
}

func (m *_UserInfoDBMgr) QueryBySQLContext(ctx context.Context, q string, args ...interface{}) (results []*UserInfo, err error) {
	rows, err := m.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("UserInfo fetch error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var result UserInfo
		err = rows.Scan(&(result.Id), &(result.Name), &(result.Mailbox), &(result.Password), &(result.Sex))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("UserInfo fetch result error: %v", err)
	}
	return
}
