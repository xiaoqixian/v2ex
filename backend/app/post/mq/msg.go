// Date:   Tue Jun 17 20:28:12 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package mq

type AddPostMessage struct {
	UserID uint   `json:"user_id"`
	PostID uint   `json:"post_id"`
	Node   string `json:"node"`
}

type ViewPostMessage struct {
	UserID uint   `json:"user_id"`
	PostID uint   `json:"post_id"`
}
