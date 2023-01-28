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
		h.GET("/search", SearchAndSort)
		h.GET("/category", Category)
	}
	//商品详情页（购物车、我的收藏、评价）
	{
		r.GET("/product/style/:product_id", SearchStyle)
		r.GET("/product/detail/:product_id", ViewProductDetail)
		r.GET("/review/view/:product_id", ViewReview)
		{
			x := r.Group("/auth")
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
	//店铺详情页
	{
		r.GET("/announcement/view/:seller_id", ViewAnnouncement)
		r.GET("/store/sort/:seller_id", StoreSort)
		r.GET("/store/category/:seller_id", StoreCategory)
		r.PUT("/auth/announcement/update/:seller_id", middleware.JWTAuthMiddleware(), middleware.SellerAuth(), UpdateAnnouncement)
	}
	//个人页面
	i := r.Group("/individual/auth")
	{
		i.Use(middleware.JWTAuthMiddleware(), middleware.UserAuth())
		i.GET("/balance/:user_id", ViewBalance)
		i.POST("/recharge/:user_id", Recharge)
		i.GET("/information/:user_id", ViewInformation)
		i.PUT("/modify/:user_id", ChangeInformation)
		i.POST("/address/add/:user_id", AddAddress)
		i.GET("/address/view/:user_id", ViewAddress)
		i.PUT("/address/update/:user_id", UpdateAddress)
		i.DELETE("/address/delete/:user_id", DeleteAddress)
	}
	//订单
	b := r.Group("/order/auth")
	{
		b.Use(middleware.JWTAuthMiddleware(), middleware.UserAuth())
		b.POST("/add/:user_id", GenOrder)
		b.POST("/settle/:user_id", SettleBill)
		b.GET("/view/:user_id", SearchOrder)
		b.PUT("/update/:user_id", UpdateOrderStatus)
		b.GET("/view/:user_id", ViewOrder)
	}
	//后台管理
	s := r.Group("/seller")
	{
		s.POST("/register", BackRegister)
		s.POST("/login", BackLogin)
		a := s.Group("/auth")
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
