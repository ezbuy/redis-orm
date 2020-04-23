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

type UserViewBaseInfo struct {
	Id       int32  `db:"id"`
	Name     string `db:"name"`
	Mailbox  string `db:"mailbox"`
	Password string `db:"password"`
	Sex      bool   `db:"sex"`
}

var UserViewBaseInfoColumns = struct {
	Id       string
	Name     string
	Mailbox  string
	Password string
	Sex      string
}{
	"id",
	"name",
	"mailbox",
	"password",
	"sex",
}

type _UserViewBaseInfoMgr struct {
}

var UserViewBaseInfoMgr *_UserViewBaseInfoMgr

func (m *_UserViewBaseInfoMgr) NewUserViewBaseInfo() *UserViewBaseInfo {
	return &UserViewBaseInfo{}
}

//! object function

func (obj *UserViewBaseInfo) GetNameSpace() string {
	return "model"
}

func (obj *UserViewBaseInfo) GetClassName() string {
	return "UserViewBaseInfo"
}

func (obj *UserViewBaseInfo) GetTableName() string {
	return ""
}

func (obj *UserViewBaseInfo) GetColumns() []string {
	columns := []string{
		".`id`",
		".`name`",
		".`mailbox`",
		".`password`",
		".`sex`",
	}
	return columns
}

func (obj *UserViewBaseInfo) GetNoneIncrementColumns() []string {
	columns := []string{
		"`id`",
		"`name`",
		"`mailbox`",
		"`password`",
		"`sex`",
	}
	return columns
}

func (obj *UserViewBaseInfo) GetPrimaryKey() PrimaryKey {
	pk := UserViewBaseInfoMgr.NewPrimaryKey()
	pk.Id = obj.Id
	return pk
}

func (obj *UserViewBaseInfo) Validate() error {
	validate := validator.New()
	return validate.Struct(obj)
}

//! primary key

type IdOfUserViewBaseInfoPK struct {
	Id int32
}

func (m *_UserViewBaseInfoMgr) NewPrimaryKey() *IdOfUserViewBaseInfoPK {
	return &IdOfUserViewBaseInfoPK{}
}

