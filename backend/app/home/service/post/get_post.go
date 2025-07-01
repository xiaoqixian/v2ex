// Date:   Sat Jun 21 20:33:10 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package post_service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/postpb"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/userpb"
)

func GetPost(ginCtx *gin.Context) {
	postID, err := strconv.Atoi(ginCtx.Param("post_id"))
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H {
			"error": fmt.Sprintf("Invalid post_id: %s", ginCtx.Param("post_id")),
		})
		return
	}

	var userID uint
	accessToken, err := ginCtx.Cookie("access_token")
	log.Println("access_token detected in cookie, get post with user id")
	if err == nil {
		parse := func(ctx context.Context, client userpb.UserServiceClient) error {
			resp, err2 := client.AuthMe(ctx, &userpb.AuthMeRequest {
				AccessToken: accessToken,
			})
			if err2 == nil && resp.Success {
				userID = uint(resp.UserId)
			}
			return nil
		}
		rpcutil.WithRPCClient("user-service", userpb.NewUserServiceClient, parse)
	}

	callback := func(ctx context.Context, client postpb.PostServiceClient) error {
		resp, err2 := client.GetPost(ctx, &postpb.GetPostRequest {
			PostId: uint64(postID),
			UserId: uint64(userID),
		})
		if err2 != nil {
			return err2
		}

		if e, ok := resp.Result.(*postpb.GetPostResponse_Err); ok {
			return fmt.Errorf("rpc GetPost failed: %s", e.Err.Message)
		}

		o, _ := resp.Result.(*postpb.GetPostResponse_Ok)
		ginCtx.JSON(http.StatusOK, o.Ok)
		return nil
	}

	err = rpcutil.WithRPCClient("post-service", postpb.NewPostServiceClient, callback)

	if err != nil {
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": err.Error(),
		})
	}
}
