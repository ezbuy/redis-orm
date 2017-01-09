package model

import (
	"fmt"
	redis "gopkg.in/redis.v5"

	"github.com/ezbuy/redis-orm/orm"
)

var (
	_ fmt.Formatter
	_ orm.VSet
)

//! relation
type UserId struct {
	Key   string `db:"key" json:"key"`
	Value int32  `db:"value" json:"value"`
}

func (relation *UserId) GetClassName() string {
	return "UserId"
}

func (relation *UserId) GetIndexes() []string {
	idx := []string{}
	return idx
}

func (relation *UserId) GetStoreType() string {
	return "list"
}

func (relation *UserId) GetPrimaryName() string {
	return "Key"
}

type _UserIdRedisMgr struct {
	*orm.RedisStore
}

func UserIdRedisMgr(stores ...*orm.RedisStore) *_UserIdRedisMgr {
	if len(stores) > 0 {
		return &_UserIdRedisMgr{stores[0]}
	}
	return &_UserIdRedisMgr{_redis_store}
}

//! pipeline write
type _UserIdRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_UserIdRedisMgr) BeginPipeline() *_UserIdRedisPipeline {
	return &_UserIdRedisPipeline{m.Pipeline(), nil}
}

func (m *_UserIdRedisMgr) NewUserId(key string) *UserId {
	return &UserId{
		Key: key,
	}
}

//! redis relation list
func (m *_UserIdRedisMgr) ListLPush(obj *UserId) error {
	return m.LPush(listOfClass(obj.GetClassName(), obj.Key), obj.Value).Err()
}

func (m *_UserIdRedisMgr) ListRPush(obj *UserId) error {
	return m.RPush(listOfClass(obj.GetClassName(), obj.Key), obj.Value).Err()
}

func (m *_UserIdRedisMgr) ListLPop(key string) (*UserId, error) {
	str, err := m.LPop(listOfClass("UserId", key)).Result()
	if err != nil {
		return nil, err
	}

	obj := m.NewUserId(key)
	if err := m.StringScan(str, &obj.Value); err != nil {
		return nil, err
	}

	return obj, nil
}

func (m *_UserIdRedisMgr) ListRPop(key string) (*UserId, error) {
	str, err := m.RPop(listOfClass("UserId", key)).Result()
	if err != nil {
		return nil, err
	}

	obj := m.NewUserId(key)
	if err := m.StringScan(str, &obj.Value); err != nil {
		return nil, err
	}

	return obj, nil
}

func (m *_UserIdRedisMgr) ListLRange(key string, start, stop int64) ([]*UserId, error) {
	strs, err := m.LRange(listOfClass("UserId", key), start, stop).Result()
	if err != nil {
		return nil, err
	}

	objs := make([]*UserId, len(strs))
	for _, str := range strs {
		obj := m.NewUserId(key)
		if err := m.StringScan(str, &obj.Value); err != nil {
			return nil, err
		}
		objs = append(objs, obj)
	}
	return objs, nil
}

func (m *_UserIdRedisMgr) ListLRem(obj *UserId) error {
	return m.LRem(listOfClass(obj.GetClassName(), obj.Key), 0, obj.Value).Err()
}

func (m *_UserIdRedisMgr) ListLLen(key string) (int64, error) {
	return m.LLen(listOfClass("UserId", key)).Result()
}

func (m *_UserIdRedisMgr) ListLDel() error {
	return nil
}
