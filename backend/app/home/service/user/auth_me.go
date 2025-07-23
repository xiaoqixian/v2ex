// Date:   Sat Jun 14 22:16:09 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package user_service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/home/util"
)
func AuthMe(ginCtx *gin.Context) {
	token, err := ginCtx.Cookie("refresh_token")
	if err != nil {
		ginCtx.JSON(http.StatusUnauthorized, gin.H {
			"error": "missing refresh token in cookie",
		})
		return
	}

	userid, valid := util.ParseToken(token)
	if !valid {
		ginCtx.JSON(http.StatusUnauthorized, gin.H {
			"error": "expired refresh token",
		})
		return
	}

	userExist, err := util.CheckUserExist(userid)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
		})
		return
	}
	if !userExist {
		ginCtx.JSON(http.StatusBadRequest, gin.H {
			"error": "bad refresh token, user not found",
		})
		return
	}

	ginCtx.Set("userid", userid)
	ginCtx.JSON(http.StatusOK, gin.H {})
}
