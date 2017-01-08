package orm

import (
	"fmt"

	redis "gopkg.in/redis.v5"
)

type RedisStore struct {
	redis.Cmdable
}

func NewRedisStore(host string, port int, password string, db int) (*RedisStore, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       db,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, err
	}

	return &RedisStore{client}, nil
}

//! functions
func (r *RedisStore) StringScan(str string, val interface{}) error {
	return redis.NewStringResult(str, nil).Scan(val)
}
