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
