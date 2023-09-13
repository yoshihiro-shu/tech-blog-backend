package mock_test

import (
	"database/sql"
	"testing"

	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/api/request"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/config"
	"github.com/yoshihiro-shu/draft-backend/backend/internal/logger"
)

func NewContext(t *testing.T, sqlDB *sql.DB) (*request.Context, error) {
	db, err := MockDBClient(sqlDB)
	if err != nil {
		return nil, err
	}
	logger := logger.New()
	cache := MockRedis(t)
	return request.NewContext(config.Configs{}, logger, db, cache), err
}

func NewMinContext() *request.Context {
	return request.NewContext(
		config.Configs{},
		logger.New(),
		nil,
		nil,
	)
}
