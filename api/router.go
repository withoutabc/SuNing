package api

import (
	"github.com/gin-gonic/gin"
	"suning/middleware"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.Cors())
	u := r.Group("/user")
	{
		u.POST("/register", Register)
		u.POST("/login", Login)
		u.POST("/logout", middleware.RefreshCookieMiddleware(), middleware.Auth(), Logout)
	}
	p := r.Group("/person")
	{
		p.GET("/balance", middleware.RefreshCookieMiddleware(), middleware.Auth(), ViewBalance)
		p.POST("/recharge", middleware.RefreshCookieMiddleware(), middleware.Auth(), Recharge)
	}
	r.Run()
}
