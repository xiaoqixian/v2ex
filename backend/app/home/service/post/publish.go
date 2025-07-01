// Date:   Wed Jun 18 23:25:22 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package post_service

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/postpb"
)

func PublishPost(ginCtx *gin.Context) {
	var req postpb.PublishPostRequest
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H {})
		return
	}

	callback := func(ctx context.Context, client postpb.PostServiceClient) error {
		resp, err2 := client.PublishPost(ctx, &req)
		if err2 != nil {
			return err2
		}

		ginCtx.JSON(http.StatusOK, gin.H {
			"message": resp.Message,
			"postid": resp.PostId,
		})
		return nil
	}

	err := rpcutil.WithRPCClient("post-service", postpb.NewPostServiceClient, callback)

	if err != nil {
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": err.Error(),
		})
	}
}
