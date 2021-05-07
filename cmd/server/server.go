package main

import (
	"log"
	"net"

	"github.com/codeedu/fc2-grpc/pb"
	"github.com/codeedu/fc2-grpc/services"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	gRPCServer := grpc.NewServer()
	pb.RegisterUserServiceServer(gRPCServer, services.NewUserService())

	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
