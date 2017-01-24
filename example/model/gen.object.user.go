package model

import (
	"database/sql"
	"fmt"
	"github.com/ezbuy/redis-orm/orm"
	redis "gopkg.in/redis.v5"
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

type User struct {
	Id          int32     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Mailbox     string    `db:"mailbox" json:"mailbox"`
	Sex         bool      `db:"sex" json:"sex"`
	Age         int32     `db:"age" json:"age"`
	Longitude   float64   `db:"longitude" json:"longitude"`
	Latitude    float64   `db:"latitude" json:"latitude"`
	Description string    `db:"description" json:"description"`
	Password    string    `db:"password" json:"password"`
	HeadUrl     string    `db:"head_url" json:"head_url"`
	Status      int32     `db:"status" json:"status"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
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
		"`id`",
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
	}
	return columns
}
func (obj *User) GetIndexes() []string {
	idx := []string{
		"Id",
		"Sex",
		"Age",
	}
	return idx
}

func (obj *User) GetStoreType() string {
	return "hash"
}

func (obj *User) GetPrimaryName() string {
	return "Id"
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
		"mailbox = ?",
		"password = ?",
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

func (u *MailboxPasswordOfUserUK) UKRelation() UniqueRelation {
	return MailboxPasswordOfUserUKRelationRedisMgr()
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
		"sex = ?",
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

func (u *SexOfUserIDX) IDXRelation() IndexRelation {
	return SexOfUserIDXRelationRedisMgr()
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
			conditions = append(conditions, fmt.Sprintf("id %s ?", u.beginOp()))
		}
		if u.IdEnd != -1 {
			conditions = append(conditions, fmt.Sprintf("id %s ?", u.endOp()))
		}
	}
	if limit {
		return fmt.Sprintf("%s %s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("Id", u.revert), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("Id", u.revert))
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

func (u *IdOfUserRNG) RNGRelation() RangeRelation {
	return IdOfUserRNGRelationRedisMgr()
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
			conditions = append(conditions, fmt.Sprintf("age %s ?", u.beginOp()))
		}
		if u.AgeEnd != -1 {
			conditions = append(conditions, fmt.Sprintf("age %s ?", u.endOp()))
		}
	}
	if limit {
		return fmt.Sprintf("%s %s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("Age", u.revert), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("Age", u.revert))
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

func (u *AgeOfUserRNG) RNGRelation() RangeRelation {
	return AgeOfUserRNGRelationRedisMgr()
}

func (m *_UserMgr) MySQL() *ReferenceResult {
	return NewReferenceResult(UserMySQLMgr())
}

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
			&(result.Age),
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
func (m *_UserMySQLMgr) Fetch(id interface{}) (*User, error) {
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

func (m *_UserMySQLMgr) FetchByIds(ids []interface{}) ([]*User, error) {
	if len(ids) == 0 {
		return []*User{}, nil
	}

	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM `users` WHERE `Id` IN (%s)", strings.Join(obj.GetColumns(), ","), orm.SliceJoin(ids, ","))
	objs, err := m.FetchBySQL(query)
	if err != nil {
		return nil, err
	}
	results := make([]*User, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*User))
	}
	return results, nil
}

func (m *_UserMySQLMgr) FindOne(unique Unique) (interface{}, error) {
	objs, err := m.queryLimit(unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return "", err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return "", fmt.Errorf("User find record not found")
}

func (m *_UserMySQLMgr) FindOneFetch(unique Unique) (*User, error) {
	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM `users` %s", strings.Join(obj.GetColumns(), ","), unique.SQLFormat(true))
	objs, err := m.FetchBySQL(query, unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*User), nil
	}
	return nil, fmt.Errorf("none record")
}

func (m *_UserMySQLMgr) Find(index Index) ([]interface{}, error) {
	return m.queryLimit(index.SQLFormat(true), index.SQLLimit(), index.SQLParams()...)
}

func (m *_UserMySQLMgr) FindFetch(index Index) ([]*User, error) {
	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM `users` %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	objs, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return nil, err
	}
	results := make([]*User, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*User))
	}
	return results, nil
}

func (m *_UserMySQLMgr) FindCount(index Index) (int64, error) {
	return m.queryCount(index.SQLFormat(false), index.SQLParams()...)
}

func (m *_UserMySQLMgr) Range(scope Range) ([]interface{}, error) {
	return m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (m *_UserMySQLMgr) RangeFetch(scope Range) ([]*User, error) {
	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM `users` %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	objs, err := m.FetchBySQL(query, scope.SQLParams()...)
	if err != nil {
		return nil, err
	}
	results := make([]*User, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*User))
	}
	return results, nil
}

func (m *_UserMySQLMgr) RangeCount(scope Range) (int64, error) {
	return m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
}

func (m *_UserMySQLMgr) RangeRevert(scope Range) ([]interface{}, error) {
	scope.Revert(true)
	return m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (m *_UserMySQLMgr) RangeRevertFetch(scope Range) ([]*User, error) {
	scope.Revert(true)
	return m.RangeFetch(scope)
}

func (m *_UserMySQLMgr) queryLimit(where string, limit int, args ...interface{}) (results []interface{}, err error) {
	query := fmt.Sprintf("SELECT `id` FROM `users` %s", where)
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
		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("User query limit result error: %v", err)
	}
	return
}

func (m *_UserMySQLMgr) queryCount(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`id`) FROM `users` %s", where)
	rows, err := m.Query(query, args...)
	if err != nil {
		return 0, fmt.Errorf("User query count error: %v", err)
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
type _UserMySQLTx struct {
	*orm.MySQLTx
	err          error
	rowsAffected int64
}

func (m *_UserMySQLMgr) BeginTx() (*_UserMySQLTx, error) {
	tx, err := m.Begin()
	if err != nil {
		return nil, err
	}
	return &_UserMySQLTx{orm.NewMySQLTx(tx), nil, 0}, nil
}

func (tx *_UserMySQLTx) BatchCreate(objs []*User) error {
	if len(objs) == 0 {
		return nil
	}

	params := make([]string, 0, len(objs))
	values := make([]interface{}, 0, len(objs)*13)
	for _, obj := range objs {
		params = append(params, fmt.Sprintf("(%s)", strings.Join(orm.NewStringSlice(13, "?"), ",")))
		values = append(values, 0)
		values = append(values, obj.Name)
		values = append(values, obj.Mailbox)
		values = append(values, obj.Sex)
		values = append(values, obj.Age)
		values = append(values, obj.Longitude)
		values = append(values, obj.Latitude)
		values = append(values, obj.Description)
		values = append(values, obj.Password)
		values = append(values, obj.HeadUrl)
		values = append(values, obj.Status)
		values = append(values, orm.TimeFormat(obj.CreatedAt))
		values = append(values, orm.TimeFormat(obj.UpdatedAt))
	}
	query := fmt.Sprintf("INSERT INTO `users`(%s) VALUES %s", strings.Join(objs[0].GetColumns(), ","), strings.Join(params, ","))
	result, err := tx.Exec(query, values...)
	if err != nil {
		tx.err = err
		return err
	}
	tx.rowsAffected, tx.err = result.RowsAffected()
	return tx.err
}

func (tx *_UserMySQLTx) BatchDelete(objs []*User) error {
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
func (tx *_UserMySQLTx) UpdateBySQL(set, where string, args ...interface{}) error {
	query := fmt.Sprintf("UPDATE `users` SET %s", set)
	if where != "" {
		query = fmt.Sprintf("UPDATE `users` SET %s WHERE %s", set, where)
	}
	result, err := tx.Exec(query, args)
	if err != nil {
		tx.err = err
		return err
	}
	tx.rowsAffected, tx.err = result.RowsAffected()
	return tx.err
}

func (tx *_UserMySQLTx) Create(obj *User) error {
	params := orm.NewStringSlice(13, "?")
	q := fmt.Sprintf("INSERT INTO `users`(%s) VALUES(%s)",
		strings.Join(obj.GetColumns(), ","),
		strings.Join(params, ","))

	result, err := tx.Exec(q, 0, obj.Name, obj.Mailbox, obj.Sex, obj.Age, obj.Longitude, obj.Latitude, obj.Description, obj.Password, obj.HeadUrl, obj.Status, orm.TimeFormat(obj.CreatedAt), orm.TimeFormat(obj.UpdatedAt))
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

func (tx *_UserMySQLTx) Update(obj *User) error {
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
	}
	q := fmt.Sprintf("UPDATE `users` SET %s WHERE `id`=?",
		strings.Join(columns, ","))
	result, err := tx.Exec(q, obj.Name, obj.Mailbox, obj.Sex, obj.Age, obj.Longitude, obj.Latitude, obj.Description, obj.Password, obj.HeadUrl, obj.Status, orm.TimeFormat(obj.CreatedAt), orm.TimeFormat(obj.UpdatedAt), obj.Id)
	if err != nil {
		tx.err = err
		return err
	}
	tx.rowsAffected, tx.err = result.RowsAffected()
	return tx.err
}

func (tx *_UserMySQLTx) Save(obj *User) error {
	err := tx.Update(obj)
	if err != nil {
		return err
	}
	if tx.rowsAffected > 0 {
		return nil
	}
	return tx.Create(obj)
}

func (tx *_UserMySQLTx) Delete(obj *User) error {
	q := fmt.Sprintf("DELETE FROM `users` WHERE `id`=?")
	result, err := tx.Exec(q, obj.Id)
	if err != nil {
		tx.err = err
		return err
	}
	tx.rowsAffected, tx.err = result.RowsAffected()
	return tx.err
}

func (tx *_UserMySQLTx) DeleteByIds(ids []interface{}) error {
	if len(ids) == 0 {
		return nil
	}

	q := fmt.Sprintf("DELETE FROM `users` WHERE `id` IN (%s)",
		orm.SliceJoin(ids, ","))
	result, err := tx.Exec(q)
	if err != nil {
		tx.err = err
		return err
	}
	tx.rowsAffected, tx.err = result.RowsAffected()
	return tx.err
}

func (tx *_UserMySQLTx) Close() error {
	if tx.err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}

//! tx read
func (tx *_UserMySQLTx) FindOne(unique Unique) (interface{}, error) {
	objs, err := tx.queryLimit(unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		tx.err = err
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	tx.err = fmt.Errorf("User find record not found")
	return nil, tx.err
}

func (tx *_UserMySQLTx) FindOneFetch(unique Unique) (*User, error) {
	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM `users` %s", strings.Join(obj.GetColumns(), ","), unique.SQLFormat(true))
	objs, err := tx.FetchBySQL(query, unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*User), nil
	}
	return nil, fmt.Errorf("none record")
}

func (tx *_UserMySQLTx) Find(index Index) ([]interface{}, error) {
	return tx.queryLimit(index.SQLFormat(true), index.SQLLimit(), index.SQLParams()...)
}

func (tx *_UserMySQLTx) FindFetch(index Index) ([]*User, error) {
	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM `users` %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	objs, err := tx.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return nil, err
	}
	results := make([]*User, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*User))
	}
	return results, nil
}

func (tx *_UserMySQLTx) FindCount(index Index) (int64, error) {
	return tx.queryCount(index.SQLFormat(false), index.SQLParams()...)
}

func (tx *_UserMySQLTx) Range(scope Range) ([]interface{}, error) {
	return tx.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (tx *_UserMySQLTx) RangeFetch(scope Range) ([]*User, error) {
	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM `users` %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	objs, err := tx.FetchBySQL(query, scope.SQLParams()...)
	if err != nil {
		return nil, err
	}
	results := make([]*User, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*User))
	}
	return results, nil
}

func (tx *_UserMySQLTx) RangeCount(scope Range) (int64, error) {
	return tx.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
}

func (tx *_UserMySQLTx) RangeRevert(scope Range) ([]interface{}, error) {
	scope.Revert(true)
	return tx.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (tx *_UserMySQLTx) RangeRevertFetch(scope Range) ([]*User, error) {
	scope.Revert(true)
	return tx.RangeFetch(scope)
}

func (tx *_UserMySQLTx) queryLimit(where string, limit int, args ...interface{}) (results []interface{}, err error) {
	query := fmt.Sprintf("SELECT `id` FROM `users`")
	if where != "" {
		query += " WHERE "
		query += where
	}

	rows, err := tx.Query(query, args...)
	if err != nil {
		tx.err = err
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
			tx.err = err
			return nil, err
		}
		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		tx.err = err
		return nil, fmt.Errorf("User query limit result error: %v", err)
	}
	return
}

func (tx *_UserMySQLTx) queryCount(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`id`) FROM `users`")
	if where != "" {
		query += " WHERE "
		query += where
	}

	rows, err := tx.Query(query, args...)
	if err != nil {
		tx.err = err
		return 0, fmt.Errorf("User query limit error: %v", err)
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

func (tx *_UserMySQLTx) Fetch(id interface{}) (*User, error) {
	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM `users` WHERE `Id` = (%s)", strings.Join(obj.GetColumns(), ","), fmt.Sprint(id))
	objs, err := tx.FetchBySQL(query)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*User), nil
	}
	return nil, fmt.Errorf("User fetch record not found")
}

func (tx *_UserMySQLTx) FetchByIds(ids []interface{}) ([]*User, error) {
	if len(ids) == 0 {
		return []*User{}, nil
	}

	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM `users` WHERE `Id` IN (%s)", strings.Join(obj.GetColumns(), ","), orm.SliceJoin(ids, ","))
	objs, err := tx.FetchBySQL(query)
	if err != nil {
		return nil, err
	}
	results := make([]*User, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*User))
	}
	return results, nil
}

func (tx *_UserMySQLTx) FetchBySQL(sql string, args ...interface{}) (results []interface{}, err error) {
	rows, err := tx.Query(sql, args...)
	if err != nil {
		tx.err = err
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
			&(result.Age),
			&(result.Longitude),
			&(result.Latitude),
			&(result.Description),
			&(result.Password),
			&(result.HeadUrl),
			&(result.Status),
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
		return nil, fmt.Errorf("User fetch result error: %v", err)
	}
	return
}

func (m *_UserMgr) Redis() *ReferenceResult {
	return NewReferenceResult(UserRedisMgr())
}

type _UserRedisMgr struct {
	*orm.RedisStore
}

func UserRedisMgr() *_UserRedisMgr {
	return &_UserRedisMgr{_redis_store}
}

func NewUserRedisMgr(cf *RedisConfig) (*_UserRedisMgr, error) {
	store, err := orm.NewRedisStore(cf.Host, cf.Port, cf.Password, 0)
	if err != nil {
		return nil, err
	}
	return &_UserRedisMgr{store}, nil
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

func (m *_UserRedisMgr) Load(db DBFetcher) error {
	if err := m.Clear(); err != nil {
		return err
	}

	return m.AddBySQL(db, "SELECT `id`,`name`,`mailbox`,`sex`, `age`, `longitude`,`latitude`,`description`,`password`,`head_url`,`status`,`created_at`, `updated_at` FROM users")

}

func (m *_UserRedisMgr) AddBySQL(db DBFetcher, sql string, args ...interface{}) error {
	objs, err := db.FetchBySQL(sql, args...)
	if err != nil {
		return err
	}

	for _, obj := range objs {
		if err := m.Save(obj.(*User)); err != nil {
			return err
		}
	}

	return nil
}
func (m *_UserRedisMgr) DelBySQL(db DBFetcher, sql string, args ...interface{}) error {
	objs, err := db.FetchBySQL(sql, args...)
	if err != nil {
		return err
	}

	for _, obj := range objs {
		if err := m.Delete(obj.(*User)); err != nil {
			return err
		}
	}
	return nil
}

//! redis model read
func (m *_UserRedisMgr) FindOne(unique Unique) (interface{}, error) {
	if relation := unique.UKRelation(); relation != nil {
		str, err := relation.FindOne(unique.Key())
		if err != nil {
			return "", err
		}
		var val int32
		if err := m.StringScan(str, &val); err != nil {
			return nil, err
		}
		return val, nil
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

func (m *_UserRedisMgr) Find(index Index) ([]interface{}, error) {
	if relation := index.IDXRelation(); relation != nil {
		strs, err := relation.Find(index.Key())
		if err != nil {
			return nil, err
		}
		results := make([]interface{}, 0, len(strs))
		for _, str := range strs {
			var val int32
			if err := m.StringScan(str, &val); err != nil {
				return nil, err
			}
			results = append(results, val)
		}
		return results, nil
	}
	return nil, fmt.Errorf("index none relation.")
}

func (m *_UserRedisMgr) FindFetch(index Index) ([]*User, error) {
	vs, err := m.Find(index)
	if err != nil {
		return nil, err
	}
	return m.FetchByIds(vs)
}

func (m *_UserRedisMgr) FindCount(index Index) (int64, error) {
	if relation := index.IDXRelation(); relation != nil {
		strs, err := relation.Find(index.Key())
		if err != nil {
			return 0, err
		}
		return int64(len(strs)), nil
	}
	return 0, fmt.Errorf("index none relation.")
}

func (m *_UserRedisMgr) Range(scope Range) ([]interface{}, error) {
	if relation := scope.RNGRelation(); relation != nil {
		strs, err := relation.Range(scope.Key(), scope.Begin(), scope.End())
		if err != nil {
			return nil, err
		}
		results := make([]interface{}, 0, len(strs))
		for _, str := range strs {
			var val int32
			if err := m.StringScan(str, &val); err != nil {
				return nil, err
			}
			results = append(results, val)
		}
		return results, nil
	}
	return nil, fmt.Errorf("range none relation.")
}

func (m *_UserRedisMgr) RangeFetch(scope Range) ([]*User, error) {
	vs, err := m.Range(scope)
	if err != nil {
		return nil, err
	}
	return m.FetchByIds(vs)
}

func (m *_UserRedisMgr) RangeCount(scope Range) (int64, error) {
	if relation := scope.RNGRelation(); relation != nil {
		strs, err := relation.Range(scope.Key(), scope.Begin(), scope.End())
		if err != nil {
			return 0, err
		}
		return int64(len(strs)), nil
	}
	return 0, fmt.Errorf("range none relation.")
}

func (m *_UserRedisMgr) RangeRevert(scope Range) ([]interface{}, error) {
	if relation := scope.RNGRelation(); relation != nil {
		scope.Revert(true)
		strs, err := relation.RangeRevert(scope.Key(), scope.Begin(), scope.End())
		if err != nil {
			return nil, err
		}
		results := make([]interface{}, 0, len(strs))
		for _, str := range strs {
			var val int32
			if err := m.StringScan(str, &val); err != nil {
				return nil, err
			}
			results = append(results, val)
		}
		return results, nil
	}
	return nil, fmt.Errorf("revert range none relation.")
}

func (m *_UserRedisMgr) RangeRevertFetch(scope Range) ([]*User, error) {
	vs, err := m.RangeRevert(scope)
	if err != nil {
		return nil, err
	}
	return m.FetchByIds(vs)
}

func (m *_UserRedisMgr) Fetch(id interface{}) (*User, error) {
	obj := UserMgr.NewUser()

	pipe := m.BeginPipeline()
	pipe.Exists(keyOfObject(obj, fmt.Sprint(id)))
	pipe.HMGet(keyOfObject(obj, fmt.Sprint(id)), "Id", "Name", "Mailbox", "Sex", "Age", "Longitude", "Latitude", "Description", "Password", "HeadUrl", "Status", "CreatedAt", "UpdatedAt")
	cmds, err := pipe.Exec()
	if err != nil {
		return nil, err
	}

	if b, err := cmds[0].(*redis.BoolCmd).Result(); err == nil {
		if !b {
			return nil, fmt.Errorf("User Id:(%v) not exist", id)
		}
	}

	strs, err := cmds[1].(*redis.SliceCmd).Result()
	if err != nil {
		return nil, err
	}
	if err := m.StringScan(strs[0].(string), &obj.Id); err != nil {
		return nil, err
	}
	if err := m.StringScan(strs[1].(string), &obj.Name); err != nil {
		return nil, err
	}
	if err := m.StringScan(strs[2].(string), &obj.Mailbox); err != nil {
		return nil, err
	}
	if err := m.StringScan(strs[3].(string), &obj.Sex); err != nil {
		return nil, err
	}
	if err := m.StringScan(strs[4].(string), &obj.Age); err != nil {
		return nil, err
	}
	if err := m.StringScan(strs[5].(string), &obj.Longitude); err != nil {
		return nil, err
	}
	if err := m.StringScan(strs[6].(string), &obj.Latitude); err != nil {
		return nil, err
	}
	if err := m.StringScan(strs[7].(string), &obj.Description); err != nil {
		return nil, err
	}
	if err := m.StringScan(strs[8].(string), &obj.Password); err != nil {
		return nil, err
	}
	if err := m.StringScan(strs[9].(string), &obj.HeadUrl); err != nil {
		return nil, err
	}
	if err := m.StringScan(strs[10].(string), &obj.Status); err != nil {
		return nil, err
	}
	var val11 string
	if err := m.StringScan(strs[11].(string), &val11); err != nil {
		return nil, err
	}
	obj.CreatedAt = orm.TimeParse(val11)
	var val12 string
	if err := m.StringScan(strs[12].(string), &val12); err != nil {
		return nil, err
	}
	obj.UpdatedAt = orm.TimeParse(val12)
	return obj, nil
}

func (m *_UserRedisMgr) FetchByIds(ids []interface{}) ([]*User, error) {
	objs := make([]*User, 0, len(ids))
	pipe := m.BeginPipeline()
	obj := UserMgr.NewUser()
	for _, id := range ids {
		pipe.Exists(keyOfObject(obj, fmt.Sprint(id)))
		pipe.HMGet(keyOfObject(obj, fmt.Sprint(id)), "Id", "Name", "Mailbox", "Sex", "Age", "Longitude", "Latitude", "Description", "Password", "HeadUrl", "Status", "CreatedAt", "UpdatedAt")
	}
	cmds, err := pipe.Exec()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(ids); i++ {
		if b, err := cmds[2*i].(*redis.BoolCmd).Result(); err == nil {
			if !b {
				return nil, fmt.Errorf("User Id:(%v) not exist", ids[i])
			}
		}

		strs, err := cmds[2*i+1].(*redis.SliceCmd).Result()
		if err != nil {
			return nil, err
		}

		obj := UserMgr.NewUser()
		if err := m.StringScan(strs[0].(string), &obj.Id); err != nil {
			return nil, err
		}
		if err := m.StringScan(strs[1].(string), &obj.Name); err != nil {
			return nil, err
		}
		if err := m.StringScan(strs[2].(string), &obj.Mailbox); err != nil {
			return nil, err
		}
		if err := m.StringScan(strs[3].(string), &obj.Sex); err != nil {
			return nil, err
		}
		if err := m.StringScan(strs[4].(string), &obj.Age); err != nil {
			return nil, err
		}
		if err := m.StringScan(strs[5].(string), &obj.Longitude); err != nil {
			return nil, err
		}
		if err := m.StringScan(strs[6].(string), &obj.Latitude); err != nil {
			return nil, err
		}
		if err := m.StringScan(strs[7].(string), &obj.Description); err != nil {
			return nil, err
		}
		if err := m.StringScan(strs[8].(string), &obj.Password); err != nil {
			return nil, err
		}
		if err := m.StringScan(strs[9].(string), &obj.HeadUrl); err != nil {
			return nil, err
		}
		if err := m.StringScan(strs[10].(string), &obj.Status); err != nil {
			return nil, err
		}
		var val11 string
		if err := m.StringScan(strs[11].(string), &val11); err != nil {
			return nil, err
		}
		obj.CreatedAt = orm.TimeParse(val11)
		var val12 string
		if err := m.StringScan(strs[12].(string), &val12); err != nil {
			return nil, err
		}
		obj.UpdatedAt = orm.TimeParse(val12)
		objs = append(objs, obj)
	}
	return objs, nil
}

func (m *_UserRedisMgr) Create(obj *User) error {
	return m.Save(obj)
}

func (m *_UserRedisMgr) Update(obj *User) error {
	return m.Save(obj)
}

func (m *_UserRedisMgr) Delete(obj *User) error {
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
	idx_rel_0.Value = obj.Id
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
	rg_rel_0.Value = obj.Id
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
	rg_rel_1.Value = obj.Id
	if err := rg_pip_1.ZSetRem(rg_rel_1); err != nil {
		return err
	}

	if err := pipe.Del(keyOfObject(obj, fmt.Sprint(obj.Id))).Err(); err != nil {
		return err
	}

	if _, err := pipe.Exec(); err != nil {
		return err
	}
	return nil
}

func (m *_UserRedisMgr) Save(obj *User) error {
	pipe := m.BeginPipeline()
	//! fields
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Id", fmt.Sprint(obj.Id))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Name", fmt.Sprint(obj.Name))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Mailbox", fmt.Sprint(obj.Mailbox))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Sex", fmt.Sprint(obj.Sex))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Age", fmt.Sprint(obj.Age))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Longitude", fmt.Sprint(obj.Longitude))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Latitude", fmt.Sprint(obj.Latitude))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Description", fmt.Sprint(obj.Description))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Password", fmt.Sprint(obj.Password))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "HeadUrl", fmt.Sprint(obj.HeadUrl))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "Status", fmt.Sprint(obj.Status))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "CreatedAt", fmt.Sprint(orm.TimeFormat(obj.CreatedAt)))
	pipe.HSet(keyOfObject(obj, fmt.Sprint(obj.Id)), "UpdatedAt", fmt.Sprint(orm.TimeFormat(obj.UpdatedAt)))

	//! uniques
	uk_key_0 := []string{
		"Mailbox",
		fmt.Sprint(obj.Mailbox),
		"Password",
		fmt.Sprint(obj.Password),
	}
	uk_pip_0 := MailboxPasswordOfUserUKRelationRedisMgr().BeginPipeline(pipe.Pipeline)
	uk_rel_0 := MailboxPasswordOfUserUKRelationRedisMgr().NewMailboxPasswordOfUserUKRelation(strings.Join(uk_key_0, ":"))
	uk_rel_0.Value = obj.Id
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
	idx_rel_0.Value = obj.Id
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
	rg_rel_0.Value = obj.Id
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
	rg_rel_1.Value = obj.Id
	if err := rg_pip_1.ZSetAdd(rg_rel_1); err != nil {
		return err
	}

	if _, err := pipe.Exec(); err != nil {
		return err
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
	Value int32  `db:"value" json:"value"`
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

func (relation *MailboxPasswordOfUserUKRelation) GetPrimaryName() string {
	return "Key"
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
	if err := m.StringScan(str, &obj.Value); err != nil {
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
	Value int32  `db:"value" json:"value"`
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

func (relation *SexOfUserIDXRelation) GetPrimaryName() string {
	return "Key"
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
		if err := m.StringScan(str, &relation.Value); err != nil {
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
	Value int32   `db:"value" json:"value"`
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

func (relation *IdOfUserRNGRelation) GetPrimaryName() string {
	return "Key"
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
		if err := m.StringScan(str, &relation.Value); err != nil {
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
		if err := m.StringScan(str, &relation.Value); err != nil {
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
	Value int32   `db:"value" json:"value"`
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

func (relation *AgeOfUserRNGRelation) GetPrimaryName() string {
	return "Key"
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
		if err := m.StringScan(str, &relation.Value); err != nil {
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
		if err := m.StringScan(str, &relation.Value); err != nil {
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
