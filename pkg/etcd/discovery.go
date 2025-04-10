package etcd

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	etcdnaming "go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"log"
	"microservices/api/proto/user"
	"microservices/pkg/config"
)

func NewUserServiceClient(c *config.Config) user.UserServiceClient {
	conn := ConnectToService(c.Services.User.Name, c.Etcd.Address)
	return user.NewUserServiceClient(conn)
}

func ConnectToService(serviceName string, url string) *grpc.ClientConn {
	cli, _ := clientv3.NewFromURL(url)

	// 创建 resolver
	resolver, err := etcdnaming.NewBuilder(cli)
	if err != nil {
		log.Fatalf("resolver error: %v", err)
	}

	// 连接 gRPC 服务（注意要启用 round_robin 策略）
	conn, err := grpc.NewClient(
		serviceName,
		grpc.WithResolvers(resolver),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("grpc dial error: %v", err)
	}

	return conn
}
