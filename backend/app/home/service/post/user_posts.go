// Date:   Wed Jun 18 20:48:11 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package post_service

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/app/home/conf"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/postpb"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/userpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PostEntry struct {
	PostId        uint64                 `json:"postid"`
	Avatar        string                 `json:"avatar"`
	Title         string                 `json:"title"`
	Author        string                 `json:"author"`
	CreatedAt     *timestamppb.Timestamp `json:"time"`
	LastReplyFrom string                 `json:"lastReplyFrom"`
	ReplyCount    int                    `json:"replyCount"`
}

func getUserInfo(posts []*postpb.PostEntry) ([]*userpb.GetUserInfoResponse, error) {
	conf := conf.GetConf()
	userIdList := make([]uint64, len(posts))
	for i, p := range posts {
		userIdList[i] = p.AuthorId
	}
	userInfoListAny, err := rpcutil.NewBuilder(
		&userpb.GetBatchUserInfoRequest {
			UserIdList: userIdList,
		},
		userpb.NewUserServiceClient,
	).WithService(conf.Consul.User).
		WithMethod("GetBatchUserInfo").
		WithMsTimeout(conf.Rpc.RpcTimeout).
		Call()
	
	if err != nil {
		return nil, err
	}
	userInfoList, ok := userInfoListAny.(*userpb.GetBatchUserInfoResponse)
	if !ok {
		return nil, fmt.Errorf("expect *userpb.GetBatchUserInfoResponse, got '%T'", userInfoListAny)
	}
	
	return userInfoList.UserInfoList, nil
}

func GetPostsForUser(ginCtx *gin.Context) {
	userid, _ := ginCtx.Get("userid")

	req := postpb.GetPostsForUserRequest {
		UserId: userid.(uint64),
	}

	conf := conf.GetConf()
	respAny, err := rpcutil.NewBuilder(&req, postpb.NewPostServiceClient).
		WithService(conf.Consul.Post).
		WithMethod("GetPostsForUser").
		WithMsTimeout(conf.Rpc.RpcTimeout).
		Call()
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, ok := respAny.(*postpb.GetPostsForUserResponse)
	if !ok {
		log.Printf("[GetPostsForUser] Expect *postpb.GetUserPostsResponse, got '%T'\n", respAny)
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": "内部错误，稍后再试",
		})
		return
	}

	posts := make([]*PostEntry, len(resp.Posts))
	for i, p := range resp.Posts {
		posts[i] = &PostEntry {
			Title: p.Title,
			CreatedAt: p.CreatedAt,
			PostId: p.PostId,
		}
	}

	// fetch user info
	var userInfoList []*userpb.GetUserInfoResponse
	{
		var wg sync.WaitGroup
		var err1 error
		wg.Add(1)
		go func() {
			defer wg.Done()
			userInfoList, err1 = getUserInfo(resp.Posts)
		}()

		//TODO: fetch comment count of post

		wg.Wait()

		if err1 != nil {
			ginCtx.JSON(http.StatusInternalServerError, gin.H {
				"error": err1.Error(),
			})
			return
		}
	}

	for i, u := range userInfoList {
		posts[i].Author = u.Name
		posts[i].Avatar = u.Avatar
	}

	ginCtx.JSON(http.StatusOK, posts)
}
