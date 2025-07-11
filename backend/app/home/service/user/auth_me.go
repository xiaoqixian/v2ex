// Date:   Sat Jun 14 22:16:09 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package user_service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/app/home/conf"
	"github.com/xiaoqixian/v2ex/backend/app/home/util"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/userpb"
)

func CheckUser(accessToken string) (uint64, error) {
	userid, valid := util.ParseToken(accessToken)
	if !valid {
		return 0, fmt.Errorf("access token expired")
	}

	req := userpb.GetUserInfoRequest {
		UserId: uint64(userid),
		JustCheckExist: true,
	}

	conf := conf.GetConf()
	respAny, err := rpcutil.NewBuilder(&req, userpb.NewUserServiceClient).
		WithService(conf.Consul.User).
		WithMethod("GetUserInfo").
		WithMsTimeout(conf.Rpc.RpcTimeout).
		Call()
	if err != nil {
		return 0, err
	}

	resp, ok := respAny.(*userpb.GetUserInfoResponse)
	if !ok {
		log.Printf("respAny convert to *userpb.LoginResponse failed, actual type is '%T'\n", respAny)
		return 0, fmt.Errorf("内部错误，稍后再试")
	}
	
	if !resp.Exist {
		return 0, fmt.Errorf("user not found")
	}

	return userid, nil
}

func AuthMe(ginCtx *gin.Context) {
	token, err := ginCtx.Cookie("access_token")
	if err != nil {
		ginCtx.JSON(http.StatusUnauthorized, gin.H {})
		log.Printf("Get 'access_token' cookie error: %s\n", err.Error())
		return
	}

	_, err = CheckUser(token)
	if err != nil {
		ginCtx.JSON(http.StatusUnauthorized, gin.H {
			"error": err.Error(),
		})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H {})
}
