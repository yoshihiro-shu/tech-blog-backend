package model

import (
	"fmt"

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
	// TODO fix later
	return c.repricas[0]
}

func initMaster(conf config.DB) *pg.DB {
	return getDBConnection(conf)
}

func initRepicas(conf []config.DB) []*pg.DB {
	dbs := make([]*pg.DB, len(conf))
	for i, v := range conf {
		dbs[i] = getDBConnection(v)
	}
	return dbs
}

func getDBConnection(c config.DB) *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", c.Host, c.Port),
		User:     c.User,
		Password: c.Password,
		Database: c.Name,
	})
}

func New(conf config.Configs) *DBContext {
	return &DBContext{
		master:   initMaster(conf.MasterDB()),
		repricas: initRepicas(conf.RepricaDB()),
	}
}
