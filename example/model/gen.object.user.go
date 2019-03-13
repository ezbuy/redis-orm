package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ezbuy/redis-orm/orm"
	"gopkg.in/go-playground/validator.v9"
	redis "gopkg.in/redis.v5"
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

type User struct {
	Id          int32      `db:"id" json:"id"`
	Name        string     `db:"name" json:"name" validate:"required"`
	Mailbox     string     `db:"mailbox" json:"mailbox" validate:"required"`
	Sex         bool       `db:"sex" json:"sex"`
	Age         int32      `db:"age" json:"age"`
	Longitude   float64    `db:"longitude" json:"longitude"`
	Latitude    float64    `db:"latitude" json:"latitude"`
	Description string     `db:"description" json:"description"`
	Password    string     `db:"password" json:"password"`
	HeadUrl     string     `db:"head_url" json:"head_url"`
	Status      int32      `db:"status" json:"status"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at"`
}

var UserColumns = struct {
	Id          string
	Name        string
	Mailbox     string
	Sex         string
	Age         string
	Longitude   string
	Latitude    string
	Description string
	Password    string
	HeadUrl     string
	Status      string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}{
	"id",
	"name",
	"mailbox",
	"sex",
	"age",
	"longitude",
	"latitude",
	"description",
	"password",
	"head_url",
	"status",
	"created_at",
	"updated_at",
	"deleted_at",
}

type _UserMgr struct {
}

var UserMgr *_UserMgr

func (m *_UserMgr) NewUser() *User {
	return &User{}
}

//! object function

func (obj *User) GetNameSpace() string {
	return "model"
}

func (obj *User) GetClassName() string {
	return "User"
}

func (obj *User) GetTableName() string {
	return "users"
}

func (obj *User) GetColumns() []string {
	columns := []string{
		"users.`id`",
		"users.`name`",
		"users.`mailbox`",
		"users.`sex`",
		"users.`age`",
		"users.`longitude`",
		"users.`latitude`",
		"users.`description`",
		"users.`password`",
		"users.`head_url`",
		"users.`status`",
		"users.`created_at`",
		"users.`updated_at`",
		"users.`deleted_at`",
	}
	return columns
}

func (obj *User) GetNoneIncrementColumns() []string {
	columns := []string{
		"`name`",
		"`mailbox`",
		"`sex`",
		"`age`",
		"`longitude`",
		"`latitude`",
		"`description`",
		"`password`",
		"`head_url`",
		"`status`",
		"`created_at`",
		"`updated_at`",
		"`deleted_at`",
	}
	return columns
}

func (obj *User) GetPrimaryKey() PrimaryKey {
	pk := UserMgr.NewPrimaryKey()
	pk.Id = obj.Id
	return pk
}

func (obj *User) Validate() error {
	validate := validator.New()
	return validate.Struct(obj)
}
func (obj *User) GetIndexes() []string {
	idx := []string{
		"Sex",
		"Age",
	}
	return idx
}

func (obj *User) GetStoreType() string {
	return "hash"
}

func (obj *User) GetPrimaryName() string {
	pk := obj.GetPrimaryKey()
	return pk.Key()
}

//! primary key

type IdOfUserPK struct {
	Id int32
}

func (m *_UserMgr) NewPrimaryKey() *IdOfUserPK {
	return &IdOfUserPK{}
}

