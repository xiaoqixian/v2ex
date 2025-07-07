// Date:   Sat Jul 05 16:51:05 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package comment_service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/commentpb"
)

func submitComment(ctx context.Context, client commentpb.CommentServiceClient, req *commentpb.AddCommentRequest) (any, error) {
	return client.AddComment(ctx, req)
}

func SubmitComment(ginCtx *gin.Context) {
	postID, err := strconv.ParseUint(ginCtx.Param("post_id"), 10, 64)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H {
			"error": fmt.Sprintf("Invalid post_id: %s", ginCtx.Param("post_id")),
		})
		return
	}

	req := commentpb.AddCommentRequest {
		PostId: postID,
	}

	if err = ginCtx.ShouldBindJSON(&req); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	respAny, err := rpcutil.WithRPCClient2(
		"comment-service",
		ctx,
		commentpb.NewCommentServiceClient,
		submitComment,
		&req,
	)
	if err != nil {
		log.Println(err.Error()) 
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": err.Error(),
		})
		return
	}
	resp := respAny.(*commentpb.AddCommentResponse)
	if !resp.Success {
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": "未知错误",
		})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H {
		"comment_id": resp.CommentId,
	})
}
