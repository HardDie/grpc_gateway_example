package server

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	pb "github.com/HardDie/grpc_gateway_example/appsecond/pkg/appsecond_service"
)

type AppSecond struct {
	pb.UnimplementedAppSecondServiceServer
	username string
}

func NewAppSecond() *AppSecond {
	return &AppSecond{
		username: "%username%",
	}
}

func (s *AppSecond) Register(conn grpc.ServiceRegistrar) {
	pb.RegisterAppSecondServiceServer(conn, s)
}

func (s *AppSecond) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Msg: fmt.Sprintf("Hello %s, from second app", s.username),
	}, nil
}
