package service

import (
	"context"

	appFirstClient "github.com/HardDie/grpc_gateway_example/appfirst/pkg/appfirst_service"
	"google.golang.org/grpc"
)

type IAppFirst interface {
	Hello(ctx context.Context, req *appFirstClient.HelloRequest) (*appFirstClient.HelloResponse, error)
}

type AppFirst struct {
	client appFirstClient.AppFirstServiceClient
}

func NewAppFirst(cc *grpc.ClientConn) *AppFirst {
	return &AppFirst{
		client: appFirstClient.NewAppFirstServiceClient(cc),
	}
}

func (s *AppFirst) Hello(ctx context.Context, req *appFirstClient.HelloRequest) (*appFirstClient.HelloResponse, error) {
	return s.client.Hello(ctx, req)
}
