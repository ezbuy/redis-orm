package orm

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisStore struct {
	redis.Cmdable
}

func NewRedisClient(host string, port int, password string, db int) (*RedisStore, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       db,
	})
	if err := client.Ping(context.TODO()).Err(); err != nil {
		return nil, err
	}

	return &RedisStore{Cmdable: client}, nil
}

func NewRedisClusterClient(opt *redis.ClusterOptions) (*RedisStore, error) {
	client := redis.NewClusterClient(opt)
	if err := client.Ping(context.TODO()).Err(); err != nil {
		return nil, err
	}

	return &RedisStore{Cmdable: client}, nil
}

func NewRedisRingClient(opt *redis.RingOptions) (*RedisStore, error) {
	client := redis.NewRing(opt)
	if err := client.Ping(context.TODO()).Err(); err != nil {
		return nil, err
	}

	return &RedisStore{Cmdable: client}, nil
}

func NewRedisFailoverClient(failoverOpt *redis.FailoverOptions) (*RedisStore, error) {
	client := redis.NewFailoverClient(failoverOpt)
	if err := client.Ping(context.TODO()).Err(); err != nil {
		return nil, err
	}

	return &RedisStore{Cmdable: client}, nil
}
