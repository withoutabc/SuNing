package middleware

import (
	"github.com/gin-gonic/gin"
	"suning/util"
	"time"
)

func RefreshUserCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		//判断cookie过没过期
		username, err := c.Cookie("username")
		if err != nil {
			util.RespUnauthorizedErr(c)
			c.Abort()
			return
		}
		if username == "" {
			util.RespUnauthorizedErr(c)
			c.Abort()
			return
		}
		//没过期
		c.Next()
		// 设置新的cookie
		expiration := time.Now().Add(time.Hour)
		c.SetCookie("username", username, int(expiration.Unix()), "/", "localhost", false, true)
	}
}
func RefreshSellerCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		//判断cookie过没过期
		seller, err := c.Cookie("seller")
		if err != nil {
			util.RespUnauthorizedErr(c)
			c.Abort()
			return
		}
		if seller == "" {
			util.RespUnauthorizedErr(c)
			c.Abort()
			return
		}
		//没过期
		c.Next()
		// 设置新的cookie
		expiration := time.Now().Add(time.Hour)
		c.SetCookie("seller", seller, int(expiration.Unix()), "/", "localhost", false, true)
	}
}
