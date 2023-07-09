package mock_test

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/yoshihiro-shu/draft-backend/backend/interfaces/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type mockDB struct {
	*gorm.DB
}

func (m mockDB) Master() *gorm.DB { return m.DB }

func (m mockDB) Reprica() *gorm.DB { return m.DB }

func (m mockDB) Close() error { return nil }

func MockDBClient(sqlDB *sql.DB) (model.DBClient, error) {
	orm, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	return mockDB{orm}, err
}

func MockGorm() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	orm, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	return orm, mock, nil
}
