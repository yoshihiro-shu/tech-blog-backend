package model

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/config"
)

type DBContext struct {
	PsqlDB *pg.DB
	// TODO
	// ReadDb  *[]sql.DB
	// WriteDb *[]sql.DB
}

func New(conf config.Configs) *DBContext {
	return &DBContext{
		PsqlDB: GetDBConnection(conf.GetDb()),
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
