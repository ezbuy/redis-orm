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

func (obj *Blog) GetPrimaryKey() PrimaryKey {
	pk := BlogMgr.NewPrimaryKey()
	pk.Id = obj.Id
	pk.UserId = obj.UserId
	return pk
}

//! primary key

type IdUserIdOfBlogPK struct {
	Id     int32
	UserId int32
}

func (m *_BlogMgr) NewPrimaryKey() *IdUserIdOfBlogPK {
	return &IdUserIdOfBlogPK{}
}

func (u *IdUserIdOfBlogPK) Key() string {
	strs := []string{
		"Id",
		fmt.Sprint(u.Id),
		"UserId",
		fmt.Sprint(u.UserId),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *IdUserIdOfBlogPK) Parse(key string) error {
	arr := strings.Split(key, ":")
	if len(arr)%2 != 0 {
		return fmt.Errorf("key (%s) format error", key)
	}
	kv := map[string]string{}
	for i := 0; i < len(arr)/2; i++ {
		kv[arr[2*i]] = arr[2*i+1]
	}
	vId, ok := kv["Id"]
	if !ok {
		return fmt.Errorf("key (%s) without (Id) field", key)
	}
	if err := orm.StringScan(vId, &(u.Id)); err != nil {
		return err
	}
	vUserId, ok := kv["UserId"]
	if !ok {
		return fmt.Errorf("key (%s) without (UserId) field", key)
	}
	if err := orm.StringScan(vUserId, &(u.UserId)); err != nil {
		return err
	}
	return nil
}

func (u *IdUserIdOfBlogPK) SQLFormat() string {
	conditions := []string{
		"id = ?",
		"user_id = ?",
	}
	return orm.SQLWhere(conditions)
}

func (u *IdUserIdOfBlogPK) SQLParams() []interface{} {
	return []interface{}{
		u.Id,
		u.UserId,
	}
}

func (u *IdUserIdOfBlogPK) Columns() []string {
	return []string{
		"`id`",
		"`user_id`",
	}
}

//! uniques

//! indexes

type StatusOfBlogIDX struct {
	Status int32
	offset int
	limit  int
}

func (u *StatusOfBlogIDX) Key() string {
	strs := []string{
		"Status",
		fmt.Sprint(u.Status),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *StatusOfBlogIDX) SQLFormat(limit bool) string {
	conditions := []string{
		"status = ?",
	}
	if limit {
		return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return orm.SQLWhere(conditions)
}

func (u *StatusOfBlogIDX) SQLParams() []interface{} {
	return []interface{}{
		u.Status,
	}
}

func (u *StatusOfBlogIDX) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *StatusOfBlogIDX) Limit(n int) {
	u.limit = n
}

func (u *StatusOfBlogIDX) Offset(n int) {
	u.offset = n
}

func (u *StatusOfBlogIDX) IDXRelation() IndexRelation {
	return nil
}

//! ranges

type _BlogDBMgr struct {
	db orm.DB
}

func (m *_BlogMgr) DB(db orm.DB) *_BlogDBMgr {
	return BlogDBMgr(db)
}

func BlogDBMgr(db orm.DB) *_BlogDBMgr {
	if db == nil {
		panic(fmt.Errorf("BlogDBMgr init need db"))
	}
	return &_BlogDBMgr{db: db}
}

func (m *_BlogDBMgr) Search(where string, orderby string, limit string, args ...interface{}) ([]*Blog, error) {
	obj := BlogMgr.NewBlog()
	conditions := []string{where, orderby, limit}
	query := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(obj.GetColumns(), ","), strings.Join(conditions, " "))
	objs, err := m.FetchBySQL(query, args...)
	if err != nil {
		return nil, err
	}
	results := make([]*Blog, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*Blog))
	}
	return results, nil
}

func (m *_BlogDBMgr) SearchConditions(conditions []string, orderby string, offset int, limit int, args ...interface{}) ([]*Blog, error) {
	obj := BlogMgr.NewBlog()
	q := fmt.Sprintf("SELECT %s FROM `blogs` %s %s %s",
		strings.Join(obj.GetColumns(), ","),
		orm.SQLWhere(conditions),
		orderby,
		orm.SQLOffsetLimit(offset, limit))
	objs, err := m.FetchBySQL(q, args...)
	if err != nil {
		return nil, err
	}
	results := make([]*Blog, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*Blog))
	}
	return results, nil
}

func (m *_BlogDBMgr) SearchCount(where string, args ...interface{}) (int64, error) {
	return m.queryCount(where, args...)
}

func (m *_BlogDBMgr) SearchConditionsCount(conditions []string, args ...interface{}) (int64, error) {
	return m.queryCount(orm.SQLWhere(conditions), args...)
}

func (m *_BlogDBMgr) FetchBySQL(q string, args ...interface{}) (results []interface{}, err error) {
	rows, err := m.db.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("Blog fetch error: %v", err)
	}
	defer rows.Close()

	var CreatedAt string
	var UpdatedAt string

	for rows.Next() {
		var result Blog
		err = rows.Scan(&(result.Id), &(result.UserId), &(result.Title), &(result.Content), &(result.Status), &(result.Readed), &CreatedAt, &UpdatedAt)
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
func (m *_BlogDBMgr) Fetch(pk PrimaryKey) (*Blog, error) {
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*Blog), nil
	}
	return nil, fmt.Errorf("Blog fetch record not found")
}

func (m *_BlogDBMgr) FetchByPrimaryKeys(pks []PrimaryKey) ([]*Blog, error) {
	results := make([]*Blog, 0, len(pks))
	for _, pk := range pks {
		obj, err := m.Fetch(pk)
		if err != nil {
			return nil, err
		}
		results = append(results, obj)
	}
	return results, nil
}

