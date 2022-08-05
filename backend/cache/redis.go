package cache

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v9"
	"github.com/yoshi429/test/config"
)

type RedisContext struct {
	RedisClient *redis.Client
	ctx         context.Context
}

func New(c config.Configs) *RedisContext {
	rds := redis.NewClient(&redis.Options{
		Addr:     c.GetRedisDNS(),
		Password: c.GetCacheRedis().Password, // no password sret
		DB:       c.GetCacheRedis().DbNumber, // use default DB
	})

	return &RedisContext{
		RedisClient: rds,
		ctx:         context.Background(),
	}
}

func (r RedisContext) SET(key string, i interface{}) error {
	b, err := json.Marshal(i)
	err = r.RedisClient.Set(r.ctx, key, b, 0).Err()
	return err
}

func (r RedisContext) GET(key string, i interface{}) error {
	str, err := r.RedisClient.Get(r.ctx, key).Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(str), i)
	return err
}

func (r RedisContext) IsNotExistKey(err error) bool {
	if err == redis.Nil {
		return true
	}
	return false
}
