package mock_test

import (
	"database/sql"

	"github.com/yoshihiro-shu/tech-blog-backend/backend/interfaces/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type mockDB struct {
	*gorm.DB
}

func (m mockDB) Master() *gorm.DB { return m.DB }

func (m mockDB) Reprica() *gorm.DB { return m.DB }

func (m mockDB) Close() error { return nil }

func MockDB(sqlDB *sql.DB) (model.DBClient, error) {
	orm, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	return mockDB{orm}, err
}
