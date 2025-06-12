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

func EmailExists(db *gorm.DB, ctx context.Context, email string) (exists bool, err error) {
	err = db.Model(&User {}).
		Select("count(*) > 0").
		Where(&User{Email: email}).
		Find(&exists).Error
	return exists, err
}

func UsernameExists(db *gorm.DB, ctx context.Context, username string) (exists bool, err error) {
	err = db.Model(&User {}).
		Select("count(*) > 0").
		Where(&User{Username: username}).
		Find(&exists).Error
	return exists, err
}

func RegisterUser(db *gorm.DB, ctx context.Context, user *User) error {
	return db.Create(user).Error
}
