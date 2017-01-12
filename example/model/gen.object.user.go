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

func (u *MailboxPasswordOfUserUK) SQLFormat() string {
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

func (u *SexOfUserIDX) SQLFormat() string {
	conditions := []string{
		"sex = ?",
	}
	return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOffsetLimit(u.offset, u.limit))
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

func (u *IdOfUserRNG) SQLFormat() string {
	conditions := []string{}
	if u.IdBegin != u.IdEnd {
		if u.IdBegin != -1 {
			conditions = append(conditions, fmt.Sprintf("id %s ?", u.beginOp()))
		}
		if u.IdEnd != -1 {
			conditions = append(conditions, fmt.Sprintf("id %s ?", u.endOp()))
		}
	}
	return fmt.Sprintf("%s %s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("Id", u.revert), orm.SQLOffsetLimit(u.offset, u.limit))
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

func (u *AgeOfUserRNG) SQLFormat() string {
	conditions := []string{}
	if u.AgeBegin != u.AgeEnd {
		if u.AgeBegin != -1 {
			conditions = append(conditions, fmt.Sprintf("age %s ?", u.beginOp()))
		}
		if u.AgeEnd != -1 {
			conditions = append(conditions, fmt.Sprintf("age %s ?", u.endOp()))
		}
	}
	return fmt.Sprintf("%s %s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("Age", u.revert), orm.SQLOffsetLimit(u.offset, u.limit))
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
	results := make([]*User, 0, len(objs))
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

func (m *_UserMySQLMgr) RevertRange(scope Range) ([]string, error) {
	scope.Revert(true)
	return m.queryLimit(scope.SQLFormat(), scope.SQLLimit(), scope.SQLParams()...)
}

func (m *_UserMySQLMgr) queryLimit(where string, limit int, args ...interface{}) (results []string, err error) {
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
		results = append(results, fmt.Sprint(result))
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("User query limit result error: %v", err)
	}
	return
}

//! object.mysql.write
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

	ids := make([]string, 0, len(objs))
	for _, obj := range objs {
		ids = append(ids, fmt.Sprint(obj.Id))
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

func (tx *_UserMySQLTx) DeleteByIds(ids []string) error {
	if len(ids) == 0 {
		return nil
	}

	q := fmt.Sprintf("DELETE FROM `users` WHERE `id` IN (%s)",
		strings.Join(ids, ","))
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
func (tx *_UserMySQLTx) FindOne(unique Unique) (string, error) {
	objs, err := tx.queryLimit(unique.SQLFormat(), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		tx.err = err
		return "", err
	}
	if len(objs) > 0 {
		return fmt.Sprint(objs[0]), nil
	}
	tx.err = fmt.Errorf("User find record not found")
	return "", tx.err
}

func (tx *_UserMySQLTx) Find(index Index) ([]string, error) {
	return tx.queryLimit(index.SQLFormat(), index.SQLLimit(), index.SQLParams()...)
}

func (tx *_UserMySQLTx) Range(scope Range) ([]string, error) {
	return tx.queryLimit(scope.SQLFormat(), scope.SQLLimit(), scope.SQLParams()...)
}

func (tx *_UserMySQLTx) RevertRange(scope Range) ([]string, error) {
	scope.Revert(true)
	return tx.queryLimit(scope.SQLFormat(), scope.SQLLimit(), scope.SQLParams()...)
}

func (tx *_UserMySQLTx) queryLimit(where string, limit int, args ...interface{}) (results []string, err error) {
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
		results = append(results, fmt.Sprint(result))
	}
	if err := rows.Err(); err != nil {
		tx.err = err
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

//! pipeline write
type _UserRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_UserRedisMgr) BeginPipeline() *_UserRedisPipeline {
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
func (m *_UserRedisMgr) FindOne(unique Unique) (string, error) {
	if relation := unique.UKRelation(); relation != nil {
		return relation.FindOne(unique.Key())
	}
	return "", nil
}

func (m *_UserRedisMgr) Find(index Index) ([]string, error) {
	if relation := index.IDXRelation(); relation != nil {
		return relation.Find(index.Key())
	}
	return nil, nil
}

func (m *_UserRedisMgr) Range(scope Range) ([]string, error) {
	if relation := scope.RNGRelation(); relation != nil {
		return relation.Range(scope.Key(), scope.Begin(), scope.End())
	}
	return nil, nil
}

func (m *_UserRedisMgr) RevertRange(scope Range) ([]string, error) {
	if relation := scope.RNGRelation(); relation != nil {
		scope.Revert(true)
		return relation.RevertRange(scope.Key(), scope.Begin(), scope.End())
	}
	return nil, nil
}

func (m *_UserRedisMgr) Fetch(id string) (*User, error) {
	obj := UserMgr.NewUser()

	b, err := m.Exists(keyOfObject(obj, id)).Result()
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, fmt.Errorf("User Id:(%s) not exist", id)
	}

	strs, err := m.HMGet(keyOfObject(obj, id), "Id", "Name", "Mailbox", "Sex", "Age", "Longitude", "Latitude", "Description", "Password", "HeadUrl", "Status", "CreatedAt", "UpdatedAt").Result()
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

func (m *_UserRedisMgr) FetchByIds(ids []string) ([]*User, error) {
	objs := make([]*User, 0, len(ids))
	for _, id := range ids {
		obj, err := m.Fetch(id)
		if err != nil {
			return objs, err
		}
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
	//! uniques
	uk_key_0 := []string{
		"Mailbox",
		fmt.Sprint(obj.Mailbox),
		"Password",
		fmt.Sprint(obj.Password),
	}
	uk_mgr_0 := MailboxPasswordOfUserUKRelationRedisMgr(m.RedisStore)
	if err := uk_mgr_0.PairRem(strings.Join(uk_key_0, ":")); err != nil {
		return err
	}

	//! indexes
	idx_key_0 := []string{
		"Sex",
		fmt.Sprint(obj.Sex),
	}
	idx_mgr_0 := SexOfUserIDXRelationRedisMgr(m.RedisStore)
	idx_rel_0 := idx_mgr_0.NewSexOfUserIDXRelation(strings.Join(idx_key_0, ":"))
	idx_rel_0.Value = obj.Id
	if err := idx_mgr_0.SetRem(idx_rel_0); err != nil {
		return err
	}

	//! ranges
	rg_key_0 := []string{
		"Id",
	}
	rg_mgr_0 := IdOfUserRNGRelationRedisMgr(m.RedisStore)
	rg_rel_0 := rg_mgr_0.NewIdOfUserRNGRelation(strings.Join(rg_key_0, ":"))
	score_rg_0, err := orm.ToFloat64(obj.Id)
	if err != nil {
		return err
	}
	rg_rel_0.Score = score_rg_0
	rg_rel_0.Value = obj.Id
	if err := rg_mgr_0.ZSetRem(rg_rel_0); err != nil {
		return err
	}
	rg_key_1 := []string{
		"Age",
	}
	rg_mgr_1 := AgeOfUserRNGRelationRedisMgr(m.RedisStore)
	rg_rel_1 := rg_mgr_1.NewAgeOfUserRNGRelation(strings.Join(rg_key_1, ":"))
	score_rg_1, err := orm.ToFloat64(obj.Age)
	if err != nil {
		return err
	}
	rg_rel_1.Score = score_rg_1
	rg_rel_1.Value = obj.Id
	if err := rg_mgr_1.ZSetRem(rg_rel_1); err != nil {
		return err
	}

	return m.Del(keyOfObject(obj, fmt.Sprint(obj.Id))).Err()
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
	if _, err := pipe.Exec(); err != nil {
		return err
	}

	//! uniques
	uk_key_0 := []string{
		"Mailbox",
		fmt.Sprint(obj.Mailbox),
		"Password",
		fmt.Sprint(obj.Password),
	}
	uk_mgr_0 := MailboxPasswordOfUserUKRelationRedisMgr(m.RedisStore)
	uk_rel_0 := uk_mgr_0.NewMailboxPasswordOfUserUKRelation(strings.Join(uk_key_0, ":"))
	uk_rel_0.Value = obj.Id
	if err := uk_mgr_0.PairAdd(uk_rel_0); err != nil {
		return err
	}

	//! indexes
	idx_key_0 := []string{
		"Sex",
		fmt.Sprint(obj.Sex),
	}
	idx_mgr_0 := SexOfUserIDXRelationRedisMgr(m.RedisStore)
	idx_rel_0 := idx_mgr_0.NewSexOfUserIDXRelation(strings.Join(idx_key_0, ":"))
	idx_rel_0.Value = obj.Id
	if err := idx_mgr_0.SetAdd(idx_rel_0); err != nil {
		return err
	}

	//! ranges
	rg_key_0 := []string{
		"Id",
	}
	rg_mgr_0 := IdOfUserRNGRelationRedisMgr(m.RedisStore)
	rg_rel_0 := rg_mgr_0.NewIdOfUserRNGRelation(strings.Join(rg_key_0, ":"))
	score_rg_0, err := orm.ToFloat64(obj.Id)
	if err != nil {
		return err
	}
	rg_rel_0.Score = score_rg_0
	rg_rel_0.Value = obj.Id
	if err := rg_mgr_0.ZSetAdd(rg_rel_0); err != nil {
		return err
	}
	rg_key_1 := []string{
		"Age",
	}
	rg_mgr_1 := AgeOfUserRNGRelationRedisMgr(m.RedisStore)
	rg_rel_1 := rg_mgr_1.NewAgeOfUserRNGRelation(strings.Join(rg_key_1, ":"))
	score_rg_1, err := orm.ToFloat64(obj.Age)
	if err != nil {
		return err
	}
	rg_rel_1.Score = score_rg_1
	rg_rel_1.Value = obj.Id
	if err := rg_mgr_1.ZSetAdd(rg_rel_1); err != nil {
		return err
	}

	return nil
}

func (m *_UserRedisMgr) Clear() error {
	if strs, err := m.Keys(pairOfClass("User", "*")).Result(); err == nil {
		m.Del(strs...)
	}
	if strs, err := m.Keys(hashOfClass("User", "object", "*")).Result(); err == nil {
		m.Del(strs...)
	}
	if strs, err := m.Keys(setOfClass("User", "*")).Result(); err == nil {
		m.Del(strs...)
	}
	if strs, err := m.Keys(zsetOfClass("User", "*")).Result(); err == nil {
		m.Del(strs...)
	}
	if strs, err := m.Keys(geoOfClass("User", "*")).Result(); err == nil {
		m.Del(strs...)
	}
	if strs, err := m.Keys(listOfClass("User", "*")).Result(); err == nil {
		m.Del(strs...)
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

//! pipeline write
type _MailboxPasswordOfUserUKRelationRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_MailboxPasswordOfUserUKRelationRedisMgr) BeginPipeline() *_MailboxPasswordOfUserUKRelationRedisPipeline {
	return &_MailboxPasswordOfUserUKRelationRedisPipeline{m.Pipeline(), nil}
}

func (m *_MailboxPasswordOfUserUKRelationRedisMgr) NewMailboxPasswordOfUserUKRelation(key string) *MailboxPasswordOfUserUKRelation {
	return &MailboxPasswordOfUserUKRelation{
		Key: key,
	}
}

//! redis relation pair
func (m *_MailboxPasswordOfUserUKRelationRedisMgr) PairAdd(obj *MailboxPasswordOfUserUKRelation) error {
	return m.Set(pairOfClass("User", obj.GetClassName(), obj.Key), obj.Value, 0).Err()
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

func (m *_MailboxPasswordOfUserUKRelationRedisMgr) FindOne(key string) (string, error) {
	return m.Get(pairOfClass("User", "MailboxPasswordOfUserUKRelation", key)).Result()
}

func (m *_MailboxPasswordOfUserUKRelationRedisMgr) Clear() error {
	strs, err := m.Keys(pairOfClass("User", "MailboxPasswordOfUserUKRelation", "*")).Result()
	if err != nil {
		return err
	}
	return m.Del(strs...).Err()
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

//! pipeline write
type _SexOfUserIDXRelationRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_SexOfUserIDXRelationRedisMgr) BeginPipeline() *_SexOfUserIDXRelationRedisPipeline {
	return &_SexOfUserIDXRelationRedisPipeline{m.Pipeline(), nil}
}

func (m *_SexOfUserIDXRelationRedisMgr) NewSexOfUserIDXRelation(key string) *SexOfUserIDXRelation {
	return &SexOfUserIDXRelation{
		Key: key,
	}
}

//! redis relation pair
func (m *_SexOfUserIDXRelationRedisMgr) SetAdd(relation *SexOfUserIDXRelation) error {
	return m.SAdd(setOfClass("User", "SexOfUserIDXRelation", relation.Key), relation.Value).Err()
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

func (m *_SexOfUserIDXRelationRedisMgr) SetDel(key string) error {
	return m.Del(setOfClass("User", "SexOfUserIDXRelation", key)).Err()
}

func (m *_SexOfUserIDXRelationRedisMgr) Find(key string) ([]string, error) {
	return m.SMembers(setOfClass("User", "SexOfUserIDXRelation", key)).Result()
}

func (m *_SexOfUserIDXRelationRedisMgr) Clear() error {
	strs, err := m.Keys(setOfClass("User", "SexOfUserIDXRelation", "*")).Result()
	if err != nil {
		return err
	}
	return m.Del(strs...).Err()
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

//! pipeline write
type _IdOfUserRNGRelationRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_IdOfUserRNGRelationRedisMgr) BeginPipeline() *_IdOfUserRNGRelationRedisPipeline {
	return &_IdOfUserRNGRelationRedisPipeline{m.Pipeline(), nil}
}

func (m *_IdOfUserRNGRelationRedisMgr) NewIdOfUserRNGRelation(key string) *IdOfUserRNGRelation {
	return &IdOfUserRNGRelation{
		Key: key,
	}
}

//! redis relation zset
func (m *_IdOfUserRNGRelationRedisMgr) ZSetAdd(relation *IdOfUserRNGRelation) error {
	return m.ZAdd(zsetOfClass("User", "IdOfUserRNGRelation", relation.Key), redis.Z{Score: relation.Score, Member: relation.Value}).Err()
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

func (m *_IdOfUserRNGRelationRedisMgr) ZSetDel(key string) error {
	return m.Del(setOfClass("User", "IdOfUserRNGRelation", key)).Err()
}

func (m *_IdOfUserRNGRelationRedisMgr) Range(key string, min, max int64) ([]string, error) {
	return m.ZRange(zsetOfClass("User", "IdOfUserRNGRelation", key), min, max).Result()
}

func (m *_IdOfUserRNGRelationRedisMgr) RevertRange(key string, min, max int64) ([]string, error) {
	return m.ZRevRange(zsetOfClass("User", "IdOfUserRNGRelation", key), min, max).Result()
}

func (m *_IdOfUserRNGRelationRedisMgr) Clear() error {
	strs, err := m.Keys(zsetOfClass("User", "IdOfUserRNGRelation", "*")).Result()
	if err != nil {
		return err
	}
	return m.Del(strs...).Err()
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

//! pipeline write
type _AgeOfUserRNGRelationRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_AgeOfUserRNGRelationRedisMgr) BeginPipeline() *_AgeOfUserRNGRelationRedisPipeline {
	return &_AgeOfUserRNGRelationRedisPipeline{m.Pipeline(), nil}
}

func (m *_AgeOfUserRNGRelationRedisMgr) NewAgeOfUserRNGRelation(key string) *AgeOfUserRNGRelation {
	return &AgeOfUserRNGRelation{
		Key: key,
	}
}

//! redis relation zset
func (m *_AgeOfUserRNGRelationRedisMgr) ZSetAdd(relation *AgeOfUserRNGRelation) error {
	return m.ZAdd(zsetOfClass("User", "AgeOfUserRNGRelation", relation.Key), redis.Z{Score: relation.Score, Member: relation.Value}).Err()
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

func (m *_AgeOfUserRNGRelationRedisMgr) ZSetDel(key string) error {
	return m.Del(setOfClass("User", "AgeOfUserRNGRelation", key)).Err()
}

func (m *_AgeOfUserRNGRelationRedisMgr) Range(key string, min, max int64) ([]string, error) {
	return m.ZRange(zsetOfClass("User", "AgeOfUserRNGRelation", key), min, max).Result()
}

func (m *_AgeOfUserRNGRelationRedisMgr) RevertRange(key string, min, max int64) ([]string, error) {
	return m.ZRevRange(zsetOfClass("User", "AgeOfUserRNGRelation", key), min, max).Result()
}

func (m *_AgeOfUserRNGRelationRedisMgr) Clear() error {
	strs, err := m.Keys(zsetOfClass("User", "AgeOfUserRNGRelation", "*")).Result()
	if err != nil {
		return err
	}
	return m.Del(strs...).Err()
}
