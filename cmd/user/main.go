package main

import (
	"go.uber.org/fx"
	"microservices/internal/user/net"
	"microservices/pkg/config"
	"microservices/pkg/db"
)

func main() {
	fx.New(
		fx.Provide(config.New),
		fx.Provide(db.New),
		fx.Provide(net.NewGRPCService),
		fx.Invoke(net.StartGRPCServer),
	).Run()
}
