// Date:   Sat Jul 05 14:22:14 2025
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

func getComments(ctx context.Context, client commentpb.CommentServiceClient, postID uint64) (any, error) {
	resp, err := client.GetComments(ctx, &commentpb.GetCommentsReqeust {
		PostId: postID,
	})
	if err != nil {
		return nil, err
	}
	return resp.Comments, nil
}

func GetComments(ginCtx *gin.Context) {
	postID, err := strconv.ParseUint(ginCtx.Param("post_id"), 10, 64)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H {
			"error": fmt.Sprintf("Invalid post_id: %s", ginCtx.Param("post_id")),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	commentsAny, err := rpcutil.WithRPCClient2(
		"comment-service",
		ctx,
		commentpb.NewCommentServiceClient,
		getComments,
		postID,
	)
	if err != nil {
		 log.Println(err.Error()) 
		 ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			 "error": err.Error(),
		 })
	}

	comments, ok := commentsAny.(*[]commentpb.Comment)
	if !ok {
		 log.Panicln("Unexpected commentsAny") 
		 ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			 "error": "未知错误，无法获取评论",
		 })
	}
	
	ginCtx.JSON(http.StatusOK, comments)
}
