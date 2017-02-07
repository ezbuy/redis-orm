package model

import (
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

type Blog struct {
	Id        int32     `db:"id"`
	UserId    int32     `db:"user_id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	Status    int32     `db:"status"`
	Readed    int32     `db:"readed"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type _BlogMgr struct {
}

var BlogMgr *_BlogMgr

func (m *_BlogMgr) NewBlog() *Blog {
	return &Blog{}
}

//! object function

func (obj *Blog) GetNameSpace() string {
	return "model"
}

func (obj *Blog) GetClassName() string {
	return "Blog"
}

func (obj *Blog) GetTableName() string {
	return "blogs"
}

func (obj *Blog) GetColumns() []string {
	columns := []string{
		"`id`",
		"`user_id`",
		"`title`",
		"`content`",
		"`status`",
		"`readed`",
		"`created_at`",
		"`updated_at`",
	}
	return columns
}

//! uniques

//! indexes

type UserIdOfBlogIDX struct {
	UserId int32
	offset int
	limit  int
}

func (u *UserIdOfBlogIDX) Key() string {
	strs := []string{
		"UserId",
		fmt.Sprint(u.UserId),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *UserIdOfBlogIDX) SQLFormat(limit bool) string {
	conditions := []string{
		"user_id = ?",
	}
	if limit {
		return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return orm.SQLWhere(conditions)
}

func (u *UserIdOfBlogIDX) SQLParams() []interface{} {
	return []interface{}{
		u.UserId,
	}
}

func (u *UserIdOfBlogIDX) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *UserIdOfBlogIDX) Limit(n int) {
	u.limit = n
}

func (u *UserIdOfBlogIDX) Offset(n int) {
	u.offset = n
}

func (u *UserIdOfBlogIDX) IDXRelation() IndexRelation {
	return nil
}

//! ranges
func (m *_BlogMgr) MySQL() *ReferenceResult {
	return NewReferenceResult(BlogMySQLMgr())
}

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

func (m *_BlogMySQLMgr) Search(where string, args ...interface{}) (results []interface{}, err error) {
	obj := BlogMgr.NewBlog()
	if where != "" {
		where = " WHERE " + where
	}
	sql := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(obj.GetColumns(), ","), where)
	return m.FetchBySQL(sql, args...)
}

func (m *_BlogMySQLMgr) SearchCount(where string, args ...interface{}) (int64, error) {
	if where != "" {
		where = " WHERE " + where
	}
	return m.queryCount(where, args...)
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
func (m *_BlogMySQLMgr) Fetch(id interface{}) (*Blog, error) {
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

func (m *_BlogMySQLMgr) FetchByIds(ids []interface{}) ([]*Blog, error) {
	if len(ids) == 0 {
		return []*Blog{}, nil
	}

	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` WHERE `Id` IN (%s)", strings.Join(obj.GetColumns(), ","), orm.SliceJoin(ids, ","))
	objs, err := m.FetchBySQL(query)
	if err != nil {
		return nil, err
	}
	results := make([]*Blog, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*Blog))
	}
	return results, nil
}

func (m *_BlogMySQLMgr) FindOne(unique Unique) (interface{}, error) {
	objs, err := m.queryLimit(unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return "", err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return "", fmt.Errorf("Blog find record not found")
}

func (m *_BlogMySQLMgr) FindOneFetch(unique Unique) (*Blog, error) {
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(obj.GetColumns(), ","), unique.SQLFormat(true))
	objs, err := m.FetchBySQL(query, unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*Blog), nil
	}
	return nil, fmt.Errorf("none record")
}

func (m *_BlogMySQLMgr) Find(index Index) ([]interface{}, error) {
	return m.queryLimit(index.SQLFormat(true), index.SQLLimit(), index.SQLParams()...)
}

func (m *_BlogMySQLMgr) FindFetch(index Index) ([]*Blog, error) {
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	objs, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return nil, err
	}
	results := make([]*Blog, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*Blog))
	}
	return results, nil
}

func (m *_BlogMySQLMgr) FindCount(index Index) (int64, error) {
	return m.queryCount(index.SQLFormat(false), index.SQLParams()...)
}

func (m *_BlogMySQLMgr) Range(scope Range) ([]interface{}, error) {
	return m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (m *_BlogMySQLMgr) RangeFetch(scope Range) ([]*Blog, error) {
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	objs, err := m.FetchBySQL(query, scope.SQLParams()...)
	if err != nil {
		return nil, err
	}
	results := make([]*Blog, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*Blog))
	}
	return results, nil
}

func (m *_BlogMySQLMgr) RangeCount(scope Range) (int64, error) {
	return m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
}

func (m *_BlogMySQLMgr) RangeRevert(scope Range) ([]interface{}, error) {
	scope.Revert(true)
	return m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (m *_BlogMySQLMgr) RangeRevertFetch(scope Range) ([]*Blog, error) {
	scope.Revert(true)
	return m.RangeFetch(scope)
}

func (m *_BlogMySQLMgr) queryLimit(where string, limit int, args ...interface{}) (results []interface{}, err error) {
	query := fmt.Sprintf("SELECT `id` FROM `blogs` %s", where)
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
		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Blog query limit result error: %v", err)
	}
	return
}

func (m *_BlogMySQLMgr) queryCount(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`id`) FROM `blogs` %s", where)
	rows, err := m.Query(query, args...)
	if err != nil {
		return 0, fmt.Errorf("Blog query count error: %v", err)
	}
	defer rows.Close()

	var count int64
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return 0, err
		}
		break
	}
	return count, nil
}

//! tx write
type _BlogMySQLTx struct {
	*orm.MySQLTx
	err          error
	rowsAffected int64
}

func (m *_BlogMySQLMgr) BeginTx() (*_BlogMySQLTx, error) {
	tx, err := m.Begin()
	if err != nil {
		return nil, err
	}
	return &_BlogMySQLTx{orm.NewMySQLTx(tx), nil, 0}, nil
}

func (tx *_BlogMySQLTx) BatchCreate(objs []*Blog) error {
	if len(objs) == 0 {
		return nil
	}

	params := make([]string, 0, len(objs))
	values := make([]interface{}, 0, len(objs)*8)
	for _, obj := range objs {
		params = append(params, fmt.Sprintf("(%s)", strings.Join(orm.NewStringSlice(8, "?"), ",")))
		values = append(values, 0)
		values = append(values, obj.UserId)
		values = append(values, obj.Title)
		values = append(values, obj.Content)
		values = append(values, obj.Status)
		values = append(values, obj.Readed)
		values = append(values, orm.TimeFormat(obj.CreatedAt))
		values = append(values, orm.TimeFormat(obj.UpdatedAt))
	}
	query := fmt.Sprintf("INSERT INTO `blogs`(%s) VALUES %s", strings.Join(objs[0].GetColumns(), ","), strings.Join(params, ","))
	result, err := tx.Exec(query, values...)
	if err != nil {
		tx.err = err
		return err
	}
	tx.rowsAffected, tx.err = result.RowsAffected()
	return tx.err
}

func (tx *_BlogMySQLTx) BatchDelete(objs []*Blog) error {
	if len(objs) == 0 {
		return nil
	}

	ids := make([]interface{}, 0, len(objs))
	for _, obj := range objs {
		ids = append(ids, obj.Id)
	}
	return tx.DeleteByIds(ids)
}

// argument example:
// set:"a=?, b=?"
// where:"c=? and d=?"
// params:[]interface{}{"a", "b", "c", "d"}...
func (tx *_BlogMySQLTx) UpdateBySQL(set, where string, args ...interface{}) error {
	query := fmt.Sprintf("UPDATE `blogs` SET %s", set)
	if where != "" {
		query = fmt.Sprintf("UPDATE `blogs` SET %s WHERE %s", set, where)
	}
	result, err := tx.Exec(query, args)
	if err != nil {
		tx.err = err
		return err
	}
	tx.rowsAffected, tx.err = result.RowsAffected()
	return tx.err
}

func (tx *_BlogMySQLTx) Create(obj *Blog) error {
	params := orm.NewStringSlice(8, "?")
	q := fmt.Sprintf("INSERT INTO `blogs`(%s) VALUES(%s)",
		strings.Join(obj.GetColumns(), ","),
		strings.Join(params, ","))

	result, err := tx.Exec(q, 0, obj.UserId, obj.Title, obj.Content, obj.Status, obj.Readed, orm.TimeFormat(obj.CreatedAt), orm.TimeFormat(obj.UpdatedAt))
	if err != nil {
		tx.err = err
		return err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		tx.err = err
		return err
	}
	obj.Id = int32(lastInsertId)
	tx.rowsAffected, tx.err = result.RowsAffected()
	return tx.err
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
		tx.err = err
		return err
	}
	tx.rowsAffected, tx.err = result.RowsAffected()
	return tx.err
}

func (tx *_BlogMySQLTx) Save(obj *Blog) error {
	err := tx.Update(obj)
	if err != nil {
		return err
	}
	if tx.rowsAffected > 0 {
		return nil
	}
	return tx.Create(obj)
}

func (tx *_BlogMySQLTx) Delete(obj *Blog) error {
	q := fmt.Sprintf("DELETE FROM `blogs` WHERE `id`=?")
	result, err := tx.Exec(q, obj.Id)
	if err != nil {
		tx.err = err
		return err
	}
	tx.rowsAffected, tx.err = result.RowsAffected()
	return tx.err
}

func (tx *_BlogMySQLTx) DeleteByIds(ids []interface{}) error {
	if len(ids) == 0 {
		return nil
	}

	q := fmt.Sprintf("DELETE FROM `blogs` WHERE `id` IN (%s)",
		orm.SliceJoin(ids, ","))
	result, err := tx.Exec(q)
	if err != nil {
		tx.err = err
		return err
	}
	tx.rowsAffected, tx.err = result.RowsAffected()
	return tx.err
}

func (tx *_BlogMySQLTx) Close() error {
	if tx.err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}

//! tx read
func (tx *_BlogMySQLTx) FindOne(unique Unique) (interface{}, error) {
	objs, err := tx.queryLimit(unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		tx.err = err
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	tx.err = fmt.Errorf("Blog find record not found")
	return nil, tx.err
}

func (tx *_BlogMySQLTx) FindOneFetch(unique Unique) (*Blog, error) {
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(obj.GetColumns(), ","), unique.SQLFormat(true))
	objs, err := tx.FetchBySQL(query, unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*Blog), nil
	}
	return nil, fmt.Errorf("none record")
}

func (tx *_BlogMySQLTx) Find(index Index) ([]interface{}, error) {
	return tx.queryLimit(index.SQLFormat(true), index.SQLLimit(), index.SQLParams()...)
}

func (tx *_BlogMySQLTx) FindFetch(index Index) ([]*Blog, error) {
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	objs, err := tx.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return nil, err
	}
	results := make([]*Blog, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*Blog))
	}
	return results, nil
}

func (tx *_BlogMySQLTx) FindCount(index Index) (int64, error) {
	return tx.queryCount(index.SQLFormat(false), index.SQLParams()...)
}

func (tx *_BlogMySQLTx) Range(scope Range) ([]interface{}, error) {
	return tx.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (tx *_BlogMySQLTx) RangeFetch(scope Range) ([]*Blog, error) {
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	objs, err := tx.FetchBySQL(query, scope.SQLParams()...)
	if err != nil {
		return nil, err
	}
	results := make([]*Blog, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*Blog))
	}
	return results, nil
}

func (tx *_BlogMySQLTx) RangeCount(scope Range) (int64, error) {
	return tx.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
}

func (tx *_BlogMySQLTx) RangeRevert(scope Range) ([]interface{}, error) {
	scope.Revert(true)
	return tx.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (tx *_BlogMySQLTx) RangeRevertFetch(scope Range) ([]*Blog, error) {
	scope.Revert(true)
	return tx.RangeFetch(scope)
}

func (tx *_BlogMySQLTx) queryLimit(where string, limit int, args ...interface{}) (results []interface{}, err error) {
	query := fmt.Sprintf("SELECT `id` FROM `blogs`")
	if where != "" {
		query += " WHERE "
		query += where
	}

	rows, err := tx.Query(query, args...)
	if err != nil {
		tx.err = err
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
			tx.err = err
			return nil, err
		}
		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		tx.err = err
		return nil, fmt.Errorf("Blog query limit result error: %v", err)
	}
	return
}

func (tx *_BlogMySQLTx) queryCount(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`id`) FROM `blogs` %s", where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		tx.err = err
		return 0, fmt.Errorf("Blog query limit error: %v", err)
	}
	defer rows.Close()

	var count int64
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			tx.err = err
			return 0, err
		}
		break
	}

	return count, nil
}

func (tx *_BlogMySQLTx) Fetch(id interface{}) (*Blog, error) {
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` WHERE `Id` = (%s)", strings.Join(obj.GetColumns(), ","), fmt.Sprint(id))
	objs, err := tx.FetchBySQL(query)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*Blog), nil
	}
	return nil, fmt.Errorf("Blog fetch record not found")
}

func (tx *_BlogMySQLTx) FetchByIds(ids []interface{}) ([]*Blog, error) {
	if len(ids) == 0 {
		return []*Blog{}, nil
	}

	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` WHERE `Id` IN (%s)", strings.Join(obj.GetColumns(), ","), orm.SliceJoin(ids, ","))
	objs, err := tx.FetchBySQL(query)
	if err != nil {
		return nil, err
	}
	results := make([]*Blog, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*Blog))
	}
	return results, nil
}

func (tx *_BlogMySQLTx) Search(where string, args ...interface{}) (results []interface{}, err error) {
	obj := BlogMgr.NewBlog()
	if where != "" {
		where = " WHERE " + where
	}
	sql := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(obj.GetColumns(), ","), where)
	return tx.FetchBySQL(sql, args...)
}

func (tx *_BlogMySQLTx) SearchCount(where string, args ...interface{}) (int64, error) {
	if where != "" {
		where = " WHERE " + where
	}
	return tx.queryCount(where, args...)
}

func (tx *_BlogMySQLTx) FetchBySQL(sql string, args ...interface{}) (results []interface{}, err error) {
	rows, err := tx.Query(sql, args...)
	if err != nil {
		tx.err = err
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
			tx.err = err
			return nil, err
		}

		result.CreatedAt = orm.TimeParse(CreatedAt)

		result.UpdatedAt = orm.TimeParse(UpdatedAt)

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		tx.err = err
		return nil, fmt.Errorf("Blog fetch result error: %v", err)
	}
	return
}
