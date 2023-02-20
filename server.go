package main

import (
	"log"
	"net"

	"github.com/jeferagudeloc/go-grpc-server-example/greeting"
	"google.golang.org/grpc"
)

func main() {
	list, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen port: %v", err)
	}

	s := greeting.Server{}

	grpcServer := grpc.NewServer()
	greeting.RegisterGreetingServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}

}
