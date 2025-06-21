// Date:   Wed Jun 11 12:50:00 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package main

import (
	"log"
	"net"

	"github.com/xiaoqixian/v2ex/backend/app/user/dal"
	"github.com/xiaoqixian/v2ex/backend/app/user/service"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/postpb"
	"google.golang.org/grpc"
)

func main() {
	dal.Init()

	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	service, err := service.NewPostService()
	if err != nil {
		panic(err)
	}
	postpb.RegisterPostServiceServer(grpcServer, service)

	log.Println("Post service listening on :8082")
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
