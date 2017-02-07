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
func (m *_UserBaseInfoMgr) MySQL() *ReferenceResult {
	return NewReferenceResult(UserBaseInfoMySQLMgr())
}

type _UserBaseInfoMySQLMgr struct {
	*orm.MySQLStore
}

func UserBaseInfoMySQLMgr() *_UserBaseInfoMySQLMgr {
	return &_UserBaseInfoMySQLMgr{_mysql_store}
}

func NewUserBaseInfoMySQLMgr(cf *MySQLConfig) (*_UserBaseInfoMySQLMgr, error) {
	store, err := orm.NewMySQLStore(cf.Host, cf.Port, cf.Database, cf.UserName, cf.Password)
	if err != nil {
		return nil, err
	}
	return &_UserBaseInfoMySQLMgr{store}, nil
}

func (m *_UserBaseInfoMySQLMgr) Search(where string, args ...interface{}) (results []interface{}, err error) {
	obj := UserBaseInfoMgr.NewUserBaseInfo()
	if where != "" {
		where = " WHERE " + where
	}
	sql := fmt.Sprintf("SELECT %s FROM `user_base_info` %s", strings.Join(obj.GetColumns(), ","), where)
	return m.FetchBySQL(sql, args...)
}

func (m *_UserBaseInfoMySQLMgr) SearchCount(where string, args ...interface{}) (int64, error) {
	if where != "" {
		where = " WHERE " + where
	}
	return m.queryCount(where, args...)
}

func (m *_UserBaseInfoMySQLMgr) FetchBySQL(q string, args ...interface{}) (results []interface{}, err error) {
	rows, err := m.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("UserBaseInfo fetch error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var result UserBaseInfo
		err = rows.Scan(&(result.Id),
			&(result.Name),
			&(result.Mailbox),
			&(result.Password),
			&(result.Sex),
		)
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
func (m *_UserBaseInfoMySQLMgr) Fetch(id interface{}) (*UserBaseInfo, error) {
	obj := UserBaseInfoMgr.NewUserBaseInfo()
	query := fmt.Sprintf("SELECT %s FROM `user_base_info` WHERE `Id` = (%s)", strings.Join(obj.GetColumns(), ","), id)
	objs, err := m.FetchBySQL(query)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0].(*UserBaseInfo), nil
	}
	return nil, fmt.Errorf("UserBaseInfo fetch record not found")
}

func (m *_UserBaseInfoMySQLMgr) FetchByIds(ids []interface{}) ([]*UserBaseInfo, error) {
	if len(ids) == 0 {
		return []*UserBaseInfo{}, nil
	}

	obj := UserBaseInfoMgr.NewUserBaseInfo()
	query := fmt.Sprintf("SELECT %s FROM `user_base_info` WHERE `Id` IN (%s)", strings.Join(obj.GetColumns(), ","), orm.SliceJoin(ids, ","))
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

func (m *_UserBaseInfoMySQLMgr) FindOne(unique Unique) (interface{}, error) {
	objs, err := m.queryLimit(unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return "", err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return "", fmt.Errorf("UserBaseInfo find record not found")
}

func (m *_UserBaseInfoMySQLMgr) FindOneFetch(unique Unique) (*UserBaseInfo, error) {
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

func (m *_UserBaseInfoMySQLMgr) Find(index Index) ([]interface{}, error) {
	return m.queryLimit(index.SQLFormat(true), index.SQLLimit(), index.SQLParams()...)
}

func (m *_UserBaseInfoMySQLMgr) FindFetch(index Index) ([]*UserBaseInfo, error) {
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

func (m *_UserBaseInfoMySQLMgr) FindCount(index Index) (int64, error) {
	return m.queryCount(index.SQLFormat(false), index.SQLParams()...)
}

func (m *_UserBaseInfoMySQLMgr) Range(scope Range) ([]interface{}, error) {
	return m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (m *_UserBaseInfoMySQLMgr) RangeFetch(scope Range) ([]*UserBaseInfo, error) {
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

func (m *_UserBaseInfoMySQLMgr) RangeCount(scope Range) (int64, error) {
	return m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
}

func (m *_UserBaseInfoMySQLMgr) RangeRevert(scope Range) ([]interface{}, error) {
	scope.Revert(true)
	return m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
}

func (m *_UserBaseInfoMySQLMgr) RangeRevertFetch(scope Range) ([]*UserBaseInfo, error) {
	scope.Revert(true)
	return m.RangeFetch(scope)
}

func (m *_UserBaseInfoMySQLMgr) queryLimit(where string, limit int, args ...interface{}) (results []interface{}, err error) {
	query := fmt.Sprintf("SELECT `id` FROM `user_base_info` %s", where)
	rows, err := m.Query(query, args...)
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

		var result int32
		if err = rows.Scan(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("UserBaseInfo query limit result error: %v", err)
	}
	return
}

func (m *_UserBaseInfoMySQLMgr) queryCount(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`id`) FROM `user_base_info` %s", where)
	rows, err := m.Query(query, args...)
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
