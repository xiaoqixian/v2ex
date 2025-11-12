// Date:   Tue Jun 17 10:12:10 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package service

import (
	"context"
	"encoding/json"
	"fmt"
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
  postKey := fmt.Sprintf("post:%d", in.PostId)
  cacheStr, err := impl.redis.Get(ctx, postKey).Result()

  if err == nil {
    var resp postpb.GetPostResponse
    if json.Unmarshal([]byte(cacheStr), &resp) == nil {
      return &resp, nil
    }
    log.Println("Something wrong with cache unmarshal, fallback to DB")
  } else if err != redis.Nil {
    log.Printf("Redis get error: %s", err.Error())
  }

  post, err := model.GetPostById(impl.db, ctx, uint(in.PostId))
  if err != nil {
    return nil, err
  }

  if post == nil {
    emptyResp := postpb.GetPostResponse{Found: false}
    data, _ := json.Marshal(&emptyResp)
    impl.redis.Set(ctx, postKey, data, 5*time.Minute)
    return &emptyResp, nil
  }

  resp := postpb.GetPostResponse{
    Found:     true,
    PostId:    uint64(post.ID),
    Title:     post.Title,
    AuthorId:  post.UserID,
    CreatedAt: timestamppb.New(post.CreatedAt),
    Content:   post.Content,
  }

  postViewRedisKey := fmt.Sprintf("post_view_cnt:%d", in.PostId)
  script := redis.NewScript(`
    local v = redis.call("INCR", KEYS[1])
    if v == 1 then
      redis.call("EXPIRE", KEYS[1], ARGV[1])
    end
    return v
  `)
  postViewCnt, err := script.Run(ctx, impl.redis, []string{postViewRedisKey}, int64(time.Hour.Seconds())).Int64()
  if err != nil {
    log.Printf("Redis incr script error: %s", err.Error())
  }

  if postViewCnt > 100 {
    respData, _ := json.Marshal(&resp)
    impl.redis.Set(ctx, postKey, respData, time.Hour)
  }

  return &resp, nil
}

func (impl *PostServiceImpl) PublishPost(
	ctx context.Context,
	in *postpb.PublishPostRequest,
) (*postpb.PublishPostResponse, error) {
	//TODO: check if user exists
	post := model.Post {
		UserID: in.UserId,
		Title: in.Title,
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
