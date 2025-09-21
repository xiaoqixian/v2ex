// Author: https://github.com/xiaoqixian

package rec_service

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/app/home/conf"
	"github.com/xiaoqixian/v2ex/backend/app/home/util"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/recpb"
)

func RecPosts(ginCtx *gin.Context) {
	var req recpb.RecRequest
	sizeStr := ginCtx.Query("size")
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		log.Printf("ShouldBindQuery error: %s", err.Error())
		ginCtx.JSON(http.StatusBadRequest, gin.H {
			"error": fmt.Sprintf("JSON bind error: %s", err.Error()),
		})
		return
	}
	req.Size = uint64(size)

	userid, _ := ginCtx.Get("userid")
	req.UserId = userid.(uint64)


	conf := conf.GetConf()
	respAny, err := rpcutil.NewBuilder(&req, recpb.NewRecServiceClient).
		WithService(conf.Consul.Rec).
		WithMethod("RecPosts").
		WithMsTimeout(conf.Rpc.RpcTimeout).
		Call()

	if err != nil {
		fmt.Printf("Error call rec service %s: %s\n", conf.Consul.Rec, err.Error())
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": err.Error(),
		})
		return
	}

	resp, ok := respAny.(*recpb.RecResponse)
	if !ok {
		log.Printf("[GetPost] Expect *recpb.RecResponse, got '%T'\n", respAny)
		ginCtx.JSON(http.StatusInternalServerError, gin.H {
			"error": "wrong rpc response type",
		})
		return
	}

	posts := util.GetPostsByIds(resp.PostIdList)

	if posts == nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H {
			"error": "Unable to get posts",
		})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H {
		"posts": posts,
	})
}
