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
	"github.com/xiaoqixian/v2ex/backend/app/home/util"
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
		parse := func(ctx context.Context, client userpb.UserServiceClient) {
			resp, err2 := client.AuthMe(ctx, &userpb.AuthMeRequest {
				AccessToken: accessToken,
			})
			if err2 == nil && resp.Success {
				userID = uint(resp.UserId)
				log.Printf("successfully get user id %d\n", userID)
			}
		}
		util.WithRPCClient("localhost:8081", userpb.NewUserServiceClient, parse)
	}

	callback := func(ctx context.Context, client postpb.PostServiceClient) {
		resp, err2 := client.GetPost(ctx, &postpb.GetPostRequest {
			PostId: uint64(postID),
			UserId: uint64(userID),
		})
		if err2 != nil {
			ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
				"error": fmt.Sprintf("RPC error: %s", err2.Error()),
			})
			return
		}

		if e, ok := resp.Result.(*postpb.GetPostResponse_Err); ok {
			ginCtx.JSON(http.StatusNotFound, gin.H {
				"error": e.Err.Message,
			})
			return
		}

		o, _ := resp.Result.(*postpb.GetPostResponse_Ok)
		ginCtx.JSON(http.StatusOK, o.Ok)
	}

	err = util.WithRPCClient("localhost:8082", postpb.NewPostServiceClient, callback)

	if err != nil {
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": fmt.Sprintf("RPC error: %s", err.Error()),
		})
	}
}
