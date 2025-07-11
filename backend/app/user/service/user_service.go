// Date:   Wed Jun 11 15:49:25 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package service

import (
	"context"
	"log"

	// "github.com/xiaoqixian/v2ex/backend/app/user/conf"
	"github.com/redis/go-redis/v9"
	"github.com/xiaoqixian/v2ex/backend/app/user/dal"
	"github.com/xiaoqixian/v2ex/backend/app/user/model"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/userpb"
	"golang.org/x/crypto/bcrypt"

	// "google.golang.org/grpc"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	userpb.UnimplementedUserServiceServer
	db *gorm.DB
	redis *redis.Client
}

func NewUserService() (*UserServiceImpl, error) {
	db := dal.MysqlDB
	rdb := dal.Redis
	db.AutoMigrate(&model.User {})

	return &UserServiceImpl {
		db: db,
		redis: rdb,
	}, nil
}

func (impl *UserServiceImpl) Register(
  ctx context.Context,
  in *userpb.RegisterRequest,
) (*userpb.RegisterResponse, error) {
	userExists, err := model.UsernameExists(impl.db, ctx, in.Username)
	if err != nil {
		return nil, err
	}
	if userExists {
		return &userpb.RegisterResponse{
			Success: false,
			Message: "用户名已存在",
		}, nil
	}

	emailExists, err := model.EmailExists(impl.db, ctx, in.Email)
	if err != nil {
		return nil, err
	}
	if emailExists {
		return &userpb.RegisterResponse{
			Success: false,
			Message: "邮箱已注册",
		}, nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User {
		Email: in.Email,
		PasswordHashed: string(hashedPassword),
		Username: in.Username,
		Avatar: "",
	}

	err = model.RegisterUser(impl.db, ctx, user)
	if err != nil {
		return nil, err
	}
	return &userpb.RegisterResponse{
		Success: true,
		Message: "",
	}, nil
}

func (impl *UserServiceImpl) Login(
  ctx context.Context,
  in *userpb.LoginRequest,
) (*userpb.LoginResponse, error) {
	user, err := model.GetUserByUsername(impl.db, ctx, in.Username)
	if err != nil {
		return nil, err
	}
	
	if user == nil {
		return &userpb.LoginResponse {
			Success: false,
			Message: "用户不存在",
		}, nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err.Error())
	}
	
	if string(hashedPassword) != user.PasswordHashed {
		return &userpb.LoginResponse {
			Success: false,
			Message: "密码错误",
		}, nil
	}

	log.Printf("User(%d, '%s') logged in\n", user.ID, user.Username)
	return &userpb.LoginResponse {
		Success: true,
		Message: "登录成功",
		User: &userpb.LoginResponse_User {
			Id: uint64(user.ID),
		},
	}, nil
}

func (impl *UserServiceImpl) GetUserInfo(
	ctx context.Context,
	in *userpb.GetUserInfoRequest,
) (*userpb.GetUserInfoResponse, error) {
	user := model.GetUserInfoById(impl.db, ctx, uint(in.UserId))
	if user == nil {
		return &userpb.GetUserInfoResponse {
			Exist: false,
		}, nil
	}

	resp := &userpb.GetUserInfoResponse {
		Exist: true,
	}
	if !in.JustCheckExist {
		resp.Name = user.Username
		resp.Avatar = user.Avatar
	}
	return resp, nil
}

func (impl *UserServiceImpl) GetBatchUserInfo(
	ctx context.Context,
	in *userpb.GetBatchUserInfoRequest,
) (*userpb.GetBatchUserInfoResponse, error) {
	userInfoList := make([]*userpb.GetUserInfoResponse, 0, len(in.UserIdList))
	
	for _, id := range in.UserIdList {
		user := model.GetUserInfoById(impl.db, ctx, uint(id))
		if user == nil {
			continue
		}
		userInfoList = append(userInfoList, &userpb.GetUserInfoResponse {
			Exist: true,
			Name: user.Username,
			Avatar: user.Avatar,
		})
	}
	return &userpb.GetBatchUserInfoResponse {
		UserInfoList: userInfoList,
	}, nil
}
