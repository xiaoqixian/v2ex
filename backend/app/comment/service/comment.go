// Date:   Sat Jun 28 10:49:56 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package service

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/xiaoqixian/v2ex/backend/app/comment/dal"
	"github.com/xiaoqixian/v2ex/backend/app/comment/model"
	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/commentpb"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/userpb"
	"gorm.io/gorm"
)

type CommentServiceImpl struct {
	commentpb.UnimplementedCommentServiceServer
	db *gorm.DB
	redis *redis.Client
}

func NewCommentService() (*CommentServiceImpl, error) {
	db := dal.MysqlDB
	rdb := dal.Redis
	db.AutoMigrate(&model.Comment {})

	return &CommentServiceImpl {
		db: db,
		redis: rdb,
	}, nil
}

func (impl *CommentServiceImpl) AddComment(
	ctx context.Context,
	in *commentpb.AddCommentRequest,
) (*commentpb.AddCommentResponse, error) {
	// check if user exists
	userExists := true
	checkUserExists := func(ctx context.Context, client userpb.UserServiceClient) error {
		getUserInfoResp, err := client.GetUserInfo(ctx, &userpb.GetUserInfoRequest {
			UserId: in.UserId,
			JustCheckExist: true,
		})
		if err != nil {
			return err
		}
		userExists = getUserInfoResp.Exist
		return nil
	}
	err := rpcutil.WithRPCClient("user-service", userpb.NewUserServiceClient, checkUserExists)
	if err != nil {
		return nil, err
	}
	if !userExists {
		return nil, fmt.Errorf("invalid user id %d", in.UserId)
	}

	commentID, err := model.AddComment(impl.db, ctx, &model.Comment {
		PostID: uint(in.PostId),
		UserID: uint(in.UserId),
		Content: in.Content,
	})

	if err != nil {
		return nil, err
	}
	
	return &commentpb.AddCommentResponse {
		Success: true,
		CommentId: uint64(commentID),
	}, nil
}
