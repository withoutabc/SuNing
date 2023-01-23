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

// BackRegister 实现了后台注册接口
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
	util.RespOK(c, "register success")
}

// BackLogin 实现了后台登录接口
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
			util.NormErr(c, 300, "user not exist")
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
	aToken, rToken, err := service.GenToken(strconv.Itoa(s.SellerId), "seller")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"info":   "generate token error",
		})
	}
	c.JSON(http.StatusOK, model.RespLogin{
		Status: 200,
		Info:   "login success",
		Data: model.Login{
			Uid:          s.SellerId,
			Token:        aToken,
			RefreshToken: rToken,
		},
	})
}

func BackLogout(c *gin.Context) {
}

//// BackRefresh 实现了后台刷新token接口
//func BackRefresh(c *gin.Context) {
//	//refresh_token
//	rToken := c.PostForm("refresh_token")
//	if rToken == "" {
//		util.RespParamErr(c)
//		return
//	}
//	_, err := service.ParseToken(rToken)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"status": 2005,
//			"info":   "无效的token",
//		})
//		return
//	}
//	//生成新的token
//	newAToken, newRToken, err := service.RefreshToken(rToken)
//	if err != nil {
//		fmt.Printf("refresh err:%v", err)
//		c.JSON(http.StatusBadRequest, gin.H{
//			"status": 400,
//			"info":   err.Error(),
//		})
//		return
//	}
//	c.JSON(http.StatusOK, model.RespToken{
//		Status: 200,
//		Info:   "refresh token success",
//		Data: model.Token{
//			Token:        newAToken,
//			RefreshToken: newRToken,
//		},
//	})
//}

// ViewProduct 实现了后台查看商品接口
func ViewProduct(c *gin.Context) {
	//获取卖家id
	SellerId := c.Param("seller_id")
	//查询卖家的商品
	products, err := service.SearchNameBySellerId(SellerId)
	if err != nil {
		fmt.Printf("view product err:%v", err)
		util.RespInternalErr(c)
		return
	}
	c.JSON(http.StatusOK, model.RespProducts{
		Status: 200,
		Info:   "view products success",
		Data:   products,
	})
}

// AddProduct 实现了后台添加商品接口
func AddProduct(c *gin.Context) {
	//获取卖家id
	SellerId := c.Param("seller_id")
	//根据sid查找seller
	s, err := service.SearchSellerBySellerId(SellerId)
	if err != nil {
		fmt.Printf("search seller:%v", err)
		util.RespInternalErr(c)
		return
	}
	//获取商品信息
	p := model.Product{
		SellerId: SellerId,
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
	products, err = service.SearchNameBySellerId(SellerId)
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
	util.RespOK(c, "add product success")
}

// UpdateProduct 实现了后台修改商品信息接口
func UpdateProduct(c *gin.Context) {
	//获取卖家id
	SellerId := c.Param("seller_id")
	//获取要修改的数据
	p := model.Product{
		SellerId: SellerId,
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
	var count int
	products, err := service.SearchNameBySellerId(SellerId)
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
	//判断修改后信息是否有变化
	//修改数据
	err = service.UpdateProduct(p)
	if err != nil {
		log.Printf("update product err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "update product success")
}

// DeleteProduct 实现了后台删除商品接口
func DeleteProduct(c *gin.Context) {
	//获取卖家id
	SellerId := c.Param("seller_id")
	//获取要删除的商品名称
	name := c.Query("name")
	//判断name是否存在
	var count int
	products, err := service.SearchNameBySellerId(SellerId)
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
	err = service.DeleteProduct(SellerId, name)
	if err != nil {
		log.Printf("delete product err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "delete product success")
}
