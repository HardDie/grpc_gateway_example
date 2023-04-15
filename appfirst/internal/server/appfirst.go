package server

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	pb "github.com/HardDie/grpc_gateway_example/appfirst/pkg/appfirst_service"
)

type AppFirst struct {
	pb.UnimplementedAppFirstServiceServer
	username string
}

func NewAppFirst() *AppFirst {
	return &AppFirst{
		username: "%username%",
	}
}

func (s *AppFirst) Register(conn grpc.ServiceRegistrar) {
	pb.RegisterAppFirstServiceServer(conn, s)
}

func (s *AppFirst) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Msg: fmt.Sprintf("Hello %s, from first app", s.username),
	}, nil
}
