package server

import (
	v1 "bookmark/api/hello/v1"
	"bookmark/internal/config"
	service "bookmark/internal/grpc"
	"bookmark/internal/pkg/log"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	conf   *config.Configuration
	server *grpc.Server
	logger log.Logger
}

func (s *GrpcServer) Run() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", s.conf.App.GrpcPort))
	if err != nil {
		return err
	}

	defer listen.Close()

	s.logger.Infof("grpc server listening on %s", s.conf.App.GrpcPort)
	return s.server.Serve(listen)
}

func (s *GrpcServer) Stop() error {
	s.server.Stop()
	return nil
}

func NewGrpcServer(conf *config.Configuration, hello *service.HelloService, logger log.Logger) *GrpcServer {

	server := grpc.NewServer()

	grpc := &GrpcServer{
		conf:   conf,
		server: server,
		logger: logger,
	}

	v1.RegisterHelloServer(server, hello)

	return grpc
}
