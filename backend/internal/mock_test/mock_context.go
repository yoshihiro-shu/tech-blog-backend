package mock_test

import (
	"testing"

	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/server/request"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/config"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/pkg/logger"
)

func NewContext(t *testing.T) *request.Context {
	conf := config.Configs{}
	db, _ := MockDB(t)
	logger := logger.New()
	cache := MockRedis(t)
	return request.NewContext(conf, logger, db, cache)
}
