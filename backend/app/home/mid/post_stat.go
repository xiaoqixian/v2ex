// Date:   Mon Jul 14 20:09:55 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package mid

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/app/home/conf"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/postpb"
)

// keep track of post stat
func TrackPostView() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if ctx.Writer.Status() != http.StatusOK {
			return
		}

		userid, exists := ctx.Get("userid")
		if !exists || userid == 0 {
			return
		}
		
		conf := conf.GetConf()
		req := postpb.AddPostViewRequest {}
		_, err := rpcutil.NewBuilder(&req, postpb.NewPostServiceClient).
			WithService(conf.Consul.Post).
			WithMethod("AddPostView").
			WithMsTimeout(conf.Rpc.RpcTimeout).
			Call()

		if err != nil {
			log.Printf("[TrackPostView] %s\n", err.Error())
			return
		}
	}
}
