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

type _BlogMySQLMgr struct {
	*orm.MySQLStore
}

func BlogMySQLMgr() *_BlogMySQLMgr {
	return &_BlogMySQLMgr{_mysql_store}
}

func NewBlogMySQLMgr(cf *MySQLConfig) (*_BlogMySQLMgr, error) {
	store, err := orm.NewMySQLStore(cf.Host, cf.Port, cf.Database, cf.UserName, cf.Password)
	if err != nil {
		return nil, err
	}
	return &_BlogMySQLMgr{store}, nil
}

func (m *_BlogMySQLMgr) FetchBySQL(sql string, args ...interface{}) (results []interface{}, err error) {
	rows, err := m.Query(sql, args...)
	if err != nil {
		return nil, fmt.Errorf("Blog fetch error: %v", err)
	}
	defer rows.Close()

	var CreatedAt string
	var UpdatedAt string

	for rows.Next() {
		var result Blog
		err = rows.Scan(&(result.Id),
			&(result.UserId),
			&(result.Title),
			&(result.Content),
			&(result.Status),
			&(result.Readed),
			&CreatedAt, &UpdatedAt)
		if err != nil {
			return nil, err
		}

		result.CreatedAt = orm.TimeParse(CreatedAt)

		result.UpdatedAt = orm.TimeParse(UpdatedAt)

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Blog fetch result error: %v", err)
	}
	return
}
func (m *_BlogMySQLMgr) Fetch(id string) (*Blog, error) {
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` WHERE `Id` = (%s)", strings.Join(obj.GetColumns(), ","), id)
	objs, err := m.FetchBySQL(query)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*Blog), nil
	}
	return nil, fmt.Errorf("Blog fetch record not found")
}

func (m *_BlogMySQLMgr) FetchByIds(ids []string) ([]*Blog, error) {
	if len(ids) == 0 {
		return []*Blog{}, nil
	}

	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` WHERE `Id` IN (%s)", strings.Join(obj.GetColumns(), ","), strings.Join(ids, ","))
	objs, err := m.FetchBySQL(query)
	if err != nil {
		return nil, err
	}
	results := make([]*Blog, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*Blog))
	}
	return results, nil
}

func (m *_BlogMySQLMgr) FindOne(unique Unique) (string, error) {
	objs, err := m.queryLimit(unique.SQLFormat(), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return "", err
	}
	if len(objs) > 0 {
		return fmt.Sprint(objs[0]), nil
	}
	return "", fmt.Errorf("Blog find record not found")
}

func (m *_BlogMySQLMgr) Find(index Index) ([]string, error) {
	return m.queryLimit(index.SQLFormat(), index.SQLLimit(), index.SQLParams()...)
}

func (m *_BlogMySQLMgr) Range(scope Range) ([]string, error) {
	return m.queryLimit(scope.SQLFormat(), scope.SQLLimit(), scope.SQLParams()...)
}

func (m *_BlogMySQLMgr) OrderBy(sort OrderBy) ([]string, error) {
	return m.queryLimit(sort.SQLFormat(), sort.SQLLimit(), sort.SQLParams()...)
}

func (m *_BlogMySQLMgr) queryLimit(where string, limit int, args ...interface{}) (results []string, err error) {
	query := fmt.Sprintf("SELECT `id` FROM `blogs`")
	if where != "" {
		query += " WHERE "
		query += where
	}

	rows, err := m.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("Blog query limit error: %v", err)
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
		return nil, fmt.Errorf("Blog query limit result error: %v", err)
	}
	return
}

//! orm.mysql.write
///////////////////////////
//! 	how to use tx
//!
//! 	tx, err := BlogMySQLMgr().BeginTx()
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
type _BlogMySQLTx struct {
	*sql.Tx
	Err          error
	RowsAffected int64
}

func (m *_BlogMySQLMgr) BeginTx() (*_BlogMySQLTx, error) {
	tx, err := m.Begin()
	if err != nil {
		return nil, err
	}
	return &_BlogMySQLTx{tx, nil, 0}, nil
}

