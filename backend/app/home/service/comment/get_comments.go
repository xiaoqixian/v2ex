// Date:   Sat Jul 05 14:22:14 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package comment_service

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/app/home/conf"
	"github.com/xiaoqixian/v2ex/backend/app/home/util"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/commentpb"
)

type CommentEntry struct {
	CommentId uint64 `json:"comment_id"`
	PostId  uint64   `json:"post_id"`
	UserId  uint64   `json:"user_id"`
	ReplyTo uint64   `json:"reply_to"`
	Content string `json:"content"`
	// Set to 0 if this is comment to the post, 
	// otherwise this is a reply to the comment that 
	// has the ID equals to ReplyTo
	Likes   uint64   `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
	UserName string `json:"username"`
	Avatar string `json:"avatar"`
}

func GetComments(ginCtx *gin.Context) {
	postid, err := strconv.ParseUint(ginCtx.Param("post_id"), 10, 64)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H {
			"error": fmt.Sprintf("Invalid post_id: %s", ginCtx.Param("post_id")),
		})
		return
	}

	req := commentpb.GetCommentsReqeust {
		PostId: postid,
	}

	conf := conf.GetConf()
	respAny, err := rpcutil.NewBuilder(&req, commentpb.NewCommentServiceClient).
		WithService(conf.Consul.Comment).
		WithMethod("GetComments").
		WithMsTimeout(conf.Rpc.RpcTimeout).
		Call()
	if err != nil {
		ginCtx.JSON(http.StatusServiceUnavailable, gin.H {
			"error": err.Error(),
		})
		return
	}
	
	resp, ok := respAny.(*commentpb.GetCommentsResponse)
	if !ok {
		log.Printf("[GetComments] Expect *commentpb.GetCommentsResponse, got '%T'\n", respAny)
		ginCtx.JSON(http.StatusInternalServerError, gin.H {
			"error": "内部错误，稍后再试",
		})
		return
	}

	userids := make([]uint64, len(resp.Comments))
	for i, cmt := range resp.Comments {
		userids[i] = cmt.UserId
	}

	entries := make([]CommentEntry, 0, len(resp.Comments))
	for _, cmt := range resp.Comments {
		entries = append(entries, CommentEntry {
			CommentId: cmt.CommentId,
			PostId: cmt.PostId,
			UserId: cmt.UserId,
			ReplyTo: cmt.ReplyTo,
			Content: cmt.Content,
			Likes: uint64(cmt.Likes),
			CreatedAt: cmt.CreatedAt.AsTime(),
		})
	}
	util.GetBatchUserInfo(userids, func(idx int, userInfo util.UserInfo) {
		if !userInfo.Exist {
			entries[idx].UserName = "用户已注销"
		} else {
			entries[idx].UserName = userInfo.Name
			entries[idx].Avatar = userInfo.Avatar
		}
	})

	ginCtx.JSON(http.StatusOK, entries)
}
