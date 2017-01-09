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

type _UserLocationMySQLMgr struct {
	*orm.MySQLStore
}

func UserLocationMySQLMgr() *_UserLocationMySQLMgr {
	return &_UserLocationMySQLMgr{_mysql_store}
}

func NewUserLocationMySQLMgr(cf *MySQLConfig) (*_UserLocationMySQLMgr, error) {
	store, err := orm.NewMySQLStore(cf.Host, cf.Port, cf.Database, cf.UserName, cf.Password)
	if err != nil {
		return nil, err
	}
	return &_UserLocationMySQLMgr{store}, nil
}

func (m *_UserLocationMySQLMgr) FetchBySQL(sql string, args ...interface{}) (results []interface{}, err error) {
	rows, err := m.Query(sql, args...)
	if err != nil {
		return nil, fmt.Errorf("UserLocation fetch error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var result UserLocation
		err = rows.Scan(&(result.Key),
			&(result.Longitude),
			&(result.Latitude),
			&(result.Value),
		)
		if err != nil {
			return nil, err
		}

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("UserLocation fetch result error: %v", err)
	}
	return
}
