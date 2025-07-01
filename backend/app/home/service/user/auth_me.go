// Date:   Sat Jun 14 22:16:09 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package user_service

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/userpb"
)

func AuthMe(ginCtx *gin.Context) {
	token, err := ginCtx.Cookie("access_token")
	if err != nil {
		ginCtx.JSON(http.StatusUnauthorized, gin.H {})
		log.Printf("Get 'access_token' cookie error: %s\n", err.Error())
		return
	}

	req := userpb.AuthMeRequest {
		AccessToken: token,
	}

	callback := func(ctx context.Context, client userpb.UserServiceClient) error {
		resp, err2 := client.AuthMe(ctx, &req)
		if err2 != nil {
			return err2
		}
		if !resp.Success {
			return fmt.Errorf("rpc AuthMe failed")
		}

		ginCtx.JSON(http.StatusOK, gin.H {})
		return nil
	}

	err = rpcutil.WithRPCClient("user-service", userpb.NewUserServiceClient, callback)

	if err != nil {
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": err.Error(),
		})
	}
}
