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
	Id          int32      `db:"id" json:"id"`
	Name        string     `db:"name" json:"name"`
	Mailbox     string     `db:"mailbox" json:"mailbox"`
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
		"`deleted_at`",
	}
	return columns
}

func (obj *User) GetPrimaryKey() PrimaryKey {
	pk := UserMgr.NewPrimaryKey()
	pk.Id = obj.Id
	return pk
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
		"id = ?",
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

type IdOfUserUK struct {
	Id int32
}

func (u *IdOfUserUK) Key() string {
	strs := []string{
		"Id",
		fmt.Sprint(u.Id),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *IdOfUserUK) SQLFormat(limit bool) string {
	conditions := []string{
		"id = ?",
	}
	return orm.SQLWhere(conditions)
}

func (u *IdOfUserUK) SQLParams() []interface{} {
	return []interface{}{
		u.Id,
	}
}

func (u *IdOfUserUK) SQLLimit() int {
	return 1
}

func (u *IdOfUserUK) Limit(n int) {
}

func (u *IdOfUserUK) Offset(n int) {
}

func (u *IdOfUserUK) UKRelation() UniqueRelation {
	return IdOfUserUKRelationRedisMgr()
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

func (m *_UserMySQLMgr) Search(where string, args ...interface{}) ([]*User, error) {
	obj := UserMgr.NewUser()
	if where != "" {
		where = " WHERE " + where
	}
	query := fmt.Sprintf("SELECT %s FROM `users` %s", strings.Join(obj.GetColumns(), ","), where)
	objs, err := m.FetchBySQL(query, args...)
	if err != nil {
		return nil, err
	}
	results := make([]*User, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*User))
	}
	return results, nil
}

func (m *_UserMySQLMgr) SearchCount(where string, args ...interface{}) (int64, error) {
	if where != "" {
		where = " WHERE " + where
	}
	return m.queryCount(where, args...)
}

func (m *_UserMySQLMgr) FetchBySQL(q string, args ...interface{}) (results []interface{}, err error) {
	rows, err := m.Query(q, args...)
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
			return nil, err
		}

		result.Description = Description.String

		result.HeadUrl = HeadUrl.String

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
		return nil, fmt.Errorf("User fetch result error: %v", err)
	}
	return
}
func (m *_UserMySQLMgr) Fetch(pk PrimaryKey) (*User, error) {
	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM `users` %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*User), nil
	}
	return nil, fmt.Errorf("User fetch record not found")
}

