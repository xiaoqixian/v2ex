// Date:   Wed Jun 18 20:48:11 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package post_service

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/postpb"
	"google.golang.org/grpc"
)

func GetPosts(ginCtx *gin.Context) {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
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

	userIDStr := ginCtx.Query("userid")
	if userIDStr != "" {
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			ginCtx.JSON(http.StatusBadRequest, gin.H {
				"error": fmt.Sprintf("Invalid user id: %s", userIDStr),
			})
			return
		}

		req := postpb.GetPostsForUserRequest {
			UserId: uint64(userID),
		}
		resp, err := client.GetPostsForUser(ctx, &req)

		if err != nil {
			ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
				"error": fmt.Sprintf("RPC call error: %s", err.Error()),
			})
			return
		}

		ginCtx.JSON(http.StatusOK, resp.Posts)
		return
	}

	ginCtx.JSON(http.StatusBadRequest, gin.H {
		"error": "Empty user id not supported",
	})
}
