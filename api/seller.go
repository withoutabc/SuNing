package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"suning/model"
	"suning/service"
	"suning/util"
	"time"
)

func BackRegister(c *gin.Context) {
	sellerName := c.PostForm("sellerName")
	password := c.PostForm("password")
	confPassword := c.PostForm("confirmPassword")
	//判断是否有效输入
	if sellerName == "" || password == "" || confPassword == "" {
		util.RespParamErr(c)
		return
	}
	//检索数据库
	u, err := service.SearchSellerByName(sellerName)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search seller error:%v", err)
		util.RespInternalErr(c)
		return
	}
	//用户是否存在
	if u.SellerName != "" {
		util.NormErr(c, 300, "seller has existed")
		return
	}
	//两次密码是否一致
	if confPassword != password {
		util.NormErr(c, 300, "different password")
		return
	}
	//用户信息写入数据库
	err = service.CreateSeller(model.Seller{
		SellerName: sellerName,
		Password:   password,
	})
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}

func BackLogin(c *gin.Context) {
	sellerName := c.PostForm("sellerName")
	password := c.PostForm("password")
	//有效输入
	if sellerName == " " || password == "" {
		util.RespParamErr(c)
		return
	}
	//查找用户
	u, err := service.SearchSellerByName(sellerName)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "user don't exist")
		} else {
			log.Printf("search seller error:%v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}
	//密码错误
	if u.Password != password {
		util.NormErr(c, 300, "wrong password")
		return
	}
	util.RespOK(c)
	//设置cookie
	c.SetCookie("sellerName", sellerName, 3600, "/", "localhost", false, true)
}

func BackLogout(c *gin.Context) {
	//检测是否登录
	sellerName, err := c.Cookie("sellerName")
	if err != nil {
		util.RespUnauthorizedErr(c)
		return
	}
	if sellerName == "" {
		util.RespUnauthorizedErr(c)
		return
	}
	//清除登陆状态cookie
	c.SetCookie("sellerName", "", -1, "/", "localhost", false, true)
	util.RespOK(c)
}

func BackRefresh(c *gin.Context) {
	//判断cookie过没过期
	sellerName, err := c.Cookie("sellerName")
	if err != nil {
		util.RespUnauthorizedErr(c)
		c.Abort()
		return
	}
	if sellerName == "" {
		util.RespUnauthorizedErr(c)
		c.Abort()
		return
	}
	//没过期
	c.Next()
	// 设置新的cookie
	expiration := time.Now().Add(time.Hour)
	c.SetCookie("sellerName", sellerName, int(expiration.Unix()), "/", "localhost", false, true)
}

func ViewProduct(c *gin.Context) {

}

func AddProduct(c *gin.Context) {

}

func UpdateProduct(c *gin.Context) {

}

func DeleteProduct(c *gin.Context) {

}
