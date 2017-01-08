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
