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

type Office struct {
	OfficeId             int32     `db:"office_id"`
	OfficeArea           string    `db:"office_area"`
	OfficeName           string    `db:"office_name"`
	SearchOriginCode     string    `db:"search_origin_code"`
	ProcessingOriginCode string    `db:"processing_origin_code"`
	CreateBy             string    `db:"create_by"`
	UpdateBy             string    `db:"update_by"`
	CreateDate           time.Time `db:"create_date"`
	UpdateDate           time.Time `db:"update_date"`
}

var OfficeColumns = struct {
	OfficeId             string
	OfficeArea           string
	OfficeName           string
	SearchOriginCode     string
	ProcessingOriginCode string
	CreateBy             string
	UpdateBy             string
	CreateDate           string
	UpdateDate           string
}{
	"office_id",
	"office_area",
	"office_name",
	"search_origin_code",
	"processing_origin_code",
	"create_by",
	"update_by",
	"create_date",
	"update_date",
}

type _OfficeMgr struct {
}

var OfficeMgr *_OfficeMgr

func (m *_OfficeMgr) NewOffice() *Office {
	return &Office{}
}

//! object function

func (obj *Office) GetNameSpace() string {
	return "model"
}

func (obj *Office) GetClassName() string {
	return "Office"
}

func (obj *Office) GetTableName() string {
	return "testCRUD"
}

func (obj *Office) GetColumns() []string {
	columns := []string{
		"testCRUD.office_id",
		"testCRUD.office_area",
		"testCRUD.office_name",
		"testCRUD.search_origin_code",
		"testCRUD.processing_origin_code",
		"testCRUD.create_by",
		"testCRUD.update_by",
		"testCRUD.create_date",
		"testCRUD.update_date",
	}
	return columns
}

func (obj *Office) GetNoneIncrementColumns() []string {
	columns := []string{
		"office_area",
		"office_name",
		"search_origin_code",
		"processing_origin_code",
		"create_by",
		"update_by",
		"create_date",
		"update_date",
	}
	return columns
}

func (obj *Office) GetPrimaryKey() PrimaryKey {
	pk := OfficeMgr.NewPrimaryKey()
	pk.OfficeId = obj.OfficeId
	return pk
}

func (obj *Office) Validate() error {
	validate := validator.New()
	return validate.Struct(obj)
}

//! primary key

type OfficeIdOfOfficePK struct {
	OfficeId int32
}

func (m *_OfficeMgr) NewPrimaryKey() *OfficeIdOfOfficePK {
	return &OfficeIdOfOfficePK{}
}

