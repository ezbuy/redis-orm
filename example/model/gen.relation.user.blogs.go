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
	Key   string `db:"key" json:"key"`
	Value int32  `db:"value" json:"value"`
}

func (relation *UserBlogs) GetClassName() string {
	return "UserBlogs"
}

func (relation *UserBlogs) GetIndexes() []string {
	idx := []string{}
	return idx
}

func (relation *UserBlogs) GetStoreType() string {
	return "set"
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

func (m *_UserBlogsRedisMgr) NewUserBlogs(key string) *UserBlogs {
	return &UserBlogs{
		Key: key,
	}
}

//! pipeline
type _UserBlogsRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_UserBlogsRedisMgr) BeginPipeline(pipes ...*redis.Pipeline) *_UserBlogsRedisPipeline {
	if len(pipes) > 0 {
		return &_UserBlogsRedisPipeline{pipes[0], nil}
	}
	return &_UserBlogsRedisPipeline{m.Pipeline(), nil}
}

//! redis relation pair
func (m *_UserBlogsRedisMgr) SetAdd(relation *UserBlogs) error {
	return m.SAdd(setOfClass("UserBlogs", "UserBlogs", relation.Key), relation.Value).Err()
}

func (pipe *_UserBlogsRedisPipeline) SetAdd(relation *UserBlogs) error {
	return pipe.SAdd(setOfClass("UserBlogs", "UserBlogs", relation.Key), relation.Value).Err()
}

func (m *_UserBlogsRedisMgr) SetGet(key string) ([]*UserBlogs, error) {
	strs, err := m.SMembers(setOfClass("UserBlogs", "UserBlogs", key)).Result()
	if err != nil {
		return nil, err
	}

	relations := make([]*UserBlogs, 0, len(strs))
	for _, str := range strs {
		relation := m.NewUserBlogs(key)
		if err := m.StringScan(str, &relation.Value); err != nil {
			return nil, err
		}
		relations = append(relations, relation)
	}
	return relations, nil
}

func (m *_UserBlogsRedisMgr) SetRem(relation *UserBlogs) error {
	return m.SRem(setOfClass("UserBlogs", "UserBlogs", relation.Key), relation.Value).Err()
}

func (pipe *_UserBlogsRedisPipeline) SetRem(relation *UserBlogs) error {
	return pipe.SRem(setOfClass("UserBlogs", "UserBlogs", relation.Key), relation.Value).Err()
}

func (m *_UserBlogsRedisMgr) SetDel(key string) error {
	return m.Del(setOfClass("UserBlogs", "UserBlogs", key)).Err()
}

func (pipe *_UserBlogsRedisPipeline) SetDel(key string) error {
	return pipe.Del(setOfClass("UserBlogs", "UserBlogs", key)).Err()
}

func (m *_UserBlogsRedisMgr) Find(key string) ([]string, error) {
	return m.SMembers(setOfClass("UserBlogs", "UserBlogs", key)).Result()
}

func (m *_UserBlogsRedisMgr) Clear() error {
	strs, err := m.Keys(setOfClass("UserBlogs", "UserBlogs", "*")).Result()
	if err != nil {
		return err
	}
	if len(strs) > 0 {
		return m.Del(strs...).Err()
	}
	return nil
}

func (m *_UserBlogsRedisMgr) Load(db DBFetcher) error {

	if err := m.Clear(); err != nil {
		return err
	}
	return m.AddBySQL(db, "SELECT `user_id`, `id` FROM blogs")

}

func (m *_UserBlogsRedisMgr) AddBySQL(db DBFetcher, sql string, args ...interface{}) error {
	objs, err := db.FetchBySQL(sql, args...)
	if err != nil {
		return err
	}

	for _, obj := range objs {
		if err := m.SetAdd(obj.(*UserBlogs)); err != nil {
			return err
		}
	}

	return nil
}
func (m *_UserBlogsRedisMgr) DelBySQL(db DBFetcher, sql string, args ...interface{}) error {
	objs, err := db.FetchBySQL(sql, args...)
	if err != nil {
		return err
	}

	for _, obj := range objs {
		if err := m.SetRem(obj.(*UserBlogs)); err != nil {
			return err
		}
	}
	return nil
}

type _UserBlogsMySQLMgr struct {
	*orm.MySQLStore
}

func UserBlogsMySQLMgr() *_UserBlogsMySQLMgr {
	return &_UserBlogsMySQLMgr{_mysql_store}
}

func NewUserBlogsMySQLMgr(cf *MySQLConfig) (*_UserBlogsMySQLMgr, error) {
	store, err := orm.NewMySQLStore(cf.Host, cf.Port, cf.Database, cf.UserName, cf.Password)
	if err != nil {
		return nil, err
	}
	return &_UserBlogsMySQLMgr{store}, nil
}

func (m *_UserBlogsMySQLMgr) FetchBySQL(q string, args ...interface{}) (results []interface{}, err error) {
	rows, err := m.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("UserBlogs fetch error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var result UserBlogs
		err = rows.Scan(&(result.Key), &(result.Value))
		if err != nil {
			return nil, err
		}

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("UserBlogs fetch result error: %v", err)
	}
	return
}
