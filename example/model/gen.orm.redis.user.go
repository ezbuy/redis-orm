package model

import (
	"github.com/ezbuy/redis-orm/orm"
	redis "gopkg.in/redis.v5"
)

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
	return m.AddBySQL(db, "SELECT `id`,`name`,`mailbox`,`sex`,`longitude`,`latitude`,`description`,`password`,`head_url`,`status`,`created_at`, `updated_at` FROM users")

	return nil
}

func (m *_UserRedisMgr) AddBySQL(db DBFetcher, sql string, args ...interface{}) error {
	objs, err := db.FetchBySQL(sql, args)
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
	objs, err := db.FetchBySQL(sql, args)
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
	return "", nil
}

func (m *_UserRedisMgr) Find(index Index) ([]string, error) {
	return nil, nil
}

func (m *_UserRedisMgr) Range(scope Range) ([]string, error) {
	return nil, nil
}

func (m *_UserRedisMgr) OrderBy(sort OrderBy) ([]string, error) {
	return nil, nil
}

func (m *_UserRedisMgr) Fetch(id string) (*User, error) {
	return nil, nil
}

func (m *_UserRedisMgr) FetchByIds(ids []string) ([]*User, error) {
	return nil, nil
}

func (m *_UserRedisMgr) Create(obj *User) error {
	return m.Save(obj)
}

func (m *_UserRedisMgr) Update(obj *User) error {
	return m.Save(obj)
}

func (m *_UserRedisMgr) Delete(obj *User) error {
	return nil
}

func (m *_UserRedisMgr) Save(obj *User) error {
	return nil
}

func (m *_UserRedisMgr) Clear() error {
	return nil
}

//! uniques

//! relation
type MailboxPasswordOfUserRelation struct {
	Key   string `db:"key" json:"key"`
	Value int32  `db:"value" json:"value"`
}

//! redis relation pair
type _MailboxPasswordOfUserRelationRedisMgr struct {
	*orm.RedisStore
}

func MailboxPasswordOfUserRelationRedisMgr() *_MailboxPasswordOfUserRelationRedisMgr {
	return &_MailboxPasswordOfUserRelationRedisMgr{_redis_store}
}

func NewMailboxPasswordOfUserRelationRedisMgr(cf *RedisConfig) (*_MailboxPasswordOfUserRelationRedisMgr, error) {
	store, err := orm.NewRedisStore(cf.Host, cf.Port, cf.Password, 0)
	if err != nil {
		return nil, err
	}
	return &_MailboxPasswordOfUserRelationRedisMgr{store}, nil
}

//! pipeline write
type _MailboxPasswordOfUserRelationRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_MailboxPasswordOfUserRelationRedisMgr) BeginPipeline() *_MailboxPasswordOfUserRelationRedisPipeline {
	return &_MailboxPasswordOfUserRelationRedisPipeline{m.Pipeline(), nil}
}

func (m *_MailboxPasswordOfUserRelationRedisMgr) PSet(obj *MailboxPasswordOfUserRelation) error {
	return nil
}

func (m *_MailboxPasswordOfUserRelationRedisMgr) PGet(obj *MailboxPasswordOfUserRelation) error {
	return nil
}

func (m *_MailboxPasswordOfUserRelationRedisMgr) PRem(obj *MailboxPasswordOfUserRelation) error {
	return nil
}

//! indexes

//! relation
type SexOfUserRelation struct {
	Key   string `db:"key" json:"key"`
	Value int32  `db:"value" json:"value"`
}

//! redis relation pair
type _SexOfUserRelationRedisMgr struct {
	*orm.RedisStore
}

func SexOfUserRelationRedisMgr() *_SexOfUserRelationRedisMgr {
	return &_SexOfUserRelationRedisMgr{_redis_store}
}

func NewSexOfUserRelationRedisMgr(cf *RedisConfig) (*_SexOfUserRelationRedisMgr, error) {
	store, err := orm.NewRedisStore(cf.Host, cf.Port, cf.Password, 0)
	if err != nil {
		return nil, err
	}
	return &_SexOfUserRelationRedisMgr{store}, nil
}

//! pipeline write
type _SexOfUserRelationRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_SexOfUserRelationRedisMgr) BeginPipeline() *_SexOfUserRelationRedisPipeline {
	return &_SexOfUserRelationRedisPipeline{m.Pipeline(), nil}
}

func (m *_SexOfUserRelationRedisMgr) SAdd(obj *SexOfUserRelation) error {
	return nil
}

func (m *_SexOfUserRelationRedisMgr) SGet(obj *SexOfUserRelation) error {
	return nil
}

func (m *_SexOfUserRelationRedisMgr) SRem(obj *SexOfUserRelation) error {
	return nil
}

//! ranges

//! orders
