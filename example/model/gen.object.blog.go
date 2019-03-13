package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/ezbuy/redis-orm/orm"
	"gopkg.in/go-playground/validator.v9"
	elastic "gopkg.in/olivere/elastic.v2"
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

type Blog struct {
	Id        int32     `db:"id" json:"id"`
	UserId    int32     `db:"user_id" json:"user_id"`
	Title     string    `db:"title" json:"title"`
	Content   string    `db:"content" json:"content"`
	Status    int32     `db:"status" json:"status"`
	Readed    int32     `db:"readed" json:"readed"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

var BlogColumns = struct {
	Id        string
	UserId    string
	Title     string
	Content   string
	Status    string
	Readed    string
	CreatedAt string
	UpdatedAt string
}{
	"id",
	"user_id",
	"title",
	"content",
	"status",
	"readed",
	"created_at",
	"updated_at",
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
		"blogs.`id`",
		"blogs.`user_id`",
		"blogs.`title`",
		"blogs.`content`",
		"blogs.`status`",
		"blogs.`readed`",
		"blogs.`created_at`",
		"blogs.`updated_at`",
	}
	return columns
}

func (obj *Blog) GetNoneIncrementColumns() []string {
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

func (obj *Blog) Validate() error {
	validate := validator.New()
	return validate.Struct(obj)
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
		"`id` = ?",
		"`user_id` = ?",
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
		"`status` = ?",
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

func (u *StatusOfBlogIDX) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *StatusOfBlogIDX) IDXRelation(store *orm.RedisStore) IndexRelation {
	return nil
}

//! ranges

type IdUserIdOfBlogRNG struct {
	Id           int32
	UserIdBegin  int64
	UserIdEnd    int64
	offset       int
	limit        int
	includeBegin bool
	includeEnd   bool
	revert       bool
}

func (u *IdUserIdOfBlogRNG) Key() string {
	strs := []string{
		"Id",
		fmt.Sprint(u.Id),
		"UserId",
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *IdUserIdOfBlogRNG) beginOp() string {
	if u.includeBegin {
		return ">="
	}
	return ">"
}
func (u *IdUserIdOfBlogRNG) endOp() string {
	if u.includeBegin {
		return "<="
	}
	return "<"
}

func (u *IdUserIdOfBlogRNG) SQLFormat(limit bool) string {
	conditions := []string{}
	conditions = append(conditions, "`id` = ?")
	if u.UserIdBegin != u.UserIdEnd {
		if u.UserIdBegin != -1 {
			conditions = append(conditions, fmt.Sprintf("`user_id` %s ?", u.beginOp()))
		}
		if u.UserIdEnd != -1 {
			conditions = append(conditions, fmt.Sprintf("`user_id` %s ?", u.endOp()))
		}
	}
	if limit {
		return fmt.Sprintf("%s %s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("`user_id`", u.revert), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("`user_id`", u.revert))
}

func (u *IdUserIdOfBlogRNG) SQLParams() []interface{} {
	params := []interface{}{
		u.Id,
	}
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

func (u *IdUserIdOfBlogRNG) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *IdUserIdOfBlogRNG) Limit(n int) {
	u.limit = n
}

func (u *IdUserIdOfBlogRNG) Offset(n int) {
	u.offset = n
}

func (u *IdUserIdOfBlogRNG) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *IdUserIdOfBlogRNG) Begin() int64 {
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

func (u *IdUserIdOfBlogRNG) End() int64 {
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

func (u *IdUserIdOfBlogRNG) Revert(b bool) {
	u.revert = b
}

func (u *IdUserIdOfBlogRNG) IncludeBegin(f bool) {
	u.includeBegin = f
}

func (u *IdUserIdOfBlogRNG) IncludeEnd(f bool) {
	u.includeEnd = f
}

func (u *IdUserIdOfBlogRNG) RNGRelation(store *orm.RedisStore) RangeRelation {
	return nil
}

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

	if limit = strings.ToUpper(strings.TrimSpace(limit)); limit != "" && !strings.HasPrefix(limit, "LIMIT") {
		limit = "LIMIT " + limit
	}

	conditions := []string{where, orderby, limit}
	query := fmt.Sprintf("SELECT %s FROM blogs %s", strings.Join(obj.GetColumns(), ","), strings.Join(conditions, " "))
	return m.FetchBySQL(query, args...)
}

func (m *_BlogDBMgr) SearchContext(ctx context.Context, where string, orderby string, limit string, args ...interface{}) ([]*Blog, error) {
	obj := BlogMgr.NewBlog()

	if limit = strings.ToUpper(strings.TrimSpace(limit)); limit != "" && !strings.HasPrefix(limit, "LIMIT") {
		limit = "LIMIT " + limit
	}

	conditions := []string{where, orderby, limit}
	query := fmt.Sprintf("SELECT %s FROM blogs %s", strings.Join(obj.GetColumns(), ","), strings.Join(conditions, " "))
	return m.FetchBySQLContext(ctx, query, args...)
}

func (m *_BlogDBMgr) SearchConditions(conditions []string, orderby string, offset int, limit int, args ...interface{}) ([]*Blog, error) {
	obj := BlogMgr.NewBlog()
	q := fmt.Sprintf("SELECT %s FROM blogs %s %s %s",
		strings.Join(obj.GetColumns(), ","),
		orm.SQLWhere(conditions),
		orderby,
		orm.SQLOffsetLimit(offset, limit))

	return m.FetchBySQL(q, args...)
}

func (m *_BlogDBMgr) SearchConditionsContext(ctx context.Context, conditions []string, orderby string, offset int, limit int, args ...interface{}) ([]*Blog, error) {
	obj := BlogMgr.NewBlog()
	q := fmt.Sprintf("SELECT %s FROM blogs %s %s %s",
		strings.Join(obj.GetColumns(), ","),
		orm.SQLWhere(conditions),
		orderby,
		orm.SQLOffsetLimit(offset, limit))

	return m.FetchBySQLContext(ctx, q, args...)
}

func (m *_BlogDBMgr) SearchCount(where string, args ...interface{}) (int64, error) {
	return m.queryCount(where, args...)
}

func (m *_BlogDBMgr) SearchCountContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	return m.queryCountContext(ctx, where, args...)
}

func (m *_BlogDBMgr) SearchConditionsCount(conditions []string, args ...interface{}) (int64, error) {
	return m.queryCount(orm.SQLWhere(conditions), args...)
}

func (m *_BlogDBMgr) SearchConditionsCountContext(ctx context.Context, conditions []string, args ...interface{}) (int64, error) {
	return m.queryCountContext(ctx, orm.SQLWhere(conditions), args...)
}

func (m *_BlogDBMgr) FetchBySQL(q string, args ...interface{}) (results []*Blog, err error) {
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
			m.db.SetError(err)
			return nil, err
		}

		result.CreatedAt = orm.TimeParse(CreatedAt)
		result.UpdatedAt = orm.TimeParse(UpdatedAt)

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("Blog fetch result error: %v", err)
	}
	return
}

func (m *_BlogDBMgr) FetchBySQLContext(ctx context.Context, q string, args ...interface{}) (results []*Blog, err error) {
	rows, err := m.db.QueryContext(ctx, q, args...)
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
			m.db.SetError(err)
			return nil, err
		}

		result.CreatedAt = orm.TimeParse(CreatedAt)
		result.UpdatedAt = orm.TimeParse(UpdatedAt)

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("Blog fetch result error: %v", err)
	}
	return
}
func (m *_BlogDBMgr) Exist(pk PrimaryKey) (bool, error) {
	c, err := m.queryCount(pk.SQLFormat(), pk.SQLParams()...)
	if err != nil {
		return false, err
	}
	return (c != 0), nil
}

// Deprecated: Use FetchByPrimaryKey instead.
func (m *_BlogDBMgr) Fetch(pk PrimaryKey) (*Blog, error) {
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM blogs %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("Blog fetch record not found")
}

// err not found check
func (m *_BlogDBMgr) IsErrNotFound(err error) bool {
	return strings.Contains(err.Error(), "not found") || err == sql.ErrNoRows
}

// primary key
func (m *_BlogDBMgr) FetchByPrimaryKey(id int32, userId int32) (*Blog, error) {
	obj := BlogMgr.NewBlog()
	pk := &IdUserIdOfBlogPK{
		Id:     id,
		UserId: userId,
	}

	query := fmt.Sprintf("SELECT %s FROM blogs %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("Blog fetch record not found")
}

func (m *_BlogDBMgr) FetchByPrimaryKeyContext(ctx context.Context, id int32, userId int32) (*Blog, error) {
	obj := BlogMgr.NewBlog()
	pk := &IdUserIdOfBlogPK{
		Id:     id,
		UserId: userId,
	}

	query := fmt.Sprintf("SELECT %s FROM blogs %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQLContext(ctx, query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("Blog fetch record not found")
}

// indexes

func (m *_BlogDBMgr) FindByStatus(status int32, limit int, offset int) ([]*Blog, error) {
	obj := BlogMgr.NewBlog()
	idx := &StatusOfBlogIDX{
		Status: status,
		limit:  limit,
		offset: offset,
	}

	query := fmt.Sprintf("SELECT %s FROM blogs %s", strings.Join(obj.GetColumns(), ","), idx.SQLFormat(true))
	return m.FetchBySQL(query, idx.SQLParams()...)
}

func (m *_BlogDBMgr) FindByStatusContext(ctx context.Context, status int32, limit int, offset int) ([]*Blog, error) {
	obj := BlogMgr.NewBlog()
	idx := &StatusOfBlogIDX{
		Status: status,
		limit:  limit,
		offset: offset,
	}
	query := fmt.Sprintf("SELECT %s FROM blogs %s", strings.Join(obj.GetColumns(), ","), idx.SQLFormat(true))
	return m.FetchBySQLContext(ctx, query, idx.SQLParams()...)
}

func (m *_BlogDBMgr) FindAllByStatus(status int32) ([]*Blog, error) {
	obj := BlogMgr.NewBlog()
	idx := &StatusOfBlogIDX{
		Status: status,
	}

	query := fmt.Sprintf("SELECT %s FROM blogs %s", strings.Join(obj.GetColumns(), ","), idx.SQLFormat(true))
	return m.FetchBySQL(query, idx.SQLParams()...)
}

func (m *_BlogDBMgr) FindAllByStatusContext(ctx context.Context, status int32) ([]*Blog, error) {
	obj := BlogMgr.NewBlog()
	idx := &StatusOfBlogIDX{
		Status: status,
	}

	query := fmt.Sprintf("SELECT %s FROM blogs %s", strings.Join(obj.GetColumns(), ","), idx.SQLFormat(true))
	return m.FetchBySQLContext(ctx, query, idx.SQLParams()...)
}

func (m *_BlogDBMgr) FindByStatusGroup(items []int32) ([]*Blog, error) {
	obj := BlogMgr.NewBlog()
	if len(items) == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, len(items))
	for _, item := range items {
		params = append(params, item)
	}
	query := fmt.Sprintf("SELECT %s FROM blogs where `status` in (?", strings.Join(obj.GetColumns(), ",")) +
		strings.Repeat(",?", len(items)-1) + ")"
	return m.FetchBySQL(query, params...)
}

func (m *_BlogDBMgr) FindByStatusGroupContext(ctx context.Context, items []int32) ([]*Blog, error) {
	obj := BlogMgr.NewBlog()
	if len(items) == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, len(items))
	for _, item := range items {
		params = append(params, item)
	}
	query := fmt.Sprintf("SELECT %s FROM blogs where `status` in (?", strings.Join(obj.GetColumns(), ",")) +
		strings.Repeat(",?", len(items)-1) + ")"
	return m.FetchBySQLContext(ctx, query, params...)
}

// uniques

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

func (m *_BlogDBMgr) FindOneContext(ctx context.Context, unique Unique) (PrimaryKey, error) {
	objs, err := m.queryLimitContext(ctx, unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("Blog find record not found")
}

// Deprecated: Use FetchByXXXUnique instead.
func (m *_BlogDBMgr) FindOneFetch(unique Unique) (*Blog, error) {
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM blogs %s", strings.Join(obj.GetColumns(), ","), unique.SQLFormat(true))
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
func (m *_BlogDBMgr) Find(index Index) (int64, []PrimaryKey, error) {
	total, err := m.queryCount(index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimit(index.SQLFormat(true), index.SQLLimit(), index.SQLParams()...)
	return total, pks, err
}

func (m *_BlogDBMgr) FindFetch(index Index) (int64, []*Blog, error) {
	total, err := m.queryCount(index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}

	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM blogs %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	results, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_BlogDBMgr) FindFetchContext(ctx context.Context, index Index) (int64, []*Blog, error) {
	total, err := m.queryCountContext(ctx, index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}

	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM blogs %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	results, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_BlogDBMgr) Range(scope Range) (int64, []PrimaryKey, error) {
	total, err := m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
	return total, pks, err
}

func (m *_BlogDBMgr) RangeContext(ctx context.Context, scope Range) (int64, []PrimaryKey, error) {
	total, err := m.queryCountContext(ctx, scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimitContext(ctx, scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
	return total, pks, err
}

func (m *_BlogDBMgr) RangeFetch(scope Range) (int64, []*Blog, error) {
	total, err := m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM blogs %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	results, err := m.FetchBySQL(query, scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_BlogDBMgr) RangeFetchContext(ctx context.Context, scope Range) (int64, []*Blog, error) {
	total, err := m.queryCountContext(ctx, scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	obj := BlogMgr.NewBlog()
	query := fmt.Sprintf("SELECT %s FROM blogs %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	results, err := m.FetchBySQLContext(ctx, query, scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_BlogDBMgr) RangeRevert(scope Range) (int64, []PrimaryKey, error) {
	scope.Revert(true)
	return m.Range(scope)
}

func (m *_BlogDBMgr) RangeRevertContext(ctx context.Context, scope Range) (int64, []PrimaryKey, error) {
	scope.Revert(true)
	return m.RangeContext(ctx, scope)
}

func (m *_BlogDBMgr) RangeRevertFetch(scope Range) (int64, []*Blog, error) {
	scope.Revert(true)
	return m.RangeFetch(scope)
}

func (m *_BlogDBMgr) RangeRevertFetchContext(ctx context.Context, scope Range) (int64, []*Blog, error) {
	scope.Revert(true)
	return m.RangeFetchContext(ctx, scope)
}

func (m *_BlogDBMgr) queryLimit(where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := BlogMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM blogs %s", strings.Join(pk.Columns(), ","), where)
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
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("Blog query limit result error: %v", err)
	}
	return
}

func (m *_BlogDBMgr) queryLimitContext(ctx context.Context, where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := BlogMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM blogs %s", strings.Join(pk.Columns(), ","), where)
	rows, err := m.db.QueryContext(ctx, query, args...)
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
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("Blog query limit result error: %v", err)
	}
	return
}

func (m *_BlogDBMgr) queryCount(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`id`) FROM blogs %s", where)
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return 0, fmt.Errorf("Blog query count error: %v", err)
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

func (m *_BlogDBMgr) queryCountContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`id`) FROM blogs %s", where)
	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("Blog query count error: %v", err)
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
	query := fmt.Sprintf("INSERT INTO blogs(%s) VALUES %s", strings.Join(objs[0].GetNoneIncrementColumns(), ","), strings.Join(params, ","))
	result, err := m.db.Exec(query, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_BlogDBMgr) BatchCreateContext(ctx context.Context, objs []*Blog) (int64, error) {
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
	query := fmt.Sprintf("INSERT INTO blogs(%s) VALUES %s", strings.Join(objs[0].GetNoneIncrementColumns(), ","), strings.Join(params, ","))
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
func (m *_BlogDBMgr) UpdateBySQL(set, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("UPDATE blogs SET %s", set)
	if where != "" {
		query = fmt.Sprintf("UPDATE blogs SET %s WHERE %s", set, where)
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
func (m *_BlogDBMgr) UpdateBySQLContext(ctx context.Context, set, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("UPDATE blogs SET %s", set)
	if where != "" {
		query = fmt.Sprintf("UPDATE blogs SET %s WHERE %s", set, where)
	}
	result, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_BlogDBMgr) Create(obj *Blog) (int64, error) {
	params := orm.NewStringSlice(8, "?")
	q := fmt.Sprintf("INSERT INTO blogs(%s) VALUES(%s)",
		strings.Join(obj.GetNoneIncrementColumns(), ","),
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

func (m *_BlogDBMgr) CreateContext(ctx context.Context, obj *Blog) (int64, error) {
	params := orm.NewStringSlice(8, "?")
	q := fmt.Sprintf("INSERT INTO blogs(%s) VALUES(%s)",
		strings.Join(obj.GetNoneIncrementColumns(), ","),
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
	result, err := m.db.ExecContext(ctx, q, values...)
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
	q := fmt.Sprintf("UPDATE blogs SET %s %s", strings.Join(columns, ","), pk.SQLFormat())
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

func (m *_BlogDBMgr) UpdateContext(ctx context.Context, obj *Blog) (int64, error) {
	columns := []string{
		"`title` = ?",
		"`content` = ?",
		"`status` = ?",
		"`readed` = ?",
		"`created_at` = ?",
		"`updated_at` = ?",
	}

	pk := obj.GetPrimaryKey()
	q := fmt.Sprintf("UPDATE blogs SET %s %s", strings.Join(columns, ","), pk.SQLFormat())
	values := make([]interface{}, 0, 8-2)
	values = append(values, obj.Title)
	values = append(values, obj.Content)
	values = append(values, obj.Status)
	values = append(values, obj.Readed)
	values = append(values, orm.TimeFormat(obj.CreatedAt))
	values = append(values, orm.TimeFormat(obj.UpdatedAt))
	values = append(values, pk.SQLParams()...)

	result, err := m.db.ExecContext(ctx, q, values...)
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

func (m *_BlogDBMgr) SaveContext(ctx context.Context, obj *Blog) (int64, error) {
	affected, err := m.UpdateContext(ctx, obj)
	if err != nil {
		return affected, err
	}
	if affected == 0 {
		return m.CreateContext(ctx, obj)
	}
	return affected, err
}

func (m *_BlogDBMgr) Delete(obj *Blog) (int64, error) {
	return m.DeleteByPrimaryKey(obj.Id, obj.UserId)
}

func (m *_BlogDBMgr) DeleteContext(ctx context.Context, obj *Blog) (int64, error) {
	return m.DeleteByPrimaryKeyContext(ctx, obj.Id, obj.UserId)
}

func (m *_BlogDBMgr) DeleteByPrimaryKey(id int32, userId int32) (int64, error) {
	pk := &IdUserIdOfBlogPK{
		Id:     id,
		UserId: userId,
	}
	q := fmt.Sprintf("DELETE FROM blogs %s", pk.SQLFormat())
	result, err := m.db.Exec(q, pk.SQLParams()...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_BlogDBMgr) DeleteByPrimaryKeyContext(ctx context.Context, id int32, userId int32) (int64, error) {
	pk := &IdUserIdOfBlogPK{
		Id:     id,
		UserId: userId,
	}
	q := fmt.Sprintf("DELETE FROM blogs %s", pk.SQLFormat())
	result, err := m.db.ExecContext(ctx, q, pk.SQLParams()...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_BlogDBMgr) DeleteBySQL(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("DELETE FROM blogs")
	if where != "" {
		query = fmt.Sprintf("DELETE FROM blogs WHERE %s", where)
	}
	result, err := m.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_BlogDBMgr) DeleteBySQLContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("DELETE FROM blogs")
	if where != "" {
		query = fmt.Sprintf("DELETE FROM blogs WHERE %s", where)
	}
	result, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

var (
	_ context.Context
)

//! orm.elastic
var BlogElasticFields = struct {
	Title     string
	Content   string
	CreatedAt string
}{
	"title",
	"content",
	"created_at",
}

var BlogElasticMgr = &_BlogElasticMgr{}

type _BlogElasticMgr struct {
	ensureMapping sync.Once
}

func (m *_BlogElasticMgr) Mapping() map[string]interface{} {
	return map[string]interface{}{
		"properties": map[string]interface{}{
			"title": map[string]interface{}{
				"type":  "string",
				"index": "not_analyzed",
			},
			"content": map[string]interface{}{
				"type":     "string",
				"index":    "analyzed",
				"analyzer": "standard",
			},
			"created_at": map[string]interface{}{
				"type": "date",
			},
		},
	}
}

func (m *_BlogElasticMgr) IndexService() (*elastic.IndexService, error) {
	var err error
	m.ensureMapping.Do(func() {
		_, err = m.PutMappingService().BodyJson(m.Mapping()).Do()
	})

	return ElasticClient().IndexService("ezorm").Type("blogs"), err
}

func (m *_BlogElasticMgr) PutMappingService() *elastic.PutMappingService {
	return ElasticClient().PutMappingService("ezorm ").Type("blogs")
}
