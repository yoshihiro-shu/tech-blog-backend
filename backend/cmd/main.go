package main

import (
	_ "github.com/lib/pq"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server"
	"github.com/yoshihiro-shu/draft-backend/internal/config"
)

func main() {
	conf := config.New()

	s := server.New(conf)

	s.SetRouters()

	s.Start()
}
