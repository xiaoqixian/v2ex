// Date:   Sun Jul 13 16:19:34 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package util

import (
	"fmt"

	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/app/home/conf"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/userpb"
)

type R = *userpb.GetUserInfoResponse

func getUserInfo(userid uint64, justCheckExist bool, withResp func(R)) error {
	req := userpb.GetUserInfoRequest {
		UserId: userid,
		JustCheckExist: justCheckExist,
	}

	conf := conf.GetConf()
	respAny, err := rpcutil.NewBuilder(&req, userpb.NewUserServiceClient).
		WithService(conf.Consul.User).
		WithMethod("GetUserInfo").
		WithMsTimeout(conf.Rpc.RpcTimeout).
		Call()

	if err != nil {
		return err
	}
	
	resp, ok := respAny.(*userpb.GetUserInfoResponse)
	if !ok {
		// log.Panicf("[GetUserInfo] expect resp type '*userpb.GetUserInfoResponse', got '%T'", respAny)
		return fmt.Errorf("[GetUserInfo] expect resp type '*userpb.GetUserInfoResponse', got '%T'", respAny)
	}

	withResp(resp)
	return nil
}

func GetUserInfo(userid uint64, withResp func(R)) error {
	return getUserInfo(userid, false, withResp)
}

func CheckUserExist(userid uint64) (bool, error) {
	var exist bool
	err := getUserInfo(userid, true, func(resp R) {
		exist = resp.Exist
	})
	return exist, err
}
