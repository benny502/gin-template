package main

import (
	"bookmark/internal/config"
	"bookmark/internal/pkg/log"
	"bookmark/internal/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

type App struct {
	engine *gin.Engine
	router *router.Router
	conf   *config.Configuration
	logger log.Logger
}

func (a *App) Run(configPath string) {

	a.router.Register()

	a.engine.Run(fmt.Sprintf(":%s", a.conf.App.Port))
}

func NewApp(r *gin.Engine, conf *config.Configuration, router *router.Router, logger log.Logger) *App {

	return &App{
		engine: r,
		conf:   conf,
		router: router,
		logger: logger,
	}
}
