// Date:   Sat Jun 21 20:33:10 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package post_service

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/app/home/conf"
	"github.com/xiaoqixian/v2ex/backend/app/home/util"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/postpb"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/userpb"
)

func GetPost(ginCtx *gin.Context) {
	postid, err := strconv.ParseUint(ginCtx.Param("post_id"), 10, 64)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H {
			"error": fmt.Sprintf("Invalid post_id: %s", ginCtx.Param("post_id")),
		})
		return
	}

	req := postpb.GetPostRequest {
		PostId: postid,
	}

	conf := conf.GetConf()
	respAny, err := rpcutil.NewBuilder(&req, postpb.NewPostServiceClient).
		WithService(conf.Consul.Post).
		WithMethod("GetPost").
		WithMsTimeout(conf.Rpc.RpcTimeout).
		Call()
	if err != nil {
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": err.Error(),
		})
		return
	}

	resp, ok := respAny.(*postpb.GetPostResponse)
	if !ok {
		log.Printf("[GetPost] Expect *postpb.GetPostResponse, got '%T'\n", respAny)
		ginCtx.JSON(http.StatusInternalServerError, gin.H {
			"error": "wrong rpc response type",
		})
		return
	}

	if !resp.Found {
		ginCtx.JSON(http.StatusNotFound, gin.H {
			"error": fmt.Sprintf("Post %d not found", postid),
		})
		return
	}

	var author, avatar string
	{
		var wg sync.WaitGroup
		var err1 error
		wg.Add(1)
		go func() {
			defer wg.Done()
			err1 = util.GetUserInfo(resp.AuthorId, func(resp *userpb.GetUserInfoResponse) {
				author = resp.Name
				avatar = resp.Avatar
			})
		}()

		wg.Wait()
		
		if err1 != nil {
			ginCtx.JSON(http.StatusInternalServerError, gin.H {
				"error": err1.Error(),
			})
			return
		}
	}
	
	ginCtx.JSON(http.StatusOK, gin.H {
		"author": author,
		"avatar": avatar,
		"title": resp.Title,
		"created_at": resp.CreatedAt,
		"content": resp.Content,
	})
}
