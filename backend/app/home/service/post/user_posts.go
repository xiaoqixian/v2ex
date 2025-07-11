// Date:   Wed Jun 18 20:48:11 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package post_service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/app/home/conf"
	user_service "github.com/xiaoqixian/v2ex/backend/app/home/service/user"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/postpb"
)

func GetPostsForUser(ginCtx *gin.Context) {
	accessToken, err := ginCtx.Cookie("access_token")
	if err != nil {
		ginCtx.JSON(http.StatusUnauthorized, gin.H{
			"error": "登录信息已失效，请重新登录",
		})
		return
	}

	userid, err := user_service.CheckUser(accessToken)
	if err != nil {
		ginCtx.JSON(http.StatusUnauthorized, gin.H{
			"error": "登录信息已失效，请重新登录",
		})
		return
	}

	req := postpb.GetPostsForUserRequest {
		UserId: userid,
	}

	conf := conf.GetConf()
	respAny, err := rpcutil.NewBuilder(&req, postpb.NewPostServiceClient).
		WithService(conf.Consul.Post).
		WithMethod("GetUserPosts").
		WithMsTimeout(conf.Rpc.RpcTimeout).
		Call()
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, ok := respAny.(*postpb.GetPostsForUserResponse)
	if !ok {
		log.Printf("[GetPostsForUser] Expect *postpb.GetUserPostsResponse, got '%T'\n", respAny)
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": "内部错误，稍后再试",
		})
		return
	}

	ginCtx.JSON(http.StatusOK, resp.Posts)
}
