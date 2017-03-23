package model

import (
	"database/sql"
	"fmt"
	"github.com/ezbuy/redis-orm/orm"
	"gopkg.in/go-playground/validator.v9"
	"strings"
	"time"
)

var (
	_ sql.DB
	_ time.Time
	_ fmt.Formatter
	_ strings.Reader
	_ orm.VSet
	_ validator.Validate
)

type UserBlogs struct {
	UserId int32 `db:"user_id"`
	BlogId int32 `db:"blog_id"`
}

type _UserBlogsMgr struct {
}

var UserBlogsMgr *_UserBlogsMgr

func (m *_UserBlogsMgr) NewUserBlogs() *UserBlogs {
	return &UserBlogs{}
}

//! object function

func (obj *UserBlogs) GetNameSpace() string {
	return "model"
}

func (obj *UserBlogs) GetClassName() string {
	return "UserBlogs"
}

func (obj *UserBlogs) GetTableName() string {
	return "user_blogs"
}

func (obj *UserBlogs) GetColumns() []string {
	columns := []string{
		"`user_id`",
		"`blog_id`",
	}
	return columns
}

func (obj *UserBlogs) GetPrimaryKey() PrimaryKey {
	pk := UserBlogsMgr.NewPrimaryKey()
	pk.UserId = obj.UserId
	pk.BlogId = obj.BlogId
	return pk
}

func (obj *UserBlogs) Validate() error {
	validate := validator.New()
	return validate.Struct(obj)
}

//! primary key

type UserIdBlogIdOfUserBlogsPK struct {
	UserId int32
	BlogId int32
}

func (m *_UserBlogsMgr) NewPrimaryKey() *UserIdBlogIdOfUserBlogsPK {
	return &UserIdBlogIdOfUserBlogsPK{}
}

