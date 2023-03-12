package cache

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v9"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/config"
)

type RedisClient interface {
	GET(key string, i interface{}) error
	SET(key string, i interface{}) error
}

type redisContext struct {
	cahceRedis *redis.Client
	ctx        context.Context
}

func New(c config.RedisCache) RedisClient {
	rds := redis.NewClient(&redis.Options{
		Addr:     c.GetRedisDNS(),
		Password: c.Password, // no password sret
		DB:       c.DbNumber, // use default DB
	})

	return &redisContext{
		cahceRedis: rds,
		ctx:        context.Background(),
	}
}

func (r redisContext) SET(key string, i interface{}) error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}
	err = r.cahceRedis.Set(r.ctx, key, b, 0).Err()
	return err
}

func (r redisContext) GET(key string, i interface{}) error {
	str, err := r.cahceRedis.Get(r.ctx, key).Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(str), i)
	return err
}

func IsNotExistKey(err error) bool {
	return err == redis.Nil
}
