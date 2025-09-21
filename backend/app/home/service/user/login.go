// Date:   Wed Jun 11 21:58:35 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package user_service

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/app/home/conf"
	"github.com/xiaoqixian/v2ex/backend/app/home/util"
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

	accessToken, err := util.GenerateToken(resp.User.Id, time.Duration(conf.JWT.AccExpTime) * time.Second)
	refreshToken, err2 := util.GenerateToken(resp.User.Id, time.Duration(conf.JWT.RefExpTime) * time.Second)

	if err != nil || err2 != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H {
			"message": "登录失败",
		})
		return
	}

	ginCtx.SetCookie(
		"access_token",
		accessToken,
		conf.JWT.AccExpTime,
		"/",
		"",
		false,
		true,
	)
	ginCtx.SetCookie(
		"refresh_token",
		refreshToken,
		conf.JWT.RefExpTime,
		"/",
		"",
		false,
		true,
	)

	ginCtx.JSON(http.StatusOK, gin.H {
		"message": "登录成功",
		"expires_in": conf.JWT.AccExpTime,
		"user": resp.User,
	})
}
