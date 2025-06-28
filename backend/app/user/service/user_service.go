// Date:   Wed Jun 11 15:49:25 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package service

import (
	"context"
	"log"
	"time"

	// "github.com/xiaoqixian/v2ex/backend/app/user/conf"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"github.com/xiaoqixian/v2ex/backend/app/user/conf"
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
		Avator: "",
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

	c := conf.GetConf()
	accessToken, err := generateToken(user.ID, time.Minute * time.Duration(c.Jwt.AccExpTime))
	if err != nil {
		log.Printf("generateToken error: %s\n", err.Error())
		return &userpb.LoginResponse {
			Success: true,
			Message: "登录成功",
		}, nil
	}

	refreshToken, err := generateToken(user.ID, time.Minute * time.Duration(c.Jwt.RefExpTime))
	if err != nil {
		log.Printf("generateToken error: %s\n", err.Error())
		return &userpb.LoginResponse {
			Success: true,
			Message: "登录成功",
		}, nil
	}
	
	log.Printf("User(%d, '%s') logged in\n", user.ID, user.Username)
	return &userpb.LoginResponse {
		Success: true,
		Message: "登录成功",
		AccessToken: accessToken,
		RefreshToken: refreshToken,
		AccessExpiresIn: int64(c.Jwt.AccExpTime),
		RefreshExpiresIn: int64(c.Jwt.RefExpTime),
		User: &userpb.LoginResponse_User {
			Id: uint64(user.ID),
		},
	}, nil
}

func (impl *UserServiceImpl) RefreshToken(
	ctx context.Context,
	in *userpb.RefreshTokenRequest,
) (*userpb.RefreshTokenResponse, error) {
	userID, valid := parseToken(in.RefreshToken)
	if !valid {
		// leave accessToken empty to represent the token is invalid
		return &userpb.RefreshTokenResponse {}, nil
	}

	if user := model.GetUserById(impl.db, ctx, userID); user == nil {
		return &userpb.RefreshTokenResponse {}, nil
	}
	
	c := conf.GetConf()
	accessToken, err := generateToken(userID, time.Minute * time.Duration(c.Jwt.AccExpTime))
	if err != nil {
		log.Fatalf("generate jwt token error: %s\n", err.Error())
		return &userpb.RefreshTokenResponse {}, nil
	}
	return &userpb.RefreshTokenResponse {
		AccessToken: accessToken,
		ExpiresIn: int64(c.Jwt.AccExpTime),
		TokenType: "Bearer",
	}, nil
}

func (impl *UserServiceImpl) AuthMe(
	ctx context.Context,
	in *userpb.AuthMeRequest,
) (*userpb.AuthMeResponse, error) {
	userID, valid := parseToken(in.AccessToken)
	return &userpb.AuthMeResponse {
		Success: valid,
		UserId: uint64(userID),
	}, nil
}

func (impl *UserServiceImpl) GetUserInfo(
	ctx context.Context,
	in *userpb.GetUserInfoRequest,
) (*userpb.GetUserInfoResponse, error) {
	user := model.GetUserById(impl.db, ctx, uint(in.UserId))
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
		resp.Avatar = user.Avator
	}
	return resp, nil
}

func generateToken(userID uint, d time.Duration) (string, error) {
	c := conf.GetConf()
	claims := jwt.MapClaims {
		"user_id": userID,
		"exp": time.Now().Add(d).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(c.Jwt.Secret))
}

func parseToken(tokenString string) (userid uint, valid bool) {
	c := conf.GetConf()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(c.Jwt.Secret), nil
	})

	if err != nil {
		log.Printf("jwt token parse error: %s\n", err.Error())
		return 0, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			return 0, false
		}
		return uint(userIDFloat), true
	}
	return 0, false
}