func (u *OfficeIdOfOfficePK) Key() string {
	strs := []string{
		"OfficeId",
		fmt.Sprint(u.OfficeId),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *OfficeIdOfOfficePK) Parse(key string) error {
	arr := strings.Split(key, ":")
	if len(arr)%2 != 0 {
		return fmt.Errorf("key (%s) format error", key)
	}
	kv := map[string]string{}
	for i := 0; i < len(arr)/2; i++ {
		kv[arr[2*i]] = arr[2*i+1]
	}
	vOfficeId, ok := kv["OfficeId"]
	if !ok {
		return fmt.Errorf("key (%s) without (OfficeId) field", key)
	}
	if err := orm.StringScan(vOfficeId, &(u.OfficeId)); err != nil {
		return err
	}
	return nil
}

func (u *OfficeIdOfOfficePK) SQLFormat() string {
	conditions := []string{
		"office_id = ?",
	}
	return orm.SQLWhere(conditions)
}

func (u *OfficeIdOfOfficePK) SQLParams() []interface{} {
	return []interface{}{
		u.OfficeId,
	}
}

func (u *OfficeIdOfOfficePK) Columns() []string {
	return []string{
		"office_id",
	}
}

//! uniques

//! indexes

//! ranges

type OfficeIdOfOfficeRNG struct {
	OfficeIdBegin int64
	OfficeIdEnd   int64
	offset        int
	limit         int
	includeBegin  bool
	includeEnd    bool
	revert        bool
}

func (u *OfficeIdOfOfficeRNG) Key() string {
	strs := []string{
		"OfficeId",
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *OfficeIdOfOfficeRNG) beginOp() string {
	if u.includeBegin {
		return ">="
	}
	return ">"
}
func (u *OfficeIdOfOfficeRNG) endOp() string {
	if u.includeBegin {
		return "<="
	}
	return "<"
}

func (u *OfficeIdOfOfficeRNG) SQLFormat(limit bool) string {
	conditions := []string{}
	if u.OfficeIdBegin != u.OfficeIdEnd {
		if u.OfficeIdBegin != -1 {
			conditions = append(conditions, fmt.Sprintf("office_id %s ?", u.beginOp()))
		}
		if u.OfficeIdEnd != -1 {
			conditions = append(conditions, fmt.Sprintf("office_id %s ?", u.endOp()))
		}
	}
	if limit {
		return fmt.Sprintf("%s %s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("office_id", u.revert), orm.MsSQLOffsetLimit(u.offset, u.limit))
	}
	return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("office_id", u.revert))
}

func (u *OfficeIdOfOfficeRNG) SQLParams() []interface{} {
	params := []interface{}{}
	if u.OfficeIdBegin != u.OfficeIdEnd {
		if u.OfficeIdBegin != -1 {
			params = append(params, u.OfficeIdBegin)
		}
		if u.OfficeIdEnd != -1 {
			params = append(params, u.OfficeIdEnd)
		}
	}
	return params
}

func (u *OfficeIdOfOfficeRNG) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *OfficeIdOfOfficeRNG) Limit(n int) {
	u.limit = n
}

func (u *OfficeIdOfOfficeRNG) Offset(n int) {
	u.offset = n
}

func (u *OfficeIdOfOfficeRNG) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *OfficeIdOfOfficeRNG) Begin() int64 {
	start := u.OfficeIdBegin
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

func (u *OfficeIdOfOfficeRNG) End() int64 {
	stop := u.OfficeIdEnd
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

func (u *OfficeIdOfOfficeRNG) Revert(b bool) {
	u.revert = b
}

func (u *OfficeIdOfOfficeRNG) IncludeBegin(f bool) {
	u.includeBegin = f
}

func (u *OfficeIdOfOfficeRNG) IncludeEnd(f bool) {
	u.includeEnd = f
}

func (u *OfficeIdOfOfficeRNG) RNGRelation(store *orm.RedisStore) RangeRelation {
	return nil
}

type _OfficeDBMgr struct {
	db orm.DB
}

func (m *_OfficeMgr) DB(db orm.DB) *_OfficeDBMgr {
	return OfficeDBMgr(db)
}

func OfficeDBMgr(db orm.DB) *_OfficeDBMgr {
	if db == nil {
		panic(fmt.Errorf("OfficeDBMgr init need db"))
	}
	return &_OfficeDBMgr{db: db}
}

func (m *_OfficeDBMgr) Search(where string, orderby string, limit string, args ...interface{}) ([]*Office, error) {
	obj := OfficeMgr.NewOffice()

	if limit = strings.ToUpper(strings.TrimSpace(limit)); limit != "" && !strings.HasPrefix(limit, "LIMIT") {
		limit = "LIMIT " + limit
	}

	conditions := []string{where, orderby, limit}
	query := fmt.Sprintf("SELECT %s FROM [dbo].[testCRUD] %s", strings.Join(obj.GetColumns(), ","), strings.Join(conditions, " "))
	return m.FetchBySQL(query, args...)
}

func (m *_OfficeDBMgr) SearchContext(ctx context.Context, where string, orderby string, limit string, args ...interface{}) ([]*Office, error) {
	obj := OfficeMgr.NewOffice()

	if limit = strings.ToUpper(strings.TrimSpace(limit)); limit != "" && !strings.HasPrefix(limit, "LIMIT") {
		limit = "LIMIT " + limit
	}

	conditions := []string{where, orderby, limit}
	query := fmt.Sprintf("SELECT %s FROM [dbo].[testCRUD] %s", strings.Join(obj.GetColumns(), ","), strings.Join(conditions, " "))
	return m.FetchBySQLContext(ctx, query, args...)
}

func (m *_OfficeDBMgr) SearchConditions(conditions []string, orderby string, offset int, limit int, args ...interface{}) ([]*Office, error) {
	obj := OfficeMgr.NewOffice()
	if orderby == "" {
		orderby = orm.SQLOrderBy("office_id", false)
	}
	q := fmt.Sprintf("SELECT %s FROM [dbo].[testCRUD] %s %s %s",
		strings.Join(obj.GetColumns(), ","),
		orm.SQLWhere(conditions),
		orderby,
		orm.MsSQLOffsetLimit(offset, limit))

	return m.FetchBySQL(q, args...)
}

func (m *_OfficeDBMgr) SearchConditionsContext(ctx context.Context, conditions []string, orderby string, offset int, limit int, args ...interface{}) ([]*Office, error) {
	obj := OfficeMgr.NewOffice()
	if orderby == "" {
		orderby = orm.SQLOrderBy("office_id", false)
	}
	q := fmt.Sprintf("SELECT %s FROM [dbo].[testCRUD] %s %s %s",
		strings.Join(obj.GetColumns(), ","),
		orm.SQLWhere(conditions),
		orderby,
		orm.MsSQLOffsetLimit(offset, limit))

	return m.FetchBySQLContext(ctx, q, args...)
}

func (m *_OfficeDBMgr) SearchCount(where string, args ...interface{}) (int64, error) {
	return m.queryCount(where, args...)
}

func (m *_OfficeDBMgr) SearchCountContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	return m.queryCountContext(ctx, where, args...)
}

func (m *_OfficeDBMgr) SearchConditionsCount(conditions []string, args ...interface{}) (int64, error) {
	return m.queryCount(orm.SQLWhere(conditions), args...)
}

func (m *_OfficeDBMgr) SearchConditionsCountContext(ctx context.Context, conditions []string, args ...interface{}) (int64, error) {
	return m.queryCountContext(ctx, orm.SQLWhere(conditions), args...)
}

func (m *_OfficeDBMgr) FetchBySQL(q string, args ...interface{}) (results []*Office, err error) {
	rows, err := m.db.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("Office fetch error: %v", err)
	}
	defer rows.Close()

	var CreateDate string
	var UpdateDate string

	for rows.Next() {
		var result Office
		err = rows.Scan(&(result.OfficeId), &(result.OfficeArea), &(result.OfficeName), &(result.SearchOriginCode), &(result.ProcessingOriginCode), &(result.CreateBy), &(result.UpdateBy), &CreateDate, &UpdateDate)
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		result.CreateDate = orm.MsSQLTimeParse(CreateDate)
		result.UpdateDate = orm.MsSQLTimeParse(UpdateDate)

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("Office fetch result error: %v", err)
	}
	return
}

