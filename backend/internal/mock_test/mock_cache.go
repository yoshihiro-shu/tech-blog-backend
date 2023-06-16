package mock_test

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/cache"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/config"
)

func MockRedis(t *testing.T) cache.RedisClient {
	redisServer := miniredis.RunT(t)
	return cache.New(config.RedisCache{
		Host: redisServer.Host(),
		Port: redisServer.Port(),
	})
}
