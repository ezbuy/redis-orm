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
type UserLocation struct {
	Key       string  `db:"key" json:"key"`
	Longitude float64 `db:"longitude" json:"longitude"`
	Latitude  float64 `db:"latitude" json:"latitude"`
	Value     int32   `db:"value" json:"value"`
}

func (relation *UserLocation) GetClassName() string {
	return "UserLocation"
}

func (relation *UserLocation) GetIndexes() []string {
	idx := []string{}
	return idx
}

func (relation *UserLocation) GetStoreType() string {
	return "geo"
}

func (relation *UserLocation) GetPrimaryName() string {
	return "Key"
}

type _UserLocationRedisMgr struct {
	*orm.RedisStore
}

func UserLocationRedisMgr(stores ...*orm.RedisStore) *_UserLocationRedisMgr {
	if len(stores) > 0 {
		return &_UserLocationRedisMgr{stores[0]}
	}
	return &_UserLocationRedisMgr{_redis_store}
}

//! pipeline write
type _UserLocationRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_UserLocationRedisMgr) BeginPipeline() *_UserLocationRedisPipeline {
	return &_UserLocationRedisPipeline{m.Pipeline(), nil}
}

func (m *_UserLocationRedisMgr) NewUserLocation(key string) *UserLocation {
	return &UserLocation{
		Key: key,
	}
}

//! redis relation pair
func (m *_UserLocationRedisMgr) LocationAdd(relation *UserLocation) error {
	return m.GeoAdd(geoOfClass("UserLocation", "UserLocation", relation.Key), &redis.GeoLocation{
		Longitude: relation.Longitude,
		Latitude:  relation.Latitude,
		Name:      fmt.Sprint(relation.Value),
	}).Err()
}

func (m *_UserLocationRedisMgr) LocationRadius(key string, longitude float64, latitude float64, query *redis.GeoRadiusQuery) ([]*UserLocation, error) {
	locations, err := m.GeoRadius(geoOfClass("UserLocation", "UserLocation", key), longitude, latitude, query).Result()
	if err != nil {
		return nil, err
	}

	relations := []*UserLocation{}
	for _, location := range locations {
		relation := m.NewUserLocation(key)
		relation.Longitude = location.Longitude
		relation.Latitude = location.Latitude
		if err := m.StringScan(location.Name, &relation.Value); err != nil {
			return nil, err
		}
		relations = append(relations, relation)
	}
	return relations, nil
}

func (m *_UserLocationRedisMgr) LocationRem(relation *UserLocation) error {
	return m.ZRem(geoOfClass("UserLocation", "UserLocation", relation.Key), fmt.Sprint(relation.Value)).Err()
}

func (m *_UserLocationRedisMgr) LocationDel(key string) error {
	return m.Del(geoOfClass("UserLocation", "UserLocation", key)).Err()
}

func (m *_UserLocationRedisMgr) Clear() error {
	strs, err := m.Keys(geoOfClass("UserLocation", "UserLocation", "*")).Result()
	if err != nil {
		return err
	}
	return m.Del(strs...).Err()
}

func (m *_UserLocationRedisMgr) Load(db DBFetcher) error {

	if err := m.Clear(); err != nil {
		return err
	}
	return m.AddBySQL(db, "SELECT 'all',`longitude`,`latitude`,`id` FROM users")

}

func (m *_UserLocationRedisMgr) AddBySQL(db DBFetcher, sql string, args ...interface{}) error {
	objs, err := db.FetchBySQL(sql, args...)
	if err != nil {
		return err
	}

	for _, obj := range objs {
		if err := m.LocationAdd(obj.(*UserLocation)); err != nil {
			return err
		}
	}

	return nil
}
func (m *_UserLocationRedisMgr) DelBySQL(db DBFetcher, sql string, args ...interface{}) error {
	objs, err := db.FetchBySQL(sql, args...)
	if err != nil {
		return err
	}

	for _, obj := range objs {
		if err := m.LocationRem(obj.(*UserLocation)); err != nil {
			return err
		}
	}
	return nil
}

type _UserLocationMySQLMgr struct {
	*orm.MySQLStore
}

func UserLocationMySQLMgr() *_UserLocationMySQLMgr {
	return &_UserLocationMySQLMgr{_mysql_store}
}

func NewUserLocationMySQLMgr(cf *MySQLConfig) (*_UserLocationMySQLMgr, error) {
	store, err := orm.NewMySQLStore(cf.Host, cf.Port, cf.Database, cf.UserName, cf.Password)
	if err != nil {
		return nil, err
	}
	return &_UserLocationMySQLMgr{store}, nil
}

func (m *_UserLocationMySQLMgr) FetchBySQL(sql string, args ...interface{}) (results []interface{}, err error) {
	rows, err := m.Query(sql, args...)
	if err != nil {
		return nil, fmt.Errorf("UserLocation fetch error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var result UserLocation
		err = rows.Scan(&(result.Key),
			&(result.Longitude),
			&(result.Latitude),
			&(result.Value),
		)
		if err != nil {
			return nil, err
		}

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("UserLocation fetch result error: %v", err)
	}
	return
}