func (u *UserIdBlogIdOfUserBlogsPK) Key() string {
	strs := []string{
		"UserId",
		fmt.Sprint(u.UserId),
		"BlogId",
		fmt.Sprint(u.BlogId),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *UserIdBlogIdOfUserBlogsPK) Parse(key string) error {
	arr := strings.Split(key, ":")
	if len(arr)%2 != 0 {
		return fmt.Errorf("key (%s) format error", key)
	}
	kv := map[string]string{}
	for i := 0; i < len(arr)/2; i++ {
		kv[arr[2*i]] = arr[2*i+1]
	}
	vUserId, ok := kv["UserId"]
	if !ok {
		return fmt.Errorf("key (%s) without (UserId) field", key)
	}
	if err := orm.StringScan(vUserId, &(u.UserId)); err != nil {
		return err
	}
	vBlogId, ok := kv["BlogId"]
	if !ok {
		return fmt.Errorf("key (%s) without (BlogId) field", key)
	}
	if err := orm.StringScan(vBlogId, &(u.BlogId)); err != nil {
		return err
	}
	return nil
}

func (u *UserIdBlogIdOfUserBlogsPK) SQLFormat() string {
	conditions := []string{
		"user_id = ?",
		"blog_id = ?",
	}
	return orm.SQLWhere(conditions)
}

func (u *UserIdBlogIdOfUserBlogsPK) SQLParams() []interface{} {
	return []interface{}{
		u.UserId,
		u.BlogId,
	}
}

func (u *UserIdBlogIdOfUserBlogsPK) Columns() []string {
	return []string{
		"`user_id`",
		"`blog_id`",
	}
}

//! uniques

type UserIdOfUserBlogsUK struct {
	UserId int32
}

func (u *UserIdOfUserBlogsUK) Key() string {
	strs := []string{
		"UserId",
		fmt.Sprint(u.UserId),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *UserIdOfUserBlogsUK) SQLFormat(limit bool) string {
	conditions := []string{
		"user_id = ?",
	}
	return orm.SQLWhere(conditions)
}

func (u *UserIdOfUserBlogsUK) SQLParams() []interface{} {
	return []interface{}{
		u.UserId,
	}
}

func (u *UserIdOfUserBlogsUK) SQLLimit() int {
	return 1
}

func (u *UserIdOfUserBlogsUK) Limit(n int) {
}

func (u *UserIdOfUserBlogsUK) Offset(n int) {
}

func (u *UserIdOfUserBlogsUK) UKRelation() UniqueRelation {
	return nil
}

type BlogIdOfUserBlogsUK struct {
	BlogId int32
}

func (u *BlogIdOfUserBlogsUK) Key() string {
	strs := []string{
		"BlogId",
		fmt.Sprint(u.BlogId),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *BlogIdOfUserBlogsUK) SQLFormat(limit bool) string {
	conditions := []string{
		"blog_id = ?",
	}
	return orm.SQLWhere(conditions)
}

func (u *BlogIdOfUserBlogsUK) SQLParams() []interface{} {
	return []interface{}{
		u.BlogId,
	}
}

func (u *BlogIdOfUserBlogsUK) SQLLimit() int {
	return 1
}

func (u *BlogIdOfUserBlogsUK) Limit(n int) {
}

func (u *BlogIdOfUserBlogsUK) Offset(n int) {
}

func (u *BlogIdOfUserBlogsUK) UKRelation() UniqueRelation {
	return nil
}

//! indexes

//! ranges

type UserIdOfUserBlogsRNG struct {
	UserIdBegin  int64
	UserIdEnd    int64
	offset       int
	limit        int
	includeBegin bool
	includeEnd   bool
	revert       bool
}

func (u *UserIdOfUserBlogsRNG) Key() string {
	strs := []string{
		"UserId",
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *UserIdOfUserBlogsRNG) beginOp() string {
	if u.includeBegin {
		return ">="
	}
	return ">"
}
func (u *UserIdOfUserBlogsRNG) endOp() string {
	if u.includeBegin {
		return "<="
	}
	return "<"
}

func (u *UserIdOfUserBlogsRNG) SQLFormat(limit bool) string {
	conditions := []string{}
	if u.UserIdBegin != u.UserIdEnd {
		if u.UserIdBegin != -1 {
			conditions = append(conditions, fmt.Sprintf("user_id %s ?", u.beginOp()))
		}
		if u.UserIdEnd != -1 {
			conditions = append(conditions, fmt.Sprintf("user_id %s ?", u.endOp()))
		}
	}
	if limit {
		return fmt.Sprintf("%s %s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("UserId", u.revert), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("UserId", u.revert))
}

func (u *UserIdOfUserBlogsRNG) SQLParams() []interface{} {
	params := []interface{}{}
	if u.UserIdBegin != u.UserIdEnd {
		if u.UserIdBegin != -1 {
			params = append(params, u.UserIdBegin)
		}
		if u.UserIdEnd != -1 {
			params = append(params, u.UserIdEnd)
		}
	}
	return params
}

func (u *UserIdOfUserBlogsRNG) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *UserIdOfUserBlogsRNG) Limit(n int) {
	u.limit = n
}

func (u *UserIdOfUserBlogsRNG) Offset(n int) {
	u.offset = n
}

func (u *UserIdOfUserBlogsRNG) Begin() int64 {
	start := u.UserIdBegin
	if start == -1 || start == 0 {
		start = 0
	}
	if start > 0 {
		if !u.includeBegin {
			start = start + 1
		}
	}
	return start
}

func (u *UserIdOfUserBlogsRNG) End() int64 {
	stop := u.UserIdEnd
	if stop == 0 || stop == -1 {
		stop = -1
	}
	if stop > 0 {
		if !u.includeBegin {
			stop = stop - 1
		}
	}
	return stop
}

func (u *UserIdOfUserBlogsRNG) Revert(b bool) {
	u.revert = b
}

func (u *UserIdOfUserBlogsRNG) IncludeBegin(f bool) {
	u.includeBegin = f
}

func (u *UserIdOfUserBlogsRNG) IncludeEnd(f bool) {
	u.includeEnd = f
}

func (u *UserIdOfUserBlogsRNG) RNGRelation() RangeRelation {
	return nil
}

type BlogIdOfUserBlogsRNG struct {
	BlogIdBegin  int64
	BlogIdEnd    int64
	offset       int
	limit        int
	includeBegin bool
	includeEnd   bool
	revert       bool
}

func (u *BlogIdOfUserBlogsRNG) Key() string {
	strs := []string{
		"BlogId",
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *BlogIdOfUserBlogsRNG) beginOp() string {
	if u.includeBegin {
		return ">="
	}
	return ">"
}
func (u *BlogIdOfUserBlogsRNG) endOp() string {
	if u.includeBegin {
		return "<="
	}
	return "<"
}

func (u *BlogIdOfUserBlogsRNG) SQLFormat(limit bool) string {
	conditions := []string{}
	if u.BlogIdBegin != u.BlogIdEnd {
		if u.BlogIdBegin != -1 {
			conditions = append(conditions, fmt.Sprintf("blog_id %s ?", u.beginOp()))
		}
		if u.BlogIdEnd != -1 {
			conditions = append(conditions, fmt.Sprintf("blog_id %s ?", u.endOp()))
		}
	}
	if limit {
		return fmt.Sprintf("%s %s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("BlogId", u.revert), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("BlogId", u.revert))
}

func (u *BlogIdOfUserBlogsRNG) SQLParams() []interface{} {
	params := []interface{}{}
	if u.BlogIdBegin != u.BlogIdEnd {
		if u.BlogIdBegin != -1 {
			params = append(params, u.BlogIdBegin)
		}
		if u.BlogIdEnd != -1 {
			params = append(params, u.BlogIdEnd)
		}
	}
	return params
}

func (u *BlogIdOfUserBlogsRNG) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *BlogIdOfUserBlogsRNG) Limit(n int) {
	u.limit = n
}

func (u *BlogIdOfUserBlogsRNG) Offset(n int) {
	u.offset = n
}

func (u *BlogIdOfUserBlogsRNG) Begin() int64 {
	start := u.BlogIdBegin
	if start == -1 || start == 0 {
		start = 0
	}
	if start > 0 {
		if !u.includeBegin {
			start = start + 1
		}
	}
	return start
}

func (u *BlogIdOfUserBlogsRNG) End() int64 {
	stop := u.BlogIdEnd
	if stop == 0 || stop == -1 {
		stop = -1
	}
	if stop > 0 {
		if !u.includeBegin {
			stop = stop - 1
		}
	}
	return stop
}

func (u *BlogIdOfUserBlogsRNG) Revert(b bool) {
	u.revert = b
}

func (u *BlogIdOfUserBlogsRNG) IncludeBegin(f bool) {
	u.includeBegin = f
}

func (u *BlogIdOfUserBlogsRNG) IncludeEnd(f bool) {
	u.includeEnd = f
}

func (u *BlogIdOfUserBlogsRNG) RNGRelation() RangeRelation {
	return nil
}

type _UserBlogsDBMgr struct {
	db orm.DB
}

func (m *_UserBlogsMgr) DB(db orm.DB) *_UserBlogsDBMgr {
	return UserBlogsDBMgr(db)
}

func UserBlogsDBMgr(db orm.DB) *_UserBlogsDBMgr {
	if db == nil {
		panic(fmt.Errorf("UserBlogsDBMgr init need db"))
	}
	return &_UserBlogsDBMgr{db: db}
}

func (m *_UserBlogsDBMgr) Search(where string, orderby string, limit string, args ...interface{}) ([]*UserBlogs, error) {
	obj := UserBlogsMgr.NewUserBlogs()
	conditions := []string{where, orderby, limit}
	query := fmt.Sprintf("SELECT %s FROM `user_blogs` %s", strings.Join(obj.GetColumns(), ","), strings.Join(conditions, " "))
	objs, err := m.FetchBySQL(query, args...)
	if err != nil {
		return nil, err
	}
	results := make([]*UserBlogs, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*UserBlogs))
	}
	return results, nil
}

func (m *_UserBlogsDBMgr) SearchConditions(conditions []string, orderby string, offset int, limit int, args ...interface{}) ([]*UserBlogs, error) {
	obj := UserBlogsMgr.NewUserBlogs()
	q := fmt.Sprintf("SELECT %s FROM `user_blogs` %s %s %s",
		strings.Join(obj.GetColumns(), ","),
		orm.SQLWhere(conditions),
		orderby,
		orm.SQLOffsetLimit(offset, limit))
	objs, err := m.FetchBySQL(q, args...)
	if err != nil {
		return nil, err
	}
	results := make([]*UserBlogs, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*UserBlogs))
	}
	return results, nil
}

func (m *_UserBlogsDBMgr) SearchCount(where string, args ...interface{}) (int64, error) {
	return m.queryCount(where, args...)
}

func (m *_UserBlogsDBMgr) SearchConditionsCount(conditions []string, args ...interface{}) (int64, error) {
	return m.queryCount(orm.SQLWhere(conditions), args...)
}

func (m *_UserBlogsDBMgr) FetchBySQL(q string, args ...interface{}) (results []interface{}, err error) {
	rows, err := m.db.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("UserBlogs fetch error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var result UserBlogs
		err = rows.Scan(&(result.UserId), &(result.BlogId))
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
func (m *_UserBlogsDBMgr) Fetch(pk PrimaryKey) (*UserBlogs, error) {
	obj := UserBlogsMgr.NewUserBlogs()
	query := fmt.Sprintf("SELECT %s FROM `user_blogs` %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*UserBlogs), nil
	}
	return nil, fmt.Errorf("UserBlogs fetch record not found")
}

func (m *_UserBlogsDBMgr) FetchByPrimaryKeys(pks []PrimaryKey) ([]*UserBlogs, error) {
	results := make([]*UserBlogs, 0, len(pks))
	for _, pk := range pks {
		obj, err := m.Fetch(pk)
		if err != nil {
			return nil, err
		}
		results = append(results, obj)
	}
	return results, nil
}

func (m *_UserBlogsDBMgr) FindOne(unique Unique) (PrimaryKey, error) {
	objs, err := m.queryLimit(unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("UserBlogs find record not found")
}

func (m *_UserBlogsDBMgr) FindOneFetch(unique Unique) (*UserBlogs, error) {
	obj := UserBlogsMgr.NewUserBlogs()
	query := fmt.Sprintf("SELECT %s FROM `user_blogs` %s", strings.Join(obj.GetColumns(), ","), unique.SQLFormat(true))
	objs, err := m.FetchBySQL(query, unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*UserBlogs), nil
	}
	return nil, fmt.Errorf("none record")
}

func (m *_UserBlogsDBMgr) Find(index Index) ([]PrimaryKey, error) {
	return m.queryLimit(index.SQLFormat(true), index.SQLLimit(), index.SQLParams()...)
}

func (m *_UserBlogsDBMgr) FindFetch(index Index) ([]*UserBlogs, error) {
	obj := UserBlogsMgr.NewUserBlogs()
	query := fmt.Sprintf("SELECT %s FROM `user_blogs` %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	objs, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return nil, err
	}
	results := make([]*UserBlogs, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*UserBlogs))
	}
	return results, nil
}

func (m *_UserBlogsDBMgr) FindCount(index Index) (int64, error) {
	return m.queryCount(index.SQLFormat(false), index.SQLParams()...)
}

func (m *_UserBlogsDBMgr) Range(scope Range) ([]PrimaryKey, error) {
	return m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (m *_UserBlogsDBMgr) RangeFetch(scope Range) ([]*UserBlogs, error) {
	obj := UserBlogsMgr.NewUserBlogs()
	query := fmt.Sprintf("SELECT %s FROM `user_blogs` %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	objs, err := m.FetchBySQL(query, scope.SQLParams()...)
	if err != nil {
		return nil, err
	}
	results := make([]*UserBlogs, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*UserBlogs))
	}
	return results, nil
}

func (m *_UserBlogsDBMgr) RangeCount(scope Range) (int64, error) {
	return m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
}

func (m *_UserBlogsDBMgr) RangeRevert(scope Range) ([]PrimaryKey, error) {
	scope.Revert(true)
	return m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (m *_UserBlogsDBMgr) RangeRevertFetch(scope Range) ([]*UserBlogs, error) {
	scope.Revert(true)
	return m.RangeFetch(scope)
}

func (m *_UserBlogsDBMgr) queryLimit(where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := UserBlogsMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM `user_blogs` %s", strings.Join(pk.Columns(), ","), where)
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("UserBlogs query limit error: %v", err)
	}
	defer rows.Close()

	offset := 0

	for rows.Next() {
		if limit >= 0 && offset >= limit {
			break
		}
		offset++

		result := UserBlogsMgr.NewPrimaryKey()
		err = rows.Scan(&(result.UserId), &(result.BlogId))
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("UserBlogs query limit result error: %v", err)
	}
	return
}

func (m *_UserBlogsDBMgr) queryCount(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`user_id`) FROM `user_blogs` %s", where)
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return 0, fmt.Errorf("UserBlogs query count error: %v", err)
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

func (m *_UserBlogsDBMgr) BatchCreate(objs []*UserBlogs) (int64, error) {
	if len(objs) == 0 {
		return 0, nil
	}

	params := make([]string, 0, len(objs))
	values := make([]interface{}, 0, len(objs)*2)
	for _, obj := range objs {
		params = append(params, fmt.Sprintf("(%s)", strings.Join(orm.NewStringSlice(2, "?"), ",")))
		values = append(values, obj.UserId)
		values = append(values, obj.BlogId)
	}
	query := fmt.Sprintf("INSERT INTO `user_blogs`(%s) VALUES %s", strings.Join(objs[0].GetColumns(), ","), strings.Join(params, ","))
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
func (m *_UserBlogsDBMgr) UpdateBySQL(set, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("UPDATE `user_blogs` SET %s", set)
	if where != "" {
		query = fmt.Sprintf("UPDATE `user_blogs` SET %s WHERE %s", set, where)
	}
	result, err := m.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_UserBlogsDBMgr) Create(obj *UserBlogs) (int64, error) {
	params := orm.NewStringSlice(2, "?")
	q := fmt.Sprintf("INSERT INTO `user_blogs`(%s) VALUES(%s)",
		strings.Join(obj.GetColumns(), ","),
		strings.Join(params, ","))

	values := make([]interface{}, 0, 2)
	values = append(values, obj.UserId)
	values = append(values, obj.BlogId)
	result, err := m.db.Exec(q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_UserBlogsDBMgr) Update(obj *UserBlogs) (int64, error) {
	columns := []string{}

	pk := obj.GetPrimaryKey()
	q := fmt.Sprintf("UPDATE `user_blogs` SET %s %s", strings.Join(columns, ","), pk.SQLFormat())
	values := make([]interface{}, 0, 2-2)
	values = append(values, pk.SQLParams()...)

	result, err := m.db.Exec(q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_UserBlogsDBMgr) Save(obj *UserBlogs) (int64, error) {
	affected, err := m.Update(obj)
	if err != nil {
		return affected, err
	}
	if affected == 0 {
		return m.Create(obj)
	}
	return affected, err
}

func (m *_UserBlogsDBMgr) Delete(obj *UserBlogs) (int64, error) {
	pk := obj.GetPrimaryKey()
	return m.DeleteByPrimaryKey(pk)
}

func (m *_UserBlogsDBMgr) DeleteByPrimaryKey(pk PrimaryKey) (int64, error) {
	q := fmt.Sprintf("DELETE FROM `user_blogs` %s", pk.SQLFormat())
	result, err := m.db.Exec(q, pk.SQLParams()...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_UserBlogsDBMgr) DeleteBySQL(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("DELETE FROM `user_blogs`")
	if where != "" {
		query = fmt.Sprintf("DELETE FROM `user_blogs` WHERE %s", where)
	}
	result, err := m.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
