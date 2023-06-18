package model

import (
	"github.com/yoshihiro-shu/draft-backend/backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func dbConnecter(conf config.DB) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(conf.GetDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectToPrimary(conf config.DB) (*gorm.DB, error) {
	return dbConnecter(conf)
}
