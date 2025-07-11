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
	"github.com/xiaoqixian/v2ex/backend/app/home/conf"
	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/app/home/util"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/postpb"
)

func getPostCallback(
	ctx context.Context, 
	client postpb.PostServiceClient,
	req *postpb.GetPostRequest,
) (any, error) {
	resp, err := client.GetPost(ctx, req)
	if err != nil {
		return nil, err
	}

	if e, ok := resp.Result.(*postpb.GetPostResponse_Err); ok {
		return nil, fmt.Errorf("rpc GetPost failed: %s", e.Err.Message)
	}

	o, _ := resp.Result.(*postpb.GetPostResponse_Ok)
	return o.Ok, nil
}


func GetPost(ginCtx *gin.Context) {
	postid, err := strconv.ParseUint(ginCtx.Param("post_id"), 10, 64)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H {
			"error": fmt.Sprintf("Invalid post_id: %s", ginCtx.Param("post_id")),
		})
		return
	}

	var userid uint64
	accessToken, err := ginCtx.Cookie("access_token")

	if err == nil {
		userid, _ = util.ParseToken(accessToken)
	}

	req := postpb.GetPostRequest {
		PostId: postid,
		UserId: userid,
	}

	conf := conf.GetConf()
	resp, err := rpcutil.NewBuilder(&req, postpb.NewPostServiceClient).
		WithService(conf.Consul.Post).
		WithCallback(getPostCallback).
		WithMsTimeout(conf.Rpc.RpcTimeout).
		Call()
	if err != nil {
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": err.Error(),
		})
		return
	}
	
	ginCtx.JSON(http.StatusOK, resp)
}
