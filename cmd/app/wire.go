//go:generate go run github.com/google/wire/cmd/wire
//go:build wireinject
// +build wireinject

package main

import (
	"bookmark/internal/biz"
	"bookmark/internal/config"
	"bookmark/internal/data"
	"bookmark/internal/grpc"
	"bookmark/internal/middleware"
	"bookmark/internal/pkg/log"
	"bookmark/internal/router"
	"bookmark/internal/server"
	"bookmark/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func wireApp(conf *config.Configuration, opts ...gin.OptionFunc) (*App, func(), error) {
	panic(wire.Build(gin.Default, NewApp, log.NewLogger, router.NewRouter, server.ProviderSet, middleware.ProviderSet, service.ProviderSet, biz.ProviderSet, data.ProviderSet, grpc.ProviderSet))
}
