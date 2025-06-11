// Date:   Wed Jun 11 15:49:25 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package service

import (
	"context"
	"fmt"

	// "github.com/xiaoqixian/v2ex/backend/app/user/conf"
	"github.com/xiaoqixian/v2ex/backend/app/user/conf"
	"github.com/xiaoqixian/v2ex/backend/app/user/model"
	"github.com/xiaoqixian/v2ex/backend/rpc_gen/userpb"
	"golang.org/x/crypto/bcrypt"
	// "google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type RegisterServiceImpl struct {
	userpb.UnimplementedUserServiceServer
	db *gorm.DB
}

func NewRegisterService() (*RegisterServiceImpl, error) {
	c := conf.GetConf()
	parseTime := "True"
	if !c.MySQL.ParseTime {
		parseTime = "False"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", 
		c.MySQL.User,
		c.MySQL.Password,
		c.MySQL.Host,
		c.MySQL.Port,
		c.MySQL.DBName,
		c.MySQL.Charset,
		parseTime,
		c.MySQL.Loc,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &RegisterServiceImpl {
		db: db,
	}, nil
}

func (impl *RegisterServiceImpl) Register(
  ctx context.Context,
  in *userpb.RegisterRequest,
) (*userpb.RegisterResponse, error) {
	user, err := model.GetByUsername(impl.db, ctx, in.Username)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return &userpb.RegisterResponse{
			Success: false,
			Message: "用户名已存在",
		}, nil
	}

	user, err = model.GetByEmail(impl.db, ctx, in.Email)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return &userpb.RegisterResponse{
			Success: false,
			Message: "邮箱已注册",
		}, nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user = &model.User {
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