func (m *_UserMySQLMgr) FetchByPrimaryKeys(pks []PrimaryKey) ([]*User, error) {
	params := make([]string, 0, len(pks))
	for _, pk := range pks {
		params = append(params, fmt.Sprint(pk.(*IdOfUserPK).Id))
	}
	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM `users` WHERE `Id` IN (%s)", strings.Join(obj.GetColumns(), ","), strings.Join(params, ","))
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

func (m *_UserMySQLMgr) FindOne(unique Unique) (PrimaryKey, error) {
	objs, err := m.queryLimit(unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("User find record not found")
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

func (m *_UserMySQLMgr) Find(index Index) ([]PrimaryKey, error) {
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

func (m *_UserMySQLMgr) Range(scope Range) ([]PrimaryKey, error) {
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

func (m *_UserMySQLMgr) RangeRevert(scope Range) ([]PrimaryKey, error) {
	scope.Revert(true)
	return m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (m *_UserMySQLMgr) RangeRevertFetch(scope Range) ([]*User, error) {
	scope.Revert(true)
	return m.RangeFetch(scope)
}

func (m *_UserMySQLMgr) queryLimit(where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := UserMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM `users` %s", strings.Join(pk.Columns(), ","), where)
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

		result := UserMgr.NewPrimaryKey()
		err = rows.Scan(&(result.Id))
		if err != nil {
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

func (m *_UserMySQLMgr) BeginTx(tx *orm.MySQLTx) (*_UserMySQLTx, error) {
	ux := tx
	if ux == nil {
		tx, err := m.MySQLStore.BeginTx()
		if err != nil {
			return nil, err
		}
		ux = tx
	}
	return &_UserMySQLTx{ux, nil, 0}, nil
}

func (tx *_UserMySQLTx) BatchCreate(objs []*User) error {
	if len(objs) == 0 {
		return nil
	}

	params := make([]string, 0, len(objs))
	values := make([]interface{}, 0, len(objs)*14)
	for _, obj := range objs {
		params = append(params, fmt.Sprintf("(%s)", strings.Join(orm.NewStringSlice(14, "?"), ",")))
		values = append(values, obj.Id)
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
		values = append(values, obj.CreatedAt.Unix())
		values = append(values, obj.UpdatedAt.Unix())
		if obj.DeletedAt == nil {
			values = append(values, nil)
		} else {
			values = append(values, obj.DeletedAt.Unix())
		}
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
	for _, obj := range objs {
		if err := tx.Delete(obj); err != nil {
			return err
		}
	}
	return nil
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
	result, err := tx.Exec(query, args...)
	if err != nil {
		tx.err = err
		return err
	}
	tx.rowsAffected, tx.err = result.RowsAffected()
	return tx.err
}

func (tx *_UserMySQLTx) Create(obj *User) error {
	params := orm.NewStringSlice(14, "?")
	q := fmt.Sprintf("INSERT INTO `users`(%s) VALUES(%s)",
		strings.Join(obj.GetColumns(), ","),
		strings.Join(params, ","))

	values := make([]interface{}, 0, 14)
	values = append(values, obj.Id)
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
	values = append(values, obj.CreatedAt.Unix())
	values = append(values, obj.UpdatedAt.Unix())
	if obj.DeletedAt == nil {
		values = append(values, nil)
	} else {
		values = append(values, obj.DeletedAt.Unix())
	}
	result, err := tx.Exec(q, values...)
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
		"`deleted_at` = ?",
	}

	pk := obj.GetPrimaryKey()
	q := fmt.Sprintf("UPDATE `users` SET %s %s", strings.Join(columns, ","), pk.SQLFormat())
	values := make([]interface{}, 0, 14-1)
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
	values = append(values, obj.CreatedAt.Unix())
	values = append(values, obj.UpdatedAt.Unix())
	if obj.DeletedAt == nil {
		values = append(values, nil)
	} else {
		values = append(values, obj.DeletedAt.Unix())
	}
	values = append(values, pk.SQLParams()...)

	result, err := tx.Exec(q, values...)
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
	pk := obj.GetPrimaryKey()
	return tx.DeleteByPrimaryKey(pk)
}

func (tx *_UserMySQLTx) DeleteByPrimaryKey(pk PrimaryKey) error {
	q := fmt.Sprintf("DELETE FROM `users` %s", pk.SQLFormat())
	result, err := tx.Exec(q, pk.SQLParams()...)
	if err != nil {
		tx.err = err
		return err
	}
	tx.rowsAffected, tx.err = result.RowsAffected()
	return tx.err
}

func (tx *_UserMySQLTx) DeleteBySQL(where string, args ...interface{}) error {
	query := fmt.Sprintf("DELETE FROM `users`")
	if where != "" {
		query = fmt.Sprintf("DELETE FROM `users` WHERE %s", where)
	}
	result, err := tx.Exec(query, args...)
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
func (tx *_UserMySQLTx) FindOne(unique Unique) (PrimaryKey, error) {
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

func (tx *_UserMySQLTx) Find(index Index) ([]PrimaryKey, error) {
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

func (tx *_UserMySQLTx) Range(scope Range) ([]PrimaryKey, error) {
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

func (tx *_UserMySQLTx) RangeRevert(scope Range) ([]PrimaryKey, error) {
	scope.Revert(true)
	return tx.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (tx *_UserMySQLTx) RangeRevertFetch(scope Range) ([]*User, error) {
	scope.Revert(true)
	return tx.RangeFetch(scope)
}

func (tx *_UserMySQLTx) queryLimit(where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := UserMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM `users` %s", strings.Join(pk.Columns(), ","), where)
	rows, err := tx.Query(query, args...)
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
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("User query limit result error: %v", err)
	}
	return
}

func (tx *_UserMySQLTx) queryCount(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`id`) FROM `users` %s", where)

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

func (tx *_UserMySQLTx) Fetch(pk PrimaryKey) (*User, error) {
	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM `users` %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := tx.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*User), nil
	}
	return nil, fmt.Errorf("User fetch record not found")
}

func (tx *_UserMySQLTx) FetchByPrimaryKeys(pks []PrimaryKey) ([]*User, error) {
	params := make([]string, 0, len(pks))
	for _, pk := range pks {
		params = append(params, fmt.Sprint(pk.(*IdOfUserPK).Id))
	}
	obj := UserMgr.NewUser()
	query := fmt.Sprintf("SELECT %s FROM `users` WHERE `Id` IN (%s)", strings.Join(obj.GetColumns(), ","), strings.Join(params, ","))
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

func (tx *_UserMySQLTx) Search(where string, args ...interface{}) ([]*User, error) {
	obj := UserMgr.NewUser()
	if where != "" {
		where = " WHERE " + where
	}
	query := fmt.Sprintf("SELECT %s FROM `users` %s", strings.Join(obj.GetColumns(), ","), where)
	objs, err := tx.FetchBySQL(query, args...)
	if err != nil {
		return nil, err
	}
	results := make([]*User, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*User))
	}
	return results, nil
}

func (tx *_UserMySQLTx) SearchCount(where string, args ...interface{}) (int64, error) {
	if where != "" {
		where = " WHERE " + where
	}
	return tx.queryCount(where, args...)
}

func (tx *_UserMySQLTx) FetchBySQL(q string, args ...interface{}) (results []interface{}, err error) {
	rows, err := tx.Query(q, args...)
	if err != nil {
		tx.err = err
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
			return nil, err
		}

		result.Description = Description.String

		result.HeadUrl = HeadUrl.String

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
		tx.err = err
		return nil, fmt.Errorf("User fetch result error: %v", err)
	}
	return
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

	return m.AddBySQL(db, "SELECT `id`,`name`,`mailbox`,`sex`, `age`, `longitude`,`latitude`,`description`,`password`,`head_url`,`status`,`created_at`, `updated_at`, `deleted_at` FROM users")

}

func (m *_UserRedisMgr) AddBySQL(db DBFetcher, sql string, args ...interface{}) error {
	objs, err := db.FetchBySQL(sql, args...)
	if err != nil {
		return err
	}

	redisObjs := make([]*User, len(objs))
	for i, obj := range objs {
		redisObjs[i] = obj.(*User)
	}

	return m.SaveBatch(redisObjs)
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
func (m *_UserRedisMgr) FindOne(unique Unique) (PrimaryKey, error) {
	if relation := unique.UKRelation(); relation != nil {
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

func (m *_UserRedisMgr) Find(index Index) ([]PrimaryKey, error) {
	if relation := index.IDXRelation(); relation != nil {
		strs, err := relation.Find(index.Key())
		if err != nil {
			return nil, err
		}
		results := make([]PrimaryKey, 0, len(strs))
		for _, str := range strs {
			pk := UserMgr.NewPrimaryKey()
			if err := pk.Parse(str); err != nil {
				return nil, err
			}
			results = append(results, pk)
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
	return m.FetchByPrimaryKeys(vs)
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

func (m *_UserRedisMgr) Range(scope Range) ([]PrimaryKey, error) {
	if relation := scope.RNGRelation(); relation != nil {
		strs, err := relation.Range(scope.Key(), scope.Begin(), scope.End())
		if err != nil {
			return nil, err
		}
		results := make([]PrimaryKey, 0, len(strs))
		for _, str := range strs {
			pk := UserMgr.NewPrimaryKey()
			if err := pk.Parse(str); err != nil {
				return nil, err
			}
			results = append(results, pk)
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
	return m.FetchByPrimaryKeys(vs)
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

func (m *_UserRedisMgr) RangeRevert(scope Range) ([]PrimaryKey, error) {
	if relation := scope.RNGRelation(); relation != nil {
		scope.Revert(true)
		strs, err := relation.RangeRevert(scope.Key(), scope.Begin(), scope.End())
		if err != nil {
			return nil, err
		}
		results := make([]PrimaryKey, 0, len(strs))
		for _, str := range strs {
			pk := UserMgr.NewPrimaryKey()
			if err := pk.Parse(str); err != nil {
				return nil, err
			}
			results = append(results, pk)
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
	return m.FetchByPrimaryKeys(vs)
}

func (m *_UserRedisMgr) Fetch(pk PrimaryKey) (*User, error) {
	obj := UserMgr.NewUser()

	pipe := m.BeginPipeline()
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
	cmds, err := pipe.Exec()
	if err != nil {
		return nil, err
	}

	if b, err := cmds[0].(*redis.BoolCmd).Result(); err == nil {
		if !b {
			return nil, fmt.Errorf("User primary key:(%s) not exist", pk.Key())
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
	var val11 int64
	if err := m.StringScan(strs[11].(string), &val11); err != nil {
		return nil, err
	}
	obj.CreatedAt = time.Unix(val11, 0)
	var val12 int64
	if err := m.StringScan(strs[12].(string), &val12); err != nil {
		return nil, err
	}
	obj.UpdatedAt = time.Unix(val12, 0)
	if strs[13].(string) == "nil" {
		obj.DeletedAt = nil
	} else {
		var val13 int64
		if err := m.StringScan(strs[13].(string), &val13); err != nil {
			return nil, err
		}
		DeletedAtValue := time.Unix(val13, 0)
		obj.DeletedAt = &DeletedAtValue
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
	for i := 0; i < len(pks); i++ {
		if b, err := cmds[2*i].(*redis.BoolCmd).Result(); err == nil {
			if !b {
				return nil, fmt.Errorf("User primary key:(%s) not exist", pks[i].Key())
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
		var val11 int64
		if err := m.StringScan(strs[11].(string), &val11); err != nil {
			return nil, err
		}
		obj.CreatedAt = time.Unix(val11, 0)
		var val12 int64
		if err := m.StringScan(strs[12].(string), &val12); err != nil {
			return nil, err
		}
		obj.UpdatedAt = time.Unix(val12, 0)
		if strs[13].(string) == "nil" {
			obj.DeletedAt = nil
		} else {
			var val13 int64
			if err := m.StringScan(strs[13].(string), &val13); err != nil {
				return nil, err
			}
			DeletedAtValue := time.Unix(val13, 0)
			obj.DeletedAt = &DeletedAtValue
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
	uk_key_1 := []string{
		"Id",
		fmt.Sprint(obj.Id),
	}
	uk_pip_1 := IdOfUserUKRelationRedisMgr().BeginPipeline(pipe.Pipeline)
	if err := uk_pip_1.PairRem(strings.Join(uk_key_1, ":")); err != nil {
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
	if len(objs) > 0 {
		pipe := m.BeginPipeline()
		for _, obj := range objs {
			m.addToPipeline(pipe, obj)
		}
		if _, err := pipe.Exec(); err != nil {
			return err
		}
	}
	return nil
}

func (m *_UserRedisMgr) Save(obj *User) error {
	if obj != nil {
		pipe := m.BeginPipeline()
		m.addToPipeline(pipe, obj)
		if _, err := pipe.Exec(); err != nil {
			return err
		}
	}
	return nil
}

func (m *_UserRedisMgr) addToPipeline(pipe *_UserRedisPipeline, obj *User) error {
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
	pipe.HSet(keyOfObject(obj, pk.Key()), "HeadUrl", fmt.Sprint(obj.HeadUrl))
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
	uk_key_1 := []string{
		"Id",
		fmt.Sprint(obj.Id),
	}
	uk_pip_1 := IdOfUserUKRelationRedisMgr().BeginPipeline(pipe.Pipeline)
	uk_rel_1 := IdOfUserUKRelationRedisMgr().NewIdOfUserUKRelation(strings.Join(uk_key_1, ":"))
	uk_rel_1.Value = pk.Key()
	if err := uk_pip_1.PairAdd(uk_rel_1); err != nil {
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

//! relation
type IdOfUserUKRelation struct {
	Key   string `db:"key" json:"key"`
	Value string `db:"value" json:"value"`
}

func (relation *IdOfUserUKRelation) GetClassName() string {
	return "IdOfUserUKRelation"
}

func (relation *IdOfUserUKRelation) GetIndexes() []string {
	idx := []string{}
	return idx
}

func (relation *IdOfUserUKRelation) GetStoreType() string {
	return "pair"
}

type _IdOfUserUKRelationRedisMgr struct {
	*orm.RedisStore
}

func IdOfUserUKRelationRedisMgr(stores ...*orm.RedisStore) *_IdOfUserUKRelationRedisMgr {
	if len(stores) > 0 {
		return &_IdOfUserUKRelationRedisMgr{stores[0]}
	}
	return &_IdOfUserUKRelationRedisMgr{_redis_store}
}

func (m *_IdOfUserUKRelationRedisMgr) NewIdOfUserUKRelation(key string) *IdOfUserUKRelation {
	return &IdOfUserUKRelation{
		Key: key,
	}
}

//! pipeline
type _IdOfUserUKRelationRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_IdOfUserUKRelationRedisMgr) BeginPipeline(pipes ...*redis.Pipeline) *_IdOfUserUKRelationRedisPipeline {
	if len(pipes) > 0 {
		return &_IdOfUserUKRelationRedisPipeline{pipes[0], nil}
	}
	return &_IdOfUserUKRelationRedisPipeline{m.Pipeline(), nil}
}

//! redis relation pair
func (m *_IdOfUserUKRelationRedisMgr) PairAdd(obj *IdOfUserUKRelation) error {
	return m.Set(pairOfClass("User", obj.GetClassName(), obj.Key), obj.Value, 0).Err()
}

func (pipe *_IdOfUserUKRelationRedisPipeline) PairAdd(obj *IdOfUserUKRelation) error {
	return pipe.Set(pairOfClass("User", obj.GetClassName(), obj.Key), obj.Value, 0).Err()
}

func (m *_IdOfUserUKRelationRedisMgr) PairGet(key string) (*IdOfUserUKRelation, error) {
	str, err := m.Get(pairOfClass("User", "IdOfUserUKRelation", key)).Result()
	if err != nil {
		return nil, err
	}

	obj := m.NewIdOfUserUKRelation(key)
	if err := m.StringScan(str, &obj.Value); err != nil {
		return nil, err
	}
	return obj, nil
}

func (m *_IdOfUserUKRelationRedisMgr) PairRem(key string) error {
	return m.Del(pairOfClass("User", "IdOfUserUKRelation", key)).Err()
}

func (pipe *_IdOfUserUKRelationRedisPipeline) PairRem(key string) error {
	return pipe.Del(pairOfClass("User", "IdOfUserUKRelation", key)).Err()
}

func (m *_IdOfUserUKRelationRedisMgr) FindOne(key string) (string, error) {
	return m.Get(pairOfClass("User", "IdOfUserUKRelation", key)).Result()
}

func (m *_IdOfUserUKRelationRedisMgr) Clear() error {
	strs, err := m.Keys(pairOfClass("User", "IdOfUserUKRelation", "*")).Result()
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
