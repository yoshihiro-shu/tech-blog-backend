package main

import (
	_ "github.com/lib/pq"
	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server"
	"github.com/yoshihiro-shu/draft-backend/internal/config"
	"github.com/yoshihiro-shu/draft-backend/internal/pkg/logger"
)

func main() {
	conf := config.New()
	logger := logger.New()

	s := server.New(conf, logger)

	s.SetRouters()

	s.Start()
}