func (m *_OfficeDBMgr) FetchBySQLContext(ctx context.Context, q string, args ...interface{}) (results []*Office, err error) {
	rows, err := m.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("Office fetch error: %v", err)
	}
	defer rows.Close()

	var CreateDate string
	var UpdateDate string

	for rows.Next() {
		var result Office
		err = rows.Scan(&(result.OfficeId), &(result.OfficeArea), &(result.OfficeName), &(result.SearchOriginCode), &(result.ProcessingOriginCode), &(result.CreateBy), &(result.UpdateBy), &CreateDate, &UpdateDate)
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		result.CreateDate = orm.MsSQLTimeParse(CreateDate)
		result.UpdateDate = orm.MsSQLTimeParse(UpdateDate)

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("Office fetch result error: %v", err)
	}
	return
}
func (m *_OfficeDBMgr) Exist(pk PrimaryKey) (bool, error) {
	c, err := m.queryCount(pk.SQLFormat(), pk.SQLParams()...)
	if err != nil {
		return false, err
	}
	return (c != 0), nil
}

// Deprecated: Use FetchByPrimaryKey instead.
func (m *_OfficeDBMgr) Fetch(pk PrimaryKey) (*Office, error) {
	obj := OfficeMgr.NewOffice()
	query := fmt.Sprintf("SELECT %s FROM [dbo].[testCRUD] %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("Office fetch record not found")
}

// err not found check
func (m *_OfficeDBMgr) IsErrNotFound(err error) bool {
	return strings.Contains(err.Error(), "not found") || err == sql.ErrNoRows
}

// primary key
func (m *_OfficeDBMgr) FetchByPrimaryKey(officeId int32) (*Office, error) {
	obj := OfficeMgr.NewOffice()
	pk := &OfficeIdOfOfficePK{
		OfficeId: officeId,
	}

	query := fmt.Sprintf("SELECT %s FROM [dbo].[testCRUD] %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("Office fetch record not found")
}

func (m *_OfficeDBMgr) FetchByPrimaryKeyContext(ctx context.Context, officeId int32) (*Office, error) {
	obj := OfficeMgr.NewOffice()
	pk := &OfficeIdOfOfficePK{
		OfficeId: officeId,
	}

	query := fmt.Sprintf("SELECT %s FROM [dbo].[testCRUD] %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQLContext(ctx, query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("Office fetch record not found")
}

func (m *_OfficeDBMgr) FetchByPrimaryKeys(officeIds []int32) ([]*Office, error) {
	size := len(officeIds)
	if size == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, size)
	for _, pk := range officeIds {
		params = append(params, pk)
	}
	obj := OfficeMgr.NewOffice()
	query := fmt.Sprintf("SELECT %s FROM [dbo].[testCRUD] WHERE office_id IN (?%s)", strings.Join(obj.GetColumns(), ","),
		strings.Repeat(",?", size-1))
	return m.FetchBySQL(query, params...)
}

func (m *_OfficeDBMgr) FetchByPrimaryKeysContext(ctx context.Context, officeIds []int32) ([]*Office, error) {
	size := len(officeIds)
	if size == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, size)
	for _, pk := range officeIds {
		params = append(params, pk)
	}
	obj := OfficeMgr.NewOffice()
	query := fmt.Sprintf("SELECT %s FROM [dbo].[testCRUD] WHERE office_id IN (?%s)", strings.Join(obj.GetColumns(), ","),
		strings.Repeat(",?", size-1))
	return m.FetchBySQLContext(ctx, query, params...)
}

// indexes

// uniques

func (m *_OfficeDBMgr) FindOne(unique Unique) (PrimaryKey, error) {
	objs, err := m.queryLimit(unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("Office find record not found")
}

func (m *_OfficeDBMgr) FindOneContext(ctx context.Context, unique Unique) (PrimaryKey, error) {
	objs, err := m.queryLimitContext(ctx, unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("Office find record not found")
}

// Deprecated: Use FetchByXXXUnique instead.
func (m *_OfficeDBMgr) FindOneFetch(unique Unique) (*Office, error) {
	obj := OfficeMgr.NewOffice()
	query := fmt.Sprintf("SELECT %s FROM [dbo].[testCRUD] %s", strings.Join(obj.GetColumns(), ","), unique.SQLFormat(true))
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
func (m *_OfficeDBMgr) Find(index Index) (int64, []PrimaryKey, error) {
	total, err := m.queryCount(index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimit(index.SQLFormat(true), index.SQLLimit(), index.SQLParams()...)
	return total, pks, err
}

func (m *_OfficeDBMgr) FindFetch(index Index) (int64, []*Office, error) {
	total, err := m.queryCount(index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}

	obj := OfficeMgr.NewOffice()
	query := fmt.Sprintf("SELECT %s FROM [dbo].[testCRUD] %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	results, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_OfficeDBMgr) FindFetchContext(ctx context.Context, index Index) (int64, []*Office, error) {
	total, err := m.queryCountContext(ctx, index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}

	obj := OfficeMgr.NewOffice()
	query := fmt.Sprintf("SELECT %s FROM [dbo].[testCRUD] %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	results, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_OfficeDBMgr) Range(scope Range) (int64, []PrimaryKey, error) {
	total, err := m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
	return total, pks, err
}

func (m *_OfficeDBMgr) RangeContext(ctx context.Context, scope Range) (int64, []PrimaryKey, error) {
	total, err := m.queryCountContext(ctx, scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimitContext(ctx, scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
	return total, pks, err
}

func (m *_OfficeDBMgr) RangeFetch(scope Range) (int64, []*Office, error) {
	total, err := m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	obj := OfficeMgr.NewOffice()
	query := fmt.Sprintf("SELECT %s FROM [dbo].[testCRUD] %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	results, err := m.FetchBySQL(query, scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_OfficeDBMgr) RangeFetchContext(ctx context.Context, scope Range) (int64, []*Office, error) {
	total, err := m.queryCountContext(ctx, scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	obj := OfficeMgr.NewOffice()
	query := fmt.Sprintf("SELECT %s FROM [dbo].[testCRUD] %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	results, err := m.FetchBySQLContext(ctx, query, scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_OfficeDBMgr) RangeRevert(scope Range) (int64, []PrimaryKey, error) {
	scope.Revert(true)
	return m.Range(scope)
}

func (m *_OfficeDBMgr) RangeRevertContext(ctx context.Context, scope Range) (int64, []PrimaryKey, error) {
	scope.Revert(true)
	return m.RangeContext(ctx, scope)
}

func (m *_OfficeDBMgr) RangeRevertFetch(scope Range) (int64, []*Office, error) {
	scope.Revert(true)
	return m.RangeFetch(scope)
}

func (m *_OfficeDBMgr) RangeRevertFetchContext(ctx context.Context, scope Range) (int64, []*Office, error) {
	scope.Revert(true)
	return m.RangeFetchContext(ctx, scope)
}

func (m *_OfficeDBMgr) queryLimit(where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := OfficeMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM [dbo].[testCRUD] %s", strings.Join(pk.Columns(), ","), where)
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("Office query limit error: %v", err)
	}
	defer rows.Close()

	offset := 0

	for rows.Next() {
		if limit >= 0 && offset >= limit {
			break
		}
		offset++

		result := OfficeMgr.NewPrimaryKey()
		err = rows.Scan(&(result.OfficeId))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("Office query limit result error: %v", err)
	}
	return
}

func (m *_OfficeDBMgr) queryLimitContext(ctx context.Context, where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := OfficeMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM [dbo].[testCRUD] %s", strings.Join(pk.Columns(), ","), where)
	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("Office query limit error: %v", err)
	}
	defer rows.Close()

	offset := 0

	for rows.Next() {
		if limit >= 0 && offset >= limit {
			break
		}
		offset++

		result := OfficeMgr.NewPrimaryKey()
		err = rows.Scan(&(result.OfficeId))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("Office query limit result error: %v", err)
	}
	return
}

func (m *_OfficeDBMgr) queryCount(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(office_id) FROM [dbo].[testCRUD] %s", where)
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return 0, fmt.Errorf("Office query count error: %v", err)
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

func (m *_OfficeDBMgr) queryCountContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(office_id) FROM [dbo].[testCRUD] %s", where)
	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("Office query count error: %v", err)
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

func (m *_OfficeDBMgr) BatchCreate(objs []*Office) (int64, error) {
	if len(objs) == 0 {
		return 0, nil
	}

	params := make([]string, 0, len(objs))
	values := make([]interface{}, 0, len(objs)*8)
	for _, obj := range objs {
		params = append(params, fmt.Sprintf("(%s)", strings.Join(orm.NewStringSlice(8, "?"), ",")))
		values = append(values, obj.OfficeArea)
		values = append(values, obj.OfficeName)
		values = append(values, obj.SearchOriginCode)
		values = append(values, obj.ProcessingOriginCode)
		values = append(values, obj.CreateBy)
		values = append(values, obj.UpdateBy)
		values = append(values, orm.MsSQLTimeFormat(obj.CreateDate))
		values = append(values, orm.MsSQLTimeFormat(obj.UpdateDate))
	}
	query := fmt.Sprintf("INSERT INTO [dbo].[testCRUD](%s) VALUES %s", strings.Join(objs[0].GetNoneIncrementColumns(), ","), strings.Join(params, ","))
	result, err := m.db.Exec(query, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_OfficeDBMgr) BatchCreateContext(ctx context.Context, objs []*Office) (int64, error) {
	if len(objs) == 0 {
		return 0, nil
	}

	params := make([]string, 0, len(objs))
	values := make([]interface{}, 0, len(objs)*8)
	for _, obj := range objs {
		params = append(params, fmt.Sprintf("(%s)", strings.Join(orm.NewStringSlice(8, "?"), ",")))
		values = append(values, obj.OfficeArea)
		values = append(values, obj.OfficeName)
		values = append(values, obj.SearchOriginCode)
		values = append(values, obj.ProcessingOriginCode)
		values = append(values, obj.CreateBy)
		values = append(values, obj.UpdateBy)
		values = append(values, orm.MsSQLTimeFormat(obj.CreateDate))
		values = append(values, orm.MsSQLTimeFormat(obj.UpdateDate))
	}
	query := fmt.Sprintf("INSERT INTO [dbo].[testCRUD](%s) VALUES %s", strings.Join(objs[0].GetNoneIncrementColumns(), ","), strings.Join(params, ","))
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
func (m *_OfficeDBMgr) UpdateBySQL(set, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("UPDATE [dbo].[testCRUD] SET %s", set)
	if where != "" {
		query = fmt.Sprintf("UPDATE [dbo].[testCRUD] SET %s WHERE %s", set, where)
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
func (m *_OfficeDBMgr) UpdateBySQLContext(ctx context.Context, set, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("UPDATE [dbo].[testCRUD] SET %s", set)
	if where != "" {
		query = fmt.Sprintf("UPDATE [dbo].[testCRUD] SET %s WHERE %s", set, where)
	}
	result, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_OfficeDBMgr) Create(obj *Office) (int64, error) {
	params := orm.NewStringSlice(8, "?")
	q := fmt.Sprintf("INSERT INTO [dbo].[testCRUD](%s) VALUES(%s)",
		strings.Join(obj.GetNoneIncrementColumns(), ","),
		strings.Join(params, ","))

	values := make([]interface{}, 0, 9)
	values = append(values, obj.OfficeArea)
	values = append(values, obj.OfficeName)
	values = append(values, obj.SearchOriginCode)
	values = append(values, obj.ProcessingOriginCode)
	values = append(values, obj.CreateBy)
	values = append(values, obj.UpdateBy)
	values = append(values, orm.MsSQLTimeFormat(obj.CreateDate))
	values = append(values, orm.MsSQLTimeFormat(obj.UpdateDate))
	result, err := m.db.Exec(q, values...)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	obj.OfficeId = int32(lastInsertId)
	return result.RowsAffected()
}

func (m *_OfficeDBMgr) CreateContext(ctx context.Context, obj *Office) (int64, error) {
	params := orm.NewStringSlice(8, "?")
	q := fmt.Sprintf("INSERT INTO [dbo].[testCRUD](%s) VALUES(%s)",
		strings.Join(obj.GetNoneIncrementColumns(), ","),
		strings.Join(params, ","))

	values := make([]interface{}, 0, 9)
	values = append(values, obj.OfficeArea)
	values = append(values, obj.OfficeName)
	values = append(values, obj.SearchOriginCode)
	values = append(values, obj.ProcessingOriginCode)
	values = append(values, obj.CreateBy)
	values = append(values, obj.UpdateBy)
	values = append(values, orm.MsSQLTimeFormat(obj.CreateDate))
	values = append(values, orm.MsSQLTimeFormat(obj.UpdateDate))
	result, err := m.db.ExecContext(ctx, q, values...)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	obj.OfficeId = int32(lastInsertId)
	return result.RowsAffected()
}

func (m *_OfficeDBMgr) Update(obj *Office) (int64, error) {
	columns := []string{
		"office_area = ?",
		"office_name = ?",
		"search_origin_code = ?",
		"processing_origin_code = ?",
		"create_by = ?",
		"update_by = ?",
		"create_date = ?",
		"update_date = ?",
	}

	pk := obj.GetPrimaryKey()
	q := fmt.Sprintf("UPDATE [dbo].[testCRUD] SET %s %s", strings.Join(columns, ","), pk.SQLFormat())
	values := make([]interface{}, 0, 9-1)
	values = append(values, obj.OfficeArea)
	values = append(values, obj.OfficeName)
	values = append(values, obj.SearchOriginCode)
	values = append(values, obj.ProcessingOriginCode)
	values = append(values, obj.CreateBy)
	values = append(values, obj.UpdateBy)
	values = append(values, orm.MsSQLTimeFormat(obj.CreateDate))
	values = append(values, orm.MsSQLTimeFormat(obj.UpdateDate))
	values = append(values, pk.SQLParams()...)

	result, err := m.db.Exec(q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_OfficeDBMgr) UpdateContext(ctx context.Context, obj *Office) (int64, error) {
	columns := []string{
		"office_area = ?",
		"office_name = ?",
		"search_origin_code = ?",
		"processing_origin_code = ?",
		"create_by = ?",
		"update_by = ?",
		"create_date = ?",
		"update_date = ?",
	}

	pk := obj.GetPrimaryKey()
	q := fmt.Sprintf("UPDATE [dbo].[testCRUD] SET %s %s", strings.Join(columns, ","), pk.SQLFormat())
	values := make([]interface{}, 0, 9-1)
	values = append(values, obj.OfficeArea)
	values = append(values, obj.OfficeName)
	values = append(values, obj.SearchOriginCode)
	values = append(values, obj.ProcessingOriginCode)
	values = append(values, obj.CreateBy)
	values = append(values, obj.UpdateBy)
	values = append(values, orm.MsSQLTimeFormat(obj.CreateDate))
	values = append(values, orm.MsSQLTimeFormat(obj.UpdateDate))
	values = append(values, pk.SQLParams()...)

	result, err := m.db.ExecContext(ctx, q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_OfficeDBMgr) Save(obj *Office) (int64, error) {
	affected, err := m.Update(obj)
	if err != nil {
		return affected, err
	}
	if affected == 0 {
		return m.Create(obj)
	}
	return affected, err
}

func (m *_OfficeDBMgr) SaveContext(ctx context.Context, obj *Office) (int64, error) {
	affected, err := m.UpdateContext(ctx, obj)
	if err != nil {
		return affected, err
	}
	if affected == 0 {
		return m.CreateContext(ctx, obj)
	}
	return affected, err
}

func (m *_OfficeDBMgr) Delete(obj *Office) (int64, error) {
	return m.DeleteByPrimaryKey(obj.OfficeId)
}

func (m *_OfficeDBMgr) DeleteContext(ctx context.Context, obj *Office) (int64, error) {
	return m.DeleteByPrimaryKeyContext(ctx, obj.OfficeId)
}

func (m *_OfficeDBMgr) DeleteByPrimaryKey(officeId int32) (int64, error) {
	pk := &OfficeIdOfOfficePK{
		OfficeId: officeId,
	}
	q := fmt.Sprintf("DELETE FROM [dbo].[testCRUD] %s", pk.SQLFormat())
	result, err := m.db.Exec(q, pk.SQLParams()...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_OfficeDBMgr) DeleteByPrimaryKeyContext(ctx context.Context, officeId int32) (int64, error) {
	pk := &OfficeIdOfOfficePK{
		OfficeId: officeId,
	}
	q := fmt.Sprintf("DELETE FROM [dbo].[testCRUD] %s", pk.SQLFormat())
	result, err := m.db.ExecContext(ctx, q, pk.SQLParams()...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_OfficeDBMgr) DeleteBySQL(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("DELETE FROM [dbo].[testCRUD]")
	if where != "" {
		query = fmt.Sprintf("DELETE FROM [dbo].[testCRUD] WHERE %s", where)
	}
	result, err := m.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_OfficeDBMgr) DeleteBySQLContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("DELETE FROM [dbo].[testCRUD]")
	if where != "" {
		query = fmt.Sprintf("DELETE FROM [dbo].[testCRUD] WHERE %s", where)
	}
	result, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
