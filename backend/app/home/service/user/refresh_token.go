// Date:   Sat Jun 14 11:59:34 2025
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

type RefreshTokenArgs struct {
	Token string `json:"refresh_token" binding:"required"`
}

func RefreshToken(ginCtx *gin.Context) {
	var args RefreshTokenArgs
	if err := ginCtx.ShouldBindJSON(&args); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H {
			"error": fmt.Sprintf("JSON bind error: %s", err.Error()),
		})
		return
	}

	req := userpb.RefreshTokenRequest {
		RefreshToken: args.Token,
	}

	callback := func(ctx context.Context, client userpb.UserServiceClient) {
		resp, err2 := client.RefreshToken(ctx, &req)
		if err2 != nil {
			ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
				"error": fmt.Sprintf("RPC error: %s", err2.Error()),
			})
			return
		}
		if resp.AccessToken == "" {
			ginCtx.JSON(http.StatusUnauthorized, gin.H {
				"error": "Refresh token expired",
			})
			return
		}

		ginCtx.JSON(http.StatusOK, gin.H {
			"access_token": resp.AccessToken,
			"expires_in": resp.ExpiresIn,
			"token_type": resp.TokenType,
		})
	}

	err := util.WithRPCClient("localhost:8081", userpb.NewUserServiceClient, callback)

	if err != nil {
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": fmt.Sprintf("RPC error: %s", err.Error()),
		})
	}
}
