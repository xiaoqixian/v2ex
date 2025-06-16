// Date:   Sat Jun 14 22:16:09 2025
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

func AuthMe(ginCtx *gin.Context) {
	token, err := ginCtx.Cookie("access_token")
	if err != nil {
		ginCtx.JSON(http.StatusUnauthorized, gin.H {})
		log.Printf("Get 'access_token' cookie error: %s\n", err.Error())
		return
	}

	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Printf("grpc.Dial: %s\n", err.Error())
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": fmt.Sprintf("RPC 连接建立超时: %s", err.Error()),
		})
		return
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	
	req := userpb.AuthMeRequest {
		AccessToken: token,
	}
	resp, err := client.AuthMe(ctx, &req)
	
	if err != nil || !resp.Success {
		if err != nil {
			log.Printf("%s\n", err.Error())
		}
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H {})
}
