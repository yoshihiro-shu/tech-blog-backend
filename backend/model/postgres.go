package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/yoshi429/test/config"
)

type DBContext struct {
	PSQLDB *sql.DB
	// TODO
	// ReadDb  *[]sql.DB
	// WriteDb *[]sql.DB
}

func New(conf config.Configs) *DBContext {
	db, err := GetDBConnection(conf.GetDb())
	if err != nil {
		log.Fatalf("Failed Connect with PostgresDB. err: %s", err.Error())
	}
	return &DBContext{
		PSQLDB: db,
	}
}

func GetDBConnection(c config.DB) (*sql.DB, error) {
	db, err := sql.Open(c.Driver, getDbDNS(c))
	if err != nil {
		log.Fatalf("Fatal get Connection with DB, err=%s\n", err.Error())
		return nil, err
	}
	return db, nil
}

func getDbDNS(c config.DB) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Name, c.Password, c.Sslmode)
}
