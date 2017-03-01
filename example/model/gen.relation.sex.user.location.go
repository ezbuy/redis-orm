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
type SexUserLocation struct {
	Key       string  `db:"key" json:"key"`
	Longitude float64 `db:"longitude" json:"longitude"`
	Latitude  float64 `db:"latitude" json:"latitude"`
	Value     int32   `db:"value" json:"value"`
}

func (relation *SexUserLocation) GetClassName() string {
	return "SexUserLocation"
}

func (relation *SexUserLocation) GetIndexes() []string {
	idx := []string{}
	return idx
}

func (relation *SexUserLocation) GetStoreType() string {
	return "geo"
}

func (relation *SexUserLocation) GetPrimaryName() string {
	return "Key"
}

type _SexUserLocationRedisMgr struct {
	*orm.RedisStore
}

func SexUserLocationRedisMgr(stores ...*orm.RedisStore) *_SexUserLocationRedisMgr {
	if len(stores) > 0 {
		return &_SexUserLocationRedisMgr{stores[0]}
	}
	return &_SexUserLocationRedisMgr{_redis_store}
}

func (m *_SexUserLocationRedisMgr) NewSexUserLocation(key string) *SexUserLocation {
	return &SexUserLocation{
		Key: key,
	}
}

//! pipeline
type _SexUserLocationRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_SexUserLocationRedisMgr) BeginPipeline(pipes ...*redis.Pipeline) *_SexUserLocationRedisPipeline {
	if len(pipes) > 0 {
		return &_SexUserLocationRedisPipeline{pipes[0], nil}
	}
	return &_SexUserLocationRedisPipeline{m.Pipeline(), nil}
}

//! redis relation pair
func (m *_SexUserLocationRedisMgr) LocationAdd(relation *SexUserLocation) error {
	return m.GeoAdd(geoOfClass("SexUserLocation", "SexUserLocation", relation.Key), &redis.GeoLocation{
		Longitude: relation.Longitude,
		Latitude:  relation.Latitude,
		Name:      fmt.Sprint(relation.Value),
	}).Err()
}

func (m *_SexUserLocationRedisMgr) LocationRadius(key string, longitude float64, latitude float64, query *redis.GeoRadiusQuery) ([]*SexUserLocation, error) {
	locations, err := m.GeoRadius(geoOfClass("SexUserLocation", "SexUserLocation", key), longitude, latitude, query).Result()
	if err != nil {
		return nil, err
	}

	relations := []*SexUserLocation{}
	for _, location := range locations {
		relation := m.NewSexUserLocation(key)
		relation.Longitude = location.Longitude
		relation.Latitude = location.Latitude
		if err := m.StringScan(location.Name, &relation.Value); err != nil {
			return nil, err
		}
		relations = append(relations, relation)
	}
	return relations, nil
}

func (m *_SexUserLocationRedisMgr) LocationRem(relation *SexUserLocation) error {
	return m.ZRem(geoOfClass("SexUserLocation", "SexUserLocation", relation.Key), fmt.Sprint(relation.Value)).Err()
}

func (m *_SexUserLocationRedisMgr) LocationDel(key string) error {
	return m.Del(geoOfClass("SexUserLocation", "SexUserLocation", key)).Err()
}

func (m *_SexUserLocationRedisMgr) Clear() error {
	strs, err := m.Keys(geoOfClass("SexUserLocation", "SexUserLocation", "*")).Result()
	if err != nil {
		return err
	}
	if len(strs) > 0 {
		return m.Del(strs...).Err()
	}
	return nil
}

func (m *_SexUserLocationRedisMgr) Load(db DBFetcher) error {

	if err := m.Clear(); err != nil {
		return err
	}
	return m.AddBySQL(db, "SELECT `sex`,`longitude`,`latitude`,`id` FROM users")

}

func (m *_SexUserLocationRedisMgr) AddBySQL(db DBFetcher, sql string, args ...interface{}) error {
	objs, err := db.FetchBySQL(sql, args...)
	if err != nil {
		return err
	}

	for _, obj := range objs {
		if err := m.LocationAdd(obj.(*SexUserLocation)); err != nil {
			return err
		}
	}

	return nil
}
func (m *_SexUserLocationRedisMgr) DelBySQL(db DBFetcher, sql string, args ...interface{}) error {
	objs, err := db.FetchBySQL(sql, args...)
	if err != nil {
		return err
	}

	for _, obj := range objs {
		if err := m.LocationRem(obj.(*SexUserLocation)); err != nil {
			return err
		}
	}
	return nil
}

type _SexUserLocationMySQLMgr struct {
	*orm.MySQLStore
}

func SexUserLocationMySQLMgr() *_SexUserLocationMySQLMgr {
	return &_SexUserLocationMySQLMgr{_mysql_store}
}

func NewSexUserLocationMySQLMgr(cf *MySQLConfig) (*_SexUserLocationMySQLMgr, error) {
	store, err := orm.NewMySQLStore(cf.Host, cf.Port, cf.Database, cf.UserName, cf.Password)
	if err != nil {
		return nil, err
	}
	return &_SexUserLocationMySQLMgr{store}, nil
}

func (m *_SexUserLocationMySQLMgr) FetchBySQL(q string, args ...interface{}) (results []interface{}, err error) {
	rows, err := m.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("SexUserLocation fetch error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var result SexUserLocation
		err = rows.Scan(&(result.Key), &(result.Longitude), &(result.Latitude), &(result.Value))
		if err != nil {
			return nil, err
		}

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("SexUserLocation fetch result error: %v", err)
	}
	return
}
