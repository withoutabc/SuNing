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
		u.GET("/refresh", middleware.UserAuth(), UserRefresh)
		u.POST("/logout", middleware.UserAuth(), Logout)
	}
	p := r.Group("/person")
	{
		p.GET("/balance", middleware.UserAuth(), ViewBalance)
		p.POST("/recharge", middleware.UserAuth(), Recharge)
		p.GET("/information", middleware.UserAuth(), ViewInformation)
		p.PUT("/modify", middleware.UserAuth(), ChangeInformation)
	}
	b := r.Group("/seller")
	{
		b.POST("/register", BackRegister)
		b.POST("/login", BackLogin)
		b.GET("/refresh", middleware.SellerAuth(), BackRefresh)
		b.POST("logout", middleware.SellerAuth(), BackLogout)
		b.GET("/view", middleware.SellerAuth(), ViewProduct)
		b.POST("/add", middleware.SellerAuth(), AddProduct)
		b.PUT("/update", middleware.SellerAuth(), UpdateProduct)
		b.DELETE("/delete", middleware.SellerAuth(), DeleteProduct)
	}
	r.Run()
}
