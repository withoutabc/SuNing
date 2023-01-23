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
	{
		r.GET("/product/style/:product_id", SearchStyle)
		r.GET("/product/detail/:product_id", ViewProductDetail)
		r.POST("/cart/add/:user_id", AddToCart)
		r.GET("/cart/view/:user_id", ViewCart)
		r.POST("/collection/add/:user_id", AddToCollection)
		r.GET("/collection/view/:user_id", ViewCollection)
		r.POST("/comment/add/:product_id", GiveComment)
		r.GET("/comment/view/:product_id", ViewComment)
	}
	p := r.Group("/individual/auth")
	{
		p.Use(middleware.JWTAuthMiddleware(), middleware.UserAuth())
		p.GET("/balance/:user_id", ViewBalance)
		p.POST("/recharge/:user_id", Recharge)
		p.GET("/information/:user_id", ViewInformation)
		p.PUT("/modify/:user_id", ChangeInformation)
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
			a.GET("/view/:seller_id", ViewProduct)
			a.POST("/add/:seller_id", AddProduct)
			a.PUT("/update/:seller_id", UpdateProduct)
			a.DELETE("/delete/:seller_id", DeleteProduct)
		}

	}

	r.Run()
}
