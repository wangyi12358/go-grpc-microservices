package main

import (
	"go.uber.org/fx"
	"microservices/internal/gateway/net"
	"microservices/internal/gateway/service"
	"microservices/pkg/config"
	"microservices/pkg/db"
	"microservices/pkg/etcd"
)

func main() {
	fx.New(
		fx.Provide(config.New),
		fx.Provide(db.New),
		fx.Provide(service.New, etcd.NewUserServiceClient),
		fx.Provide(net.NewHTTPServer),
		fx.Invoke(net.StartHTTPServer),
	).Run()
}
