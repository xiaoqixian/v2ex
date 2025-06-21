// Date:   Sat Jun 21 20:33:10 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package post_service

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/home/util"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/postpb"
)

func GetPost(ginCtx *gin.Context) {
	postID, err := strconv.Atoi(ginCtx.Param("post_id"))
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H {
			"error": fmt.Sprintf("Invalid post_id: %s", ginCtx.Param("post_id")),
		})
		return
	}

	req := postpb.GetPostRequest {
		PostId: uint64(postID),
	}

	callback := func(ctx context.Context, client postpb.PostServiceClient) {
		resp, err2 := client.GetPost(ctx, &req)
		if err2 != nil {
			ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
				"error": fmt.Sprintf("RPC error: %s", err2.Error()),
			})
			return
		}

		ginCtx.JSON(http.StatusOK, resp)
	}

	err = util.WithRPCClient(":8082", postpb.NewPostServiceClient, callback)

	if err != nil {
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": fmt.Sprintf("RPC error: %s", err.Error()),
		})
	}
}
