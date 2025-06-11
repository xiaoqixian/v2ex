// Date:   Wed Jun 11 17:32:29 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package model

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `json:"email" gorm:"unique;unique_index"`
	PasswordHashed string `json:"passwordHashed" gorm:"type:varchar(255) not null"`
	Username       string `json:"username"`
	Avator         string `json:"avator"`
}

func GetByEmail(db *gorm.DB, ctx context.Context, email string) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where(&User{Email: email}).First(&user).Error
	return user, err
}

func GetByUsername(db *gorm.DB, ctx context.Context, username string) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where(&User{Username: username}).First(&user).Error
	return user, err
}

func RegisterUser(db *gorm.DB, ctx context.Context, user *User) error {
	return db.Create(user).Error
}
