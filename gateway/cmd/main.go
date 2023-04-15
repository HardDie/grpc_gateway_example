package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/HardDie/grpc_gateway_example/gateway/internal/server"
	"github.com/HardDie/grpc_gateway_example/gateway/internal/service"
)

const (
	httpPort = ":9003"
)

func main() {
	ctx := context.Background()

	// Create a TCP connection
	lis, err := net.Listen("tcp", httpPort)
	if err != nil {
		log.Fatal(err)
	}

	// Create grpc mux
	grpcMux := runtime.NewServeMux()

	// Create http mux and register grpc mux
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	// Connect to appFirst
	log.Println("Connecting to appFirst...")
	appFirstConn, err := grpc.DialContext(ctx, "localhost:9001", grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connecting to appFirst:", err.Error())
	}
	log.Println("appFirst connected!")

	// Connect to appSecond
	log.Println("Connecting to appSecond...")
	appSecondConn, err := grpc.DialContext(ctx, "localhost:9002", grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connecting to appSecond:", err.Error())
	}
	log.Println("appSecond connected!")

	// Create services
	firstService := service.NewAppFirst(appFirstConn)
	secondService := service.NewAppSecond(appSecondConn)

	// Create servers
	firstServer := server.NewAppFirst(firstService)
	secondServer := server.NewAppSecond(secondService)

	// Register servers on grpc mux
	err = firstServer.RegisterHTTP(ctx, grpcMux)
	if err != nil {
		log.Fatal("error register http handler:", err.Error())
	}
	err = secondServer.RegisterHTTP(ctx, grpcMux)
	if err != nil {
		log.Fatal("error register http handler:", err.Error())
	}

	// Serving the HTTP server on a created TCP socket
	log.Println("HTTP server listening on " + httpPort)
	err = http.Serve(lis, mux)
	if err != nil {
		log.Fatal(err)
	}
}
