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

type UserBaseInfo struct {
	Id       int32  `db:"id"`
	Name     string `db:"name"`
	Mailbox  string `db:"mailbox"`
	Password string `db:"password"`
	Sex      bool   `db:"sex"`
}

type _UserBaseInfoMgr struct {
}

var UserBaseInfoMgr *_UserBaseInfoMgr

func (m *_UserBaseInfoMgr) NewUserBaseInfo() *UserBaseInfo {
	return &UserBaseInfo{}
}

//! object function

func (obj *UserBaseInfo) GetNameSpace() string {
	return "model"
}

func (obj *UserBaseInfo) GetClassName() string {
	return "UserBaseInfo"
}

func (obj *UserBaseInfo) GetTableName() string {
	return ""
}

func (obj *UserBaseInfo) GetColumns() []string {
	columns := []string{
		"`id`",
		"`name`",
		"`mailbox`",
		"`password`",
		"`sex`",
	}
	return columns
}

func (obj *UserBaseInfo) GetPrimaryKey() PrimaryKey {
	pk := UserBaseInfoMgr.NewPrimaryKey()
	pk.Id = obj.Id
	return pk
}

func (obj *UserBaseInfo) Validate() error {
	validate := validator.New()
	return validate.Struct(obj)
}

//! primary key

type IdOfUserBaseInfoPK struct {
	Id int32
}

func (m *_UserBaseInfoMgr) NewPrimaryKey() *IdOfUserBaseInfoPK {
	return &IdOfUserBaseInfoPK{}
}

