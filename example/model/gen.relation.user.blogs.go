package model

import (
	"fmt"
	"github.com/ezbuy/redis-orm/orm"
	redis "gopkg.in/redis.v5"
	"strings"
	"time"
)

var (
	_ time.Time
	_ fmt.Formatter
	_ strings.Reader
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
func (m *_UserBlogsRedisMgr) ZSetAdd(relation *UserBlogs) error {
	return m.ZAdd(zsetOfClass("UserBlogs", "UserBlogs", relation.Key), redis.Z{Score: relation.Score, Member: relation.Value}).Err()
}

func (m *_UserBlogsRedisMgr) ZSetRange(key string, min, max int64) ([]*UserBlogs, error) {
	strs, err := m.ZRange(zsetOfClass("UserBlogs", key), min, max).Result()
	if err != nil {
		return nil, err
	}

	relations := make([]*UserBlogs, len(strs))
	for _, str := range strs {
		relation := m.NewUserBlogs(key)
		if err := m.StringScan(str, &relation.Value); err != nil {
			return nil, err
		}
		relations = append(relations, relation)
	}
	return relations, nil
}

func (m *_UserBlogsRedisMgr) ZSetRem(relation *UserBlogs) error {
	return m.ZRem(zsetOfClass("UserBlogs", "UserBlogs", relation.Key), redis.Z{Score: relation.Score, Member: relation.Value}).Err()
}

func (m *_UserBlogsRedisMgr) ZSetDel(key string) error {
	return m.Del(setOfClass("UserBlogs", "UserBlogs", key)).Err()
}

func (m *_UserBlogsRedisMgr) Range(key string, min, max int64) ([]string, error) {
	return m.ZRange(zsetOfClass("UserBlogs", "UserBlogs", key), min, max).Result()
}

func (m *_UserBlogsRedisMgr) OrderBy(key string, asc bool) ([]string, error) {
	//! TODO revert
	return m.ZRange(zsetOfClass("UserBlogs", "UserBlogs", key), 0, -1).Result()
}
