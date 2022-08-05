package main

import (
	_ "github.com/lib/pq"
	"github.com/yoshi429/draft-backend/config"
	"github.com/yoshi429/draft-backend/server"
)

func main() {
	conf := config.New()

	s := server.New(conf)

	s.Start()
}
