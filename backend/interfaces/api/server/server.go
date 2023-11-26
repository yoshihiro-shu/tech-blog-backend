package server

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/yoshihiro-shu/tech-blog-backend/backend/infrastructure/persistence/cache"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/interfaces/api/router"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/interfaces/api/user_api"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/interfaces/model"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/internal/config"
	"github.com/yoshihiro-shu/tech-blog-backend/backend/internal/logger"
	"go.uber.org/zap"
)

const (
	banner = `
____________________________________O/_______
                                    O\
Server is Started
____________________________________O/_______
                                    O\
`
)

type Server struct {
	*http.Server
	conf   config.Configs
	logger logger.Logger
	db     model.DBClient
	cache  cache.RedisClient
}

func New(conf config.Configs, logger logger.Logger, db model.DBClient, cache cache.RedisClient) *Server {
	return &Server{
		Server: &http.Server{
			Addr:           conf.GetUserAddr(),
			Handler:        router.New(),
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		conf:   conf,
		logger: logger,
		db:     db,
		cache:  cache,
	}
}

func (s Server) SetRouters() {
	user_api.Apply(
		s.Handler.(router.Router),
		s.conf,
		s.logger,
		s.db,
		s.cache,
	)
}

func (s Server) Start() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := s.ListenAndServe(); err != nil {
			s.logger.Error("failed at listening server.", zap.Error(err))
		}
	}()

	s.logger.Info(banner)

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	s.Shutdown(ctx)
	// Optionally, you could run s.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	s.logger.Info("shutting down")
	os.Exit(0)
}
