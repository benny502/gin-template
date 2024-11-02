package main

import (
	"bookmark/internal/config"
	"bookmark/internal/pkg/log"
	"bookmark/internal/server"
	"runtime"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

type App struct {
	conf    *config.Configuration
	logger  log.Logger
	servers []server.Server
}

func (a *App) Run() {

	for _, srv := range a.servers {
		g.Go(func() error {
			defer func() {
				if err := recover(); err != nil {
					buf := make([]byte, 64<<10) //nolint:gomnd
					n := runtime.Stack(buf, false)
					buf = buf[:n]
					a.logger.Errorf("%v:\n%s\n", err, buf)
					panic(err)
				}
			}()
			return srv.Run()
		})
	}
	if err := g.Wait(); err != nil {
		a.logger.Error(err)
	}
}

func (a *App) Register(server server.Server) {
	a.servers = append(a.servers, server)
}

func NewApp(httpServer *server.HttpServer, grpcServer *server.GrpcServer, conf *config.Configuration, logger log.Logger) *App {

	app := &App{
		conf:   conf,
		logger: logger,
	}

	app.Register(httpServer)
	app.Register(grpcServer)

	return app
}
