package cache

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v9"
	"github.com/yoshihiro-shu/draft-backend/internal/config"
)

type RedisContext struct {
	redisClient *redis.Client
	ctx         context.Context
}

func New(c config.Configs) *RedisContext {
	rds := redis.NewClient(&redis.Options{
		Addr:     c.GetRedisDNS(),
		Password: c.GetCacheRedis().Password, // no password sret
		DB:       c.GetCacheRedis().DbNumber, // use default DB
	})

	return &RedisContext{
		redisClient: rds,
		ctx:         context.Background(),
	}
}

func (r RedisContext) SET(key string, i interface{}) error {
	b, _ := json.Marshal(i)
	err := r.redisClient.Set(r.ctx, key, b, 0).Err()
	return err
}

func (r RedisContext) GET(key string, i interface{}) error {
	str, err := r.redisClient.Get(r.ctx, key).Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(str), i)
	return err
}

func IsNotExistKey(err error) bool {
	return err == redis.Nil
}
