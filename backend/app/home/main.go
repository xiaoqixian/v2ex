// Date:   Wed Jun 11 21:18:29 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/home/mid"
	"github.com/xiaoqixian/v2ex/backend/app/home/service/comment"
	"github.com/xiaoqixian/v2ex/backend/app/home/service/post"
	"github.com/xiaoqixian/v2ex/backend/app/home/service/user"
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
	r.POST("/auth/refresh", mid.JWTGenAccess(), user_service.AuthMe)
	r.GET("/auth/me", mid.JWTGenAccess(), user_service.AuthMe)

	r.GET("/home_posts", mid.JWTParse(), post_service.GetPostsForUser)
	r.GET("/posts/:post_id", mid.JWTParse(), post_service.GetPost)
	r.POST("/posts", mid.JWTAuth(), post_service.PublishPost)

	r.GET("/comments/:post_id", comment_service.GetComments)
	r.POST("/comments/:post_id", mid.JWTAuth(), comment_service.SubmitComment)

	r.Run(":8080")
}
