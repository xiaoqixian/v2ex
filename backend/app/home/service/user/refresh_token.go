// Date:   Sat Jun 14 11:59:34 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package user_service

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/home/conf"
	"github.com/xiaoqixian/v2ex/backend/app/home/util"
)

func RefreshToken(ginCtx *gin.Context) {
	token, err := ginCtx.Cookie("refresh_token")
	if err != nil {
		ginCtx.JSON(http.StatusUnauthorized, gin.H {})
		log.Printf("Get 'refresh_token' cookie error: %s\n", err.Error())
		return
	}

	userid, err := CheckUser(token)
	if err != nil {
		ginCtx.JSON(http.StatusUnauthorized, gin.H {
			"error": err.Error(),
		})
		return
	}
	
	conf := conf.GetConf()
	accessToken, err := util.GenerateToken(userid, time.Duration(conf.JWT.AccExpTime) * time.Second)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
		})
		return
	}

	ginCtx.SetCookie(
		"access_token",
		accessToken,
		conf.JWT.AccExpTime,
		"/",
		"localhost",
		false,
		true,
	)
	ginCtx.JSON(http.StatusOK, gin.H {
		"expires_in": conf.JWT.AccExpTime,
	})
}
