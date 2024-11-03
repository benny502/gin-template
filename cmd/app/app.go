package main

import (
	"bookmark/internal/config"
	"bookmark/internal/pkg/log"
	"bookmark/internal/server"
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type App struct {
	conf    *config.Configuration
	logger  log.Logger
	servers []server.Server
}

func (a *App) Run() error {
	g, ctx := errgroup.WithContext(context.Background())

	for _, srv := range a.servers {
		g.Go(func() error {
			return srv.Run()
		})
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	select {
	case <-ctx.Done():
		a.logger.Info("canceled or expired")
	case sig := <-sigChan:
		a.logger.Infof("signal: %s", sig)
		for _, srv := range a.servers {
			srv.Stop()
		}
	}

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		a.logger.Error(err)
		return err
	}
	return nil
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
