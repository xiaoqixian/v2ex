// Date:   Wed Jul 09 19:57:10 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package util

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/xiaoqixian/v2ex/backend/app/home/conf"
)

func GenerateToken(userID uint64, d time.Duration) (string, error) {
	c := conf.GetConf()
	claims := jwt.MapClaims {
		"user_id": userID,
		"exp": time.Now().Add(d).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(c.JWT.Secret))
}

// parse a jwt access token, return userid, true if its still valid
// otherwise return 0, false
func ParseToken(tokenString string) (userid uint64, valid bool) {
	c := conf.GetConf()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(c.JWT.Secret), nil
	})

	if err != nil {
		log.Printf("jwt token parse error: %s\n", err.Error())
		return 0, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			return 0, false
		}
		return uint64(userIDFloat), true
	}
	return 0, false
}
