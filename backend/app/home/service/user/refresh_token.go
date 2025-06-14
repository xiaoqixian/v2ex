// Date:   Sat Jun 14 11:59:34 2025
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

type RefreshTokenArgs struct {
	Token string `json:"refresh_token" binding:"required"`
}

func RefreshToken(gin_ctx *gin.Context) {
	log.Println("New RefreshToken request")

	var args RefreshTokenArgs
	if err := gin_ctx.ShouldBindJSON(&args); err != nil {
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
	
	req := userpb.RefreshTokenRequest {
		RefreshToken: args.Token,
	}
	resp, err := client.RefreshToken(ctx, &req)
	
	if err != nil {
		gin_ctx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": fmt.Sprintf("RPC 请求错误: %s", err.Error()),
		})
		return
	}

	if resp.AccessToken == "" {
		gin_ctx.JSON(http.StatusUnauthorized, gin.H {
			"error": "Refresh token expired",
		})
		return
	}

	gin_ctx.JSON(http.StatusOK, gin.H {
		"access_token": resp.AccessToken,
		"expires_in": resp.ExpiresIn,
		"token_type": resp.TokenType,
	})
}
