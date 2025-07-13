// Date:   Wed Jun 18 23:25:22 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package post_service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/app/home/conf"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/postpb"
)

func PublishPost(ginCtx *gin.Context) {
	var req postpb.PublishPostRequest
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})
		return
	}

	userid, exists := ginCtx.Get("userid")
	if !exists {
		ginCtx.JSON(http.StatusUnauthorized, gin.H {
			"error": "登录信息已失效，请重新登录",
		})
		return
	}
	req.UserId = userid.(uint64)

	conf := conf.GetConf()
	respAny, err := rpcutil.NewBuilder(&req, postpb.NewPostServiceClient).
		WithService(conf.Consul.Post).
		WithMethod("PublishPost").
		WithMsTimeout(conf.Rpc.RpcTimeout).
		Call()
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
		})
		return
	}

	resp, ok := respAny.(*postpb.PublishPostResponse)
	if !ok {
		log.Printf("[PublishPost] Expect *postpb.PublishPostResponse, got '%T'\n", respAny)
		ginCtx.JSON(http.StatusInternalServerError, gin.H {
			"error": "内部错误，稍后再试",
		})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H {
		"message": resp.Message,
		"postid": resp.PostId,
	})
}
