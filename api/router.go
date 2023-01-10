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
		u.GET("/refresh", middleware.Auth(), Refresh)
		u.POST("/logout", middleware.Auth(), Logout)
	}
	p := r.Group("/person")
	{
		p.GET("/balance", middleware.Auth(), ViewBalance)
		p.POST("/recharge", middleware.Auth(), Recharge)
		p.GET("/information", middleware.Auth(), ViewInformation)
		p.PUT("/modify", middleware.Auth(), ChangeInformation)
	}
	r.Run()
}
