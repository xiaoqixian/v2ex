// Date:   Wed Jun 11 21:58:35 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package user_service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/userpb"
	"google.golang.org/grpc"
)

type LoginArgs struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserLogin(gin_ctx *gin.Context) {
	log.Println("New RegisterUser request")

	var args LoginArgs
	err := gin_ctx.ShouldBindJSON(&args)
	if err != nil {
		gin_ctx.JSON(http.StatusBadRequest, gin.H {
			"error": "JSON 请求解析错误",
		})
		return
	}

	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		gin_ctx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": fmt.Sprintf("RPC 连接建立超时: %s", err.Error()),
		})
		return
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	req := userpb.LoginRequest {
		Username: args.Username,
		Password: args.Password,
	}

	resp, err := client.Login(ctx, &req)
	if err != nil {
		gin_ctx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": fmt.Sprintf("RPC 请求错误: %s", err.Error()),
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
		"access_token": resp.AccessToken,
		"refresh_token": resp.RefreshToken,
		"expires_in": resp.ExpiresIn,
		"token_type": resp.TokenType,
	})
}
