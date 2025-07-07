// Date:   Wed Jun 11 21:18:29 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/home/service/post"
	"github.com/xiaoqixian/v2ex/backend/app/home/service/user"
	"github.com/xiaoqixian/v2ex/backend/app/home/service/comment"
)

func main() {
	r := gin.Default()

	// CORS(Cross-Origin Resource Sharing) protection
	r.Use(cors.New(cors.Config {
		AllowOrigins:     []string{"http://localhost:8000"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
	}))

	r.POST("/register", user_service.RegisterUser)
	r.POST("/login", user_service.UserLogin)
	r.POST("/auth/refresh", user_service.RefreshToken)
	r.GET("/auth/me", user_service.AuthMe)

	r.GET("/user_posts/:user_id", post_service.GetPostsForUser)
	r.GET("/posts/:post_id", post_service.GetPost)
	r.POST("/posts", post_service.PublishPost)

	r.GET("/comments/:post_id", comment_service.GetComments)
	r.POST("/comments/:post_id", comment_service.SubmitComment)

	r.Run(":8080")
}
