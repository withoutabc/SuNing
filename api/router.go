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
		a := u.Group("/auth")
		{
			a.Use(middleware.JWTAuthMiddleware(), middleware.UserAuth())
			a.POST("/refresh", Refresh)
			a.POST("/logout", Logout)
		}

	}
	h := r.Group("/home")
	{
		h.GET("/search", Search)
		h.GET("/sort", Sort)
		h.GET("/category", Category)
	}
	p := r.Group("/individual/auth")
	{
		p.Use(middleware.JWTAuthMiddleware(), middleware.UserAuth())
		p.GET("/balance/:uid", ViewBalance)
		p.POST("/recharge/:uid", Recharge)
		p.GET("/information/:uid", ViewInformation)
		p.PUT("/modify/:uid", ChangeInformation)
	}
	b := r.Group("/seller")
	{
		b.POST("/register", BackRegister)
		b.POST("/login", BackLogin)
		a := b.Group("/auth")
		{
			a.Use(middleware.JWTAuthMiddleware(), middleware.SellerAuth())
			a.GET("/refresh", Refresh)
			a.POST("/logout", BackLogout)
			a.GET("/view/:sid", ViewProduct)
			a.POST("/add/:sid", AddProduct)
			a.PUT("/update/:sid", UpdateProduct)
			a.DELETE("/delete/:sid", DeleteProduct)
		}

	}

	r.Run()
}
