package server

import (
	"bookmark/internal/config"
	"bookmark/internal/pkg/log"
	"bookmark/internal/router"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	engine *gin.Engine
	router *router.Router
	conf   *config.Configuration
	logger log.Logger
}

func (s *HttpServer) Run() error {
	s.router.Register()

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", s.conf.App.Port),
		Handler: s.engine,
	}
	s.logger.Infof("http server listening on %s", s.conf.App.Port)
	return server.ListenAndServe()
}

func NewHttpServer(engine *gin.Engine, router *router.Router, conf *config.Configuration, logger log.Logger) *HttpServer {
	engine.Use()
	return &HttpServer{
		engine: engine,
		router: router,
		conf:   conf,
		logger: logger,
	}
}
