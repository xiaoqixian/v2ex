// Date:   Mon Jun 16 23:34:42 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID  uint64 `json:"user_id"  gorm:"not null;index"`
	Title   string `json:"title"    gorm:"type:varchar(255);not null"`
	Node    string `json:"node"     gorm:"type:varchar(100);not null"`
	Content string `json:"content"  gorm:"type:text;not null"`
}

func AddPost(db *gorm.DB, ctx context.Context, post *Post) error {
	err := db.Create(post).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPostById(db *gorm.DB, ctx context.Context, postID uint) (*Post, error) {
	var post Post
	err := db.Model(&Post{}).
		Where("id = ?", postID).
		First(&post).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &post, err
}

func GetPostsByUserID(db *gorm.DB, ctx context.Context, userID uint64) ([]Post, error) {
	var posts []Post
	err := db.Where("user_id = ?", userID).Order("created_at DESC").Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}
