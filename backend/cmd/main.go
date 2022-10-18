package main

import (
	_ "github.com/lib/pq"
	"github.com/yoshihiro-shu/draft-backend/config"
	"github.com/yoshihiro-shu/draft-backend/server"
)

func main() {
	conf := config.New()

	s := server.New(conf)

	s.SetRouters()

	s.Start()
}
