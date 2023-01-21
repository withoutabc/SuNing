package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"suning/util"
)

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get("role")
		if ok != true {
			fmt.Println("role not exist")
			util.NormErr(c, 400, "role not exist")
			c.Abort()
		}
		if role == "user" {
			c.Next()
		}
		c.Abort()
	}
}

func SellerAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get("role")
		if ok != true {
			fmt.Println("role not exist")
			util.NormErr(c, 400, "role not exist")
			c.Abort()
		}
		if role == "seller" {
			c.Next()
		}
		c.Abort()
	}

}
