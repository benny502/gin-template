package server

import (
	"bookmark/internal/config"
	"bookmark/internal/pkg/log"
	"bookmark/internal/router"
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	engine *gin.Engine
	router *router.Router
	conf   *config.Configuration
	logger log.Logger
	server *http.Server
}

func (s *HttpServer) Run() error {
	s.router.Register()

	s.logger.Infof("http server listening on %s", s.conf.App.Port)

	err := s.server.ListenAndServe()

	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *HttpServer) Stop() error {
	return s.server.Shutdown(context.Background())
}

func NewHttpServer(engine *gin.Engine, router *router.Router, conf *config.Configuration, logger log.Logger) *HttpServer {

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", conf.App.Port),
		Handler: engine,
	}

	return &HttpServer{
		server: server,
		engine: engine,
		router: router,
		conf:   conf,
		logger: logger,
	}
}
