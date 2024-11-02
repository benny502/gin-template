package grpc

import (
	v1 "bookmark/api/hello/v1"
	"context"
)

type HelloService struct {
	v1.UnimplementedHelloServer
}

func (s *HelloService) SayHello(ctx context.Context, req *v1.HelloRequest) (*v1.HelloResponse, error) {
	return &v1.HelloResponse{
		Message: "hello " + req.Name,
	}, nil
}
func NewHelloService() *HelloService {
	return &HelloService{}
}
