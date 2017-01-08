package model

import (
	"database/sql"
	"fmt"
	"github.com/ezbuy/redis-orm/orm"
	"strings"
)

var (
	_ sql.DB
)

type _UserMySQLMgr struct {
	*orm.MySQLStore
}

func UserMySQLMgr() *_UserMySQLMgr {
	return &_UserMySQLMgr{_mysql_store}
}

func NewUserMySQLMgr(cf *MySQLConfig) (*_UserMySQLMgr, error) {
	store, err := orm.NewMySQLStore(cf.Host, cf.Port, cf.Database, cf.UserName, cf.Password)
	if err != nil {
		return nil, err
	}
	return &_UserMySQLMgr{store}, nil
}

func (m *_UserMySQLMgr) FetchBySQL(sql string, args ...interface{}) (results []interface{}, err error) {
	rows, err := m.Query(sql, args...)
	if err != nil {
		return nil, fmt.Errorf("User fetch error: %v", err)
	}
	defer rows.Close()

	var CreatedAt string
	var UpdatedAt string

	for rows.Next() {
		var result User
		err = rows.Scan(&(result.Id),
			&(result.Name),
			&(result.Mailbox),
			&(result.Sex),
			&(result.Longitude),
			&(result.Latitude),
			&(result.Description),
			&(result.Password),
			&(result.HeadUrl),
			&(result.Status),
			&CreatedAt, &UpdatedAt)
		if err != nil {
			return nil, err
		}

		result.CreatedAt = orm.TimeParse(CreatedAt)

		result.UpdatedAt = orm.TimeParse(UpdatedAt)

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("User fetch result error: %v", err)
	}
	return
}
func (m *_UserMySQLMgr) Fetch(id string) (*User, error) {
	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM `users` WHERE `Id` = (%s)", strings.Join(obj.GetColumns(), ","), id)
	objs, err := m.FetchBySQL(query)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*User), nil
	}
	return nil, fmt.Errorf("User fetch record not found")
}

func (m *_UserMySQLMgr) FetchByIds(ids []string) ([]*User, error) {
	if len(ids) == 0 {
		return []*User{}, nil
	}

	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM `users` WHERE `Id` IN (%s)", strings.Join(obj.GetColumns(), ","), strings.Join(ids, ","))
	objs, err := m.FetchBySQL(query)
	if err != nil {
		return nil, err
	}
	results := make([]*User, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*User))
	}
	return results, nil
}

func (m *_UserMySQLMgr) FindOne(unique Unique) (string, error) {
	objs, err := m.queryLimit(unique.SQLFormat(), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return "", err
	}
	if len(objs) > 0 {
		return fmt.Sprint(objs[0]), nil
	}
	return "", fmt.Errorf("User find record not found")
}

func (m *_UserMySQLMgr) Find(index Index) ([]string, error) {
	return m.queryLimit(index.SQLFormat(), index.SQLLimit(), index.SQLParams()...)
}

func (m *_UserMySQLMgr) Range(scope Range) ([]string, error) {
	return m.queryLimit(scope.SQLFormat(), scope.SQLLimit(), scope.SQLParams()...)
}

func (m *_UserMySQLMgr) OrderBy(sort OrderBy) ([]string, error) {
	return m.queryLimit(sort.SQLFormat(), sort.SQLLimit(), sort.SQLParams()...)
}

func (m *_UserMySQLMgr) queryLimit(where string, limit int, args ...interface{}) (results []string, err error) {
	query := fmt.Sprintf("SELECT `id` FROM `users`")
	if where != "" {
		query += " WHERE "
		query += where
	}

	rows, err := m.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("User query limit error: %v", err)
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
		return nil, fmt.Errorf("User query limit result error: %v", err)
	}
	return
}

//! orm.mysql.write
///////////////////////////
//! 	how to use tx
//!
//! 	tx, err := UserMySQLMgr().BeginTx()
//! 	if err != nil {
//! 		return err
//! 	}
//! 	defer tx.Close()
//!
//! 	tx.Create(obj)
//! 	tx.Update(obj)
//! 	tx.Delete(obj)
///////////////////////////

//! tx write
type _UserMySQLTx struct {
	*sql.Tx
	Err          error
	RowsAffected int64
}

func (m *_UserMySQLMgr) BeginTx() (*_UserMySQLTx, error) {
	tx, err := m.Begin()
	if err != nil {
		return nil, err
	}
	return &_UserMySQLTx{tx, nil, 0}, nil
}

func (tx *_UserMySQLTx) Create(obj *User) error {
	params := orm.NewStringSlice(12, "?")
	q := fmt.Sprintf("INSERT INTO `users`(%s) VALUES(%s)",
		strings.Join(obj.GetColumns(), ","),
		strings.Join(params, ","))

	result, err := tx.Exec(q, obj.Id, obj.Name, obj.Mailbox, obj.Sex, obj.Longitude, obj.Latitude, obj.Description, obj.Password, obj.HeadUrl, obj.Status, orm.TimeFormat(obj.CreatedAt), orm.TimeFormat(obj.UpdatedAt))
	if err != nil {
		tx.Err = err
		return err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		tx.Err = err
		return err
	}
	obj.Id = int32(lastInsertId)
	tx.RowsAffected, tx.Err = result.RowsAffected()
	return tx.Err
}

