package model

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ezbuy/redis-orm/orm"
	"github.com/go-redis/redis/v8"
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
	redis.Pipeliner
	Err error
}

func (m *_SexUserLocationRedisMgr) BeginPipeline(pipes ...redis.Pipeliner) *_SexUserLocationRedisPipeline {
	if len(pipes) > 0 {
		return &_SexUserLocationRedisPipeline{pipes[0], nil}
	}
	return &_SexUserLocationRedisPipeline{m.Pipeline(), nil}
}

//! redis relation pair
func (m *_SexUserLocationRedisMgr) LocationAdd(relation *SexUserLocation) error {
	return m.GeoAdd(context.TODO(), geoOfClass("SexUserLocation", "SexUserLocation", relation.Key), &redis.GeoLocation{
		Longitude: relation.Longitude,
		Latitude:  relation.Latitude,
		Name:      fmt.Sprint(relation.Value),
	}).Err()
}

func (m *_SexUserLocationRedisMgr) LocationRadius(key string, longitude float64, latitude float64, query *redis.GeoRadiusQuery) ([]*SexUserLocation, error) {
	locations, err := m.GeoRadius(context.TODO(), geoOfClass("SexUserLocation", "SexUserLocation", key), longitude, latitude, query).Result()
	if err != nil {
		return nil, err
	}

	relations := []*SexUserLocation{}
	for _, location := range locations {
		relation := m.NewSexUserLocation(key)
		relation.Longitude = location.Longitude
		relation.Latitude = location.Latitude
		if err := orm.StringScan(location.Name, &relation.Value); err != nil {
			return nil, err
		}
		relations = append(relations, relation)
	}
	return relations, nil
}

func (m *_SexUserLocationRedisMgr) LocationRem(relation *SexUserLocation) error {
	return m.ZRem(context.TODO(), geoOfClass("SexUserLocation", "SexUserLocation", relation.Key), fmt.Sprint(relation.Value)).Err()
}

func (m *_SexUserLocationRedisMgr) LocationDel(key string) error {
	return m.Del(context.TODO(), geoOfClass("SexUserLocation", "SexUserLocation", key)).Err()
}

func (m *_SexUserLocationRedisMgr) Clear() error {
	ctx := context.TODO()
	strs, err := m.Keys(ctx, geoOfClass("SexUserLocation", "SexUserLocation", "*")).Result()
	if err != nil {
		return err
	}
	if len(strs) > 0 {
		return m.Del(ctx, strs...).Err()
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

type _SexUserLocationDBMgr struct {
	db orm.DB
}

func SexUserLocationDBMgr(db orm.DB) *_SexUserLocationDBMgr {
	if db == nil {
		panic(fmt.Errorf("SexUserLocationDBMgr init need db"))
	}
	return &_SexUserLocationDBMgr{db: db}
}

func (m *_SexUserLocationDBMgr) FetchBySQL(q string, args ...interface{}) (results []interface{}, err error) {
	rows, err := m.db.Query(q, args...)
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
