package middleware

import (
	"github.com/gin-gonic/gin"
	"suning/util"
)

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := c.Cookie("uid")
		if err != nil {
			util.RespUnauthorizedErr(c)
			c.Abort()
			return
		}
		if uid == "" {
			util.RespUnauthorizedErr(c)
			c.Abort()
			return
		}
		c.Next()
	}
}

func SellerAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid, err := c.Cookie("sid")
		if err != nil {
			util.RespUnauthorizedErr(c)
			c.Abort()
			return
		}
		if sid == "" {
			util.RespUnauthorizedErr(c)
			c.Abort()
			return
		}
		c.Next()
	}
}
