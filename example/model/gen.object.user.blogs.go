package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/ezbuy/redis-orm/orm"
	"gopkg.in/go-playground/validator.v9"
)

var (
	_ sql.DB
	_ time.Time
	_ fmt.Formatter
	_ strings.Reader
	_ orm.VSet
	_ validator.Validate
	_ context.Context
)

type UserBlogs struct {
	UserId int32 `db:"user_id"`
	BlogId int32 `db:"blog_id"`
}

var UserBlogsColumns = struct {
	UserId string
	BlogId string
}{
	"user_id",
	"blog_id",
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
		"user_blogs.`user_id`",
		"user_blogs.`blog_id`",
	}
	return columns
}

func (obj *UserBlogs) GetNoneIncrementColumns() []string {
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
		"`user_id` = ?",
		"`blog_id` = ?",
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

//! indexes

//! ranges

type UserIdBlogIdOfUserBlogsRNG struct {
	UserId       int32
	BlogIdBegin  int64
	BlogIdEnd    int64
	offset       int
	limit        int
	includeBegin bool
	includeEnd   bool
	revert       bool
}

func (u *UserIdBlogIdOfUserBlogsRNG) Key() string {
	strs := []string{
		"UserId",
		fmt.Sprint(u.UserId),
		"BlogId",
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *UserIdBlogIdOfUserBlogsRNG) beginOp() string {
	if u.includeBegin {
		return ">="
	}
	return ">"
}
func (u *UserIdBlogIdOfUserBlogsRNG) endOp() string {
	if u.includeBegin {
		return "<="
	}
	return "<"
}

func (u *UserIdBlogIdOfUserBlogsRNG) SQLFormat(limit bool) string {
	conditions := []string{}
	conditions = append(conditions, "`user_id` = ?")
	if u.BlogIdBegin != u.BlogIdEnd {
		if u.BlogIdBegin != -1 {
			conditions = append(conditions, fmt.Sprintf("`blog_id` %s ?", u.beginOp()))
		}
		if u.BlogIdEnd != -1 {
			conditions = append(conditions, fmt.Sprintf("`blog_id` %s ?", u.endOp()))
		}
	}
	if limit {
		return fmt.Sprintf("%s %s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("`blog_id`", u.revert), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("`blog_id`", u.revert))
}

func (u *UserIdBlogIdOfUserBlogsRNG) SQLParams() []interface{} {
	params := []interface{}{
		u.UserId,
	}
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

func (u *UserIdBlogIdOfUserBlogsRNG) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *UserIdBlogIdOfUserBlogsRNG) Limit(n int) {
	u.limit = n
}

func (u *UserIdBlogIdOfUserBlogsRNG) Offset(n int) {
	u.offset = n
}

func (u *UserIdBlogIdOfUserBlogsRNG) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *UserIdBlogIdOfUserBlogsRNG) Begin() int64 {
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

func (u *UserIdBlogIdOfUserBlogsRNG) End() int64 {
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

func (u *UserIdBlogIdOfUserBlogsRNG) Revert(b bool) {
	u.revert = b
}

func (u *UserIdBlogIdOfUserBlogsRNG) IncludeBegin(f bool) {
	u.includeBegin = f
}

func (u *UserIdBlogIdOfUserBlogsRNG) IncludeEnd(f bool) {
	u.includeEnd = f
}

func (u *UserIdBlogIdOfUserBlogsRNG) RNGRelation(store *orm.RedisStore) RangeRelation {
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

	if limit = strings.ToUpper(strings.TrimSpace(limit)); limit != "" && !strings.HasPrefix(limit, "LIMIT") {
		limit = "LIMIT " + limit
	}

	conditions := []string{where, orderby, limit}
	query := fmt.Sprintf("SELECT %s FROM user_blogs %s", strings.Join(obj.GetColumns(), ","), strings.Join(conditions, " "))
	return m.FetchBySQL(query, args...)
}

func (m *_UserBlogsDBMgr) SearchContext(ctx context.Context, where string, orderby string, limit string, args ...interface{}) ([]*UserBlogs, error) {
	obj := UserBlogsMgr.NewUserBlogs()

	if limit = strings.ToUpper(strings.TrimSpace(limit)); limit != "" && !strings.HasPrefix(limit, "LIMIT") {
		limit = "LIMIT " + limit
	}

	conditions := []string{where, orderby, limit}
	query := fmt.Sprintf("SELECT %s FROM user_blogs %s", strings.Join(obj.GetColumns(), ","), strings.Join(conditions, " "))
	return m.FetchBySQLContext(ctx, query, args...)
}

func (m *_UserBlogsDBMgr) SearchConditions(conditions []string, orderby string, offset int, limit int, args ...interface{}) ([]*UserBlogs, error) {
	obj := UserBlogsMgr.NewUserBlogs()
	q := fmt.Sprintf("SELECT %s FROM user_blogs %s %s %s",
		strings.Join(obj.GetColumns(), ","),
		orm.SQLWhere(conditions),
		orderby,
		orm.SQLOffsetLimit(offset, limit))

	return m.FetchBySQL(q, args...)
}

func (m *_UserBlogsDBMgr) SearchConditionsContext(ctx context.Context, conditions []string, orderby string, offset int, limit int, args ...interface{}) ([]*UserBlogs, error) {
	obj := UserBlogsMgr.NewUserBlogs()
	q := fmt.Sprintf("SELECT %s FROM user_blogs %s %s %s",
		strings.Join(obj.GetColumns(), ","),
		orm.SQLWhere(conditions),
		orderby,
		orm.SQLOffsetLimit(offset, limit))

	return m.FetchBySQLContext(ctx, q, args...)
}

func (m *_UserBlogsDBMgr) SearchCount(where string, args ...interface{}) (int64, error) {
	return m.queryCount(where, args...)
}

func (m *_UserBlogsDBMgr) SearchCountContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	return m.queryCountContext(ctx, where, args...)
}

func (m *_UserBlogsDBMgr) SearchConditionsCount(conditions []string, args ...interface{}) (int64, error) {
	return m.queryCount(orm.SQLWhere(conditions), args...)
}

func (m *_UserBlogsDBMgr) SearchConditionsCountContext(ctx context.Context, conditions []string, args ...interface{}) (int64, error) {
	return m.queryCountContext(ctx, orm.SQLWhere(conditions), args...)
}

func (m *_UserBlogsDBMgr) FetchBySQL(q string, args ...interface{}) (results []*UserBlogs, err error) {
	rows, err := m.db.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("UserBlogs fetch error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var result UserBlogs
		err = rows.Scan(&(result.UserId), &(result.BlogId))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("UserBlogs fetch result error: %v", err)
	}
	return
}

func (m *_UserBlogsDBMgr) FetchBySQLContext(ctx context.Context, q string, args ...interface{}) (results []*UserBlogs, err error) {
	rows, err := m.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("UserBlogs fetch error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var result UserBlogs
		err = rows.Scan(&(result.UserId), &(result.BlogId))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("UserBlogs fetch result error: %v", err)
	}
	return
}
func (m *_UserBlogsDBMgr) Exist(pk PrimaryKey) (bool, error) {
	c, err := m.queryCount(pk.SQLFormat(), pk.SQLParams()...)
	if err != nil {
		return false, err
	}
	return (c != 0), nil
}

// Deprecated: Use FetchByPrimaryKey instead.
func (m *_UserBlogsDBMgr) Fetch(pk PrimaryKey) (*UserBlogs, error) {
	obj := UserBlogsMgr.NewUserBlogs()
	query := fmt.Sprintf("SELECT %s FROM user_blogs %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("UserBlogs fetch record not found")
}

// err not found check
func (m *_UserBlogsDBMgr) IsErrNotFound(err error) bool {
	return strings.Contains(err.Error(), "not found") || err == sql.ErrNoRows
}

// primary key
func (m *_UserBlogsDBMgr) FetchByPrimaryKey(userId int32, blogId int32) (*UserBlogs, error) {
	obj := UserBlogsMgr.NewUserBlogs()
	pk := &UserIdBlogIdOfUserBlogsPK{
		UserId: userId,
		BlogId: blogId,
	}

	query := fmt.Sprintf("SELECT %s FROM user_blogs %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("UserBlogs fetch record not found")
}

func (m *_UserBlogsDBMgr) FetchByPrimaryKeyContext(ctx context.Context, userId int32, blogId int32) (*UserBlogs, error) {
	obj := UserBlogsMgr.NewUserBlogs()
	pk := &UserIdBlogIdOfUserBlogsPK{
		UserId: userId,
		BlogId: blogId,
	}

	query := fmt.Sprintf("SELECT %s FROM user_blogs %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQLContext(ctx, query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("UserBlogs fetch record not found")
}

// indexes

// uniques

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

func (m *_UserBlogsDBMgr) FindOneContext(ctx context.Context, unique Unique) (PrimaryKey, error) {
	objs, err := m.queryLimitContext(ctx, unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("UserBlogs find record not found")
}

// Deprecated: Use FetchByXXXUnique instead.
func (m *_UserBlogsDBMgr) FindOneFetch(unique Unique) (*UserBlogs, error) {
	obj := UserBlogsMgr.NewUserBlogs()
	query := fmt.Sprintf("SELECT %s FROM user_blogs %s", strings.Join(obj.GetColumns(), ","), unique.SQLFormat(true))
	objs, err := m.FetchBySQL(query, unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("none record")
}

// Deprecated: Use FindByXXXUnique instead.
func (m *_UserBlogsDBMgr) Find(index Index) (int64, []PrimaryKey, error) {
	total, err := m.queryCount(index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimit(index.SQLFormat(true), index.SQLLimit(), index.SQLParams()...)
	return total, pks, err
}

func (m *_UserBlogsDBMgr) FindFetch(index Index) (int64, []*UserBlogs, error) {
	total, err := m.queryCount(index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}

	obj := UserBlogsMgr.NewUserBlogs()
	query := fmt.Sprintf("SELECT %s FROM user_blogs %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	results, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_UserBlogsDBMgr) FindFetchContext(ctx context.Context, index Index) (int64, []*UserBlogs, error) {
	total, err := m.queryCountContext(ctx, index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}

	obj := UserBlogsMgr.NewUserBlogs()
	query := fmt.Sprintf("SELECT %s FROM user_blogs %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	results, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_UserBlogsDBMgr) Range(scope Range) (int64, []PrimaryKey, error) {
	total, err := m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
	return total, pks, err
}

func (m *_UserBlogsDBMgr) RangeContext(ctx context.Context, scope Range) (int64, []PrimaryKey, error) {
	total, err := m.queryCountContext(ctx, scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimitContext(ctx, scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
	return total, pks, err
}

func (m *_UserBlogsDBMgr) RangeFetch(scope Range) (int64, []*UserBlogs, error) {
	total, err := m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	obj := UserBlogsMgr.NewUserBlogs()
	query := fmt.Sprintf("SELECT %s FROM user_blogs %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	results, err := m.FetchBySQL(query, scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_UserBlogsDBMgr) RangeFetchContext(ctx context.Context, scope Range) (int64, []*UserBlogs, error) {
	total, err := m.queryCountContext(ctx, scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	obj := UserBlogsMgr.NewUserBlogs()
	query := fmt.Sprintf("SELECT %s FROM user_blogs %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	results, err := m.FetchBySQLContext(ctx, query, scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_UserBlogsDBMgr) RangeRevert(scope Range) (int64, []PrimaryKey, error) {
	scope.Revert(true)
	return m.Range(scope)
}

func (m *_UserBlogsDBMgr) RangeRevertContext(ctx context.Context, scope Range) (int64, []PrimaryKey, error) {
	scope.Revert(true)
	return m.RangeContext(ctx, scope)
}

func (m *_UserBlogsDBMgr) RangeRevertFetch(scope Range) (int64, []*UserBlogs, error) {
	scope.Revert(true)
	return m.RangeFetch(scope)
}

func (m *_UserBlogsDBMgr) RangeRevertFetchContext(ctx context.Context, scope Range) (int64, []*UserBlogs, error) {
	scope.Revert(true)
	return m.RangeFetchContext(ctx, scope)
}

func (m *_UserBlogsDBMgr) queryLimit(where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := UserBlogsMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM user_blogs %s", strings.Join(pk.Columns(), ","), where)
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
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("UserBlogs query limit result error: %v", err)
	}
	return
}

func (m *_UserBlogsDBMgr) queryLimitContext(ctx context.Context, where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := UserBlogsMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM user_blogs %s", strings.Join(pk.Columns(), ","), where)
	rows, err := m.db.QueryContext(ctx, query, args...)
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
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("UserBlogs query limit result error: %v", err)
	}
	return
}

func (m *_UserBlogsDBMgr) queryCount(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`user_id`) FROM user_blogs %s", where)
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return 0, fmt.Errorf("UserBlogs query count error: %v", err)
	}
	defer rows.Close()

	var count int64
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			m.db.SetError(err)
			return 0, err
		}
		break
	}
	return count, nil
}

func (m *_UserBlogsDBMgr) queryCountContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`user_id`) FROM user_blogs %s", where)
	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("UserBlogs query count error: %v", err)
	}
	defer rows.Close()

	var count int64
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			m.db.SetError(err)
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
	query := fmt.Sprintf("INSERT INTO user_blogs(%s) VALUES %s", strings.Join(objs[0].GetNoneIncrementColumns(), ","), strings.Join(params, ","))
	result, err := m.db.Exec(query, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_UserBlogsDBMgr) BatchCreateContext(ctx context.Context, objs []*UserBlogs) (int64, error) {
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
	query := fmt.Sprintf("INSERT INTO user_blogs(%s) VALUES %s", strings.Join(objs[0].GetNoneIncrementColumns(), ","), strings.Join(params, ","))
	result, err := m.db.ExecContext(ctx, query, values...)
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
	query := fmt.Sprintf("UPDATE user_blogs SET %s", set)
	if where != "" {
		query = fmt.Sprintf("UPDATE user_blogs SET %s WHERE %s", set, where)
	}
	result, err := m.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// argument example:
// set:"a=?, b=?"
// where:"c=? and d=?"
// params:[]interface{}{"a", "b", "c", "d"}...
func (m *_UserBlogsDBMgr) UpdateBySQLContext(ctx context.Context, set, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("UPDATE user_blogs SET %s", set)
	if where != "" {
		query = fmt.Sprintf("UPDATE user_blogs SET %s WHERE %s", set, where)
	}
	result, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_UserBlogsDBMgr) Create(obj *UserBlogs) (int64, error) {
	params := orm.NewStringSlice(2, "?")
	q := fmt.Sprintf("INSERT INTO user_blogs(%s) VALUES(%s)",
		strings.Join(obj.GetNoneIncrementColumns(), ","),
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

func (m *_UserBlogsDBMgr) CreateContext(ctx context.Context, obj *UserBlogs) (int64, error) {
	params := orm.NewStringSlice(2, "?")
	q := fmt.Sprintf("INSERT INTO user_blogs(%s) VALUES(%s)",
		strings.Join(obj.GetNoneIncrementColumns(), ","),
		strings.Join(params, ","))

	values := make([]interface{}, 0, 2)
	values = append(values, obj.UserId)
	values = append(values, obj.BlogId)
	result, err := m.db.ExecContext(ctx, q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_UserBlogsDBMgr) Update(obj *UserBlogs) (int64, error) {
	columns := []string{}

	pk := obj.GetPrimaryKey()
	q := fmt.Sprintf("UPDATE user_blogs SET %s %s", strings.Join(columns, ","), pk.SQLFormat())
	values := make([]interface{}, 0, 2-2)
	values = append(values, pk.SQLParams()...)

	result, err := m.db.Exec(q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_UserBlogsDBMgr) UpdateContext(ctx context.Context, obj *UserBlogs) (int64, error) {
	columns := []string{}

	pk := obj.GetPrimaryKey()
	q := fmt.Sprintf("UPDATE user_blogs SET %s %s", strings.Join(columns, ","), pk.SQLFormat())
	values := make([]interface{}, 0, 2-2)
	values = append(values, pk.SQLParams()...)

	result, err := m.db.ExecContext(ctx, q, values...)
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

func (m *_UserBlogsDBMgr) SaveContext(ctx context.Context, obj *UserBlogs) (int64, error) {
	affected, err := m.UpdateContext(ctx, obj)
	if err != nil {
		return affected, err
	}
	if affected == 0 {
		return m.CreateContext(ctx, obj)
	}
	return affected, err
}

func (m *_UserBlogsDBMgr) Delete(obj *UserBlogs) (int64, error) {
	return m.DeleteByPrimaryKey(obj.UserId, obj.BlogId)
}

func (m *_UserBlogsDBMgr) DeleteContext(ctx context.Context, obj *UserBlogs) (int64, error) {
	return m.DeleteByPrimaryKeyContext(ctx, obj.UserId, obj.BlogId)
}

func (m *_UserBlogsDBMgr) DeleteByPrimaryKey(userId int32, blogId int32) (int64, error) {
	pk := &UserIdBlogIdOfUserBlogsPK{
		UserId: userId,
		BlogId: blogId,
	}
	q := fmt.Sprintf("DELETE FROM user_blogs %s", pk.SQLFormat())
	result, err := m.db.Exec(q, pk.SQLParams()...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_UserBlogsDBMgr) DeleteByPrimaryKeyContext(ctx context.Context, userId int32, blogId int32) (int64, error) {
	pk := &UserIdBlogIdOfUserBlogsPK{
		UserId: userId,
		BlogId: blogId,
	}
	q := fmt.Sprintf("DELETE FROM user_blogs %s", pk.SQLFormat())
	result, err := m.db.ExecContext(ctx, q, pk.SQLParams()...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_UserBlogsDBMgr) DeleteBySQL(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("DELETE FROM user_blogs")
	if where != "" {
		query = fmt.Sprintf("DELETE FROM user_blogs WHERE %s", where)
	}
	result, err := m.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_UserBlogsDBMgr) DeleteBySQLContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("DELETE FROM user_blogs")
	if where != "" {
		query = fmt.Sprintf("DELETE FROM user_blogs WHERE %s", where)
	}
	result, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
