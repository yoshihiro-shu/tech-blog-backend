package server

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/yoshihiro-shu/draft-backend/interfaces/api/server/router"
	"github.com/yoshihiro-shu/draft-backend/internal/config"
	"github.com/yoshihiro-shu/draft-backend/internal/pkg/logger"
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
	logger logger.Logger
}

func New(conf config.Configs, logger logger.Logger) *Server {
	return &Server{
		Server: &http.Server{
			Addr:           conf.GetUserAddr(),
			Handler:        router.New(conf),
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		logger: logger,
	}
}

func (s Server) SetRouters() {
	s.Server.Handler.(*router.Router).Apply(s.logger)
}

func (srv Server) Start() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			srv.logger.Error("failed at listening server.", zap.Error(err))
		}
	}()

	srv.logger.Info(banner)

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
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	srv.logger.Info("shutting down")
	os.Exit(0)
}