func (u *IdOfUserBaseInfoPK) Key() string {
	strs := []string{
		"Id",
		fmt.Sprint(u.Id),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *IdOfUserBaseInfoPK) Parse(key string) error {
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

func (u *IdOfUserBaseInfoPK) SQLFormat() string {
	conditions := []string{
		"id = ?",
	}
	return orm.SQLWhere(conditions)
}

func (u *IdOfUserBaseInfoPK) SQLParams() []interface{} {
	return []interface{}{
		u.Id,
	}
}

func (u *IdOfUserBaseInfoPK) Columns() []string {
	return []string{
		"`id`",
	}
}

//! uniques

type MailboxPasswordOfUserBaseInfoUK struct {
	Mailbox  string
	Password string
}

func (u *MailboxPasswordOfUserBaseInfoUK) Key() string {
	strs := []string{
		"Mailbox",
		fmt.Sprint(u.Mailbox),
		"Password",
		fmt.Sprint(u.Password),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *MailboxPasswordOfUserBaseInfoUK) SQLFormat(limit bool) string {
	conditions := []string{
		"mailbox = ?",
		"password = ?",
	}
	return orm.SQLWhere(conditions)
}

func (u *MailboxPasswordOfUserBaseInfoUK) SQLParams() []interface{} {
	return []interface{}{
		u.Mailbox,
		u.Password,
	}
}

func (u *MailboxPasswordOfUserBaseInfoUK) SQLLimit() int {
	return 1
}

func (u *MailboxPasswordOfUserBaseInfoUK) Limit(n int) {
}

func (u *MailboxPasswordOfUserBaseInfoUK) Offset(n int) {
}

func (u *MailboxPasswordOfUserBaseInfoUK) UKRelation() UniqueRelation {
	return nil
}

//! indexes

type NameOfUserBaseInfoIDX struct {
	Name   string
	offset int
	limit  int
}

func (u *NameOfUserBaseInfoIDX) Key() string {
	strs := []string{
		"Name",
		fmt.Sprint(u.Name),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *NameOfUserBaseInfoIDX) SQLFormat(limit bool) string {
	conditions := []string{
		"name = ?",
	}
	if limit {
		return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return orm.SQLWhere(conditions)
}

func (u *NameOfUserBaseInfoIDX) SQLParams() []interface{} {
	return []interface{}{
		u.Name,
	}
}

func (u *NameOfUserBaseInfoIDX) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *NameOfUserBaseInfoIDX) Limit(n int) {
	u.limit = n
}

func (u *NameOfUserBaseInfoIDX) Offset(n int) {
	u.offset = n
}

func (u *NameOfUserBaseInfoIDX) IDXRelation() IndexRelation {
	return nil
}

//! ranges

type IdOfUserBaseInfoRNG struct {
	IdBegin      int64
	IdEnd        int64
	offset       int
	limit        int
	includeBegin bool
	includeEnd   bool
	revert       bool
}

func (u *IdOfUserBaseInfoRNG) Key() string {
	strs := []string{
		"Id",
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *IdOfUserBaseInfoRNG) beginOp() string {
	if u.includeBegin {
		return ">="
	}
	return ">"
}
func (u *IdOfUserBaseInfoRNG) endOp() string {
	if u.includeBegin {
		return "<="
	}
	return "<"
}

func (u *IdOfUserBaseInfoRNG) SQLFormat(limit bool) string {
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

func (u *IdOfUserBaseInfoRNG) SQLParams() []interface{} {
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

func (u *IdOfUserBaseInfoRNG) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *IdOfUserBaseInfoRNG) Limit(n int) {
	u.limit = n
}

func (u *IdOfUserBaseInfoRNG) Offset(n int) {
	u.offset = n
}

func (u *IdOfUserBaseInfoRNG) Begin() int64 {
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

func (u *IdOfUserBaseInfoRNG) End() int64 {
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

func (u *IdOfUserBaseInfoRNG) Revert(b bool) {
	u.revert = b
}

func (u *IdOfUserBaseInfoRNG) IncludeBegin(f bool) {
	u.includeBegin = f
}

func (u *IdOfUserBaseInfoRNG) IncludeEnd(f bool) {
	u.includeEnd = f
}

func (u *IdOfUserBaseInfoRNG) RNGRelation() RangeRelation {
	return nil
}

type _UserBaseInfoDBMgr struct {
	db orm.DB
}

func (m *_UserBaseInfoMgr) DB(db orm.DB) *_UserBaseInfoDBMgr {
	return UserBaseInfoDBMgr(db)
}

func UserBaseInfoDBMgr(db orm.DB) *_UserBaseInfoDBMgr {
	if db == nil {
		panic(fmt.Errorf("UserBaseInfoDBMgr init need db"))
	}
	return &_UserBaseInfoDBMgr{db: db}
}

func (m *_UserBaseInfoDBMgr) Search(where string, orderby string, limit string, args ...interface{}) ([]*UserBaseInfo, error) {
	obj := UserBaseInfoMgr.NewUserBaseInfo()
	conditions := []string{where, orderby, limit}
	query := fmt.Sprintf("SELECT %s FROM `user_base_info` %s", strings.Join(obj.GetColumns(), ","), strings.Join(conditions, " "))
	objs, err := m.FetchBySQL(query, args...)
	if err != nil {
		return nil, err
	}
	results := make([]*UserBaseInfo, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*UserBaseInfo))
	}
	return results, nil
}

func (m *_UserBaseInfoDBMgr) SearchConditions(conditions []string, orderby string, offset int, limit int, args ...interface{}) ([]*UserBaseInfo, error) {
	obj := UserBaseInfoMgr.NewUserBaseInfo()
	q := fmt.Sprintf("SELECT %s FROM `user_base_info` %s %s %s",
		strings.Join(obj.GetColumns(), ","),
		orm.SQLWhere(conditions),
		orderby,
		orm.SQLOffsetLimit(offset, limit))
	objs, err := m.FetchBySQL(q, args...)
	if err != nil {
		return nil, err
	}
	results := make([]*UserBaseInfo, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*UserBaseInfo))
	}
	return results, nil
}

func (m *_UserBaseInfoDBMgr) SearchCount(where string, args ...interface{}) (int64, error) {
	return m.queryCount(where, args...)
}

func (m *_UserBaseInfoDBMgr) SearchConditionsCount(conditions []string, args ...interface{}) (int64, error) {
	return m.queryCount(orm.SQLWhere(conditions), args...)
}

func (m *_UserBaseInfoDBMgr) FetchBySQL(q string, args ...interface{}) (results []interface{}, err error) {
	rows, err := m.db.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("UserBaseInfo fetch error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var result UserBaseInfo
		err = rows.Scan(&(result.Id), &(result.Name), &(result.Mailbox), &(result.Password), &(result.Sex))
		if err != nil {
			return nil, err
		}

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("UserBaseInfo fetch result error: %v", err)
	}
	return
}
func (m *_UserBaseInfoDBMgr) Fetch(pk PrimaryKey) (*UserBaseInfo, error) {
	obj := UserBaseInfoMgr.NewUserBaseInfo()
	query := fmt.Sprintf("SELECT %s FROM `user_base_info` %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*UserBaseInfo), nil
	}
	return nil, fmt.Errorf("UserBaseInfo fetch record not found")
}

func (m *_UserBaseInfoDBMgr) FetchByPrimaryKeys(pks []PrimaryKey) ([]*UserBaseInfo, error) {
	params := make([]string, 0, len(pks))
	for _, pk := range pks {
		params = append(params, fmt.Sprint(pk.(*IdOfUserBaseInfoPK).Id))
	}
	obj := UserBaseInfoMgr.NewUserBaseInfo()
	query := fmt.Sprintf("SELECT %s FROM `user_base_info` WHERE `Id` IN (%s)", strings.Join(obj.GetColumns(), ","), strings.Join(params, ","))
	objs, err := m.FetchBySQL(query)
	if err != nil {
		return nil, err
	}
	results := make([]*UserBaseInfo, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*UserBaseInfo))
	}
	return results, nil
}

func (m *_UserBaseInfoDBMgr) FindOne(unique Unique) (PrimaryKey, error) {
	objs, err := m.queryLimit(unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("UserBaseInfo find record not found")
}

func (m *_UserBaseInfoDBMgr) FindOneFetch(unique Unique) (*UserBaseInfo, error) {
	obj := UserBaseInfoMgr.NewUserBaseInfo()
	query := fmt.Sprintf("SELECT %s FROM `user_base_info` %s", strings.Join(obj.GetColumns(), ","), unique.SQLFormat(true))
	objs, err := m.FetchBySQL(query, unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*UserBaseInfo), nil
	}
	return nil, fmt.Errorf("none record")
}

func (m *_UserBaseInfoDBMgr) Find(index Index) ([]PrimaryKey, error) {
	return m.queryLimit(index.SQLFormat(true), index.SQLLimit(), index.SQLParams()...)
}

func (m *_UserBaseInfoDBMgr) FindFetch(index Index) ([]*UserBaseInfo, error) {
	obj := UserBaseInfoMgr.NewUserBaseInfo()
	query := fmt.Sprintf("SELECT %s FROM `user_base_info` %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	objs, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return nil, err
	}
	results := make([]*UserBaseInfo, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*UserBaseInfo))
	}
	return results, nil
}

