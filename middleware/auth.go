package middleware

import (
	"github.com/gin-gonic/gin"
	"suning/util"
)

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		c.Next()
	}
}

func SellerAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		c.Next()
	}
}
