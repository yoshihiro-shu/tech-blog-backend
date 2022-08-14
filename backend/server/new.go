package server

import (
	"log"
	"net/http"
	"time"

	"github.com/yoshihiro-shu/draft-backend/config"
	"github.com/yoshihiro-shu/draft-backend/router"
)

type Server struct {
	*http.Server
}

func New(conf config.Configs) *Server {
	r := router.New(conf)

	r.ApplyRouters()

	return &Server{
		Server: &http.Server{
			Addr:           conf.GetUserAddr(),
			Handler:        r,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s Server) Start() {
	log.Fatalln(s.ListenAndServe())
}
