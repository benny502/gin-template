package server

import "github.com/google/wire"

type Server interface {
	Run() error
	Stop() error
}

var ProviderSet = wire.NewSet(NewHttpServer, NewGrpcServer)
