// Date:   Wed Jul 09 23:29:33 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package rpcutil

import (
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"

	consul "github.com/hashicorp/consul/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
  consulClient *consul.Client
  once         sync.Once
)

// A for user data, C for client
type Builder[A any, C any] struct {
	service string
	ctx context.Context
	newClientFunc func(grpc.ClientConnInterface) C
	callback func(context.Context, C, A) (any, error)
	data A
	cancel context.CancelFunc
}

func NewBuilder[A any, C any](
	data A,
	newClientFunc func(grpc.ClientConnInterface) C, 
) *Builder[A, C] {
	return &Builder[A, C] {
		newClientFunc: newClientFunc,
		data: data,
	}
}

func (b *Builder[A, C]) WithCallback(
	callback func(context.Context, C, A) (any, error), 
) *Builder[A, C] {
	b.callback = callback
	return b
}

func (b *Builder[A, C]) WithCtx(ctx context.Context) *Builder[A, C] {
	b.ctx = ctx
	return b
}

func (b *Builder[A, C]) WithMsTimeout(d int) *Builder[A, C] {
	return b.WithTimeout(time.Duration(d) * time.Millisecond)
}

func (b *Builder[A, C]) WithTimeout(d time.Duration) *Builder[A, C] {
	if b.ctx == nil {
		b.ctx = context.Background()
	}
	b.ctx, b.cancel = context.WithTimeout(b.ctx, d)
	return b
}

func (b *Builder[A, C]) WithService(name string) *Builder[A, C] {
	b.service = name
	return b
}

func (b *Builder[A, C]) WithMethod(name string) *Builder[A, C] {
	b.callback = func(ctx context.Context, client C, req A) (any, error) {
		method := reflect.ValueOf(client).MethodByName(name)
		if !method.IsValid() {
			log.Panicf("method %s not found on client", name)
			return nil, fmt.Errorf("method %s not found on client", name)
		}

		in := []reflect.Value{
			reflect.ValueOf(ctx),
			reflect.ValueOf(req),
		}

		out := method.Call(in)

		if len(out) != 2 {
			return nil, fmt.Errorf("method %s must return (any, error)", name)
		}

		result := out[0].Interface()
		var err error
		if !out[1].IsNil() {
			err = out[1].Interface().(error)
		}

		return result, err
	}
	return b
}

func (b *Builder[A, C]) Call() (any, error) {
	if b.cancel != nil {
		defer b.cancel()
	}
	
	consulCli := getConsulClient()
	services, _, err := consulCli.Health().Service(b.service, "", true, nil)
	if err != nil {
		return nil, err
	}

	if len(services) == 0 {
		return nil, fmt.Errorf("service with name '%s' not found", b.service)
	}

	addr := fmt.Sprintf("%s:%d", services[0].Service.Address, services[0].Service.Port)
	
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	
	return b.callback(b.ctx, b.newClientFunc(conn), b.data)
}

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