func (m *_UserBaseInfoDBMgr) FindCount(index Index) (int64, error) {
	return m.queryCount(index.SQLFormat(false), index.SQLParams()...)
}

func (m *_UserBaseInfoDBMgr) Range(scope Range) ([]PrimaryKey, error) {
	return m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (m *_UserBaseInfoDBMgr) RangeFetch(scope Range) ([]*UserBaseInfo, error) {
	obj := UserBaseInfoMgr.NewUserBaseInfo()
	query := fmt.Sprintf("SELECT %s FROM `user_base_info` %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	objs, err := m.FetchBySQL(query, scope.SQLParams()...)
	if err != nil {
		return nil, err
	}
	results := make([]*UserBaseInfo, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*UserBaseInfo))
	}
	return results, nil
}

func (m *_UserBaseInfoDBMgr) RangeCount(scope Range) (int64, error) {
	return m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
}

func (m *_UserBaseInfoDBMgr) RangeRevert(scope Range) ([]PrimaryKey, error) {
	scope.Revert(true)
	return m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (m *_UserBaseInfoDBMgr) RangeRevertFetch(scope Range) ([]*UserBaseInfo, error) {
	scope.Revert(true)
	return m.RangeFetch(scope)
}

func (m *_UserBaseInfoDBMgr) queryLimit(where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := UserBaseInfoMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM `user_base_info` %s", strings.Join(pk.Columns(), ","), where)
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("UserBaseInfo query limit error: %v", err)
	}
	defer rows.Close()

	offset := 0

	for rows.Next() {
		if limit >= 0 && offset >= limit {
			break
		}
		offset++

		result := UserBaseInfoMgr.NewPrimaryKey()
		err = rows.Scan(&(result.Id))
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("UserBaseInfo query limit result error: %v", err)
	}
	return
}

func (m *_UserBaseInfoDBMgr) queryCount(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`id`) FROM `user_base_info` %s", where)
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return 0, fmt.Errorf("UserBaseInfo query count error: %v", err)
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
