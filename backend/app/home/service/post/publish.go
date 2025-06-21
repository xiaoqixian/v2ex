// Date:   Wed Jun 18 23:25:22 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package post_service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/postpb"
	"google.golang.org/grpc"
)

func PublishPost(ginCtx *gin.Context) {
	var req postpb.PublishPostRequest
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H {})
		return
	}

	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": fmt.Sprintf("RPC 连接建立超时: %s", err.Error()),
		})
		return
	}
	defer conn.Close()

	client := postpb.NewPostServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	
	resp, err := client.PublishPost(ctx, &req)
	if err != nil {
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": fmt.Sprintf("RPC call error: %s", err.Error()),
		})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H {
		"message": resp.Message,
		"postid": resp.PostId,
	})
}
