package middleware

import (
	"github.com/gin-gonic/gin"
	"suning/util"
)

func Auth() gin.HandlerFunc {
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
