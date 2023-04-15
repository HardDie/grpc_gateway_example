package server

import (
	"context"

	appFirstClient "github.com/HardDie/grpc_gateway_example/appfirst/pkg/appfirst_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/HardDie/grpc_gateway_example/gateway/internal/service"
	pb "github.com/HardDie/grpc_gateway_example/gateway/pkg/gateway_service"
)

type AppFirst struct {
	pb.UnimplementedAppFirstServer
	service service.IAppFirst
}

func NewAppFirst(service service.IAppFirst) *AppFirst {
	return &AppFirst{
		service: service,
	}
}

func (s *AppFirst) RegisterGRPC(conn grpc.ServiceRegistrar) {
	pb.RegisterAppFirstServer(conn, s)
}
func (s *AppFirst) RegisterHTTP(ctx context.Context, mux *runtime.ServeMux) error {
	return pb.RegisterAppFirstHandlerServer(ctx, mux, s)
}

func (s *AppFirst) Hello(ctx context.Context, req *appFirstClient.HelloRequest) (*appFirstClient.HelloResponse, error) {
	return s.service.Hello(ctx, req)
}
