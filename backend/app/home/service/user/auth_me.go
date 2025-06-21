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
	"github.com/xiaoqixian/v2ex/backend/app/home/util"
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

	callback := func(ctx context.Context, client userpb.UserServiceClient) {
		resp, err2 := client.AuthMe(ctx, &req)
		if err2 != nil || !resp.Success {
			ginCtx.JSON(http.StatusServiceUnavailable, gin.H {})
			return
		}

		ginCtx.JSON(http.StatusOK, gin.H {})
	}

	err = util.WithRPCClient(":8081", userpb.NewUserServiceClient, callback)

	if err != nil {
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": fmt.Sprintf("RPC error: %s", err.Error()),
		})
	}
}
