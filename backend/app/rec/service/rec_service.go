// Date:   Sun Sep 14 08:08:41 PM 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package service

import (
	"context"

	redis "github.com/redis/go-redis/v9"
	"github.com/xiaoqixian/v2ex/backend/app/rec/dal"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/recpb"
	"gorm.io/gorm"
)

type RecServiceImpl struct {
	recpb.UnimplementedRecServiceServer
	db *gorm.DB
	redis *redis.Client
}

func NewRecService() (*RecServiceImpl, error) {
	db := dal.MysqlDB
	rdb := dal.Redis

	return &RecServiceImpl {
		db: db,
		redis: rdb,
	}, nil
}

func (impl *RecServiceImpl) RecPosts(
	ctx context.Context,
	in *recpb.RecRequest,
) (*recpb.RecResponse, error) {
	var ids []uint64
	err := impl.db.Table("posts").Limit(int(in.Size)).Pluck("id", &ids).Error
	if err != nil {
		return nil, err
	}
	return &recpb.RecResponse {
		PostIdList: ids,
	}, nil
}
