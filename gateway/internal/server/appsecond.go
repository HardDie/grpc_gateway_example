package server

import (
	"context"

	appSecondClient "github.com/HardDie/grpc_gateway_example/appsecond/pkg/appsecond_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/HardDie/grpc_gateway_example/gateway/internal/service"
	pb "github.com/HardDie/grpc_gateway_example/gateway/pkg/gateway_service"
)

type AppSecond struct {
	pb.UnimplementedAppSecondServer
	service service.IAppSecond
}

func NewAppSecond(service service.IAppSecond) *AppSecond {
	return &AppSecond{
		service: service,
	}
}

func (s *AppSecond) RegisterGRPC(conn grpc.ServiceRegistrar) {
	pb.RegisterAppSecondServer(conn, s)
}
func (s *AppSecond) RegisterHTTP(ctx context.Context, mux *runtime.ServeMux) error {
	return pb.RegisterAppSecondHandlerServer(ctx, mux, s)
}

func (s *AppSecond) Hello(ctx context.Context, req *appSecondClient.HelloRequest) (*appSecondClient.HelloResponse, error) {
	return s.service.Hello(ctx, req)
}
