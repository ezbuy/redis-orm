package model

import (
	"database/sql"
	"fmt"
	"github.com/ezbuy/redis-orm/orm"
	"strings"
)

var (
	_ sql.DB
	_ fmt.Formatter
	_ strings.Reader
	_ orm.VSet
)

type _UserBlogsMySQLMgr struct {
	*orm.MySQLStore
}

func UserBlogsMySQLMgr() *_UserBlogsMySQLMgr {
	return &_UserBlogsMySQLMgr{_mysql_store}
}

func NewUserBlogsMySQLMgr(cf *MySQLConfig) (*_UserBlogsMySQLMgr, error) {
	store, err := orm.NewMySQLStore(cf.Host, cf.Port, cf.Database, cf.UserName, cf.Password)
	if err != nil {
		return nil, err
	}
	return &_UserBlogsMySQLMgr{store}, nil
}

func (m *_UserBlogsMySQLMgr) FetchBySQL(sql string, args ...interface{}) (results []interface{}, err error) {
	rows, err := m.Query(sql, args...)
	if err != nil {
		return nil, fmt.Errorf("UserBlogs fetch error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var result UserBlogs
		err = rows.Scan()
		if err != nil {
			return nil, err
		}

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("UserBlogs fetch result error: %v", err)
	}
	return
}
