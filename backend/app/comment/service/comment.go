// Date:   Sat Jun 28 10:49:56 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package service

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/xiaoqixian/v2ex/backend/app/comment/conf"
	"github.com/xiaoqixian/v2ex/backend/app/comment/dal"
	"github.com/xiaoqixian/v2ex/backend/app/comment/model"
	"github.com/xiaoqixian/v2ex/backend/app/common/rpcutil"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/commentpb"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/userpb"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	getUserInfoReq := userpb.GetUserInfoRequest {
		UserId: in.UserId,
		JustCheckExist: true,
	}
	conf := conf.GetConf()
	getUserInfoRespAny, err := rpcutil.NewBuilder(&getUserInfoReq, userpb.NewUserServiceClient).
		WithService(conf.Consul.User).
		WithMethod("GetUserInfo").
		WithMsTimeout(conf.Rpc.RpcTimeout).
		Call()

	if err != nil {
		return nil, err
	}
	getUserInfoResp, ok := getUserInfoRespAny.(*userpb.GetUserInfoResponse)
	if !ok {
		return nil, fmt.Errorf("RPC error: expect response type '*userpb.GetUserInfoResponse', got '%T'", getUserInfoRespAny)
	}

	if !getUserInfoResp.Exist {
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

func (impl *CommentServiceImpl) GetComments(
	ctx context.Context,
	in *commentpb.GetCommentsReqeust,
) (*commentpb.GetCommentsResponse, error) {
	comments, err := model.GetCommentsByPostID(impl.db, ctx, uint(in.PostId))
	if err != nil {
		return nil, err
	}
	
	userIDList := make([]uint64, len(comments))
	for i, c := range comments {
		userIDList[i] = uint64(c.UserID)
	}
	getBatchUserInfoReq := userpb.GetBatchUserInfoRequest {
		UserIdList: userIDList,
	}

	conf := conf.GetConf()
	userInfoListAny, err := rpcutil.NewBuilder(&getBatchUserInfoReq, userpb.NewUserServiceClient).
		WithService(conf.Consul.User).
		WithMethod("GetBatchUserInfo").
		WithMsTimeout(conf.Rpc.RpcTimeout).
		Call()
	if err != nil {
		return nil, err
	}

	userInfoListResp, ok := userInfoListAny.(*userpb.GetBatchUserInfoResponse)
	if !ok {
		return nil, fmt.Errorf("[GetComments] expect type '*userpb.GetBatchUserInfoResponse', got '%T'", userInfoListAny)
	}

	userInfoList := userInfoListResp.UserInfoList

	respComments := make([]*commentpb.Comment, 0, len(comments))
	for i, c := range comments {
		if !userInfoList[i].Exist {
			continue
		}

		respComments = append(respComments, &commentpb.Comment {
			CommentId: uint64(c.ID),
			UserName: userInfoList[i].Name,
			Content: c.Content,
			Likes: uint32(c.Likes),
			CreatedAt: timestamppb.New(c.CreatedAt),
			Avatar: userInfoList[i].Avatar,
		})
	}

	return &commentpb.GetCommentsResponse {
		Comments: respComments,
	}, nil
}
