// Date:   Sat Jul 05 16:51:05 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package comment_service

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/app/home/conf"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/commentpb"
)

func SubmitComment(ginCtx *gin.Context) {
	postid, err := strconv.ParseUint(ginCtx.Param("post_id"), 10, 64)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H {
			"error": fmt.Sprintf("Invalid post_id: %s", ginCtx.Param("post_id")),
		})
		return
	}

	userid, exists := ginCtx.Get("userid")
	if !exists {
		log.Panicln("userid not found in gin context, authentication should be done before submitting a comment")
		ginCtx.JSON(http.StatusUnauthorized, gin.H {
			"error": "unauthorized",
		})
		return
	}

	req := commentpb.AddCommentRequest {
		PostId: postid,
		UserId: userid.(uint64),
	}

	if err = ginCtx.ShouldBindJSON(&req); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})
		return
	}

	conf := conf.GetConf()
	respAny, err := rpcutil.NewBuilder(&req, commentpb.NewCommentServiceClient).
		WithService(conf.Consul.Comment).
		WithMethod("AddComment").
		WithMsTimeout(conf.Rpc.RpcTimeout).
		Call()
	if err != nil {
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": err.Error(),
		})
		return
	}
	
	resp, ok := respAny.(*commentpb.AddCommentResponse)
	if !ok {
		log.Printf("[SubmitComment] Expect *commentpb.GetCommentsResponse, got '%T'\n", respAny)
		ginCtx.JSON(http.StatusInternalServerError, gin.H {
			"error": "内部错误，稍后再试",
		})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H {
		"comment_id": resp.CommentId,
	})
}
