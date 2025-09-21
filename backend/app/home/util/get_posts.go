// Date:   Sat Sep 20 08:55:23 PM 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package util

import (
	"log"
	"sync"
	"time"

	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/app/home/conf"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/postpb"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/userpb"
)

type PostEntry struct {
  PostId    uint64    `json:"post_id"`
  AuthorId  uint64    `json:"author_id"`
  Title     string    `json:"title"`
  Node      string    `json:"node"`
  CreatedAt time.Time `json:"created_at"`
  Author    string    `json:"author"`
  Avatar    string    `json:"avatar"`
}

func GetPostsByIds(postids []uint64) []PostEntry {
	conf := conf.GetConf()
	var getUserInfoWG sync.WaitGroup
	posts := make([]PostEntry, len(postids))

	for i, postid := range postids {
		getPostReq := postpb.GetPostRequest {
			PostId: postid,
		}
		
		respAny, err := rpcutil.NewBuilder(&getPostReq, postpb.NewPostServiceClient).
			WithService(conf.Consul.Post).
			WithMethod("GetPost").
			WithMsTimeout(conf.Rpc.RpcTimeout).
			Call()
		if err != nil {
			log.Printf("GetPost error: %s", err.Error())
			continue
		}

		getPostResp, ok := respAny.(*postpb.GetPostResponse)
		if !ok {
			log.Printf("[GetPost] Expect *postpb.GetPostResponse, got '%T'\n", respAny)
			return nil
		}
		
		if !getPostResp.Found {
			log.Printf("Post %d is just deleted", postid)
			continue
		}

		getUserInfoWG.Add(1)
		go func(idx int) {
			defer getUserInfoWG.Done()
			GetUserInfo(getPostResp.AuthorId, func(resp *userpb.GetUserInfoResponse) {
				posts[idx].Author = resp.Name
				posts[idx].Avatar = resp.Avatar
			})
		}(i)

		posts[i].AuthorId = getPostResp.AuthorId
		posts[i].PostId   = getPostResp.PostId
		posts[i].Title    = getPostResp.Title
		posts[i].Node     = getPostResp.Node
		posts[i].CreatedAt = getPostResp.CreatedAt.AsTime()
	}

	getUserInfoWG.Wait()

	return posts
}
