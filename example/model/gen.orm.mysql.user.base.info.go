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

type _UserBaseInfoMySQLMgr struct {
	*orm.MySQLStore
}

func UserBaseInfoMySQLMgr() *_UserBaseInfoMySQLMgr {
	return &_UserBaseInfoMySQLMgr{_mysql_store}
}

func NewUserBaseInfoMySQLMgr(cf *MySQLConfig) (*_UserBaseInfoMySQLMgr, error) {
	store, err := orm.NewMySQLStore(cf.Host, cf.Port, cf.Database, cf.UserName, cf.Password)
	if err != nil {
		return nil, err
	}
	return &_UserBaseInfoMySQLMgr{store}, nil
}

func (m *_UserBaseInfoMySQLMgr) FetchBySQL(sql string, args ...interface{}) (results []interface{}, err error) {
	rows, err := m.Query(sql, args...)
	if err != nil {
		return nil, fmt.Errorf("UserBaseInfo fetch error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var result UserBaseInfo
		err = rows.Scan(&(result.Id),
			&(result.Name),
			&(result.Mailbox),
			&(result.Password),
			&(result.Sex),
		)
		if err != nil {
			return nil, err
		}

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("UserBaseInfo fetch result error: %v", err)
	}
	return
}
func (m *_UserBaseInfoMySQLMgr) Fetch(id string) (*UserBaseInfo, error) {
	obj := UserBaseInfoMgr.NewUserBaseInfo()
	query := fmt.Sprintf("SELECT %s FROM `user_base_info` WHERE `Id` = (%s)", strings.Join(obj.GetColumns(), ","), id)
	objs, err := m.FetchBySQL(query)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*UserBaseInfo), nil
	}
	return nil, fmt.Errorf("UserBaseInfo fetch record not found")
}

func (m *_UserBaseInfoMySQLMgr) FetchByIds(ids []string) ([]*UserBaseInfo, error) {
	if len(ids) == 0 {
		return []*UserBaseInfo{}, nil
	}

	obj := UserBaseInfoMgr.NewUserBaseInfo()
	query := fmt.Sprintf("SELECT %s FROM `user_base_info` WHERE `Id` IN (%s)", strings.Join(obj.GetColumns(), ","), strings.Join(ids, ","))
	objs, err := m.FetchBySQL(query)
	if err != nil {
		return nil, err
	}
	results := make([]*UserBaseInfo, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*UserBaseInfo))
	}
	return results, nil
}

func (m *_UserBaseInfoMySQLMgr) FindOne(unique Unique) (string, error) {
	objs, err := m.queryLimit(unique.SQLFormat(), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return "", err
	}
	if len(objs) > 0 {
		return fmt.Sprint(objs[0]), nil
	}
	return "", fmt.Errorf("UserBaseInfo find record not found")
}

func (m *_UserBaseInfoMySQLMgr) Find(index Index) ([]string, error) {
	return m.queryLimit(index.SQLFormat(), index.SQLLimit(), index.SQLParams()...)
}

func (m *_UserBaseInfoMySQLMgr) Range(scope Range) ([]string, error) {
	return m.queryLimit(scope.SQLFormat(), scope.SQLLimit(), scope.SQLParams()...)
}

func (m *_UserBaseInfoMySQLMgr) OrderBy(sort OrderBy) ([]string, error) {
	return m.queryLimit(sort.SQLFormat(), sort.SQLLimit(), sort.SQLParams()...)
}

func (m *_UserBaseInfoMySQLMgr) queryLimit(where string, limit int, args ...interface{}) (results []string, err error) {
	query := fmt.Sprintf("SELECT `id` FROM `user_base_info`")
	if where != "" {
		query += " WHERE "
		query += where
	}

	rows, err := m.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("UserBaseInfo query limit error: %v", err)
	}
	defer rows.Close()

	offset := 0
	for rows.Next() {
		if limit >= 0 && offset >= limit {
			break
		}
		offset++

		var result int32
		if err = rows.Scan(&result); err != nil {
			return nil, err
		}
		results = append(results, fmt.Sprint(result))
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("UserBaseInfo query limit result error: %v", err)
	}
	return
}
