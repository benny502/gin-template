package server

import "github.com/google/wire"

type Server interface {
	Run() error
}

var ProviderSet = wire.NewSet(NewHttpServer, NewGrpcServer)
