package model

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/internal/config"
)

type DBContext struct {
	master   *pg.DB
	repricas []*pg.DB
}

func (c DBContext) Master() *pg.DB {
	return c.master
}

func (c DBContext) Reprica() *pg.DB {
	numOfDB := big.NewInt(int64(len(c.repricas)))
	n, err := rand.Int(rand.Reader, numOfDB)
	if err != nil {
		return c.repricas[0]
	}
	return c.repricas[n.Int64()]
}

func (c DBContext) Close() {
	c.master.Close()

	for i := range c.repricas {
		c.repricas[i].Close()
	}
}

// TODO add ping
func connectToMaster(conf config.DB) *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		User:     conf.User,
		Password: conf.Password,
		Database: conf.Name,
	})
}

// TODO add ping
func connectToRepricas(conf []config.DB) []*pg.DB {
	repricas := make([]*pg.DB, len(conf))
	for i, v := range conf {
		repricas[i] = pg.Connect(&pg.Options{
			Addr:     fmt.Sprintf("%s:%s", v.Host, v.Port),
			User:     v.User,
			Password: v.Password,
			Database: v.Name,
		})
	}
	return repricas
}

func New(conf config.Configs) *DBContext {
	return &DBContext{
		master:   connectToMaster(conf.MasterDB()),
		repricas: connectToRepricas(conf.RepricaDB()),
	}
}
