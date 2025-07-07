// Date:   Sat Jun 28 11:25:42 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package rpcutil

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	consul "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
  consulClient *consul.Client
  once         sync.Once
)

func getConsulClient() *consul.Client {
  once.Do(func() {
    config := consul.DefaultConfig()
    config.Address = "127.0.0.1:8500"

    var err error
    consulClient, err = consul.NewClient(config)
    if err != nil {
      log.Fatalf("failed to create consul client: %v", err)
    }
  })
  return consulClient
}

func WithRPCClient[T any](
	serviceName string,
	newClient func(conn grpc.ClientConnInterface) T,
	callback func(ctx context.Context, client T) error,
) error {
	cli := getConsulClient()
	services, _, err := cli.Health().Service(serviceName, "", true, nil)
	
	if err != nil {
		return err
	}

	if len(services) == 0 {
		return fmt.Errorf("services for '%s' not found", serviceName)
	}
	
	addr := fmt.Sprintf("%s:%d", services[0].Service.Address, services[0].Service.Port)

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	return callback(ctx, newClient(conn))
}

func WithRPCClient2[T any, A any](
	serviceName string,
	ctx context.Context,
	newClient func(conn grpc.ClientConnInterface) T,
	callback func(ctx context.Context, client T, data A) (any, error),
	data A,
) (any, error) {
	cli := getConsulClient()
	services, _, err := cli.Health().Service(serviceName, "", true, nil)
	
	if err != nil {
		return nil, err
	}

	if len(services) == 0 {
		return nil, fmt.Errorf("services for '%s' not found", serviceName)
	}
	
	addr := fmt.Sprintf("%s:%d", services[0].Service.Address, services[0].Service.Port)

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return callback(ctx, newClient(conn), data)
}
