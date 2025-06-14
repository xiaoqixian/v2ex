// Date:   Wed Jun 11 21:18:29 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/home/service/user"
)

func main() {
	r := gin.Default()
	r.POST("/register", user_service.RegisterUser)
	r.POST("/login", user_service.UserLogin)
	r.POST("/refresh-token", user_service.RefreshToken)
	r.Run(":8080")
}