func (tx *_UserMySQLTx) Update(obj *User) error {
	columns := []string{
		"`name` = ?",
		"`mailbox` = ?",
		"`sex` = ?",
		"`longitude` = ?",
		"`latitude` = ?",
		"`description` = ?",
		"`password` = ?",
		"`head_url` = ?",
		"`status` = ?",
		"`created_at` = ?",
		"`updated_at` = ?",
	}
	q := fmt.Sprintf("UPDATE `users` SET %s WHERE `id`=?",
		strings.Join(columns, ","))
	result, err := tx.Exec(q, obj.Name, obj.Mailbox, obj.Sex, obj.Longitude, obj.Latitude, obj.Description, obj.Password, obj.HeadUrl, obj.Status, orm.TimeFormat(obj.CreatedAt), orm.TimeFormat(obj.UpdatedAt), obj.Id)
	if err != nil {
		tx.Err = err
		return err
	}
	tx.RowsAffected, tx.Err = result.RowsAffected()
	return tx.Err
}

func (tx *_UserMySQLTx) Save(obj *User) error {
	err := tx.Update(obj)
	if err != nil {
		return err
	}
	if tx.RowsAffected > 0 {
		return nil
	}
	return tx.Create(obj)
}

func (tx *_UserMySQLTx) Delete(obj *User) error {
	q := fmt.Sprintf("DELETE FROM `users` WHERE `id`=?")
	result, err := tx.Exec(q, obj.Id)
	if err != nil {
		tx.Err = err
		return err
	}
	tx.RowsAffected, tx.Err = result.RowsAffected()
	return tx.Err
}

func (tx *_UserMySQLTx) DeleteByIds(ids []string) error {
	if len(ids) == 0 {
		return nil
	}

	q := fmt.Sprintf("DELETE FROM `users` WHERE `id` IN (%s)",
		strings.Join(ids, ","))
	result, err := tx.Exec(q)
	if err != nil {
		tx.Err = err
		return err
	}
	tx.RowsAffected, tx.Err = result.RowsAffected()
	return tx.Err
}

func (tx *_UserMySQLTx) Close() error {
	if tx.Err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}

//! tx read
func (tx *_UserMySQLTx) FindOne(unique Unique) (string, error) {
	objs, err := tx.queryLimit(unique.SQLFormat(), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		tx.Err = err
		return "", err
	}
	if len(objs) > 0 {
		return fmt.Sprint(objs[0]), nil
	}
	tx.Err = fmt.Errorf("User find record not found")
	return "", tx.Err
}

func (tx *_UserMySQLTx) Find(index Index) ([]string, error) {
	return tx.queryLimit(index.SQLFormat(), index.SQLLimit(), index.SQLParams()...)
}

func (tx *_UserMySQLTx) Range(scope Range) ([]string, error) {
	return tx.queryLimit(scope.SQLFormat(), scope.SQLLimit(), scope.SQLParams()...)
}

func (tx *_UserMySQLTx) OrderBy(sort OrderBy) ([]string, error) {
	return tx.queryLimit(sort.SQLFormat(), sort.SQLLimit(), sort.SQLParams()...)
}

func (tx *_UserMySQLTx) queryLimit(where string, limit int, args ...interface{}) (results []string, err error) {
	query := fmt.Sprintf("SELECT `id` FROM `users`")
	if where != "" {
		query += " WHERE "
		query += where
	}

	rows, err := tx.Query(query, args...)
	if err != nil {
		tx.Err = err
		return nil, fmt.Errorf("User query limit error: %v", err)
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
			tx.Err = err
			return nil, err
		}
		results = append(results, fmt.Sprint(result))
	}
	if err := rows.Err(); err != nil {
		tx.Err = err
		return nil, fmt.Errorf("User query limit result error: %v", err)
	}
	return
}

func (tx *_UserMySQLTx) Fetch(id interface{}) (*User, error) {
	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM `users` WHERE `Id` = (%s)", strings.Join(obj.GetColumns(), ","), fmt.Sprint(id))
	objs, err := tx.FetchBySQL(query)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("User fetch record not found")
}

func (tx *_UserMySQLTx) FetchByIds(ids []string) ([]*User, error) {
	if len(ids) == 0 {
		return []*User{}, nil
	}

	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM `users` WHERE `Id` IN (%s)", strings.Join(obj.GetColumns(), ","), strings.Join(ids, ","))
	return tx.FetchBySQL(query)
}

func (tx *_UserMySQLTx) FetchBySQL(sql string, args ...interface{}) (results []*User, err error) {
	rows, err := tx.Query(sql, args...)
	if err != nil {
		tx.Err = err
		return nil, fmt.Errorf("User fetch error: %v", err)
	}
	defer rows.Close()

	var CreatedAt string
	var UpdatedAt string

	for rows.Next() {
		var result User
		err = rows.Scan(&(result.Id),
			&(result.Name),
			&(result.Mailbox),
			&(result.Sex),
			&(result.Longitude),
			&(result.Latitude),
			&(result.Description),
			&(result.Password),
			&(result.HeadUrl),
			&(result.Status),
			&CreatedAt, &UpdatedAt)
		if err != nil {
			tx.Err = err
			return nil, err
		}

		result.CreatedAt = orm.TimeParse(CreatedAt)

		result.UpdatedAt = orm.TimeParse(UpdatedAt)

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		tx.Err = err
		return nil, fmt.Errorf("User fetch result error: %v", err)
	}
	return
}