func (u *IdOfUserPK) Key() string {
	strs := []string{
		"Id",
		fmt.Sprint(u.Id),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *IdOfUserPK) Parse(key string) error {
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

func (u *IdOfUserPK) SQLFormat() string {
	conditions := []string{
		"`id` = ?",
	}
	return orm.SQLWhere(conditions)
}

func (u *IdOfUserPK) SQLParams() []interface{} {
	return []interface{}{
		u.Id,
	}
}

func (u *IdOfUserPK) Columns() []string {
	return []string{
		"`id`",
	}
}

//! uniques

type MailboxPasswordOfUserUK struct {
	Mailbox  string
	Password string
}

func (u *MailboxPasswordOfUserUK) Key() string {
	strs := []string{
		"Mailbox",
		fmt.Sprint(u.Mailbox),
		"Password",
		fmt.Sprint(u.Password),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *MailboxPasswordOfUserUK) SQLFormat(limit bool) string {
	conditions := []string{
		"`mailbox` = ?",
		"`password` = ?",
	}
	return orm.SQLWhere(conditions)
}

func (u *MailboxPasswordOfUserUK) SQLParams() []interface{} {
	return []interface{}{
		u.Mailbox,
		u.Password,
	}
}

func (u *MailboxPasswordOfUserUK) SQLLimit() int {
	return 1
}

func (u *MailboxPasswordOfUserUK) Limit(n int) {
}

func (u *MailboxPasswordOfUserUK) Offset(n int) {
}

func (u *MailboxPasswordOfUserUK) UKRelation(store *orm.RedisStore) UniqueRelation {
	return MailboxPasswordOfUserUKRelationRedisMgr(store)
}

//! indexes

type SexOfUserIDX struct {
	Sex    bool
	offset int
	limit  int
}

func (u *SexOfUserIDX) Key() string {
	strs := []string{
		"Sex",
		fmt.Sprint(u.Sex),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *SexOfUserIDX) SQLFormat(limit bool) string {
	conditions := []string{
		"`sex` = ?",
	}
	if limit {
		return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return orm.SQLWhere(conditions)
}

func (u *SexOfUserIDX) SQLParams() []interface{} {
	return []interface{}{
		u.Sex,
	}
}

func (u *SexOfUserIDX) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *SexOfUserIDX) Limit(n int) {
	u.limit = n
}

func (u *SexOfUserIDX) Offset(n int) {
	u.offset = n
}

func (u *SexOfUserIDX) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *SexOfUserIDX) IDXRelation(store *orm.RedisStore) IndexRelation {
	return SexOfUserIDXRelationRedisMgr(store)
}

//! ranges

type IdOfUserRNG struct {
	IdBegin      int64
	IdEnd        int64
	offset       int
	limit        int
	includeBegin bool
	includeEnd   bool
	revert       bool
}

func (u *IdOfUserRNG) Key() string {
	strs := []string{
		"Id",
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *IdOfUserRNG) beginOp() string {
	if u.includeBegin {
		return ">="
	}
	return ">"
}
func (u *IdOfUserRNG) endOp() string {
	if u.includeBegin {
		return "<="
	}
	return "<"
}

func (u *IdOfUserRNG) SQLFormat(limit bool) string {
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

func (u *IdOfUserRNG) SQLParams() []interface{} {
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

func (u *IdOfUserRNG) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *IdOfUserRNG) Limit(n int) {
	u.limit = n
}

func (u *IdOfUserRNG) Offset(n int) {
	u.offset = n
}

func (u *IdOfUserRNG) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *IdOfUserRNG) Begin() int64 {
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

func (u *IdOfUserRNG) End() int64 {
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

func (u *IdOfUserRNG) Revert(b bool) {
	u.revert = b
}

func (u *IdOfUserRNG) IncludeBegin(f bool) {
	u.includeBegin = f
}

func (u *IdOfUserRNG) IncludeEnd(f bool) {
	u.includeEnd = f
}

func (u *IdOfUserRNG) RNGRelation(store *orm.RedisStore) RangeRelation {
	return IdOfUserRNGRelationRedisMgr(store)
}

type AgeOfUserRNG struct {
	AgeBegin     int64
	AgeEnd       int64
	offset       int
	limit        int
	includeBegin bool
	includeEnd   bool
	revert       bool
}

func (u *AgeOfUserRNG) Key() string {
	strs := []string{
		"Age",
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *AgeOfUserRNG) beginOp() string {
	if u.includeBegin {
		return ">="
	}
	return ">"
}
func (u *AgeOfUserRNG) endOp() string {
	if u.includeBegin {
		return "<="
	}
	return "<"
}

func (u *AgeOfUserRNG) SQLFormat(limit bool) string {
	conditions := []string{}
	if u.AgeBegin != u.AgeEnd {
		if u.AgeBegin != -1 {
			conditions = append(conditions, fmt.Sprintf("`age` %s ?", u.beginOp()))
		}
		if u.AgeEnd != -1 {
			conditions = append(conditions, fmt.Sprintf("`age` %s ?", u.endOp()))
		}
	}
	if limit {
		return fmt.Sprintf("%s %s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("`age`", u.revert), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("`age`", u.revert))
}

func (u *AgeOfUserRNG) SQLParams() []interface{} {
	params := []interface{}{}
	if u.AgeBegin != u.AgeEnd {
		if u.AgeBegin != -1 {
			params = append(params, u.AgeBegin)
		}
		if u.AgeEnd != -1 {
			params = append(params, u.AgeEnd)
		}
	}
	return params
}

func (u *AgeOfUserRNG) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *AgeOfUserRNG) Limit(n int) {
	u.limit = n
}

func (u *AgeOfUserRNG) Offset(n int) {
	u.offset = n
}

func (u *AgeOfUserRNG) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *AgeOfUserRNG) Begin() int64 {
	start := u.AgeBegin
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

func (u *AgeOfUserRNG) End() int64 {
	stop := u.AgeEnd
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

func (u *AgeOfUserRNG) Revert(b bool) {
	u.revert = b
}

func (u *AgeOfUserRNG) IncludeBegin(f bool) {
	u.includeBegin = f
}

func (u *AgeOfUserRNG) IncludeEnd(f bool) {
	u.includeEnd = f
}

func (u *AgeOfUserRNG) RNGRelation(store *orm.RedisStore) RangeRelation {
	return AgeOfUserRNGRelationRedisMgr(store)
}

type _UserDBMgr struct {
	db orm.DB
}

func (m *_UserMgr) DB(db orm.DB) *_UserDBMgr {
	return UserDBMgr(db)
}

func UserDBMgr(db orm.DB) *_UserDBMgr {
	if db == nil {
		panic(fmt.Errorf("UserDBMgr init need db"))
	}
	return &_UserDBMgr{db: db}
}

func (m *_UserDBMgr) Search(where string, orderby string, limit string, args ...interface{}) ([]*User, error) {
	obj := UserMgr.NewUser()

	if limit = strings.ToUpper(strings.TrimSpace(limit)); limit != "" && !strings.HasPrefix(limit, "LIMIT") {
		limit = "LIMIT " + limit
	}

	conditions := []string{where, orderby, limit}
	query := fmt.Sprintf("SELECT %s FROM users %s", strings.Join(obj.GetColumns(), ","), strings.Join(conditions, " "))
	return m.FetchBySQL(query, args...)
}

func (m *_UserDBMgr) SearchContext(ctx context.Context, where string, orderby string, limit string, args ...interface{}) ([]*User, error) {
	obj := UserMgr.NewUser()

	if limit = strings.ToUpper(strings.TrimSpace(limit)); limit != "" && !strings.HasPrefix(limit, "LIMIT") {
		limit = "LIMIT " + limit
	}

	conditions := []string{where, orderby, limit}
	query := fmt.Sprintf("SELECT %s FROM users %s", strings.Join(obj.GetColumns(), ","), strings.Join(conditions, " "))
	return m.FetchBySQLContext(ctx, query, args...)
}

func (m *_UserDBMgr) SearchConditions(conditions []string, orderby string, offset int, limit int, args ...interface{}) ([]*User, error) {
	obj := UserMgr.NewUser()
	q := fmt.Sprintf("SELECT %s FROM users %s %s %s",
		strings.Join(obj.GetColumns(), ","),
		orm.SQLWhere(conditions),
		orderby,
		orm.SQLOffsetLimit(offset, limit))

	return m.FetchBySQL(q, args...)
}

func (m *_UserDBMgr) SearchConditionsContext(ctx context.Context, conditions []string, orderby string, offset int, limit int, args ...interface{}) ([]*User, error) {
	obj := UserMgr.NewUser()
	q := fmt.Sprintf("SELECT %s FROM users %s %s %s",
		strings.Join(obj.GetColumns(), ","),
		orm.SQLWhere(conditions),
		orderby,
		orm.SQLOffsetLimit(offset, limit))

	return m.FetchBySQLContext(ctx, q, args...)
}

func (m *_UserDBMgr) SearchCount(where string, args ...interface{}) (int64, error) {
	return m.queryCount(where, args...)
}

func (m *_UserDBMgr) SearchCountContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	return m.queryCountContext(ctx, where, args...)
}

func (m *_UserDBMgr) SearchConditionsCount(conditions []string, args ...interface{}) (int64, error) {
	return m.queryCount(orm.SQLWhere(conditions), args...)
}

func (m *_UserDBMgr) SearchConditionsCountContext(ctx context.Context, conditions []string, args ...interface{}) (int64, error) {
	return m.queryCountContext(ctx, orm.SQLWhere(conditions), args...)
}

func (m *_UserDBMgr) FetchBySQL(q string, args ...interface{}) (results []*User, err error) {
	rows, err := m.db.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("User fetch error: %v", err)
	}
	defer rows.Close()

	var Description sql.NullString
	var HeadUrl sql.NullString
	var CreatedAt int64
	var UpdatedAt int64
	var DeletedAt sql.NullInt64

	for rows.Next() {
		var result User
		err = rows.Scan(&(result.Id), &(result.Name), &(result.Mailbox), &(result.Sex), &(result.Age), &(result.Longitude), &(result.Latitude), &Description, &(result.Password), &HeadUrl, &(result.Status), &CreatedAt, &UpdatedAt, &DeletedAt)
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		result.Description = Description.String

		result.HeadUrl = HeadUrl.String
		result.HeadUrl = orm.Decode(result.HeadUrl)

		result.CreatedAt = time.Unix(CreatedAt, 0)
		result.UpdatedAt = time.Unix(UpdatedAt, 0)
		if DeletedAt.Valid {
			DeletedAtValue := DeletedAt.Int64
			DeletedAtPoint := time.Unix(DeletedAtValue, 0)
			result.DeletedAt = &DeletedAtPoint
		} else {
			result.DeletedAt = nil
		}

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("User fetch result error: %v", err)
	}
	return
}

func (m *_UserDBMgr) FetchBySQLContext(ctx context.Context, q string, args ...interface{}) (results []*User, err error) {
	rows, err := m.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("User fetch error: %v", err)
	}
	defer rows.Close()

	var Description sql.NullString
	var HeadUrl sql.NullString
	var CreatedAt int64
	var UpdatedAt int64
	var DeletedAt sql.NullInt64

	for rows.Next() {
		var result User
		err = rows.Scan(&(result.Id), &(result.Name), &(result.Mailbox), &(result.Sex), &(result.Age), &(result.Longitude), &(result.Latitude), &Description, &(result.Password), &HeadUrl, &(result.Status), &CreatedAt, &UpdatedAt, &DeletedAt)
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		result.Description = Description.String

		result.HeadUrl = HeadUrl.String
		result.HeadUrl = orm.Decode(result.HeadUrl)

		result.CreatedAt = time.Unix(CreatedAt, 0)
		result.UpdatedAt = time.Unix(UpdatedAt, 0)
		if DeletedAt.Valid {
			DeletedAtValue := DeletedAt.Int64
			DeletedAtPoint := time.Unix(DeletedAtValue, 0)
			result.DeletedAt = &DeletedAtPoint
		} else {
			result.DeletedAt = nil
		}

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("User fetch result error: %v", err)
	}
	return
}
func (m *_UserDBMgr) Exist(pk PrimaryKey) (bool, error) {
	c, err := m.queryCount(pk.SQLFormat(), pk.SQLParams()...)
	if err != nil {
		return false, err
	}
	return (c != 0), nil
}

// Deprecated: Use FetchByPrimaryKey instead.
func (m *_UserDBMgr) Fetch(pk PrimaryKey) (*User, error) {
	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM users %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("User fetch record not found")
}

// err not found check
func (m *_UserDBMgr) IsErrNotFound(err error) bool {
	return strings.Contains(err.Error(), "not found") || err == sql.ErrNoRows
}

// primary key
func (m *_UserDBMgr) FetchByPrimaryKey(id int32) (*User, error) {
	obj := UserMgr.NewUser()
	pk := &IdOfUserPK{
		Id: id,
	}

	query := fmt.Sprintf("SELECT %s FROM users %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("User fetch record not found")
}

func (m *_UserDBMgr) FetchByPrimaryKeyContext(ctx context.Context, id int32) (*User, error) {
	obj := UserMgr.NewUser()
	pk := &IdOfUserPK{
		Id: id,
	}

	query := fmt.Sprintf("SELECT %s FROM users %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQLContext(ctx, query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("User fetch record not found")
}

func (m *_UserDBMgr) FetchByPrimaryKeys(ids []int32) ([]*User, error) {
	size := len(ids)
	if size == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, size)
	for _, pk := range ids {
		params = append(params, pk)
	}
	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM users WHERE `id` IN (?%s)", strings.Join(obj.GetColumns(), ","),
		strings.Repeat(",?", size-1))
	return m.FetchBySQL(query, params...)
}

func (m *_UserDBMgr) FetchByPrimaryKeysContext(ctx context.Context, ids []int32) ([]*User, error) {
	size := len(ids)
	if size == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, size)
	for _, pk := range ids {
		params = append(params, pk)
	}
	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM users WHERE `id` IN (?%s)", strings.Join(obj.GetColumns(), ","),
		strings.Repeat(",?", size-1))
	return m.FetchBySQLContext(ctx, query, params...)
}

// indexes

func (m *_UserDBMgr) FindBySex(sex bool, limit int, offset int) ([]*User, error) {
	obj := UserMgr.NewUser()
	idx := &SexOfUserIDX{
		Sex:    sex,
		limit:  limit,
		offset: offset,
	}

	query := fmt.Sprintf("SELECT %s FROM users %s", strings.Join(obj.GetColumns(), ","), idx.SQLFormat(true))
	return m.FetchBySQL(query, idx.SQLParams()...)
}

func (m *_UserDBMgr) FindBySexContext(ctx context.Context, sex bool, limit int, offset int) ([]*User, error) {
	obj := UserMgr.NewUser()
	idx := &SexOfUserIDX{
		Sex:    sex,
		limit:  limit,
		offset: offset,
	}
	query := fmt.Sprintf("SELECT %s FROM users %s", strings.Join(obj.GetColumns(), ","), idx.SQLFormat(true))
	return m.FetchBySQLContext(ctx, query, idx.SQLParams()...)
}

func (m *_UserDBMgr) FindAllBySex(sex bool) ([]*User, error) {
	obj := UserMgr.NewUser()
	idx := &SexOfUserIDX{
		Sex: sex,
	}

	query := fmt.Sprintf("SELECT %s FROM users %s", strings.Join(obj.GetColumns(), ","), idx.SQLFormat(true))
	return m.FetchBySQL(query, idx.SQLParams()...)
}

func (m *_UserDBMgr) FindAllBySexContext(ctx context.Context, sex bool) ([]*User, error) {
	obj := UserMgr.NewUser()
	idx := &SexOfUserIDX{
		Sex: sex,
	}

	query := fmt.Sprintf("SELECT %s FROM users %s", strings.Join(obj.GetColumns(), ","), idx.SQLFormat(true))
	return m.FetchBySQLContext(ctx, query, idx.SQLParams()...)
}

func (m *_UserDBMgr) FindBySexGroup(items []bool) ([]*User, error) {
	obj := UserMgr.NewUser()
	if len(items) == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, len(items))
	for _, item := range items {
		params = append(params, item)
	}
	query := fmt.Sprintf("SELECT %s FROM users where `sex` in (?", strings.Join(obj.GetColumns(), ",")) +
		strings.Repeat(",?", len(items)-1) + ")"
	return m.FetchBySQL(query, params...)
}

func (m *_UserDBMgr) FindBySexGroupContext(ctx context.Context, items []bool) ([]*User, error) {
	obj := UserMgr.NewUser()
	if len(items) == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, len(items))
	for _, item := range items {
		params = append(params, item)
	}
	query := fmt.Sprintf("SELECT %s FROM users where `sex` in (?", strings.Join(obj.GetColumns(), ",")) +
		strings.Repeat(",?", len(items)-1) + ")"
	return m.FetchBySQLContext(ctx, query, params...)
}

// uniques

func (m *_UserDBMgr) FetchByMailboxPassword(mailbox string, password string) (*User, error) {
	obj := UserMgr.NewUser()
	uniq := &MailboxPasswordOfUserUK{
		Mailbox:  mailbox,
		Password: password,
	}

	query := fmt.Sprintf("SELECT %s FROM users %s", strings.Join(obj.GetColumns(), ","), uniq.SQLFormat(true))
	objs, err := m.FetchBySQL(query, uniq.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("User fetch record not found")
}

func (m *_UserDBMgr) FetchByMailboxPasswordContext(ctx context.Context, mailbox string, password string) (*User, error) {
	obj := UserMgr.NewUser()
	uniq := &MailboxPasswordOfUserUK{
		Mailbox:  mailbox,
		Password: password,
	}

	query := fmt.Sprintf("SELECT %s FROM users %s", strings.Join(obj.GetColumns(), ","), uniq.SQLFormat(true))
	objs, err := m.FetchBySQLContext(ctx, query, uniq.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("User fetch record not found")
}

func (m *_UserDBMgr) FindOne(unique Unique) (PrimaryKey, error) {
	objs, err := m.queryLimit(unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("User find record not found")
}

func (m *_UserDBMgr) FindOneContext(ctx context.Context, unique Unique) (PrimaryKey, error) {
	objs, err := m.queryLimitContext(ctx, unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("User find record not found")
}

// Deprecated: Use FetchByXXXUnique instead.
func (m *_UserDBMgr) FindOneFetch(unique Unique) (*User, error) {
	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM users %s", strings.Join(obj.GetColumns(), ","), unique.SQLFormat(true))
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
func (m *_UserDBMgr) Find(index Index) (int64, []PrimaryKey, error) {
	total, err := m.queryCount(index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimit(index.SQLFormat(true), index.SQLLimit(), index.SQLParams()...)
	return total, pks, err
}

func (m *_UserDBMgr) FindFetch(index Index) (int64, []*User, error) {
	total, err := m.queryCount(index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}

	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM users %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	results, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_UserDBMgr) FindFetchContext(ctx context.Context, index Index) (int64, []*User, error) {
	total, err := m.queryCountContext(ctx, index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}

	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM users %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	results, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_UserDBMgr) Range(scope Range) (int64, []PrimaryKey, error) {
	total, err := m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
	return total, pks, err
}

func (m *_UserDBMgr) RangeContext(ctx context.Context, scope Range) (int64, []PrimaryKey, error) {
	total, err := m.queryCountContext(ctx, scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimitContext(ctx, scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
	return total, pks, err
}

func (m *_UserDBMgr) RangeFetch(scope Range) (int64, []*User, error) {
	total, err := m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM users %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	results, err := m.FetchBySQL(query, scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_UserDBMgr) RangeFetchContext(ctx context.Context, scope Range) (int64, []*User, error) {
	total, err := m.queryCountContext(ctx, scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM users %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	results, err := m.FetchBySQLContext(ctx, query, scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_UserDBMgr) RangeRevert(scope Range) (int64, []PrimaryKey, error) {
	scope.Revert(true)
	return m.Range(scope)
}

func (m *_UserDBMgr) RangeRevertContext(ctx context.Context, scope Range) (int64, []PrimaryKey, error) {
	scope.Revert(true)
	return m.RangeContext(ctx, scope)
}

func (m *_UserDBMgr) RangeRevertFetch(scope Range) (int64, []*User, error) {
	scope.Revert(true)
	return m.RangeFetch(scope)
}

func (m *_UserDBMgr) RangeRevertFetchContext(ctx context.Context, scope Range) (int64, []*User, error) {
	scope.Revert(true)
	return m.RangeFetchContext(ctx, scope)
}

func (m *_UserDBMgr) queryLimit(where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := UserMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM users %s", strings.Join(pk.Columns(), ","), where)
	rows, err := m.db.Query(query, args...)
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

		result := UserMgr.NewPrimaryKey()
		err = rows.Scan(&(result.Id))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("User query limit result error: %v", err)
	}
	return
}

func (m *_UserDBMgr) queryLimitContext(ctx context.Context, where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := UserMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM users %s", strings.Join(pk.Columns(), ","), where)
	rows, err := m.db.QueryContext(ctx, query, args...)
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

		result := UserMgr.NewPrimaryKey()
		err = rows.Scan(&(result.Id))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("User query limit result error: %v", err)
	}
	return
}

func (m *_UserDBMgr) queryCount(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`id`) FROM users %s", where)
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return 0, fmt.Errorf("User query count error: %v", err)
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

func (m *_UserDBMgr) queryCountContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`id`) FROM users %s", where)
	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("User query count error: %v", err)
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

func (m *_UserDBMgr) BatchCreate(objs []*User) (int64, error) {
	if len(objs) == 0 {
		return 0, nil
	}

	params := make([]string, 0, len(objs))
	values := make([]interface{}, 0, len(objs)*13)
	for _, obj := range objs {
		params = append(params, fmt.Sprintf("(%s)", strings.Join(orm.NewStringSlice(13, "?"), ",")))
		values = append(values, obj.Name)
		values = append(values, obj.Mailbox)
		values = append(values, obj.Sex)
		values = append(values, obj.Age)
		values = append(values, obj.Longitude)
		values = append(values, obj.Latitude)
		values = append(values, obj.Description)
		values = append(values, obj.Password)
		values = append(values, orm.Encode(obj.HeadUrl))
		values = append(values, obj.Status)
		values = append(values, obj.CreatedAt.Unix())
		values = append(values, obj.UpdatedAt.Unix())
		if obj.DeletedAt == nil {
			values = append(values, nil)
		} else {
			values = append(values, obj.DeletedAt.Unix())
		}
	}
	query := fmt.Sprintf("INSERT INTO users(%s) VALUES %s", strings.Join(objs[0].GetNoneIncrementColumns(), ","), strings.Join(params, ","))
	result, err := m.db.Exec(query, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_UserDBMgr) BatchCreateContext(ctx context.Context, objs []*User) (int64, error) {
	if len(objs) == 0 {
		return 0, nil
	}

	params := make([]string, 0, len(objs))
	values := make([]interface{}, 0, len(objs)*13)
	for _, obj := range objs {
		params = append(params, fmt.Sprintf("(%s)", strings.Join(orm.NewStringSlice(13, "?"), ",")))
		values = append(values, obj.Name)
		values = append(values, obj.Mailbox)
		values = append(values, obj.Sex)
		values = append(values, obj.Age)
		values = append(values, obj.Longitude)
		values = append(values, obj.Latitude)
		values = append(values, obj.Description)
		values = append(values, obj.Password)
		values = append(values, orm.Encode(obj.HeadUrl))
		values = append(values, obj.Status)
		values = append(values, obj.CreatedAt.Unix())
		values = append(values, obj.UpdatedAt.Unix())
		if obj.DeletedAt == nil {
			values = append(values, nil)
		} else {
			values = append(values, obj.DeletedAt.Unix())
		}
	}
	query := fmt.Sprintf("INSERT INTO users(%s) VALUES %s", strings.Join(objs[0].GetNoneIncrementColumns(), ","), strings.Join(params, ","))
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
func (m *_UserDBMgr) UpdateBySQL(set, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("UPDATE users SET %s", set)
	if where != "" {
		query = fmt.Sprintf("UPDATE users SET %s WHERE %s", set, where)
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
func (m *_UserDBMgr) UpdateBySQLContext(ctx context.Context, set, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("UPDATE users SET %s", set)
	if where != "" {
		query = fmt.Sprintf("UPDATE users SET %s WHERE %s", set, where)
	}
	result, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_UserDBMgr) Create(obj *User) (int64, error) {
	params := orm.NewStringSlice(13, "?")
	q := fmt.Sprintf("INSERT INTO users(%s) VALUES(%s)",
		strings.Join(obj.GetNoneIncrementColumns(), ","),
		strings.Join(params, ","))

	values := make([]interface{}, 0, 14)
	values = append(values, obj.Name)
	values = append(values, obj.Mailbox)
	values = append(values, obj.Sex)
	values = append(values, obj.Age)
	values = append(values, obj.Longitude)
	values = append(values, obj.Latitude)
	values = append(values, obj.Description)
	values = append(values, obj.Password)
	values = append(values, orm.Encode(obj.HeadUrl))
	values = append(values, obj.Status)
	values = append(values, obj.CreatedAt.Unix())
	values = append(values, obj.UpdatedAt.Unix())
	if obj.DeletedAt == nil {
		values = append(values, nil)
	} else {
		values = append(values, obj.DeletedAt.Unix())
	}
	result, err := m.db.Exec(q, values...)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	obj.Id = int32(lastInsertId)
	return result.RowsAffected()
}

func (m *_UserDBMgr) CreateContext(ctx context.Context, obj *User) (int64, error) {
	params := orm.NewStringSlice(13, "?")
	q := fmt.Sprintf("INSERT INTO users(%s) VALUES(%s)",
		strings.Join(obj.GetNoneIncrementColumns(), ","),
		strings.Join(params, ","))

	values := make([]interface{}, 0, 14)
	values = append(values, obj.Name)
	values = append(values, obj.Mailbox)
	values = append(values, obj.Sex)
	values = append(values, obj.Age)
	values = append(values, obj.Longitude)
	values = append(values, obj.Latitude)
	values = append(values, obj.Description)
	values = append(values, obj.Password)
	values = append(values, orm.Encode(obj.HeadUrl))
	values = append(values, obj.Status)
	values = append(values, obj.CreatedAt.Unix())
	values = append(values, obj.UpdatedAt.Unix())
	if obj.DeletedAt == nil {
		values = append(values, nil)
	} else {
		values = append(values, obj.DeletedAt.Unix())
	}
	result, err := m.db.ExecContext(ctx, q, values...)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	obj.Id = int32(lastInsertId)
	return result.RowsAffected()
}

func (m *_UserDBMgr) Update(obj *User) (int64, error) {
	columns := []string{
		"`name` = ?",
		"`mailbox` = ?",
		"`sex` = ?",
		"`age` = ?",
		"`longitude` = ?",
		"`latitude` = ?",
		"`description` = ?",
		"`password` = ?",
		"`head_url` = ?",
		"`status` = ?",
		"`created_at` = ?",
		"`updated_at` = ?",
		"`deleted_at` = ?",
	}

	pk := obj.GetPrimaryKey()
	q := fmt.Sprintf("UPDATE users SET %s %s", strings.Join(columns, ","), pk.SQLFormat())
	values := make([]interface{}, 0, 14-1)
	values = append(values, obj.Name)
	values = append(values, obj.Mailbox)
	values = append(values, obj.Sex)
	values = append(values, obj.Age)
	values = append(values, obj.Longitude)
	values = append(values, obj.Latitude)
	values = append(values, obj.Description)
	values = append(values, obj.Password)
	values = append(values, orm.Encode(obj.HeadUrl))
	values = append(values, obj.Status)
	values = append(values, obj.CreatedAt.Unix())
	values = append(values, obj.UpdatedAt.Unix())
	if obj.DeletedAt == nil {
		values = append(values, nil)
	} else {
		values = append(values, obj.DeletedAt.Unix())
	}
	values = append(values, pk.SQLParams()...)

	result, err := m.db.Exec(q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_UserDBMgr) UpdateContext(ctx context.Context, obj *User) (int64, error) {
	columns := []string{
		"`name` = ?",
		"`mailbox` = ?",
		"`sex` = ?",
		"`age` = ?",
		"`longitude` = ?",
		"`latitude` = ?",
		"`description` = ?",
		"`password` = ?",
		"`head_url` = ?",
		"`status` = ?",
		"`created_at` = ?",
		"`updated_at` = ?",
		"`deleted_at` = ?",
	}

	pk := obj.GetPrimaryKey()
	q := fmt.Sprintf("UPDATE users SET %s %s", strings.Join(columns, ","), pk.SQLFormat())
	values := make([]interface{}, 0, 14-1)
	values = append(values, obj.Name)
	values = append(values, obj.Mailbox)
	values = append(values, obj.Sex)
	values = append(values, obj.Age)
	values = append(values, obj.Longitude)
	values = append(values, obj.Latitude)
	values = append(values, obj.Description)
	values = append(values, obj.Password)
	values = append(values, orm.Encode(obj.HeadUrl))
	values = append(values, obj.Status)
	values = append(values, obj.CreatedAt.Unix())
	values = append(values, obj.UpdatedAt.Unix())
	if obj.DeletedAt == nil {
		values = append(values, nil)
	} else {
		values = append(values, obj.DeletedAt.Unix())
	}
	values = append(values, pk.SQLParams()...)

	result, err := m.db.ExecContext(ctx, q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_UserDBMgr) Save(obj *User) (int64, error) {
	affected, err := m.Update(obj)
	if err != nil {
		return affected, err
	}
	if affected == 0 {
		return m.Create(obj)
	}
	return affected, err
}

func (m *_UserDBMgr) SaveContext(ctx context.Context, obj *User) (int64, error) {
	affected, err := m.UpdateContext(ctx, obj)
	if err != nil {
		return affected, err
	}
	if affected == 0 {
		return m.CreateContext(ctx, obj)
	}
	return affected, err
}

func (m *_UserDBMgr) Delete(obj *User) (int64, error) {
	return m.DeleteByPrimaryKey(obj.Id)
}

func (m *_UserDBMgr) DeleteContext(ctx context.Context, obj *User) (int64, error) {
	return m.DeleteByPrimaryKeyContext(ctx, obj.Id)
}

func (m *_UserDBMgr) DeleteByPrimaryKey(id int32) (int64, error) {
	pk := &IdOfUserPK{
		Id: id,
	}
	q := fmt.Sprintf("DELETE FROM users %s", pk.SQLFormat())
	result, err := m.db.Exec(q, pk.SQLParams()...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_UserDBMgr) DeleteByPrimaryKeyContext(ctx context.Context, id int32) (int64, error) {
	pk := &IdOfUserPK{
		Id: id,
	}
	q := fmt.Sprintf("DELETE FROM users %s", pk.SQLFormat())
	result, err := m.db.ExecContext(ctx, q, pk.SQLParams()...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_UserDBMgr) DeleteBySQL(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("DELETE FROM users")
	if where != "" {
		query = fmt.Sprintf("DELETE FROM users WHERE %s", where)
	}
	result, err := m.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_UserDBMgr) DeleteBySQLContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("DELETE FROM users")
	if where != "" {
		query = fmt.Sprintf("DELETE FROM users WHERE %s", where)
	}
	result, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

type _UserRedisMgr struct {
	*orm.RedisStore
}

func (m *_UserMgr) Redis(store *orm.RedisStore) *_UserRedisMgr {
	return UserRedisMgr(store)
}

func UserRedisMgr(store *orm.RedisStore) *_UserRedisMgr {
	if store == nil {
		panic(fmt.Errorf("UserRedisMgr init need redis store"))
	}
	return &_UserRedisMgr{RedisStore: store}
}

//! pipeline
type _UserRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_UserRedisMgr) BeginPipeline(pipes ...*redis.Pipeline) *_UserRedisPipeline {
	if len(pipes) > 0 {
		return &_UserRedisPipeline{pipes[0], nil}
	}
	return &_UserRedisPipeline{m.Pipeline(), nil}
}

func (m *_UserRedisMgr) Load(db *_UserDBMgr) error {
	if err := m.Clear(); err != nil {
		return err
	}

	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM `users`", strings.Join(obj.GetColumns(), ","))
	return m.AddBySQL(db, query)

}

func (m *_UserRedisMgr) AddBySQL(db *_UserDBMgr, sql string, args ...interface{}) error {
	objs, err := db.FetchBySQL(sql, args...)
	if err != nil {
		return err
	}

	return m.SaveBatch(objs)
}
func (m *_UserRedisMgr) DelBySQL(db *_UserDBMgr, sql string, args ...interface{}) error {
	objs, err := db.FetchBySQL(sql, args...)
	if err != nil {
		return err
	}

	for _, obj := range objs {
		if err := m.Delete(obj); err != nil {
			return err
		}
	}
	return nil
}

var newUserObj = UserMgr.NewUser()

// get redis key of User, PrimaryKeys: id int32
func RedisKeyOfPrimaryUser(Id int32) string {
	strs := []string{
		"Id",
		fmt.Sprint(Id),
	}
	return keyOfObject(newUserObj, fmt.Sprintf("%s", strings.Join(strs, ":")))
}

//! redis model read
func (m *_UserRedisMgr) FindOne(unique Unique) (PrimaryKey, error) {
	if relation := unique.UKRelation(m.RedisStore); relation != nil {
		str, err := relation.FindOne(unique.Key())
		if err != nil {
			return nil, err
		}

		pk := UserMgr.NewPrimaryKey()
		if err := pk.Parse(str); err != nil {
			return nil, err
		}
		return pk, nil
	}
	return nil, fmt.Errorf("unique none relation.")
}

func (m *_UserRedisMgr) FindOneFetch(unique Unique) (*User, error) {
	v, err := m.FindOne(unique)
	if err != nil {
		return nil, err
	}
	return m.Fetch(v)
}

func (m *_UserRedisMgr) Find(index Index) (int64, []PrimaryKey, error) {
	if relation := index.IDXRelation(m.RedisStore); relation != nil {
		strs, err := relation.Find(index.Key())
		if err != nil {
			return 0, nil, err
		}
		total := int64(len(strs))
		p1, p2 := index.PositionOffsetLimit(len(strs))
		strs = strs[p1:p2]

		results := make([]PrimaryKey, 0, len(strs))
		for _, str := range strs {
			pk := UserMgr.NewPrimaryKey()
			if err := pk.Parse(str); err != nil {
				total--
				continue
			}
			results = append(results, pk)
		}
		return total, results, nil
	}
	return 0, nil, fmt.Errorf("index none relation.")
}

func (m *_UserRedisMgr) FindFetch(index Index) (int64, []*User, error) {
	total, vs, err := m.Find(index)
	if err != nil {
		return 0, nil, err
	}
	objs, err := m.FetchByPrimaryKeys(vs)
	return total, objs, err
}

func (m *_UserRedisMgr) Range(scope Range) (int64, []PrimaryKey, error) {
	if relation := scope.RNGRelation(m.RedisStore); relation != nil {
		strs, err := relation.Range(scope.Key(), scope.Begin(), scope.End())
		if err != nil {
			return 0, nil, err
		}
		total := int64(len(strs))
		p1, p2 := scope.PositionOffsetLimit(len(strs))
		strs = strs[p1:p2]

		results := make([]PrimaryKey, 0, len(strs))
		for _, str := range strs {
			pk := UserMgr.NewPrimaryKey()
			if err := pk.Parse(str); err != nil {
				total--
				continue
			}
			results = append(results, pk)
		}
		return total, results, nil
	}
	return 0, nil, fmt.Errorf("range none relation.")
}

func (m *_UserRedisMgr) RangeFetch(scope Range) (int64, []*User, error) {
	total, vs, err := m.Range(scope)
	if err != nil {
		return 0, nil, err
	}
	objs, err := m.FetchByPrimaryKeys(vs)
	return total, objs, err
}

func (m *_UserRedisMgr) RangeRevert(scope Range) (int64, []PrimaryKey, error) {
	if relation := scope.RNGRelation(m.RedisStore); relation != nil {
		scope.Revert(true)
		strs, err := relation.RangeRevert(scope.Key(), scope.Begin(), scope.End())
		if err != nil {
			return 0, nil, err
		}

		total := int64(len(strs))
		p1, p2 := scope.PositionOffsetLimit(len(strs))
		strs = strs[p1:p2]

		results := make([]PrimaryKey, 0, len(strs))
		for _, str := range strs {
			pk := UserMgr.NewPrimaryKey()
			if err := pk.Parse(str); err != nil {
				total--
				continue
			}
			results = append(results, pk)
		}
		return total, results, nil
	}
	return 0, nil, fmt.Errorf("revert range none relation.")
}

func (m *_UserRedisMgr) RangeRevertFetch(scope Range) (int64, []*User, error) {
	total, vs, err := m.RangeRevert(scope)
	if err != nil {
		return 0, nil, err
	}
	objs, err := m.FetchByPrimaryKeys(vs)
	return total, objs, err
}

func (m *_UserRedisMgr) Fetch(pk PrimaryKey) (*User, error) {
	key := keyOfObject(newUserObj, pk.Key())
	return m.FetchByKey(key)
}

func (m *_UserRedisMgr) FetchByKey(key string) (*User, error) {
	obj := UserMgr.NewUser()

	pipe := m.BeginPipeline()
	pipe.Exists(key)
	pipe.HMGet(key,
		"Id",
		"Name",
		"Mailbox",
		"Sex",
		"Age",
		"Longitude",
		"Latitude",
		"Description",
		"Password",
		"HeadUrl",
		"Status",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt")
	cmds, err := pipe.Exec()
	if err != nil {
		return nil, err
	}

	if b, err := cmds[0].(*redis.BoolCmd).Result(); err == nil {
		if !b {
			return nil, fmt.Errorf("User primary key:(%s) not exist", key)
		}
	}

	strs, err := cmds[1].(*redis.SliceCmd).Result()
	if err != nil {
		return nil, err
	}

	var sv string
	if strs[0] != nil {
		sv, _ = strs[0].(string)
		if err := orm.StringScan(sv, &obj.Id); err != nil {
			return nil, err
		}
	}
	if strs[1] != nil {
		sv, _ = strs[1].(string)
		if err := orm.StringScan(sv, &obj.Name); err != nil {
			return nil, err
		}
	}
	if strs[2] != nil {
		sv, _ = strs[2].(string)
		if err := orm.StringScan(sv, &obj.Mailbox); err != nil {
			return nil, err
		}
	}
	if strs[3] != nil {
		sv, _ = strs[3].(string)
		if err := orm.StringScan(sv, &obj.Sex); err != nil {
			return nil, err
		}
	}
	if strs[4] != nil {
		sv, _ = strs[4].(string)
		if err := orm.StringScan(sv, &obj.Age); err != nil {
			return nil, err
		}
	}
	if strs[5] != nil {
		sv, _ = strs[5].(string)
		if err := orm.StringScan(sv, &obj.Longitude); err != nil {
			return nil, err
		}
	}
	if strs[6] != nil {
		sv, _ = strs[6].(string)
		if err := orm.StringScan(sv, &obj.Latitude); err != nil {
			return nil, err
		}
	}
	if strs[7] != nil {
		sv, _ = strs[7].(string)
		if err := orm.StringScan(sv, &obj.Description); err != nil {
			return nil, err
		}
	}
	if strs[8] != nil {
		sv, _ = strs[8].(string)
		if err := orm.StringScan(sv, &obj.Password); err != nil {
			return nil, err
		}
	}
	if strs[9] != nil {
		sv, _ = strs[9].(string)
		if err := orm.StringScan(sv, &obj.HeadUrl); err != nil {
			return nil, err
		}
		obj.HeadUrl = orm.Decode(obj.HeadUrl)
	}
	if strs[10] != nil {
		sv, _ = strs[10].(string)
		if err := orm.StringScan(sv, &obj.Status); err != nil {
			return nil, err
		}
	}
	if strs[11] != nil {
		sv, _ = strs[11].(string)
		var val11 int64
		if err := orm.StringScan(sv, &val11); err != nil {
			return nil, err
		}
		obj.CreatedAt = time.Unix(val11, 0)
	}
	if strs[12] != nil {
		sv, _ = strs[12].(string)
		var val12 int64
		if err := orm.StringScan(sv, &val12); err != nil {
			return nil, err
		}
		obj.UpdatedAt = time.Unix(val12, 0)
	}
	if strs[13] != nil {
		sv, _ = strs[13].(string)
		if sv == "nil" {
			obj.DeletedAt = nil
		} else {
			var val13 int64
			if err := orm.StringScan(sv, &val13); err != nil {
				return nil, err
			}
			DeletedAtValue := time.Unix(val13, 0)
			obj.DeletedAt = &DeletedAtValue
		}
	}
	return obj, nil
}

func (m *_UserRedisMgr) FetchByPrimaryKeys(pks []PrimaryKey) ([]*User, error) {
	objs := make([]*User, 0, len(pks))
	pipe := m.BeginPipeline()
	obj := UserMgr.NewUser()
	for _, pk := range pks {
		pipe.Exists(keyOfObject(obj, pk.Key()))
		pipe.HMGet(keyOfObject(obj, pk.Key()),
			"Id",
			"Name",
			"Mailbox",
			"Sex",
			"Age",
			"Longitude",
			"Latitude",
			"Description",
			"Password",
			"HeadUrl",
			"Status",
			"CreatedAt",
			"UpdatedAt",
			"DeletedAt")
	}
	cmds, err := pipe.Exec()
	if err != nil {
		return nil, err
	}
	errall := []string{}
	sv := ""
	ok := true
	for i := 0; i < len(pks); i++ {
		if b, err := cmds[2*i].(*redis.BoolCmd).Result(); err == nil {
			if !b {
				errall = append(errall, fmt.Sprintf("User primary key:(%s) not exist", pks[i].Key()))
				continue
			}
		}

		strs, err := cmds[2*i+1].(*redis.SliceCmd).Result()
		if err != nil {
			errall = append(errall, fmt.Sprintf("key:%v,err:%v", pks[i].Key(), err.Error()))
			continue
		}

		obj := UserMgr.NewUser()
		if strs[0] != nil {
			sv, ok = strs[0].(string)
			if !ok {
				errall = append(errall, fmt.Sprintf("convert %v to string error", strs[0]))
				continue
			}
			if err := orm.StringScan(sv, &obj.Id); err != nil {
				errall = append(errall, fmt.Sprintf("key:%v,err:%v", pks[i].Key(), err.Error()))
				continue
			}
		}
		if strs[1] != nil {
			sv, ok = strs[1].(string)
			if !ok {
				errall = append(errall, fmt.Sprintf("convert %v to string error", strs[1]))
				continue
			}
			if err := orm.StringScan(sv, &obj.Name); err != nil {
				errall = append(errall, fmt.Sprintf("key:%v,err:%v", pks[i].Key(), err.Error()))
				continue
			}
		}
		if strs[2] != nil {
			sv, ok = strs[2].(string)
			if !ok {
				errall = append(errall, fmt.Sprintf("convert %v to string error", strs[2]))
				continue
			}
			if err := orm.StringScan(sv, &obj.Mailbox); err != nil {
				errall = append(errall, fmt.Sprintf("key:%v,err:%v", pks[i].Key(), err.Error()))
				continue
			}
		}
		if strs[3] != nil {
			sv, ok = strs[3].(string)
			if !ok {
				errall = append(errall, fmt.Sprintf("convert %v to string error", strs[3]))
				continue
			}
			if err := orm.StringScan(sv, &obj.Sex); err != nil {
				errall = append(errall, fmt.Sprintf("key:%v,err:%v", pks[i].Key(), err.Error()))
				continue
			}
		}
		if strs[4] != nil {
			sv, ok = strs[4].(string)
			if !ok {
				errall = append(errall, fmt.Sprintf("convert %v to string error", strs[4]))
				continue
			}
			if err := orm.StringScan(sv, &obj.Age); err != nil {
				errall = append(errall, fmt.Sprintf("key:%v,err:%v", pks[i].Key(), err.Error()))
				continue
			}
		}
		if strs[5] != nil {
			sv, ok = strs[5].(string)
			if !ok {
				errall = append(errall, fmt.Sprintf("convert %v to string error", strs[5]))
				continue
			}
			if err := orm.StringScan(sv, &obj.Longitude); err != nil {
				errall = append(errall, fmt.Sprintf("key:%v,err:%v", pks[i].Key(), err.Error()))
				continue
			}
		}
		if strs[6] != nil {
			sv, ok = strs[6].(string)
			if !ok {
				errall = append(errall, fmt.Sprintf("convert %v to string error", strs[6]))
				continue
			}
			if err := orm.StringScan(sv, &obj.Latitude); err != nil {
				errall = append(errall, fmt.Sprintf("key:%v,err:%v", pks[i].Key(), err.Error()))
				continue
			}
		}
		if strs[7] != nil {
			sv, ok = strs[7].(string)
			if !ok {
				errall = append(errall, fmt.Sprintf("convert %v to string error", strs[7]))
				continue
			}
			if err := orm.StringScan(sv, &obj.Description); err != nil {
				errall = append(errall, fmt.Sprintf("key:%v,err:%v", pks[i].Key(), err.Error()))
				continue
			}
		}
		if strs[8] != nil {
			sv, ok = strs[8].(string)
			if !ok {
				errall = append(errall, fmt.Sprintf("convert %v to string error", strs[8]))
				continue
			}
			if err := orm.StringScan(sv, &obj.Password); err != nil {
				errall = append(errall, fmt.Sprintf("key:%v,err:%v", pks[i].Key(), err.Error()))
				continue
			}
		}
		if strs[9] != nil {
			sv, ok = strs[9].(string)
			if !ok {
				errall = append(errall, fmt.Sprintf("convert %v to string error", strs[9]))
				continue
			}
			if err := orm.StringScan(sv, &obj.HeadUrl); err != nil {
				errall = append(errall, fmt.Sprintf("key:%v,err:%v", pks[i].Key(), err.Error()))
				continue
			}
			obj.HeadUrl = orm.Decode(obj.HeadUrl)
		}
		if strs[10] != nil {
			sv, ok = strs[10].(string)
			if !ok {
				errall = append(errall, fmt.Sprintf("convert %v to string error", strs[10]))
				continue
			}
			if err := orm.StringScan(sv, &obj.Status); err != nil {
				errall = append(errall, fmt.Sprintf("key:%v,err:%v", pks[i].Key(), err.Error()))
				continue
			}
		}
		if strs[11] != nil {
			sv, ok = strs[11].(string)
			var val11 int64
			if !ok {
				errall = append(errall, fmt.Sprintf("convert %v to string error", strs[11]))
				continue
			}
			if err := orm.StringScan(sv, &val11); err != nil {
				errall = append(errall, fmt.Sprintf("key:%v,err:%v", pks[i].Key(), err.Error()))
				continue
			}
			obj.CreatedAt = time.Unix(val11, 0)
		}
		if strs[12] != nil {
			sv, ok = strs[12].(string)
			var val12 int64
			if !ok {
				errall = append(errall, fmt.Sprintf("convert %v to string error", strs[12]))
				continue
			}
			if err := orm.StringScan(sv, &val12); err != nil {
				errall = append(errall, fmt.Sprintf("key:%v,err:%v", pks[i].Key(), err.Error()))
				continue
			}
			obj.UpdatedAt = time.Unix(val12, 0)
		}
		if strs[13] != nil {
			sv, ok = strs[13].(string)
			if sv == "nil" {
				obj.DeletedAt = nil
			} else {
				var val13 int64
				if !ok {
					errall = append(errall, fmt.Sprintf("convert %v to string error", strs[13]))
					continue
				}
				if err := orm.StringScan(sv, &val13); err != nil {
					errall = append(errall, fmt.Sprintf("key:%v,err:%v", pks[i].Key(), err.Error()))
					continue
				}
				DeletedAtValue := time.Unix(val13, 0)
				obj.DeletedAt = &DeletedAtValue
			}
		}
		objs = append(objs, obj)
	}
	if len(errall) > 0 {
		return objs, errors.New(strings.Join(errall, ERROR_SPLIT))
	}
	return objs, nil
}

func (m *_UserRedisMgr) Create(obj *User) error {
	return m.Save(obj)
}

func (m *_UserRedisMgr) Update(obj *User) error {
	return m.Save(obj)
}

func (m *_UserRedisMgr) CreateWithExpire(obj *User, expire time.Duration) error {
	return m.SaveWithExpire(obj, expire)
}

func (m *_UserRedisMgr) UpdateWithExpire(obj *User, expire time.Duration) error {
	return m.SaveWithExpire(obj, expire)
}

func (m *_UserRedisMgr) Delete(obj *User) error {
	pk := obj.GetPrimaryKey()
	pipe := m.BeginPipeline()
	//! uniques
	uk_key_0 := []string{
		"Mailbox",
		fmt.Sprint(obj.Mailbox),
		"Password",
		fmt.Sprint(obj.Password),
	}
	uk_pip_0 := MailboxPasswordOfUserUKRelationRedisMgr().BeginPipeline(pipe.Pipeline)
	if err := uk_pip_0.PairRem(strings.Join(uk_key_0, ":")); err != nil {
		return err
	}

	//! indexes
	idx_key_0 := []string{
		"Sex",
		fmt.Sprint(obj.Sex),
	}
	idx_pip_0 := SexOfUserIDXRelationRedisMgr().BeginPipeline(pipe.Pipeline)
	idx_rel_0 := SexOfUserIDXRelationRedisMgr().NewSexOfUserIDXRelation(strings.Join(idx_key_0, ":"))
	idx_rel_0.Value = pk.Key()
	if err := idx_pip_0.SetRem(idx_rel_0); err != nil {
		return err
	}

	//! ranges
	rg_key_0 := []string{
		"Id",
	}
	rg_pip_0 := IdOfUserRNGRelationRedisMgr().BeginPipeline(pipe.Pipeline)
	rg_rel_0 := IdOfUserRNGRelationRedisMgr().NewIdOfUserRNGRelation(strings.Join(rg_key_0, ":"))
	score_rg_0, err := orm.ToFloat64(obj.Id)
	if err != nil {
		return err
	}
	rg_rel_0.Score = score_rg_0
	rg_rel_0.Value = pk.Key()
	if err := rg_pip_0.ZSetRem(rg_rel_0); err != nil {
		return err
	}
	rg_key_1 := []string{
		"Age",
	}
	rg_pip_1 := AgeOfUserRNGRelationRedisMgr().BeginPipeline(pipe.Pipeline)
	rg_rel_1 := AgeOfUserRNGRelationRedisMgr().NewAgeOfUserRNGRelation(strings.Join(rg_key_1, ":"))
	score_rg_1, err := orm.ToFloat64(obj.Age)
	if err != nil {
		return err
	}
	rg_rel_1.Score = score_rg_1
	rg_rel_1.Value = pk.Key()
	if err := rg_pip_1.ZSetRem(rg_rel_1); err != nil {
		return err
	}

	if err := pipe.Del(keyOfObject(obj, pk.Key())).Err(); err != nil {
		return err
	}

	if _, err := pipe.Exec(); err != nil {
		return err
	}
	return nil
}

func (m *_UserRedisMgr) SaveBatch(objs []*User) error {
	return m.SaveBatchWithExpire(objs, 0)
}

func (m *_UserRedisMgr) Save(obj *User) error {
	return m.SaveWithExpire(obj, 0)
}

func (m *_UserRedisMgr) SaveBatchWithExpire(objs []*User, expire time.Duration) error {
	if len(objs) > 0 {
		pipe := m.BeginPipeline()
		for _, obj := range objs {
			err := m.addToPipeline(pipe, obj, expire)
			if err != nil {
				pipe.Close()
				return err
			}
		}
		if _, err := pipe.Exec(); err != nil {
			pipe.Close()
			return err
		}
	}
	return nil
}

func (m *_UserRedisMgr) SaveWithExpire(obj *User, expire time.Duration) error {
	if obj != nil {
		pipe := m.BeginPipeline()
		err := m.addToPipeline(pipe, obj, expire)
		if err != nil {
			pipe.Close()
			return err
		}
		if _, err = pipe.Exec(); err != nil {
			pipe.Close()
			return err
		}
	}
	return nil
}

func (m *_UserRedisMgr) addToPipeline(pipe *_UserRedisPipeline, obj *User, expire time.Duration) error {
	pk := obj.GetPrimaryKey()
	//! fields
	pipe.HSet(keyOfObject(obj, pk.Key()), "Id", fmt.Sprint(obj.Id))
	pipe.HSet(keyOfObject(obj, pk.Key()), "Name", fmt.Sprint(obj.Name))
	pipe.HSet(keyOfObject(obj, pk.Key()), "Mailbox", fmt.Sprint(obj.Mailbox))
	pipe.HSet(keyOfObject(obj, pk.Key()), "Sex", fmt.Sprint(obj.Sex))
	pipe.HSet(keyOfObject(obj, pk.Key()), "Age", fmt.Sprint(obj.Age))
	pipe.HSet(keyOfObject(obj, pk.Key()), "Longitude", fmt.Sprint(obj.Longitude))
	pipe.HSet(keyOfObject(obj, pk.Key()), "Latitude", fmt.Sprint(obj.Latitude))
	pipe.HSet(keyOfObject(obj, pk.Key()), "Description", fmt.Sprint(obj.Description))
	pipe.HSet(keyOfObject(obj, pk.Key()), "Password", fmt.Sprint(obj.Password))
	pipe.HSet(keyOfObject(obj, pk.Key()), "HeadUrl", orm.Encode(fmt.Sprint(obj.HeadUrl)))
	pipe.HSet(keyOfObject(obj, pk.Key()), "Status", fmt.Sprint(obj.Status))
	pipe.HSet(keyOfObject(obj, pk.Key()), "CreatedAt", fmt.Sprint(obj.CreatedAt.Unix()))
	pipe.HSet(keyOfObject(obj, pk.Key()), "UpdatedAt", fmt.Sprint(obj.UpdatedAt.Unix()))
	if obj.DeletedAt != nil {
		pipe.HSet(keyOfObject(obj, pk.Key()), "DeletedAt", fmt.Sprint(obj.DeletedAt.Unix()))
	} else {
		pipe.HSet(keyOfObject(obj, pk.Key()), "DeletedAt", "nil")
	}

	//! uniques
	uk_key_0 := []string{
		"Mailbox",
		fmt.Sprint(obj.Mailbox),
		"Password",
		fmt.Sprint(obj.Password),
	}
	uk_pip_0 := MailboxPasswordOfUserUKRelationRedisMgr().BeginPipeline(pipe.Pipeline)
	uk_rel_0 := MailboxPasswordOfUserUKRelationRedisMgr().NewMailboxPasswordOfUserUKRelation(strings.Join(uk_key_0, ":"))
	uk_rel_0.Value = pk.Key()
	if err := uk_pip_0.PairAdd(uk_rel_0); err != nil {
		return err
	}

	//! indexes
	idx_key_0 := []string{
		"Sex",
		fmt.Sprint(obj.Sex),
	}
	idx_pip_0 := SexOfUserIDXRelationRedisMgr().BeginPipeline(pipe.Pipeline)
	idx_rel_0 := SexOfUserIDXRelationRedisMgr().NewSexOfUserIDXRelation(strings.Join(idx_key_0, ":"))
	idx_rel_0.Value = pk.Key()
	if err := idx_pip_0.SetAdd(idx_rel_0); err != nil {
		return err
	}

	//! ranges
	rg_key_0 := []string{
		"Id",
	}
	rg_pip_0 := IdOfUserRNGRelationRedisMgr().BeginPipeline(pipe.Pipeline)
	rg_rel_0 := IdOfUserRNGRelationRedisMgr().NewIdOfUserRNGRelation(strings.Join(rg_key_0, ":"))
	score_rg_0, err := orm.ToFloat64(obj.Id)
	if err != nil {
		return err
	}
	rg_rel_0.Score = score_rg_0
	rg_rel_0.Value = pk.Key()
	if err := rg_pip_0.ZSetAdd(rg_rel_0); err != nil {
		return err
	}
	rg_key_1 := []string{
		"Age",
	}
	rg_pip_1 := AgeOfUserRNGRelationRedisMgr().BeginPipeline(pipe.Pipeline)
	rg_rel_1 := AgeOfUserRNGRelationRedisMgr().NewAgeOfUserRNGRelation(strings.Join(rg_key_1, ":"))
	score_rg_1, err := orm.ToFloat64(obj.Age)
	if err != nil {
		return err
	}
	rg_rel_1.Score = score_rg_1
	rg_rel_1.Value = pk.Key()
	if err := rg_pip_1.ZSetAdd(rg_rel_1); err != nil {
		return err
	}
	if expire > 0 {
		pipe.Expire(keyOfObject(obj, pk.Key()), expire)
	}

	return nil
}

func (m *_UserRedisMgr) Clear() error {
	if strs, err := m.Keys(pairOfClass("User", "*")).Result(); err == nil {
		if len(strs) > 0 {
			m.Del(strs...)
		}
	}
	if strs, err := m.Keys(hashOfClass("User", "object", "*")).Result(); err == nil {
		if len(strs) > 0 {
			m.Del(strs...)
		}
	}
	if strs, err := m.Keys(setOfClass("User", "*")).Result(); err == nil {
		if len(strs) > 0 {
			m.Del(strs...)
		}
	}
	if strs, err := m.Keys(zsetOfClass("User", "*")).Result(); err == nil {
		if len(strs) > 0 {
			m.Del(strs...)
		}
	}
	if strs, err := m.Keys(geoOfClass("User", "*")).Result(); err == nil {
		if len(strs) > 0 {
			m.Del(strs...)
		}
	}
	if strs, err := m.Keys(listOfClass("User", "*")).Result(); err == nil {
		if len(strs) > 0 {
			m.Del(strs...)
		}
	}
	return nil
}

//! uniques

//! relation
type MailboxPasswordOfUserUKRelation struct {
	Key   string `db:"key" json:"key"`
	Value string `db:"value" json:"value"`
}

func (relation *MailboxPasswordOfUserUKRelation) GetClassName() string {
	return "MailboxPasswordOfUserUKRelation"
}

func (relation *MailboxPasswordOfUserUKRelation) GetIndexes() []string {
	idx := []string{}
	return idx
}

func (relation *MailboxPasswordOfUserUKRelation) GetStoreType() string {
	return "pair"
}

type _MailboxPasswordOfUserUKRelationRedisMgr struct {
	*orm.RedisStore
}

func MailboxPasswordOfUserUKRelationRedisMgr(stores ...*orm.RedisStore) *_MailboxPasswordOfUserUKRelationRedisMgr {
	if len(stores) > 0 {
		return &_MailboxPasswordOfUserUKRelationRedisMgr{stores[0]}
	}
	return &_MailboxPasswordOfUserUKRelationRedisMgr{_redis_store}
}

func (m *_MailboxPasswordOfUserUKRelationRedisMgr) NewMailboxPasswordOfUserUKRelation(key string) *MailboxPasswordOfUserUKRelation {
	return &MailboxPasswordOfUserUKRelation{
		Key: key,
	}
}

//! pipeline
type _MailboxPasswordOfUserUKRelationRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_MailboxPasswordOfUserUKRelationRedisMgr) BeginPipeline(pipes ...*redis.Pipeline) *_MailboxPasswordOfUserUKRelationRedisPipeline {
	if len(pipes) > 0 {
		return &_MailboxPasswordOfUserUKRelationRedisPipeline{pipes[0], nil}
	}
	return &_MailboxPasswordOfUserUKRelationRedisPipeline{m.Pipeline(), nil}
}

//! redis relation pair
func (m *_MailboxPasswordOfUserUKRelationRedisMgr) PairAdd(obj *MailboxPasswordOfUserUKRelation) error {
	return m.Set(pairOfClass("User", obj.GetClassName(), obj.Key), obj.Value, 0).Err()
}

func (pipe *_MailboxPasswordOfUserUKRelationRedisPipeline) PairAdd(obj *MailboxPasswordOfUserUKRelation) error {
	return pipe.Set(pairOfClass("User", obj.GetClassName(), obj.Key), obj.Value, 0).Err()
}

func (m *_MailboxPasswordOfUserUKRelationRedisMgr) PairGet(key string) (*MailboxPasswordOfUserUKRelation, error) {
	str, err := m.Get(pairOfClass("User", "MailboxPasswordOfUserUKRelation", key)).Result()
	if err != nil {
		return nil, err
	}

	obj := m.NewMailboxPasswordOfUserUKRelation(key)
	if err := orm.StringScan(str, &obj.Value); err != nil {
		return nil, err
	}
	return obj, nil
}

func (m *_MailboxPasswordOfUserUKRelationRedisMgr) PairRem(key string) error {
	return m.Del(pairOfClass("User", "MailboxPasswordOfUserUKRelation", key)).Err()
}

func (pipe *_MailboxPasswordOfUserUKRelationRedisPipeline) PairRem(key string) error {
	return pipe.Del(pairOfClass("User", "MailboxPasswordOfUserUKRelation", key)).Err()
}

func (m *_MailboxPasswordOfUserUKRelationRedisMgr) FindOne(key string) (string, error) {
	return m.Get(pairOfClass("User", "MailboxPasswordOfUserUKRelation", key)).Result()
}

func (m *_MailboxPasswordOfUserUKRelationRedisMgr) Clear() error {
	strs, err := m.Keys(pairOfClass("User", "MailboxPasswordOfUserUKRelation", "*")).Result()
	if err != nil {
		return err
	}
	if len(strs) > 0 {
		return m.Del(strs...).Err()
	}
	return nil
}

//! indexes

//! relation
type SexOfUserIDXRelation struct {
	Key   string `db:"key" json:"key"`
	Value string `db:"value" json:"value"`
}

func (relation *SexOfUserIDXRelation) GetClassName() string {
	return "SexOfUserIDXRelation"
}

func (relation *SexOfUserIDXRelation) GetIndexes() []string {
	idx := []string{}
	return idx
}

func (relation *SexOfUserIDXRelation) GetStoreType() string {
	return "set"
}

type _SexOfUserIDXRelationRedisMgr struct {
	*orm.RedisStore
}

func SexOfUserIDXRelationRedisMgr(stores ...*orm.RedisStore) *_SexOfUserIDXRelationRedisMgr {
	if len(stores) > 0 {
		return &_SexOfUserIDXRelationRedisMgr{stores[0]}
	}
	return &_SexOfUserIDXRelationRedisMgr{_redis_store}
}

func (m *_SexOfUserIDXRelationRedisMgr) NewSexOfUserIDXRelation(key string) *SexOfUserIDXRelation {
	return &SexOfUserIDXRelation{
		Key: key,
	}
}

//! pipeline
type _SexOfUserIDXRelationRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_SexOfUserIDXRelationRedisMgr) BeginPipeline(pipes ...*redis.Pipeline) *_SexOfUserIDXRelationRedisPipeline {
	if len(pipes) > 0 {
		return &_SexOfUserIDXRelationRedisPipeline{pipes[0], nil}
	}
	return &_SexOfUserIDXRelationRedisPipeline{m.Pipeline(), nil}
}

//! redis relation pair
func (m *_SexOfUserIDXRelationRedisMgr) SetAdd(relation *SexOfUserIDXRelation) error {
	return m.SAdd(setOfClass("User", "SexOfUserIDXRelation", relation.Key), relation.Value).Err()
}

func (pipe *_SexOfUserIDXRelationRedisPipeline) SetAdd(relation *SexOfUserIDXRelation) error {
	return pipe.SAdd(setOfClass("User", "SexOfUserIDXRelation", relation.Key), relation.Value).Err()
}

func (m *_SexOfUserIDXRelationRedisMgr) SetGet(key string) ([]*SexOfUserIDXRelation, error) {
	strs, err := m.SMembers(setOfClass("User", "SexOfUserIDXRelation", key)).Result()
	if err != nil {
		return nil, err
	}

	relations := make([]*SexOfUserIDXRelation, 0, len(strs))
	for _, str := range strs {
		relation := m.NewSexOfUserIDXRelation(key)
		if err := orm.StringScan(str, &relation.Value); err != nil {
			return nil, err
		}
		relations = append(relations, relation)
	}
	return relations, nil
}

func (m *_SexOfUserIDXRelationRedisMgr) SetRem(relation *SexOfUserIDXRelation) error {
	return m.SRem(setOfClass("User", "SexOfUserIDXRelation", relation.Key), relation.Value).Err()
}

func (pipe *_SexOfUserIDXRelationRedisPipeline) SetRem(relation *SexOfUserIDXRelation) error {
	return pipe.SRem(setOfClass("User", "SexOfUserIDXRelation", relation.Key), relation.Value).Err()
}

func (m *_SexOfUserIDXRelationRedisMgr) SetDel(key string) error {
	return m.Del(setOfClass("User", "SexOfUserIDXRelation", key)).Err()
}

func (pipe *_SexOfUserIDXRelationRedisPipeline) SetDel(key string) error {
	return pipe.Del(setOfClass("User", "SexOfUserIDXRelation", key)).Err()
}

func (m *_SexOfUserIDXRelationRedisMgr) Find(key string) ([]string, error) {
	return m.SMembers(setOfClass("User", "SexOfUserIDXRelation", key)).Result()
}

func (m *_SexOfUserIDXRelationRedisMgr) Clear() error {
	strs, err := m.Keys(setOfClass("User", "SexOfUserIDXRelation", "*")).Result()
	if err != nil {
		return err
	}
	if len(strs) > 0 {
		return m.Del(strs...).Err()
	}
	return nil
}

//! ranges

//! relation
type IdOfUserRNGRelation struct {
	Key   string  `db:"key" json:"key"`
	Score float64 `db:"score" json:"score"`
	Value string  `db:"value" json:"value"`
}

func (relation *IdOfUserRNGRelation) GetClassName() string {
	return "IdOfUserRNGRelation"
}

func (relation *IdOfUserRNGRelation) GetIndexes() []string {
	idx := []string{}
	return idx
}

func (relation *IdOfUserRNGRelation) GetStoreType() string {
	return "zset"
}

type _IdOfUserRNGRelationRedisMgr struct {
	*orm.RedisStore
}

func IdOfUserRNGRelationRedisMgr(stores ...*orm.RedisStore) *_IdOfUserRNGRelationRedisMgr {
	if len(stores) > 0 {
		return &_IdOfUserRNGRelationRedisMgr{stores[0]}
	}
	return &_IdOfUserRNGRelationRedisMgr{_redis_store}
}

func (m *_IdOfUserRNGRelationRedisMgr) NewIdOfUserRNGRelation(key string) *IdOfUserRNGRelation {
	return &IdOfUserRNGRelation{
		Key: key,
	}
}

//! pipeline
type _IdOfUserRNGRelationRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_IdOfUserRNGRelationRedisMgr) BeginPipeline(pipes ...*redis.Pipeline) *_IdOfUserRNGRelationRedisPipeline {
	if len(pipes) > 0 {
		return &_IdOfUserRNGRelationRedisPipeline{pipes[0], nil}
	}
	return &_IdOfUserRNGRelationRedisPipeline{m.Pipeline(), nil}
}

//! redis relation zset
func (m *_IdOfUserRNGRelationRedisMgr) ZSetAdd(relation *IdOfUserRNGRelation) error {
	return m.ZAdd(zsetOfClass("User", "IdOfUserRNGRelation", relation.Key), redis.Z{Score: relation.Score, Member: relation.Value}).Err()
}

func (pipe *_IdOfUserRNGRelationRedisPipeline) ZSetAdd(relation *IdOfUserRNGRelation) error {
	return pipe.ZAdd(zsetOfClass("User", "IdOfUserRNGRelation", relation.Key), redis.Z{Score: relation.Score, Member: relation.Value}).Err()
}

func (m *_IdOfUserRNGRelationRedisMgr) ZSetRange(key string, min, max int64) ([]*IdOfUserRNGRelation, error) {
	strs, err := m.ZRange(zsetOfClass("IdOfUserRNGRelation", key), min, max).Result()
	if err != nil {
		return nil, err
	}

	relations := make([]*IdOfUserRNGRelation, 0, len(strs))
	for _, str := range strs {
		relation := m.NewIdOfUserRNGRelation(key)
		if err := orm.StringScan(str, &relation.Value); err != nil {
			return nil, err
		}
		relations = append(relations, relation)
	}
	return relations, nil
}

func (m *_IdOfUserRNGRelationRedisMgr) ZSetRevertRange(key string, min, max int64) ([]*IdOfUserRNGRelation, error) {
	strs, err := m.ZRevRange(zsetOfClass("IdOfUserRNGRelation", key), min, max).Result()
	if err != nil {
		return nil, err
	}

	relations := make([]*IdOfUserRNGRelation, 0, len(strs))
	for _, str := range strs {
		relation := m.NewIdOfUserRNGRelation(key)
		if err := orm.StringScan(str, &relation.Value); err != nil {
			return nil, err
		}
		relations = append(relations, relation)
	}
	return relations, nil
}

func (m *_IdOfUserRNGRelationRedisMgr) ZSetRem(relation *IdOfUserRNGRelation) error {
	return m.ZRem(zsetOfClass("User", "IdOfUserRNGRelation", relation.Key), relation.Value).Err()
}

func (pipe *_IdOfUserRNGRelationRedisPipeline) ZSetRem(relation *IdOfUserRNGRelation) error {
	return pipe.ZRem(zsetOfClass("User", "IdOfUserRNGRelation", relation.Key), relation.Value).Err()
}

func (m *_IdOfUserRNGRelationRedisMgr) ZSetDel(key string) error {
	return m.Del(setOfClass("User", "IdOfUserRNGRelation", key)).Err()
}

func (pipe *_IdOfUserRNGRelationRedisPipeline) ZSetDel(key string) error {
	return pipe.Del(setOfClass("User", "IdOfUserRNGRelation", key)).Err()
}

func (m *_IdOfUserRNGRelationRedisMgr) Range(key string, min, max int64) ([]string, error) {
	return m.ZRange(zsetOfClass("User", "IdOfUserRNGRelation", key), min, max).Result()
}

func (m *_IdOfUserRNGRelationRedisMgr) RangeRevert(key string, min, max int64) ([]string, error) {
	return m.ZRevRange(zsetOfClass("User", "IdOfUserRNGRelation", key), min, max).Result()
}

func (m *_IdOfUserRNGRelationRedisMgr) Clear() error {
	strs, err := m.Keys(zsetOfClass("User", "IdOfUserRNGRelation", "*")).Result()
	if err != nil {
		return err
	}
	if len(strs) > 0 {
		return m.Del(strs...).Err()
	}
	return nil
}

//! relation
type AgeOfUserRNGRelation struct {
	Key   string  `db:"key" json:"key"`
	Score float64 `db:"score" json:"score"`
	Value string  `db:"value" json:"value"`
}

func (relation *AgeOfUserRNGRelation) GetClassName() string {
	return "AgeOfUserRNGRelation"
}

func (relation *AgeOfUserRNGRelation) GetIndexes() []string {
	idx := []string{}
	return idx
}

func (relation *AgeOfUserRNGRelation) GetStoreType() string {
	return "zset"
}

type _AgeOfUserRNGRelationRedisMgr struct {
	*orm.RedisStore
}

func AgeOfUserRNGRelationRedisMgr(stores ...*orm.RedisStore) *_AgeOfUserRNGRelationRedisMgr {
	if len(stores) > 0 {
		return &_AgeOfUserRNGRelationRedisMgr{stores[0]}
	}
	return &_AgeOfUserRNGRelationRedisMgr{_redis_store}
}

func (m *_AgeOfUserRNGRelationRedisMgr) NewAgeOfUserRNGRelation(key string) *AgeOfUserRNGRelation {
	return &AgeOfUserRNGRelation{
		Key: key,
	}
}

//! pipeline
type _AgeOfUserRNGRelationRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_AgeOfUserRNGRelationRedisMgr) BeginPipeline(pipes ...*redis.Pipeline) *_AgeOfUserRNGRelationRedisPipeline {
	if len(pipes) > 0 {
		return &_AgeOfUserRNGRelationRedisPipeline{pipes[0], nil}
	}
	return &_AgeOfUserRNGRelationRedisPipeline{m.Pipeline(), nil}
}

//! redis relation zset
func (m *_AgeOfUserRNGRelationRedisMgr) ZSetAdd(relation *AgeOfUserRNGRelation) error {
	return m.ZAdd(zsetOfClass("User", "AgeOfUserRNGRelation", relation.Key), redis.Z{Score: relation.Score, Member: relation.Value}).Err()
}

func (pipe *_AgeOfUserRNGRelationRedisPipeline) ZSetAdd(relation *AgeOfUserRNGRelation) error {
	return pipe.ZAdd(zsetOfClass("User", "AgeOfUserRNGRelation", relation.Key), redis.Z{Score: relation.Score, Member: relation.Value}).Err()
}

func (m *_AgeOfUserRNGRelationRedisMgr) ZSetRange(key string, min, max int64) ([]*AgeOfUserRNGRelation, error) {
	strs, err := m.ZRange(zsetOfClass("AgeOfUserRNGRelation", key), min, max).Result()
	if err != nil {
		return nil, err
	}

	relations := make([]*AgeOfUserRNGRelation, 0, len(strs))
	for _, str := range strs {
		relation := m.NewAgeOfUserRNGRelation(key)
		if err := orm.StringScan(str, &relation.Value); err != nil {
			return nil, err
		}
		relations = append(relations, relation)
	}
	return relations, nil
}

func (m *_AgeOfUserRNGRelationRedisMgr) ZSetRevertRange(key string, min, max int64) ([]*AgeOfUserRNGRelation, error) {
	strs, err := m.ZRevRange(zsetOfClass("AgeOfUserRNGRelation", key), min, max).Result()
	if err != nil {
		return nil, err
	}

	relations := make([]*AgeOfUserRNGRelation, 0, len(strs))
	for _, str := range strs {
		relation := m.NewAgeOfUserRNGRelation(key)
		if err := orm.StringScan(str, &relation.Value); err != nil {
			return nil, err
		}
		relations = append(relations, relation)
	}
	return relations, nil
}

func (m *_AgeOfUserRNGRelationRedisMgr) ZSetRem(relation *AgeOfUserRNGRelation) error {
	return m.ZRem(zsetOfClass("User", "AgeOfUserRNGRelation", relation.Key), relation.Value).Err()
}

func (pipe *_AgeOfUserRNGRelationRedisPipeline) ZSetRem(relation *AgeOfUserRNGRelation) error {
	return pipe.ZRem(zsetOfClass("User", "AgeOfUserRNGRelation", relation.Key), relation.Value).Err()
}

func (m *_AgeOfUserRNGRelationRedisMgr) ZSetDel(key string) error {
	return m.Del(setOfClass("User", "AgeOfUserRNGRelation", key)).Err()
}

func (pipe *_AgeOfUserRNGRelationRedisPipeline) ZSetDel(key string) error {
	return pipe.Del(setOfClass("User", "AgeOfUserRNGRelation", key)).Err()
}

func (m *_AgeOfUserRNGRelationRedisMgr) Range(key string, min, max int64) ([]string, error) {
	return m.ZRange(zsetOfClass("User", "AgeOfUserRNGRelation", key), min, max).Result()
}

func (m *_AgeOfUserRNGRelationRedisMgr) RangeRevert(key string, min, max int64) ([]string, error) {
	return m.ZRevRange(zsetOfClass("User", "AgeOfUserRNGRelation", key), min, max).Result()
}

func (m *_AgeOfUserRNGRelationRedisMgr) Clear() error {
	strs, err := m.Keys(zsetOfClass("User", "AgeOfUserRNGRelation", "*")).Result()
	if err != nil {
		return err
	}
	if len(strs) > 0 {
		return m.Del(strs...).Err()
	}
	return nil
}
