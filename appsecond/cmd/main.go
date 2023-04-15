package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/HardDie/grpc_gateway_example/appsecond/internal/server"
)

const (
	grpcPort = ":9002"
)

func main() {
	// Create a TCP connection
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal(err)
	}

	// Create the GRPC server
	grpcServer := grpc.NewServer()

	//// Allows us to use a 'list' call to list all available APIs
	//reflection.Register(grpcServer)

	// Create server object and register
	server.NewAppSecond().Register(grpcServer)

	// Serving the GRPC server on a created TCP socket
	log.Println("GRPC server listening on " + grpcPort)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
