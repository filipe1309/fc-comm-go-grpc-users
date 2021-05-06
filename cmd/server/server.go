package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	gRPCServer := grpc.NewServer()

	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