func (m *_BlogDBMgr) FindOne(unique Unique) (PrimaryKey, error) {
	objs, err := m.queryLimit(unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("Blog find record not found")
}

func (m *_BlogDBMgr) FindOneFetch(unique Unique) (*Blog, error) {
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

func (m *_BlogDBMgr) Find(index Index) ([]PrimaryKey, error) {
	return m.queryLimit(index.SQLFormat(true), index.SQLLimit(), index.SQLParams()...)
}

func (m *_BlogDBMgr) FindFetch(index Index) ([]*Blog, error) {
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

func (m *_BlogDBMgr) FindCount(index Index) (int64, error) {
	return m.queryCount(index.SQLFormat(false), index.SQLParams()...)
}

func (m *_BlogDBMgr) Range(scope Range) ([]PrimaryKey, error) {
	return m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (m *_BlogDBMgr) RangeFetch(scope Range) ([]*Blog, error) {
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

func (m *_BlogDBMgr) RangeCount(scope Range) (int64, error) {
	return m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
}

func (m *_BlogDBMgr) RangeRevert(scope Range) ([]PrimaryKey, error) {
	scope.Revert(true)
	return m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (m *_BlogDBMgr) RangeRevertFetch(scope Range) ([]*Blog, error) {
	scope.Revert(true)
	return m.RangeFetch(scope)
}

func (m *_BlogDBMgr) queryLimit(where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := BlogMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM `blogs` %s", strings.Join(pk.Columns(), ","), where)
	rows, err := m.db.Query(query, args...)
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

		result := BlogMgr.NewPrimaryKey()
		err = rows.Scan(&(result.Id), &(result.UserId))
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Blog query limit result error: %v", err)
	}
	return
}

func (m *_BlogDBMgr) queryCount(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`id`) FROM `blogs` %s", where)
	rows, err := m.db.Query(query, args...)
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

func (m *_BlogDBMgr) BatchCreate(objs []*Blog) (int64, error) {
	if len(objs) == 0 {
		return 0, nil
	}

	params := make([]string, 0, len(objs))
	values := make([]interface{}, 0, len(objs)*8)
	for _, obj := range objs {
		params = append(params, fmt.Sprintf("(%s)", strings.Join(orm.NewStringSlice(8, "?"), ",")))
		values = append(values, obj.Id)
		values = append(values, obj.UserId)
		values = append(values, obj.Title)
		values = append(values, obj.Content)
		values = append(values, obj.Status)
		values = append(values, obj.Readed)
		values = append(values, orm.TimeFormat(obj.CreatedAt))
		values = append(values, orm.TimeFormat(obj.UpdatedAt))
	}
	query := fmt.Sprintf("INSERT INTO `blogs`(%s) VALUES %s", strings.Join(objs[0].GetColumns(), ","), strings.Join(params, ","))
	result, err := m.db.Exec(query, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// argument example:
// set:"a=?, b=?"
// where:"c=? and d=?"
// params:[]interface{}{"a", "b", "c", "d"}...
func (m *_BlogDBMgr) UpdateBySQL(set, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("UPDATE `blogs` SET %s", set)
	if where != "" {
		query = fmt.Sprintf("UPDATE `blogs` SET %s WHERE %s", set, where)
	}
	result, err := m.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_BlogDBMgr) Create(obj *Blog) (int64, error) {
	params := orm.NewStringSlice(8, "?")
	q := fmt.Sprintf("INSERT INTO `blogs`(%s) VALUES(%s)",
		strings.Join(obj.GetColumns(), ","),
		strings.Join(params, ","))

	values := make([]interface{}, 0, 8)
	values = append(values, obj.Id)
	values = append(values, obj.UserId)
	values = append(values, obj.Title)
	values = append(values, obj.Content)
	values = append(values, obj.Status)
	values = append(values, obj.Readed)
	values = append(values, orm.TimeFormat(obj.CreatedAt))
	values = append(values, orm.TimeFormat(obj.UpdatedAt))
	result, err := m.db.Exec(q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_BlogDBMgr) Update(obj *Blog) (int64, error) {
	columns := []string{
		"`title` = ?",
		"`content` = ?",
		"`status` = ?",
		"`readed` = ?",
		"`created_at` = ?",
		"`updated_at` = ?",
	}

	pk := obj.GetPrimaryKey()
	q := fmt.Sprintf("UPDATE `blogs` SET %s %s", strings.Join(columns, ","), pk.SQLFormat())
	values := make([]interface{}, 0, 8-2)
	values = append(values, obj.Title)
	values = append(values, obj.Content)
	values = append(values, obj.Status)
	values = append(values, obj.Readed)
	values = append(values, orm.TimeFormat(obj.CreatedAt))
	values = append(values, orm.TimeFormat(obj.UpdatedAt))
	values = append(values, pk.SQLParams()...)

	result, err := m.db.Exec(q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_BlogDBMgr) Save(obj *Blog) (int64, error) {
	affected, err := m.Update(obj)
	if err != nil {
		return affected, err
	}
	if affected == 0 {
		return m.Create(obj)
	}
	return affected, err
}

func (m *_BlogDBMgr) Delete(obj *Blog) (int64, error) {
	pk := obj.GetPrimaryKey()
	return m.DeleteByPrimaryKey(pk)
}

func (m *_BlogDBMgr) DeleteByPrimaryKey(pk PrimaryKey) (int64, error) {
	q := fmt.Sprintf("DELETE FROM `blogs` %s", pk.SQLFormat())
	result, err := m.db.Exec(q, pk.SQLParams()...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_BlogDBMgr) DeleteBySQL(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("DELETE FROM `blogs`")
	if where != "" {
		query = fmt.Sprintf("DELETE FROM `blogs` WHERE %s", where)
	}
	result, err := m.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
