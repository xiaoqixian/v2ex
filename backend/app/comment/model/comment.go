// Date:   Fri Jun 27 22:29:51 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package model

import (
	"context"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	PostID  uint   `json:"post_id"  gorm:"not null;index"`
	UserID  uint   `json:"user_id"  gorm:"not null;index"`
	Content string `json:"content"  gorm:"type:text;not null"`
	// Set to 0 if this is comment to the post, 
	// otherwise this is a reply to the comment that 
	// has the ID equals to ReplyTo
	ReplyTo uint   `json:"reply_to" gorm:"default:0;index"`
	Likes   uint   `json:"likes"    gorm:"default:0"`
}

func AddComment(db *gorm.DB, ctx context.Context, comment *Comment) (uint, error) {
	err := db.Create(comment).Error
	if err != nil {
		return 0, err
	}
	return comment.ID, nil
}

func GetCommentsByPostID(db *gorm.DB, ctx context.Context, postID uint) ([]*Comment, error) {
  var comments []*Comment
  if err := db.WithContext(ctx).
    Where("post_id = ?", postID).
    Order("created_at ASC").
    Find(&comments).Error; err != nil {
    return nil, err
  }
  return comments, nil
}
