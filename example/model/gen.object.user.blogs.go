package model

import (
	"github.com/ezbuy/redis-orm/orm"
	redis "gopkg.in/redis.v5"
)

var (
	_ orm.VSet
)

//! relation
type UserBlogs struct {
	Key   string `db:"key" json:"key"`
	Value int32  `db:"value" json:"value"`
}

//! redis relation pair
type _UserBlogsRedisMgr struct {
	*orm.RedisStore
}

func UserBlogsRedisMgr() *_UserBlogsRedisMgr {
	return &_UserBlogsRedisMgr{_redis_store}
}

func NewUserBlogsRedisMgr(cf *RedisConfig) (*_UserBlogsRedisMgr, error) {
	store, err := orm.NewRedisStore(cf.Host, cf.Port, cf.Password, 0)
	if err != nil {
		return nil, err
	}
	return &_UserBlogsRedisMgr{store}, nil
}

//! pipeline write
type _UserBlogsRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_UserBlogsRedisMgr) BeginPipeline() *_UserBlogsRedisPipeline {
	return &_UserBlogsRedisPipeline{m.Pipeline(), nil}
}

func (m *_UserBlogsRedisMgr) SAdd(obj *UserBlogs) error {
	return nil
}

func (m *_UserBlogsRedisMgr) SGet(obj *UserBlogs) error {
	return nil
}

func (m *_UserBlogsRedisMgr) SRem(obj *UserBlogs) error {
	return nil
}
