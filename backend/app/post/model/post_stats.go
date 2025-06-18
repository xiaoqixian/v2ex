// Date:   Tue Jun 17 12:23:32 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package model

type PostStats struct {
	PostID        uint64 `gorm:"primaryKey"`
	ViewCount     uint64
	CommentCount  uint64
	LikeCount     uint64
}


