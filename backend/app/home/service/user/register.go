// Date:   Wed Jun 11 21:58:35 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package user_service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/app/home/conf"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/userpb"
)

func RegisterUser(ginCtx *gin.Context) {
	log.Println("New RegisterUser request")

	var req userpb.RegisterRequest
	err := ginCtx.ShouldBindJSON(&req)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H {
			"error": fmt.Sprintf("JSON bind error: %s", err.Error()),
		})
		return
	}

	conf := conf.GetConf()
	respAny, err := rpcutil.NewBuilder(&req, userpb.NewUserServiceClient).
		WithService(conf.Consul.User).
		WithMethod("Register").
		WithMsTimeout(conf.Rpc.RpcTimeout).
		Call()
	if err != nil {
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": err.Error(),
		})
		return
	}

	resp, ok := respAny.(*userpb.RegisterResponse)
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

	ginCtx.JSON(http.StatusOK, gin.H {
		"message": "注册成功",
	})
}
