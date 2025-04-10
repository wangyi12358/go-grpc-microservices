package etcd

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	etcdnaming "go.etcd.io/etcd/client/v3/naming/endpoints"
)

type Registrar struct {
	Client       *clientv3.Client
	ServiceName  string
	Port         string
	InstanceAddr string
	leaseID      clientv3.LeaseID
}

func NewEtcdRegistrar(endpoints []string, serviceName, port string) (*Registrar, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	return &Registrar{
		Client:      client,
		ServiceName: serviceName,
		Port:        port,
		// TODO: get the actual IP address of the instance
		InstanceAddr: "0.0.0.0" + port,
	}, nil
}

func (r *Registrar) Register(ctx context.Context) error {
	// 创建租约
	leaseResp, err := r.Client.Grant(ctx, 10)
	if err != nil {
		return err
	}
	r.leaseID = leaseResp.ID

	// 服务注册
	manager, err := etcdnaming.NewManager(r.Client, r.ServiceName)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s/%s", r.ServiceName, r.InstanceAddr)

	err = manager.AddEndpoint(ctx, key, etcdnaming.Endpoint{
		Addr: r.InstanceAddr,
	}, clientv3.WithLease(r.leaseID))
	if err != nil {
		return err
	}

	// 启动 keepalive
	go func() {
		ch, err := r.Client.KeepAlive(ctx, r.leaseID)
		if err != nil {
			log.Printf("etcd keepalive error: %v", err)
			return
		}
		for {
			select {
			case <-ctx.Done():
				return
			case _, ok := <-ch:
				if !ok {
					log.Println("etcd keepalive channel closed")
					return
				}
			}
		}
	}()
	log.Printf("✅ Registered service %s -> %s to etcd", r.ServiceName, r.InstanceAddr)
	return nil
}

func (r *Registrar) Unregister(ctx context.Context) error {
	key := fmt.Sprintf("%s/%s", r.ServiceName, r.InstanceAddr)
	manager, err := etcdnaming.NewManager(r.Client, r.ServiceName)
	if err != nil {
		return err
	}
	err = manager.DeleteEndpoint(ctx, key)
	if err != nil {
		return err
	}
	log.Printf("❌ Unregistered service %s -> %s from etcd", r.ServiceName, r.InstanceAddr)
	return nil
}
