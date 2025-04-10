package net

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"log"
	"microservices/api/proto/user"
	"microservices/internal/user/handler"
	"microservices/pkg/config"
	"microservices/pkg/etcd"
	"net"
)

func NewGRPCService() *grpc.Server {

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, handler.NewUserServiceHandler())
	return grpcServer
}

func StartGRPCServer(lc fx.Lifecycle, config *config.Config, grpcServer *grpc.Server) {
	// etcd 服务注册器

	registrar, err := etcd.NewEtcdRegistrar(
		[]string{config.Etcd.Address},
		config.Services.User.Name,
		config.Services.User.Port,
	)
	if err != nil {
		log.Fatalf("failed to create etcd registrar: %v", err)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				lis, err := net.Listen("tcp", config.Services.User.Port)
				if err != nil {
					log.Fatalf("failed to listen: %v", err)
				}

				// 注册到 etcd
				if err := registrar.Register(ctx); err != nil {
					log.Fatalf("failed to register service: %v", err)
				}

				log.Printf("starting gRPC server on %s", config.Services.User.Port)
				if err := grpcServer.Serve(lis); err != nil {
					log.Fatalf("failed to serve: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Server stopping...")
			return nil
		},
	})
}
