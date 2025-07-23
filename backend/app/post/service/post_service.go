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
	"github.com/xiaoqixian/v2ex/backend/app/post/conf"
	"github.com/xiaoqixian/v2ex/backend/app/post/dal"
	"github.com/xiaoqixian/v2ex/backend/app/post/model"
	"github.com/xiaoqixian/v2ex/backend/app/post/mq"
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
	post, err := model.GetPostById(impl.db, ctx, uint(in.PostId))
	if err != nil {
		return nil, err
	}
	if post == nil {
		return &postpb.GetPostResponse {
			Found: false,
		}, nil
	}

	return &postpb.GetPostResponse {
		Found: true,
		PostId: uint64(post.ID),
		Title: post.Title,
		AuthorId: post.UserID,
		Node: post.Node,
		CreatedAt: timestamppb.New(post.CreatedAt),
		Content: post.Content,
	}, nil
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
	log.Println("Write a AddPost message to kafka")
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
	err := impl.db.Select("id, created_at, user_id, title, node").
		Order("created_at DESC").Limit(N).Find(&posts).Error

	if err != nil {
		return nil, err
	}

	size := len(posts)
	respPosts := make([]*postpb.PostEntry, size)
	for i, p := range posts {
		respPosts[i] = &postpb.PostEntry {
			PostId: uint64(p.ID),
			AuthorId: p.UserID,
			Title: p.Title,
			Node: p.Node,
			CreatedAt: timestamppb.New(p.CreatedAt),
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

func (impl *PostServiceImpl) writeViewPostMsg(ctx context.Context, postID uint, userID uint) error {
	data, err := json.Marshal(&mq.ViewPostMessage {
		UserID: userID,
		PostID: postID,
	})
	if err != nil {
		return err
	}

	msg := kafka.Message {
		Key: []byte("view-post"),
		Value: data,
		Time: time.Now(),
	}

	return impl.kafkaProducer.WriteMessages(ctx, msg)
}
