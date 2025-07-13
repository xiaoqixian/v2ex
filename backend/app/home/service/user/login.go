// Date:   Wed Jun 11 21:58:35 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package user_service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/app/home/conf"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/userpb"
)

func UserLogin(ginCtx *gin.Context) {
	var req userpb.LoginRequest
	err := ginCtx.ShouldBindJSON(&req)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H {
			"error": "JSON 请求解析错误",
		})
		return
	}

	conf := conf.GetConf()

	respAny, err := rpcutil.NewBuilder(&req, userpb.NewUserServiceClient).
		WithService(conf.Consul.User).
		WithMethod("Login").
		WithMsTimeout(conf.Rpc.RpcTimeout).
		Call()
	if err != nil {
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": err.Error(),
		})
		return
	}

	resp, ok := respAny.(*userpb.LoginResponse)
	if !ok {
		log.Printf("respAny convert to *userpb.LoginResponse failed, actual type is '%T'\n", respAny)
		ginCtx.JSON(http.StatusInternalServerError, gin.H {
			"error": "内部错误，稍后再试",
		})
		return
	}

	if !resp.Success {
		ginCtx.JSON(http.StatusUnauthorized, gin.H {
			"error": resp.Message,
		})
		return
	}

	ginCtx.Set("userid", resp.User.Id)

	ginCtx.JSON(http.StatusOK, gin.H {
		"message": "登录成功",
		"user": resp.User,
	})
}