func (u *IdOfUserViewBaseInfoPK) Key() string {
	strs := []string{
		"Id",
		fmt.Sprint(u.Id),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *IdOfUserViewBaseInfoPK) Parse(key string) error {
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
	return nil
}

func (u *IdOfUserViewBaseInfoPK) SQLFormat() string {
	conditions := []string{
		"`id` = ?",
	}
	return orm.SQLWhere(conditions)
}

func (u *IdOfUserViewBaseInfoPK) SQLParams() []interface{} {
	return []interface{}{
		u.Id,
	}
}

func (u *IdOfUserViewBaseInfoPK) Columns() []string {
	return []string{
		"`id`",
	}
}

//! uniques

type MailboxPasswordOfUserViewBaseInfoUK struct {
	Mailbox  string
	Password string
}

func (u *MailboxPasswordOfUserViewBaseInfoUK) Key() string {
	strs := []string{
		"Mailbox",
		fmt.Sprint(u.Mailbox),
		"Password",
		fmt.Sprint(u.Password),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *MailboxPasswordOfUserViewBaseInfoUK) SQLFormat(limit bool) string {
	conditions := []string{
		"`mailbox` = ?",
		"`password` = ?",
	}
	return orm.SQLWhere(conditions)
}

func (u *MailboxPasswordOfUserViewBaseInfoUK) SQLParams() []interface{} {
	return []interface{}{
		u.Mailbox,
		u.Password,
	}
}

func (u *MailboxPasswordOfUserViewBaseInfoUK) SQLLimit() int {
	return 1
}

func (u *MailboxPasswordOfUserViewBaseInfoUK) Limit(n int) {
}

func (u *MailboxPasswordOfUserViewBaseInfoUK) Offset(n int) {
}

func (u *MailboxPasswordOfUserViewBaseInfoUK) UKRelation(store *orm.RedisStore) UniqueRelation {
	return nil
}

//! indexes

type NameOfUserViewBaseInfoIDX struct {
	Name   string
	offset int
	limit  int
}

func (u *NameOfUserViewBaseInfoIDX) Key() string {
	strs := []string{
		"Name",
		fmt.Sprint(u.Name),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *NameOfUserViewBaseInfoIDX) SQLFormat(limit bool) string {
	conditions := []string{
		"`name` = ?",
	}
	if limit {
		return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return orm.SQLWhere(conditions)
}

func (u *NameOfUserViewBaseInfoIDX) SQLParams() []interface{} {
	return []interface{}{
		u.Name,
	}
}

func (u *NameOfUserViewBaseInfoIDX) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *NameOfUserViewBaseInfoIDX) Limit(n int) {
	u.limit = n
}

func (u *NameOfUserViewBaseInfoIDX) Offset(n int) {
	u.offset = n
}

func (u *NameOfUserViewBaseInfoIDX) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *NameOfUserViewBaseInfoIDX) IDXRelation(store *orm.RedisStore) IndexRelation {
	return nil
}

//! ranges

type IdOfUserViewBaseInfoRNG struct {
	IdBegin      int64
	IdEnd        int64
	offset       int
	limit        int
	includeBegin bool
	includeEnd   bool
	revert       bool
}

func (u *IdOfUserViewBaseInfoRNG) Key() string {
	strs := []string{
		"Id",
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *IdOfUserViewBaseInfoRNG) beginOp() string {
	if u.includeBegin {
		return ">="
	}
	return ">"
}
func (u *IdOfUserViewBaseInfoRNG) endOp() string {
	if u.includeBegin {
		return "<="
	}
	return "<"
}

func (u *IdOfUserViewBaseInfoRNG) SQLFormat(limit bool) string {
	conditions := []string{}
	if u.IdBegin != u.IdEnd {
		if u.IdBegin != -1 {
			conditions = append(conditions, fmt.Sprintf("`id` %s ?", u.beginOp()))
		}
		if u.IdEnd != -1 {
			conditions = append(conditions, fmt.Sprintf("`id` %s ?", u.endOp()))
		}
	}
	if limit {
		return fmt.Sprintf("%s %s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("`id`", u.revert), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("`id`", u.revert))
}

func (u *IdOfUserViewBaseInfoRNG) SQLParams() []interface{} {
	params := []interface{}{}
	if u.IdBegin != u.IdEnd {
		if u.IdBegin != -1 {
			params = append(params, u.IdBegin)
		}
		if u.IdEnd != -1 {
			params = append(params, u.IdEnd)
		}
	}
	return params
}

func (u *IdOfUserViewBaseInfoRNG) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *IdOfUserViewBaseInfoRNG) Limit(n int) {
	u.limit = n
}

func (u *IdOfUserViewBaseInfoRNG) Offset(n int) {
	u.offset = n
}

func (u *IdOfUserViewBaseInfoRNG) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *IdOfUserViewBaseInfoRNG) Begin() int64 {
	start := u.IdBegin
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

func (u *IdOfUserViewBaseInfoRNG) End() int64 {
	stop := u.IdEnd
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

func (u *IdOfUserViewBaseInfoRNG) Revert(b bool) {
	u.revert = b
}

func (u *IdOfUserViewBaseInfoRNG) IncludeBegin(f bool) {
	u.includeBegin = f
}

func (u *IdOfUserViewBaseInfoRNG) IncludeEnd(f bool) {
	u.includeEnd = f
}

func (u *IdOfUserViewBaseInfoRNG) RNGRelation(store *orm.RedisStore) RangeRelation {
	return nil
}

type _UserViewBaseInfoDBMgr struct {
	db orm.DB
}

func (m *_UserViewBaseInfoMgr) DB(db orm.DB) *_UserViewBaseInfoDBMgr {
	return UserViewBaseInfoDBMgr(db)
}

func UserViewBaseInfoDBMgr(db orm.DB) *_UserViewBaseInfoDBMgr {
	if db == nil {
		panic(fmt.Errorf("UserViewBaseInfoDBMgr init need db"))
	}
	return &_UserViewBaseInfoDBMgr{db: db}
}

func (m *_UserViewBaseInfoDBMgr) Search(where string, orderby string, limit string, args ...interface{}) ([]*UserViewBaseInfo, error) {
	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()

	if limit = strings.ToUpper(strings.TrimSpace(limit)); limit != "" && !strings.HasPrefix(limit, "LIMIT") {
		limit = "LIMIT " + limit
	}

	conditions := []string{where, orderby, limit}
	query := fmt.Sprintf("SELECT %s FROM user_view_base_info %s", strings.Join(obj.GetColumns(), ","), strings.Join(conditions, " "))
	return m.FetchBySQL(query, args...)
}

func (m *_UserViewBaseInfoDBMgr) SearchContext(ctx context.Context, where string, orderby string, limit string, args ...interface{}) ([]*UserViewBaseInfo, error) {
	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()

	if limit = strings.ToUpper(strings.TrimSpace(limit)); limit != "" && !strings.HasPrefix(limit, "LIMIT") {
		limit = "LIMIT " + limit
	}

	conditions := []string{where, orderby, limit}
	query := fmt.Sprintf("SELECT %s FROM user_view_base_info %s", strings.Join(obj.GetColumns(), ","), strings.Join(conditions, " "))
	return m.FetchBySQLContext(ctx, query, args...)
}

func (m *_UserViewBaseInfoDBMgr) SearchConditions(conditions []string, orderby string, offset int, limit int, args ...interface{}) ([]*UserViewBaseInfo, error) {
	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()
	q := fmt.Sprintf("SELECT %s FROM user_view_base_info %s %s %s",
		strings.Join(obj.GetColumns(), ","),
		orm.SQLWhere(conditions),
		orderby,
		orm.SQLOffsetLimit(offset, limit))

	return m.FetchBySQL(q, args...)
}

func (m *_UserViewBaseInfoDBMgr) SearchConditionsContext(ctx context.Context, conditions []string, orderby string, offset int, limit int, args ...interface{}) ([]*UserViewBaseInfo, error) {
	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()
	q := fmt.Sprintf("SELECT %s FROM user_view_base_info %s %s %s",
		strings.Join(obj.GetColumns(), ","),
		orm.SQLWhere(conditions),
		orderby,
		orm.SQLOffsetLimit(offset, limit))

	return m.FetchBySQLContext(ctx, q, args...)
}

func (m *_UserViewBaseInfoDBMgr) SearchCount(where string, args ...interface{}) (int64, error) {
	return m.queryCount(where, args...)
}

func (m *_UserViewBaseInfoDBMgr) SearchCountContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	return m.queryCountContext(ctx, where, args...)
}

func (m *_UserViewBaseInfoDBMgr) SearchConditionsCount(conditions []string, args ...interface{}) (int64, error) {
	return m.queryCount(orm.SQLWhere(conditions), args...)
}

func (m *_UserViewBaseInfoDBMgr) SearchConditionsCountContext(ctx context.Context, conditions []string, args ...interface{}) (int64, error) {
	return m.queryCountContext(ctx, orm.SQLWhere(conditions), args...)
}

func (m *_UserViewBaseInfoDBMgr) FetchBySQL(q string, args ...interface{}) (results []*UserViewBaseInfo, err error) {
	rows, err := m.db.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("UserViewBaseInfo fetch error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var result UserViewBaseInfo
		err = rows.Scan(&(result.Id), &(result.Name), &(result.Mailbox), &(result.Password), &(result.Sex))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("UserViewBaseInfo fetch result error: %v", err)
	}
	return
}

func (m *_UserViewBaseInfoDBMgr) FetchBySQLContext(ctx context.Context, q string, args ...interface{}) (results []*UserViewBaseInfo, err error) {
	rows, err := m.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("UserViewBaseInfo fetch error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var result UserViewBaseInfo
		err = rows.Scan(&(result.Id), &(result.Name), &(result.Mailbox), &(result.Password), &(result.Sex))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("UserViewBaseInfo fetch result error: %v", err)
	}
	return
}
func (m *_UserViewBaseInfoDBMgr) Exist(pk PrimaryKey) (bool, error) {
	c, err := m.queryCount(pk.SQLFormat(), pk.SQLParams()...)
	if err != nil {
		return false, err
	}
	return (c != 0), nil
}

// Deprecated: Use FetchByPrimaryKey instead.
func (m *_UserViewBaseInfoDBMgr) Fetch(pk PrimaryKey) (*UserViewBaseInfo, error) {
	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()
	query := fmt.Sprintf("SELECT %s FROM user_view_base_info %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("UserViewBaseInfo fetch record not found")
}

// err not found check
func (m *_UserViewBaseInfoDBMgr) IsErrNotFound(err error) bool {
	return strings.Contains(err.Error(), "not found") || err == sql.ErrNoRows
}

// primary key
func (m *_UserViewBaseInfoDBMgr) FetchByPrimaryKey(id int32) (*UserViewBaseInfo, error) {
	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()
	pk := &IdOfUserViewBaseInfoPK{
		Id: id,
	}

	query := fmt.Sprintf("SELECT %s FROM user_view_base_info %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("UserViewBaseInfo fetch record not found")
}

func (m *_UserViewBaseInfoDBMgr) FetchByPrimaryKeyContext(ctx context.Context, id int32) (*UserViewBaseInfo, error) {
	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()
	pk := &IdOfUserViewBaseInfoPK{
		Id: id,
	}

	query := fmt.Sprintf("SELECT %s FROM user_view_base_info %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQLContext(ctx, query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("UserViewBaseInfo fetch record not found")
}

func (m *_UserViewBaseInfoDBMgr) FetchByPrimaryKeys(ids []int32) ([]*UserViewBaseInfo, error) {
	size := len(ids)
	if size == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, size)
	for _, pk := range ids {
		params = append(params, pk)
	}
	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()
	query := fmt.Sprintf("SELECT %s FROM user_view_base_info WHERE `id` IN (?%s)", strings.Join(obj.GetColumns(), ","),
		strings.Repeat(",?", size-1))
	return m.FetchBySQL(query, params...)
}

func (m *_UserViewBaseInfoDBMgr) FetchByPrimaryKeysContext(ctx context.Context, ids []int32) ([]*UserViewBaseInfo, error) {
	size := len(ids)
	if size == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, size)
	for _, pk := range ids {
		params = append(params, pk)
	}
	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()
	query := fmt.Sprintf("SELECT %s FROM user_view_base_info WHERE `id` IN (?%s)", strings.Join(obj.GetColumns(), ","),
		strings.Repeat(",?", size-1))
	return m.FetchBySQLContext(ctx, query, params...)
}

// indexes

func (m *_UserViewBaseInfoDBMgr) FindByName(name string, limit int, offset int) ([]*UserViewBaseInfo, error) {
	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()
	idx_ := &NameOfUserViewBaseInfoIDX{
		Name:   name,
		limit:  limit,
		offset: offset,
	}

	query := fmt.Sprintf("SELECT %s FROM user_view_base_info %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQL(query, idx_.SQLParams()...)
}

func (m *_UserViewBaseInfoDBMgr) FindByNameContext(ctx context.Context, name string, limit int, offset int) ([]*UserViewBaseInfo, error) {
	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()
	idx_ := &NameOfUserViewBaseInfoIDX{
		Name:   name,
		limit:  limit,
		offset: offset,
	}
	query := fmt.Sprintf("SELECT %s FROM user_view_base_info %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQLContext(ctx, query, idx_.SQLParams()...)
}

func (m *_UserViewBaseInfoDBMgr) FindAllByName(name string) ([]*UserViewBaseInfo, error) {
	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()
	idx_ := &NameOfUserViewBaseInfoIDX{
		Name: name,
	}

	query := fmt.Sprintf("SELECT %s FROM user_view_base_info %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQL(query, idx_.SQLParams()...)
}

func (m *_UserViewBaseInfoDBMgr) FindAllByNameContext(ctx context.Context, name string) ([]*UserViewBaseInfo, error) {
	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()
	idx_ := &NameOfUserViewBaseInfoIDX{
		Name: name,
	}

	query := fmt.Sprintf("SELECT %s FROM user_view_base_info %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQLContext(ctx, query, idx_.SQLParams()...)
}

func (m *_UserViewBaseInfoDBMgr) FindByNameGroup(items []string) ([]*UserViewBaseInfo, error) {
	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()
	if len(items) == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, len(items))
	for _, item := range items {
		params = append(params, item)
	}
	query := fmt.Sprintf("SELECT %s FROM user_view_base_info where `name` in (?", strings.Join(obj.GetColumns(), ",")) +
		strings.Repeat(",?", len(items)-1) + ")"
	return m.FetchBySQL(query, params...)
}

func (m *_UserViewBaseInfoDBMgr) FindByNameGroupContext(ctx context.Context, items []string) ([]*UserViewBaseInfo, error) {
	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()
	if len(items) == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, len(items))
	for _, item := range items {
		params = append(params, item)
	}
	query := fmt.Sprintf("SELECT %s FROM user_view_base_info where `name` in (?", strings.Join(obj.GetColumns(), ",")) +
		strings.Repeat(",?", len(items)-1) + ")"
	return m.FetchBySQLContext(ctx, query, params...)
}

// uniques

func (m *_UserViewBaseInfoDBMgr) FetchByMailboxPassword(mailbox string, password string) (*UserViewBaseInfo, error) {
	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()
	uniq := &MailboxPasswordOfUserViewBaseInfoUK{
		Mailbox:  mailbox,
		Password: password,
	}

	query := fmt.Sprintf("SELECT %s FROM user_view_base_info %s", strings.Join(obj.GetColumns(), ","), uniq.SQLFormat(true))
	objs, err := m.FetchBySQL(query, uniq.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("UserViewBaseInfo fetch record not found")
}

func (m *_UserViewBaseInfoDBMgr) FetchByMailboxPasswordContext(ctx context.Context, mailbox string, password string) (*UserViewBaseInfo, error) {
	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()
	uniq := &MailboxPasswordOfUserViewBaseInfoUK{
		Mailbox:  mailbox,
		Password: password,
	}

	query := fmt.Sprintf("SELECT %s FROM user_view_base_info %s", strings.Join(obj.GetColumns(), ","), uniq.SQLFormat(true))
	objs, err := m.FetchBySQLContext(ctx, query, uniq.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("UserViewBaseInfo fetch record not found")
}

func (m *_UserViewBaseInfoDBMgr) FindOne(unique Unique) (PrimaryKey, error) {
	objs, err := m.queryLimit(unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("UserViewBaseInfo find record not found")
}

func (m *_UserViewBaseInfoDBMgr) FindOneContext(ctx context.Context, unique Unique) (PrimaryKey, error) {
	objs, err := m.queryLimitContext(ctx, unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("UserViewBaseInfo find record not found")
}

// Deprecated: Use FetchByXXXUnique instead.
func (m *_UserViewBaseInfoDBMgr) FindOneFetch(unique Unique) (*UserViewBaseInfo, error) {
	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()
	query := fmt.Sprintf("SELECT %s FROM user_view_base_info %s", strings.Join(obj.GetColumns(), ","), unique.SQLFormat(true))
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
func (m *_UserViewBaseInfoDBMgr) Find(index Index) (int64, []PrimaryKey, error) {
	total, err := m.queryCount(index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimit(index.SQLFormat(true), index.SQLLimit(), index.SQLParams()...)
	return total, pks, err
}

func (m *_UserViewBaseInfoDBMgr) FindFetch(index Index) (int64, []*UserViewBaseInfo, error) {
	total, err := m.queryCount(index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}

	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()
	query := fmt.Sprintf("SELECT %s FROM user_view_base_info %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	results, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_UserViewBaseInfoDBMgr) FindFetchContext(ctx context.Context, index Index) (int64, []*UserViewBaseInfo, error) {
	total, err := m.queryCountContext(ctx, index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}

	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()
	query := fmt.Sprintf("SELECT %s FROM user_view_base_info %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	results, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_UserViewBaseInfoDBMgr) Range(scope Range) (int64, []PrimaryKey, error) {
	total, err := m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
	return total, pks, err
}

func (m *_UserViewBaseInfoDBMgr) RangeContext(ctx context.Context, scope Range) (int64, []PrimaryKey, error) {
	total, err := m.queryCountContext(ctx, scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimitContext(ctx, scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
	return total, pks, err
}

func (m *_UserViewBaseInfoDBMgr) RangeFetch(scope Range) (int64, []*UserViewBaseInfo, error) {
	total, err := m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()
	query := fmt.Sprintf("SELECT %s FROM user_view_base_info %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	results, err := m.FetchBySQL(query, scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_UserViewBaseInfoDBMgr) RangeFetchContext(ctx context.Context, scope Range) (int64, []*UserViewBaseInfo, error) {
	total, err := m.queryCountContext(ctx, scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	obj := UserViewBaseInfoMgr.NewUserViewBaseInfo()
	query := fmt.Sprintf("SELECT %s FROM user_view_base_info %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	results, err := m.FetchBySQLContext(ctx, query, scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_UserViewBaseInfoDBMgr) RangeRevert(scope Range) (int64, []PrimaryKey, error) {
	scope.Revert(true)
	return m.Range(scope)
}

func (m *_UserViewBaseInfoDBMgr) RangeRevertContext(ctx context.Context, scope Range) (int64, []PrimaryKey, error) {
	scope.Revert(true)
	return m.RangeContext(ctx, scope)
}

func (m *_UserViewBaseInfoDBMgr) RangeRevertFetch(scope Range) (int64, []*UserViewBaseInfo, error) {
	scope.Revert(true)
	return m.RangeFetch(scope)
}

func (m *_UserViewBaseInfoDBMgr) RangeRevertFetchContext(ctx context.Context, scope Range) (int64, []*UserViewBaseInfo, error) {
	scope.Revert(true)
	return m.RangeFetchContext(ctx, scope)
}

func (m *_UserViewBaseInfoDBMgr) queryLimit(where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := UserViewBaseInfoMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM user_view_base_info %s", strings.Join(pk.Columns(), ","), where)
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("UserViewBaseInfo query limit error: %v", err)
	}
	defer rows.Close()

	offset := 0

	for rows.Next() {
		if limit >= 0 && offset >= limit {
			break
		}
		offset++

		result := UserViewBaseInfoMgr.NewPrimaryKey()
		err = rows.Scan(&(result.Id))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("UserViewBaseInfo query limit result error: %v", err)
	}
	return
}

func (m *_UserViewBaseInfoDBMgr) queryLimitContext(ctx context.Context, where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := UserViewBaseInfoMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM user_view_base_info %s", strings.Join(pk.Columns(), ","), where)
	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("UserViewBaseInfo query limit error: %v", err)
	}
	defer rows.Close()

	offset := 0

	for rows.Next() {
		if limit >= 0 && offset >= limit {
			break
		}
		offset++

		result := UserViewBaseInfoMgr.NewPrimaryKey()
		err = rows.Scan(&(result.Id))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("UserViewBaseInfo query limit result error: %v", err)
	}
	return
}

func (m *_UserViewBaseInfoDBMgr) queryCount(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`id`) FROM user_view_base_info %s", where)
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return 0, fmt.Errorf("UserViewBaseInfo query count error: %v", err)
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

func (m *_UserViewBaseInfoDBMgr) queryCountContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`id`) FROM user_view_base_info %s", where)
	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("UserViewBaseInfo query count error: %v", err)
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
