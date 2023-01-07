package middleware

import (
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	// 从 cookie 中获取登录状态
	ifLogin, err := c.Cookie("if_login")
	if err != nil {
		ifLogin = "false"
	}
	// 用户未登录
	if ifLogin != "true" {
		c.Abort()
		return
	}
	// 继续执行后续的处理函数
	c.Next()
}
