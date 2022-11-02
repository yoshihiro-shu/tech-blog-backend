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
	return &Server{
		Server: &http.Server{
			Addr:           conf.GetUserAddr(),
			Handler:        router.New(conf),
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s Server) SetRouters() {
	s.Server.Handler.(*router.Router).ApplyRouters()
}

func (s Server) Start() {
	log.Fatalln(s.ListenAndServe())
}
