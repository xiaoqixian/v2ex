// Date:   Wed Jun 11 12:50:00 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package main

import (
	"log"
	"net"

	"github.com/xiaoqixian/v2ex/backend/app/common/util"
	"github.com/xiaoqixian/v2ex/backend/app/post/dal"
	"github.com/xiaoqixian/v2ex/backend/app/post/service"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/postpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	dal.Init()

	const addr = "localhost:8082"

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	util.RegisterService("post-service", "post-service", addr)

	grpcServer := grpc.NewServer()
	service, err := service.NewPostService()
	if err != nil {
		panic(err)
	}
	postpb.RegisterPostServiceServer(grpcServer, service)

	// register grpc health check service
	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)

	log.Printf("Post service listening on %s\n", addr)
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
