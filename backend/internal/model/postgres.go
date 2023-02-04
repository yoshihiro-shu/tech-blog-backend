package model

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/internal/config"
)

type DBContext struct {
	master *pg.DB
	// TODO
	// ReadDb  *[]sql.DB
	// WriteDb *[]sql.DB
}

func (c DBContext) Master() *pg.DB {
	return c.master
}

func New(conf config.Configs) *DBContext {
	return &DBContext{
		master: GetDBConnection(conf.GetDb()),
	}
}

func GetDBConnection(c config.DB) *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", c.Host, c.Port),
		User:     c.User,
		Password: c.Password,
		Database: c.Name,
	})
}
