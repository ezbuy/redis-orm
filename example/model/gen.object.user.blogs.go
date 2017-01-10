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
type UserBlogs struct {
	Key   string  `db:"key" json:"key"`
	Score float64 `db:"score" json:"score"`
	Value int32   `db:"value" json:"value"`
}

func (relation *UserBlogs) GetClassName() string {
	return "UserBlogs"
}

func (relation *UserBlogs) GetIndexes() []string {
	idx := []string{}
	return idx
}

func (relation *UserBlogs) GetStoreType() string {
	return "zset"
}

func (relation *UserBlogs) GetPrimaryName() string {
	return "Key"
}

type _UserBlogsRedisMgr struct {
	*orm.RedisStore
}

func UserBlogsRedisMgr(stores ...*orm.RedisStore) *_UserBlogsRedisMgr {
	if len(stores) > 0 {
		return &_UserBlogsRedisMgr{stores[0]}
	}
	return &_UserBlogsRedisMgr{_redis_store}
}

//! pipeline write
type _UserBlogsRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_UserBlogsRedisMgr) BeginPipeline() *_UserBlogsRedisPipeline {
	return &_UserBlogsRedisPipeline{m.Pipeline(), nil}
}

func (m *_UserBlogsRedisMgr) NewUserBlogs(key string) *UserBlogs {
	return &UserBlogs{
		Key: key,
	}
}

//! redis relation zset
func (m *_UserBlogsRedisMgr) ZSetAdd(obj *UserBlogs) error {
	return m.ZAdd(zsetOfClass(obj.GetClassName(), obj.Key), redis.Z{Score: obj.Score, Member: obj.Value}).Err()
}

func (m *_UserBlogsRedisMgr) ZSetRange(key string, min, max int64) ([]*UserBlogs, error) {
	strs, err := m.ZRange(zsetOfClass("UserBlogs", key), min, max).Result()
	if err != nil {
		return nil, err
	}

	objs := make([]*UserBlogs, len(strs))
	for _, str := range strs {
		obj := m.NewUserBlogs(key)
		if err := m.StringScan(str, &obj.Value); err != nil {
			return nil, err
		}
		objs = append(objs, obj)
	}
	return objs, nil
}

func (m *_UserBlogsRedisMgr) ZSetRem(obj *UserBlogs) error {
	return m.ZRem(zsetOfClass(obj.GetClassName(), obj.Key), redis.Z{Score: obj.Score, Member: obj.Value}).Err()
}

func (m *_UserBlogsRedisMgr) Range(key string, min, max int64) ([]string, error) {
	return m.ZRange(zsetOfClass("UserBlogs", key), min, max).Result()
}

func (m *_UserBlogsRedisMgr) OrderBy(key string, asc bool) ([]string, error) {
	//! TODO revert
	return m.ZRange(zsetOfClass("UserBlogs", key), 0, -1).Result()
}
