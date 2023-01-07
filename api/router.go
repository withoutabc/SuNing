package api

import (
	"github.com/gin-gonic/gin"
	"suning/middleware"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.Cors)
	u := r.Group("/user")
	{
		u.POST("/register", Register)
		u.POST("/login", Login)
	}
	r.Use(middleware.Auth)
	r.Run()
}
