package api

import (
	"github.com/gin-gonic/gin"
	"suning/middleware"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.Cors())
	//登陆注册
	u := r.Group("/user")
	{
		u.POST("/register", Register)
		u.POST("/login", Login)
		a := u.Group("/auth")
		{
			a.Use(middleware.JWTAuthMiddleware(), middleware.UserAuth())
			a.POST("/refresh", Refresh)
		}
	}
	//主页
	h := r.Group("/home")
	{
		h.GET("/search", Search)
		h.GET("/sort", Sort)
		h.GET("/category", Category)
	}
	//商品详情页
	{
		r.GET("/product/style/:product_id", SearchStyle)
		r.GET("/product/detail/:product_id", ViewProductDetail)
		r.GET("/review/view/:product_id", ViewReview)
		{
			x := r.Group("")
			x.Use(middleware.JWTAuthMiddleware(), middleware.UserAuth())
			x.POST("/cart/add/:user_id", AddToCart)
			x.GET("/cart/view/:user_id", ViewCart)
			x.DELETE("/cart/delete/:user_id", DeleteCart)
			x.POST("/collection/add/:user_id", AddToCollection)
			x.GET("/collection/view/:user_id", ViewCollection)
			x.DELETE("/collection/delete/:user_id", DeleteCollection)
			x.POST("/review/add/:user_id", GiveReview)
		}
	}
	//个人页面
	i := r.Group("/individual/auth")
	{
		i.Use(middleware.JWTAuthMiddleware(), middleware.UserAuth())
		i.GET("/balance/:user_id", ViewBalance)
		i.POST("/recharge/:user_id", Recharge)
		i.GET("/information/:user_id", ViewInformation)
		i.PUT("/modify/:user_id", ChangeInformation)
	}
	//后台管理
	b := r.Group("/seller")
	{
		b.POST("/register", BackRegister)
		b.POST("/login", BackLogin)
		a := b.Group("/auth")
		{
			a.Use(middleware.JWTAuthMiddleware(), middleware.SellerAuth())
			a.GET("/refresh", Refresh)
			a.GET("/view/:seller_id", ViewProduct)
			a.POST("/add/:seller_id", AddProduct)
			a.PUT("/update/:seller_id", UpdateProduct)
			a.DELETE("/delete/:seller_id", DeleteProduct)
		}

	}

	r.Run()
}
