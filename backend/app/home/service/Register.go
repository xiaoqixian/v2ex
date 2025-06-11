// Date:   Wed Jun 11 21:58:35 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package service

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/userpb"
	"google.golang.org/grpc"
)

type RegisterArgs struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func RegisterUser(gin_ctx *gin.Context) {

	var args RegisterArgs
	err := gin_ctx.ShouldBindJSON(&args)
	if err != nil {
		gin_ctx.JSON(http.StatusBadRequest, gin.H {
			"error": "服务暂不可用，请稍后再试",
		})
		return
	}

	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		gin_ctx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": "服务暂不可用，请稍后再试",
		})
		return
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	req := userpb.RegisterRequest {
		Username: args.Username,
		Email: args.Email,
		Password: args.Password,
	}

	resp, err := client.Register(ctx, &req)
	if err != nil {
		gin_ctx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": "服务暂不可用，请稍后再试",
		})
		return
	}
	if !resp.Success {
		gin_ctx.JSON(http.StatusBadRequest, gin.H {
			"error": resp.Message,
		})
		return
	}
	
	gin_ctx.JSON(http.StatusOK, gin.H {
		"message": "注册成功",
	})
}
