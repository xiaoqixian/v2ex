// Date:   Wed Jun 11 21:58:35 2025
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

type RegisterArgs struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func RegisterUser(ginCtx *gin.Context) {
	log.Println("New RegisterUser request")

	var args RegisterArgs
	err := ginCtx.ShouldBindJSON(&args)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H {
			"error": fmt.Sprintf("JSON bind error: %s", err.Error()),
		})
		return
	}

	req := userpb.RegisterRequest {
		Username: args.Username,
		Email: args.Email,
		Password: args.Password,
	}

	callback := func(ctx context.Context, client userpb.UserServiceClient) {
		resp, err2 := client.Register(ctx, &req)
		if err2 != nil || !resp.Success {
			ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
				"error": fmt.Sprintf("RPC error: %s", err2.Error()),
			})
			return
		}
		if !resp.Success {
			ginCtx.JSON(http.StatusBadRequest, gin.H {
				"error": resp.Message,
			})
			return
		}
		
		ginCtx.JSON(http.StatusOK, gin.H {
			"message": "注册成功",
		})
	}

	err = util.WithRPCClient(":8081", userpb.NewUserServiceClient, callback)

	if err != nil {
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": fmt.Sprintf("RPC error: %s", err.Error()),
		})
	}
}
