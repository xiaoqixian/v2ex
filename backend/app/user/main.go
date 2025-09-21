// Date:   Wed Jun 11 12:50:00 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package main

import (
	"log"
	"net"

	"github.com/xiaoqixian/v2ex/backend/app/common/util"
	"github.com/xiaoqixian/v2ex/backend/app/user/dal"
	"github.com/xiaoqixian/v2ex/backend/app/user/service"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	dal.Init()

	// const addr = "0.0.0.0:8081"
	addr := util.GetEnv("BINDADDR", "")
	if addr == "" {
		log.Fatalf("BINDADDR env var not found")
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	util.ConsulRegisterService("user-service", "user-service", addr)

	grpcServer := grpc.NewServer()
	service, err := service.NewUserService()
	if err != nil {
		panic(err)
	}
	userpb.RegisterUserServiceServer(grpcServer, service)

	// register grpc health check service
	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)

	log.Printf("user service listening on %s\n", addr)
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