func (tx *_BlogMySQLTx) Create(obj *Blog) error {
	params := orm.NewStringSlice(8, "?")
	q := fmt.Sprintf("INSERT INTO `blogs`(%s) VALUES(%s)",
		strings.Join(obj.GetColumns(), ","),
		strings.Join(params, ","))

	result, err := tx.Exec(q, obj.Id, obj.UserId, obj.Title, obj.Content, obj.Status, obj.Readed, orm.TimeFormat(obj.CreatedAt), orm.TimeFormat(obj.UpdatedAt))
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

func (tx *_BlogMySQLTx) Update(obj *Blog) error {
	columns := []string{
		"`user_id` = ?",
		"`title` = ?",
		"`content` = ?",
		"`status` = ?",
		"`readed` = ?",
		"`created_at` = ?",
		"`updated_at` = ?",
	}
	q := fmt.Sprintf("UPDATE `blogs` SET %s WHERE `id`=?",
		strings.Join(columns, ","))
	result, err := tx.Exec(q, obj.UserId, obj.Title, obj.Content, obj.Status, obj.Readed, orm.TimeFormat(obj.CreatedAt), orm.TimeFormat(obj.UpdatedAt), obj.Id)
	if err != nil {
		tx.Err = err
		return err
	}
	tx.RowsAffected, tx.Err = result.RowsAffected()
	return tx.Err
}

func (tx *_BlogMySQLTx) Save(obj *Blog) error {
	err := tx.Update(obj)
	if err != nil {
		return err
	}
	if tx.RowsAffected > 0 {
		return nil
	}
	return tx.Create(obj)
}

func (tx *_BlogMySQLTx) Delete(obj *Blog) error {
	q := fmt.Sprintf("DELETE FROM `blogs` WHERE `id`=?")
	result, err := tx.Exec(q, obj.Id)
	if err != nil {
		tx.Err = err
		return err
	}
	tx.RowsAffected, tx.Err = result.RowsAffected()
	return tx.Err
}

func (tx *_BlogMySQLTx) DeleteByIds(ids []string) error {
	if len(ids) == 0 {
		return nil
	}

	q := fmt.Sprintf("DELETE FROM `blogs` WHERE `id` IN (%s)",
		strings.Join(ids, ","))
	result, err := tx.Exec(q)
	if err != nil {
		tx.Err = err
		return err
	}
	tx.RowsAffected, tx.Err = result.RowsAffected()
	return tx.Err
}

func (tx *_BlogMySQLTx) Close() error {
	if tx.Err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}

//! tx read
func (tx *_BlogMySQLTx) FindOne(unique Unique) (string, error) {
	objs, err := tx.queryLimit(unique.SQLFormat(), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		tx.Err = err
		return "", err
	}
	if len(objs) > 0 {
		return fmt.Sprint(objs[0]), nil
	}
	tx.Err = fmt.Errorf("Blog find record not found")
	return "", tx.Err
}

func (tx *_BlogMySQLTx) Find(index Index) ([]string, error) {
	return tx.queryLimit(index.SQLFormat(), index.SQLLimit(), index.SQLParams()...)
}

func (tx *_BlogMySQLTx) Range(scope Range) ([]string, error) {
	return tx.queryLimit(scope.SQLFormat(), scope.SQLLimit(), scope.SQLParams()...)
}

func (tx *_BlogMySQLTx) OrderBy(sort OrderBy) ([]string, error) {
	return tx.queryLimit(sort.SQLFormat(), sort.SQLLimit(), sort.SQLParams()...)
}

func (tx *_BlogMySQLTx) queryLimit(where string, limit int, args ...interface{}) (results []string, err error) {
	query := fmt.Sprintf("SELECT `id` FROM `blogs`")
	if where != "" {
		query += " WHERE "
		query += where
	}

	rows, err := tx.Query(query, args...)
	if err != nil {
		tx.Err = err
		return nil, fmt.Errorf("Blog query limit error: %v", err)
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
		return nil, fmt.Errorf("Blog query limit result error: %v", err)
	}
	return
}

func (tx *_BlogMySQLTx) Fetch(id interface{}) (*Blog, error) {
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` WHERE `Id` = (%s)", strings.Join(obj.GetColumns(), ","), fmt.Sprint(id))
	objs, err := tx.FetchBySQL(query)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("Blog fetch record not found")
}

func (tx *_BlogMySQLTx) FetchByIds(ids []string) ([]*Blog, error) {
	if len(ids) == 0 {
		return []*Blog{}, nil
	}

	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` WHERE `Id` IN (%s)", strings.Join(obj.GetColumns(), ","), strings.Join(ids, ","))
	return tx.FetchBySQL(query)
}

func (tx *_BlogMySQLTx) FetchBySQL(sql string, args ...interface{}) (results []*Blog, err error) {
	rows, err := tx.Query(sql, args...)
	if err != nil {
		tx.Err = err
		return nil, fmt.Errorf("Blog fetch error: %v", err)
	}
	defer rows.Close()

	var CreatedAt string
	var UpdatedAt string

	for rows.Next() {
		var result Blog
		err = rows.Scan(&(result.Id),
			&(result.UserId),
			&(result.Title),
			&(result.Content),
			&(result.Status),
			&(result.Readed),
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
		return nil, fmt.Errorf("Blog fetch result error: %v", err)
	}
	return
}
