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

type _UserIdMySQLMgr struct {
	*orm.MySQLStore
}

func UserIdMySQLMgr() *_UserIdMySQLMgr {
	return &_UserIdMySQLMgr{_mysql_store}
}

func NewUserIdMySQLMgr(cf *MySQLConfig) (*_UserIdMySQLMgr, error) {
	store, err := orm.NewMySQLStore(cf.Host, cf.Port, cf.Database, cf.UserName, cf.Password)
	if err != nil {
		return nil, err
	}
	return &_UserIdMySQLMgr{store}, nil
}

func (m *_UserIdMySQLMgr) FetchBySQL(sql string, args ...interface{}) (results []interface{}, err error) {
	rows, err := m.Query(sql, args...)
	if err != nil {
		return nil, fmt.Errorf("UserId fetch error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var result UserId
		err = rows.Scan(&(result.Key),
			&(result.Value),
		)
		if err != nil {
			return nil, err
		}

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("UserId fetch result error: %v", err)
	}
	return
}
