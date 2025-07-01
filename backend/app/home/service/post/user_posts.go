// Date:   Wed Jun 18 20:48:11 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package post_service

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/postpb"
)

func GetPostsForUser(ginCtx *gin.Context) {
	userIDStr := ginCtx.Param("userid")
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

	callback := func(ctx context.Context, client postpb.PostServiceClient) error {
		resp, err2 := client.GetPostsForUser(ctx, &req)
		if err2 != nil {
			return err2
		}

		ginCtx.JSON(http.StatusOK, resp.Posts)
		return nil
	}

	err = rpcutil.WithRPCClient("localhost:8082", postpb.NewPostServiceClient, callback)

	if err != nil {
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": err.Error(),
		})
	}
}
