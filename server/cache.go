package server

import (
	"auth-service/helpers"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func createCacheClient() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     helpers.GetEnvParam("REDIS_HOST", ""),
		Password: helpers.GetEnvParam("REDIS_PASS", ""),
		DB:       0,
	})
}

func closeCacheClient() {
	rdb.Close()
}

func (cache *Cache) Set(key string, value string, expiration time.Duration) {
	ctx := context.Background()

	rdb.Set(ctx, key, value, expiration)
}

func (cache *Cache) Get(key string) (string, error) {
	ctx := context.Background()
	val := rdb.Get(ctx, key)
	result, err := val.Result()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (cache *Cache) Remove(key string) {
	ctx := context.Background()
	rdb.Del(ctx, key)
}

func (cache *Cache) Ping() string {

	ctx := context.Background()

	status, err := rdb.Ping(ctx).Result()
	cacheStatus := ""
	if err != nil {
		cacheStatus = fmt.Sprintf("Cannot connect to cache server, error : %v", err.Error())
	} else {
		cacheStatus = fmt.Sprintf("Ping - %s", status)
	}

	return cacheStatus
}
