// Date:   Fri Jun 27 22:44:31 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package main

import (
	"log"
	"net"

	"github.com/xiaoqixian/v2ex/backend/app/comment/dal"
	"github.com/xiaoqixian/v2ex/backend/app/comment/service"
	"github.com/xiaoqixian/v2ex/backend/app/common/util"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/commentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	dal.Init()

	const addr = "localhost:8083"

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	util.RegisterService("comment-service", "comment-service", addr)

	grpcServer := grpc.NewServer()
	service, err := service.NewCommentService()
	if err != nil {
		panic(err)
	}
	commentpb.RegisterCommentServiceServer(grpcServer, service)

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
