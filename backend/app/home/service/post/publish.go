// Date:   Wed Jun 18 23:25:22 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package post_service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/home/util"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/postpb"
)

func PublishPost(ginCtx *gin.Context) {
	var req postpb.PublishPostRequest
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H {})
		return
	}

	callback := func(ctx context.Context, client postpb.PostServiceClient) {
		resp, err2 := client.PublishPost(ctx, &req)
		if err2 != nil {
			ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
				"error": fmt.Sprintf("RPC error: %s", err2.Error()),
			})
			return
		}

		ginCtx.JSON(http.StatusOK, gin.H {
			"message": resp.Message,
			"postid": resp.PostId,
		})
	}

	err := util.WithRPCClient("localhost:8082", postpb.NewPostServiceClient, callback)

	if err != nil {
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": fmt.Sprintf("RPC error: %s", err.Error()),
		})
	}
}
