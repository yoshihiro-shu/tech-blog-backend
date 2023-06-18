package model

import (
	"crypto/rand"
	"math/big"

	"github.com/yoshihiro-shu/draft-backend/backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClient interface {
	Master() *gorm.DB
	Reprica() *gorm.DB
	Close() error
}

type context struct {
	master   *gorm.DB
	repricas []*gorm.DB
}

func (c context) Master() *gorm.DB {
	return c.master
}

func (c context) Reprica() *gorm.DB {
	numOfDB := big.NewInt(int64(len(c.repricas)))
	n, err := rand.Int(rand.Reader, numOfDB)
	if err != nil {
		return c.repricas[0]
	}
	return c.repricas[n.Int64()]
}

func (c context) Close() error {
	err := close(c.master)
	if err != nil {
		return nil
	}
	for i := range c.repricas {
		err = close(c.repricas[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func close(gormDB *gorm.DB) error {
	db, err := gormDB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

// TODO add ping
func connectToMaster(conf config.DB) (*gorm.DB, error) {
	return connenctToDB(conf)
}

// TODO add ping
func connectToRepricas(conf []config.DB) ([]*gorm.DB, error) {
	repricas := make([]*gorm.DB, len(conf))
	for i, v := range conf {
		db, err := connenctToDB(v)
		if err != nil {
			return nil, err
		}
		repricas[i] = db
	}
	return repricas, nil
}

func connenctToDB(conf config.DB) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(conf.GetDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func New(conf config.Configs) (DBClient, error) {
	master, err := connectToMaster(conf.MasterDB())
	if err != nil {
		return nil, err
	}
	repricas, err := connectToRepricas(conf.RepricaDB())
	if err != nil {
		return nil, err
	}
	return &context{
		master:   master,
		repricas: repricas,
	}, nil
}
