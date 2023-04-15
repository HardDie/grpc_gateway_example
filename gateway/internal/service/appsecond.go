package service

import (
	"context"

	appSecondClient "github.com/HardDie/grpc_gateway_example/appsecond/pkg/appsecond_service"
	"google.golang.org/grpc"
)

type IAppSecond interface {
	Hello(ctx context.Context, req *appSecondClient.HelloRequest) (*appSecondClient.HelloResponse, error)
}

type AppSecond struct {
	client appSecondClient.AppSecondServiceClient
}

func NewAppSecond(cc *grpc.ClientConn) *AppSecond {
	return &AppSecond{
		client: appSecondClient.NewAppSecondServiceClient(cc),
	}
}

func (s *AppSecond) Hello(ctx context.Context, req *appSecondClient.HelloRequest) (*appSecondClient.HelloResponse, error) {
	return s.client.Hello(ctx, req)
}
