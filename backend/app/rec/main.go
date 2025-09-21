// Date:   Sun Sep 14 07:36:43 PM 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package main

import (
	"log"
	"net"

	"github.com/xiaoqixian/v2ex/backend/app/common/util"
	"github.com/xiaoqixian/v2ex/backend/app/rec/dal"
	"github.com/xiaoqixian/v2ex/backend/app/rec/service"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/recpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	dal.Init()

	// const addr = "0.0.0.0:8087"
	addr := util.GetEnv("BINDADDR", "")
	if addr == "" {
		log.Fatalf("BINDADDR env var not found")
	}
	
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	util.ConsulRegisterService("rec-service", "rec-service", addr)

	grpcServer := grpc.NewServer()
	service, err := service.NewRecService()
	if err != nil {
		panic(err)
	}
	recpb.RegisterRecServiceServer(grpcServer, service)

	// register grpc health check service
	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)

	log.Printf("rec service listening on %s\n", addr)
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
