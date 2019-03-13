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

type IndexedBlog struct {
	Id        int32     `json:"id"`
	UserId    int32     `json:"user_id"`
	Hash      string    `json:"hash"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Readed    int32     `json:"readed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type _IndexedBlogMgr struct {
}

var IndexedBlogMgr *_IndexedBlogMgr

func (m *_IndexedBlogMgr) NewIndexedBlog() *IndexedBlog {
	return &IndexedBlog{}
}

//! object function

func (obj *IndexedBlog) GetNameSpace() string {
	return "model"
}

func (obj *IndexedBlog) GetClassName() string {
	return "IndexedBlog"
}

func (obj *IndexedBlog) GetTableName() string {
	return "indexed_blog"
}

func (obj *IndexedBlog) GetColumns() []string {
	columns := []string{
		"indexed_blog.id",
		"indexed_blog.user_id",
		"indexed_blog.hash",
		"indexed_blog.title",
		"indexed_blog.content",
		"indexed_blog.readed",
		"indexed_blog.created_at",
		"indexed_blog.updated_at",
	}
	return columns
}

func (obj *IndexedBlog) GetNoneIncrementColumns() []string {
	columns := []string{
		"id",
		"user_id",
		"hash",
		"title",
		"content",
		"readed",
		"created_at",
		"updated_at",
	}
	return columns
}

func (obj *IndexedBlog) GetPrimaryKey() PrimaryKey {
	pk := IndexedBlogMgr.NewPrimaryKey()
	pk.Id = obj.Id
	return pk
}

func (obj *IndexedBlog) Validate() error {
	validate := validator.New()
	return validate.Struct(obj)
}

//! primary key

type IdOfIndexedBlogPK struct {
	Id int32
}

func (m *_IndexedBlogMgr) NewPrimaryKey() *IdOfIndexedBlogPK {
	return &IdOfIndexedBlogPK{}
}

func (u *IdOfIndexedBlogPK) Key() string {
	strs := []string{
		"Id",
		fmt.Sprint(u.Id),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *IdOfIndexedBlogPK) Parse(key string) error {
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

func (u *IdOfIndexedBlogPK) SQLFormat() string {
	conditions := []string{
		"id = ?",
	}
	return orm.SQLWhere(conditions)
}

func (u *IdOfIndexedBlogPK) SQLParams() []interface{} {
	return []interface{}{
		u.Id,
	}
}

func (u *IdOfIndexedBlogPK) Columns() []string {
	return []string{
		"id",
	}
}

//! uniques

//! indexes

//! ranges

type IdOfIndexedBlogRNG struct {
	IdBegin      int64
	IdEnd        int64
	offset       int
	limit        int
	includeBegin bool
	includeEnd   bool
	revert       bool
}

func (u *IdOfIndexedBlogRNG) Key() string {
	strs := []string{
		"Id",
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *IdOfIndexedBlogRNG) beginOp() string {
	if u.includeBegin {
		return ">="
	}
	return ">"
}
func (u *IdOfIndexedBlogRNG) endOp() string {
	if u.includeBegin {
		return "<="
	}
	return "<"
}

func (u *IdOfIndexedBlogRNG) SQLFormat(limit bool) string {
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
		return fmt.Sprintf("%s %s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("id", u.revert), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("id", u.revert))
}

func (u *IdOfIndexedBlogRNG) SQLParams() []interface{} {
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

func (u *IdOfIndexedBlogRNG) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *IdOfIndexedBlogRNG) Limit(n int) {
	u.limit = n
}

func (u *IdOfIndexedBlogRNG) Offset(n int) {
	u.offset = n
}

func (u *IdOfIndexedBlogRNG) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *IdOfIndexedBlogRNG) Begin() int64 {
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

func (u *IdOfIndexedBlogRNG) End() int64 {
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

func (u *IdOfIndexedBlogRNG) Revert(b bool) {
	u.revert = b
}

func (u *IdOfIndexedBlogRNG) IncludeBegin(f bool) {
	u.includeBegin = f
}

func (u *IdOfIndexedBlogRNG) IncludeEnd(f bool) {
	u.includeEnd = f
}

func (u *IdOfIndexedBlogRNG) RNGRelation(store *orm.RedisStore) RangeRelation {
	return nil
}

var (
	_ context.Context
)

//! orm.elastic
var IndexedBlogElasticFields = struct {
	UserId    string
	Hash      string
	Title     string
	Content   string
	Readed    string
	CreatedAt string
	UpdatedAt string
}{
	"user_id",
	"hash",
	"title",
	"content",
	"readed",
	"created_at",
	"updated_at",
}

var IndexedBlogElasticMgr = &_IndexedBlogElasticMgr{}

type _IndexedBlogElasticMgr struct {
	ensureMapping sync.Once
}

func (m *_IndexedBlogElasticMgr) Mapping() map[string]interface{} {
	return map[string]interface{}{
		"properties": map[string]interface{}{
			"user_id": map[string]interface{}{
				"type": "integer",
			},
			"hash": map[string]interface{}{
				"type":  "string",
				"index": "not_analyzed",
			},
			"title": map[string]interface{}{
				"type":  "string",
				"index": "analyzed",
			},
			"content": map[string]interface{}{
				"type":     "string",
				"index":    "analyzed",
				"analyzer": "standard",
			},
			"readed": map[string]interface{}{
				"type": "integer",
			},
			"created_at": map[string]interface{}{
				"type":   "date",
				"format": "yyyy-MM-dd HH:mm:ss",
			},
			"updated_at": map[string]interface{}{
				"type":   "date",
				"format": "yyyy-MM-dd HH:mm:ss",
			},
		},
	}
}

func (m *_IndexedBlogElasticMgr) IndexService() (*elastic.IndexService, error) {
	var err error
	m.ensureMapping.Do(func() {
		_, err = m.PutMappingService().BodyJson(m.Mapping()).Do()
	})

	return ElasticClient().IndexService("ezsearch").Type("indexed_blog"), err
}

func (m *_IndexedBlogElasticMgr) PutMappingService() *elastic.PutMappingService {
	return ElasticClient().PutMappingService("ezsearch ").Type("indexed_blog")
}
