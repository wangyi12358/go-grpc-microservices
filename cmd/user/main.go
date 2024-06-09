package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"microservices/api/proto/user"
	"microservices/internal/user/handler"
	"microservices/pkg/config"
	"net"
)

func init() {
	config.Setup()
}

func main() {
	address := fmt.Sprintf(":%s", config.Config.Services.User.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, &handler.UserServiceHandler{})

	log.Printf("starting gRPC server on %s", config.Config.Services.User.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
