package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"suning/model"
	"suning/service"
	"suning/util"
)

func BackRegister(c *gin.Context) {
	seller := c.PostForm("seller")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")
	//判断是否输入
	if seller == "" || password == "" || confirmPassword == "" {
		util.RespParamErr(c)
		return
	}
	//检索数据库
	u, err := service.SearchSellerByName(seller)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search seller error:%v", err)
		util.RespInternalErr(c)
		return
	}
	//用户是否存在
	if u.Seller != "" {
		util.NormErr(c, 300, "seller has existed")
		return
	}
	//两次密码是否一致
	if confirmPassword != password {
		util.NormErr(c, 300, "different password")
		return
	}
	//用户信息写入数据库
	err = service.CreateSeller(model.Seller{
		Seller:   seller,
		Password: password,
	})
	if err != nil {
		fmt.Printf("create seller err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}

func BackLogin(c *gin.Context) {
	seller := c.PostForm("seller")
	password := c.PostForm("password")
	//有效输入
	if seller == " " || password == "" {
		util.RespParamErr(c)
		return
	}
	//查找用户
	s, err := service.SearchSellerByName(seller)
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
	if s.Password != password {
		util.NormErr(c, 300, "wrong password")
		return
	}
	//密码正确
	aToken, rToken, _ := service.GenToken(strconv.Itoa(s.Sid), "seller")
	c.JSON(http.StatusOK, model.RespLogin{
		Status: 200,
		Info:   "login success",
		Data: model.Login{
			Uid:          s.Sid,
			Token:        aToken,
			RefreshToken: rToken,
		},
	})
}

func BackLogout(c *gin.Context) {
	//检测是否登录
	sid, err := c.Cookie("sid")
	if err != nil {
		util.RespUnauthorizedErr(c)
		return
	}
	if sid == "" {
		util.RespUnauthorizedErr(c)
		return
	}
	//清除登陆状态cookie
	c.SetCookie("sid", "", -1, "/", "localhost", false, true)
	util.RespOK(c)
}

func BackRefresh(c *gin.Context) {
	//refresh_token
	rToken := c.PostForm("refresh_token")
	if rToken == "" {
		util.RespParamErr(c)
		return
	}
	_, err := service.ParseToken(rToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 2005,
			"info":   "无效的token",
		})
		return
	}
	//生成新的token
	newAToken, newRToken, err := service.RefreshToken(rToken)
	if err != nil {
		fmt.Printf("refresh err:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"info":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.RespToken{
		Status: 200,
		Info:   "refresh token success",
		Data: model.Token{
			Token:        newAToken,
			RefreshToken: newRToken,
		},
	})
}

func ViewProduct(c *gin.Context) {
	//获取卖家名称
	sid, err := c.Cookie("sid")
	if err != nil {
		util.RespUnauthorizedErr(c)
		return
	}
	//查询卖家的商品
	var products []model.Product
	products, err = service.SearchNameBySid(sid)
	if err != nil {
		fmt.Printf("view product err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.ViewProducts(c, "view product successfully", products)
}

func AddProduct(c *gin.Context) {
	//获取卖家名称
	sid, err := c.Cookie("sid")
	if err != nil {
		util.RespUnauthorizedErr(c)
		return
	}
	if sid == "" {
		util.RespUnauthorizedErr(c)
		return
	}
	//根据uid查找seller
	var s model.Seller
	s, err = service.SearchSellerBySid(sid)
	if err != nil {
		fmt.Printf("search seller:%v", err)
		util.RespInternalErr(c)
		return
	}
	//获取商品信息
	p := model.Product{
		Sid:      sid,
		Seller:   s.Seller,
		Name:     c.PostForm("name"),
		Price:    c.PostForm("price"),
		Sales:    c.PostForm("sales"),
		Rating:   c.PostForm("rating"),
		Category: c.PostForm("category"),
		Image:    c.PostForm("image"),
	}
	if p.Name == "" || p.Price == "" || p.Sales == "" || p.Rating == "" || p.Category == "" {
		util.RespParamErr(c)
		return
	}
	//同一商家的商品名不可重复
	var products []model.Product
	products, err = service.SearchNameBySid(sid)
	for _, product := range products {
		if product.Name == p.Name {
			util.NormErr(c, 400, "product has existed")
			return
		}
	}
	//插入商品信息
	err = service.AddProduct(p)
	if err != nil {
		fmt.Printf("add product:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}

func UpdateProduct(c *gin.Context) {
	//获取卖家名称
	sid, err := c.Cookie("sid")
	if err != nil {
		util.RespUnauthorizedErr(c)
		return
	}
	if sid == "" {
		util.RespUnauthorizedErr(c)
		return
	}
	//获取要修改的数据
	p := model.Product{
		Sid:      sid,
		Name:     c.PostForm("name"),
		Price:    c.PostForm("price"),
		Sales:    c.PostForm("sales"),
		Rating:   c.PostForm("rating"),
		Category: c.PostForm("category"),
		Image:    c.PostForm("image"),
	}
	//判断name是否写了
	if p.Name == "" {
		util.NormErr(c, 400, "unknown name")
		return
	}
	//判断name是否存在
	var products []model.Product
	var count int
	products, err = service.SearchNameBySid(sid)
	for _, product := range products {
		if product.Name == p.Name {
			count = 1
			break
		}
	}
	if count != 1 {
		util.NormErr(c, 400, "product not exist")
		return
	}
	//判断是否有修改
	if p.Price == "" && p.Sales == "" && p.Rating == "" && p.Category == "" && p.Image == "" {
		util.NormErr(c, 400, "fail to update")
		return
	}
	//修改数据
	err = service.UpdateProduct(p)
	if err != nil {
		log.Printf("update product err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}

func DeleteProduct(c *gin.Context) {
	//获取卖家名称
	sid, err := c.Cookie("sid")
	if err != nil {
		util.RespUnauthorizedErr(c)
		return
	}
	if sid == "" {
		util.RespUnauthorizedErr(c)
		return
	}
	//获取要删除的商品名称
	name := c.Query("name")
	//判断name是否存在
	var products []model.Product
	var count int
	products, err = service.SearchNameBySid(sid)
	for _, product := range products {
		if product.Name == name {
			count = 1
			break
		}
	}
	if count != 1 {
		util.NormErr(c, 400, "product not exist")
		return
	}
	//删除商品
	err = service.DeleteProduct(sid, name)
	if err != nil {
		log.Printf("delete product err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}
