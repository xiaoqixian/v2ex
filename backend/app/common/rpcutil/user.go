// Date:   Sat Jul 05 11:45:30 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package rpcutil

import (
	"context"

	"github.com/xiaoqixian/v2ex/backend/rpc_gen/userpb"
)

// Check if a user exists by calling user RPC
func CheckUserExists(ctx context.Context, client userpb.UserServiceClient, userID uint64) (any, error) {
	resp, err := client.GetUserInfo(ctx, &userpb.GetUserInfoRequest {
		UserId: userID,
		JustCheckExist: true,
	})
	if err != nil {
		return nil, err
	}
	return resp.Exist, nil
}

func GetUserInfoById(ctx context.Context, client userpb.UserServiceClient, userID uint64) (any, error) {
	resp, err := client.GetUserInfo(ctx, &userpb.GetUserInfoRequest {
		UserId: userID,
		JustCheckExist: false,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetBatchUserInfoById(ctx context.Context, client userpb.UserServiceClient, userIDList []uint64) (any, error) {
	resp, err := client.GetBatchUserInfo(ctx, &userpb.GetBatchUserInfoRequest {
		UserIdList: userIDList,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
