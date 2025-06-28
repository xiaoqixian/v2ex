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
)

func main() {
	dal.Init()

	const addr = "localhost:8081"

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	util.RegisterService("user-service", "user-service", addr)

	grpcServer := grpc.NewServer()
	service, err := service.NewUserService()
	if err != nil {
		panic(err)
	}
	userpb.RegisterUserServiceServer(grpcServer, service)

	log.Printf("user service listening on %s\n", addr)
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
