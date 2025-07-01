// Date:   Sat Jun 14 11:59:34 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package user_service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
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

	callback := func(ctx context.Context, client userpb.UserServiceClient) error {
		resp, err2 := client.RefreshToken(ctx, &req)
		if err2 != nil {
			return err2
		}
		if resp.AccessToken == "" {
			return fmt.Errorf("RefreshToken expired")
		}

		ginCtx.JSON(http.StatusOK, gin.H {
			"access_token": resp.AccessToken,
			"expires_in": resp.ExpiresIn,
			"token_type": resp.TokenType,
		})
		return nil
	}

	err := rpcutil.WithRPCClient("user-service", userpb.NewUserServiceClient, callback)

	if err != nil {
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": err.Error(),
		})
	}
}
