// Date:   Tue Jun 17 10:12:10 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package service

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"github.com/xiaoqixian/v2ex/backend/app/user/conf"
	"github.com/xiaoqixian/v2ex/backend/app/user/dal"
	"github.com/xiaoqixian/v2ex/backend/app/user/model"
	"github.com/xiaoqixian/v2ex/backend/app/user/mq"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/postpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type PostServiceImpl struct {
	postpb.UnimplementedPostServiceServer
	db *gorm.DB
	redis *redis.Client
	kafkaProducer *kafka.Writer
}

func NewPostService() (*PostServiceImpl, error) {
	db := dal.MysqlDB
	rdb := dal.Redis
	db.AutoMigrate(&model.Post {})

	c := conf.GetConf()

	writer := kafka.NewWriter(kafka.WriterConfig {
		Brokers: c.Kafka.Brokers,
		Topic: c.Kafka.Topic,
	})

	return &PostServiceImpl {
		db: db,
		redis: rdb,
		kafkaProducer: writer,
	}, nil
}

func (impl *PostServiceImpl) GetPost(
	ctx context.Context,
	in *postpb.GetPostRequest,
) (*postpb.GetPostResponse, error) {
	
}

func (impl *PostServiceImpl) PublishPost(
	ctx context.Context,
	in *postpb.PublishPostRequest,
) (*postpb.PublishPostResponse, error) {
	//TODO: check if user exists
	post := model.Post {
		UserID: in.UserId,
		Title: in.Title,
		Node: in.Node,
		Content: in.Content,
	}

	err := model.AddPost(impl.db, ctx, &post)
	if err != nil {
		return nil, err
	}

	// Write AddPost message to kafka mq
	err = impl.writeAddPostMsg(ctx, post.ID, uint(in.UserId), in.Node)
	if err != nil {
		log.Printf("Write kafka error: %s\n", err.Error())
	}

	return &postpb.PublishPostResponse {
		PostId: uint64(post.ID),
		Message: "发布成功",
	}, nil
}

func (impl *PostServiceImpl) GetPostsForUser(
	ctx context.Context,
	in *postpb.GetPostsForUserRequest,
) (*postpb.GetPostsForUserResponse, error) {
	//TODO: customized user posts recommendations
	const N = 20
	var posts []model.Post
	err := impl.db.Order("created_at DESC").Limit(N).Find(&posts).Error

	if err != nil {
		return nil, err
	}

	size := len(posts)
	respPosts := make([]*postpb.Post, size)
	for i := size - 1; i >= 0; i-- {
		respPosts[size-i] = &postpb.Post {
			AuthorId: posts[i].UserID,
			Title: posts[i].Title,
			Content: posts[i].Content,
			Node: posts[i].Node,
			CreatedAt: timestamppb.New(posts[i].CreatedAt),
		}
	}

	return &postpb.GetPostsForUserResponse {
		Success: true,
		Posts: respPosts,
	}, nil
}

func (impl *PostServiceImpl) writeAddPostMsg(ctx context.Context, postID uint, userID uint, node string) error {
	data, err := json.Marshal(&mq.AddPostMessage {
		UserID: userID,
		PostID: postID,
		Node: node,
	})
	if err != nil {
		return err
	}

	msg := kafka.Message {
		Key: []byte("add-post"),
		Value: data,
		Time: time.Now(),
	}

	return impl.kafkaProducer.WriteMessages(ctx, msg)
}
