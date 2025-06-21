// Date:   Wed Jun 11 21:58:35 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package user_service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/home/util"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/userpb"
)

type LoginArgs struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserLogin(ginCtx *gin.Context) {
	var args LoginArgs
	err := ginCtx.ShouldBindJSON(&args)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H {
			"error": "JSON 请求解析错误",
		})
		return
	}

	req := userpb.LoginRequest {
		Username: args.Username,
		Password: args.Password,
	}

	callback := func(ctx context.Context, client userpb.UserServiceClient) {
		resp, err2 := client.Login(ctx, &req)
		if err2 != nil {
			ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
				"error": fmt.Sprintf("RPC error: %s", err.Error()),
			})
			return
		}

		if !resp.Success {
			ginCtx.JSON(http.StatusBadRequest, gin.H {
				"error": resp.Message,
			})
			return
		}

		ginCtx.SetCookie(
			"access_token",
			resp.AccessToken,
			int(resp.AccessExpiresIn),
			"/",
			"localhost",
			false,
			true,
		)
		ginCtx.SetCookie(
			"refresh_token",
			resp.RefreshToken,
			int(resp.RefreshExpiresIn),
			"/",
			"localhost",
			false,
			true,
		)
		ginCtx.JSON(http.StatusOK, gin.H {
			"message": "注册成功",
			"expires_in": resp.AccessExpiresIn,
			"user": resp.User,
		})
	}

	err = util.WithRPCClient(":8081", userpb.NewUserServiceClient, callback)

	if err != nil {
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": fmt.Sprintf("RPC error: %s", err.Error()),
		})
	}
}
