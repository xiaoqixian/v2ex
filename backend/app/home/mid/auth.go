// Date:   Sat Jul 12 11:28:07 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package mid

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/home/conf"
	"github.com/xiaoqixian/v2ex/backend/app/home/util"
)

func getUseridInCookie(ctx *gin.Context) (uint64, error) {
	accessToken, err := ctx.Cookie("access_token")
	if err != nil {
		return 0, fmt.Errorf("missing token")
	}

	userid, valid := util.ParseToken(accessToken)
	if !valid {
		return 0, fmt.Errorf("invalid or expired token")
	}
	return userid, nil
}

func JWTParse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userid, _ := getUseridInCookie(ctx)
		ctx.Set("userid", userid)
	}
}

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userid, err := getUseridInCookie(ctx)
		if err != nil {
			log.Printf("JWT auth failed: %s\n", err.Error())
			ctx.JSON(http.StatusUnauthorized, gin.H {
				"error": fmt.Sprintf("unauthorized: %s", err.Error()),
			})
			return
		}

		ctx.Set("userid", userid)
		ctx.Next()
	}
}

func JWTGen() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if ctx.Writer.Status() != http.StatusOK {
			return
		}
		
		userid, exists := ctx.Get("userid")
		if !exists {
			log.Println("[Middleware JWTGen] userid not found in gin context, unable to generate JWT token")
			return
		}
		conf := conf.GetConf()
		accessToken, err := util.GenerateToken(userid.(uint64), time.Duration(conf.JWT.AccExpTime) * time.Second)
		if err != nil {
			log.Printf("[Middleware JWTGen] generate accessToken failed: %s\n", err.Error())
			return
		}
		refreshToken, err := util.GenerateToken(userid.(uint64), time.Duration(conf.JWT.RefExpTime) * time.Second)
		if err != nil {
			log.Printf("[Middleware JWTGen] generate refreshToken failed: %s\n", err.Error())
			return
		}

		ctx.SetCookie(
			"access_token",
			accessToken,
			conf.JWT.AccExpTime,
			"/",
			"localhost",
			false,
			true,
		)
		ctx.SetCookie(
			"refresh_token",
			refreshToken,
			conf.JWT.RefExpTime,
			"/",
			"localhost",
			false,
			true,
		)
		ctx.Writer.Header().Set("expires_in", strconv.Itoa(conf.JWT.AccExpTime))
	}
}
