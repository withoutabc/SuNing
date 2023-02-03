package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
	"suning/model"
	"suning/service"
	"suning/util"
)

// SearchStyle 实现查看商品款式接口
func SearchStyle(c *gin.Context) {
	//获取商品id
	productId := c.Param("product_id")
	if productId == "" {
		util.RespParamErr(c)
		return
	}
	//查找款式
	Styles, err := service.SearchStyleByProductId(productId)
	if err != nil {
		fmt.Printf("search styles err:%v", err)
		util.RespInternalErr(c)
		return
	}
	c.JSON(http.StatusOK, model.RespStyles{
		Status: 200,
		Info:   "search style success",
		Data:   Styles,
	})
}

// AddToCart 实现加入购物车接口
func AddToCart(c *gin.Context) {
	//获取用户id
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	//获取商品,数量
	name := c.Query("name")
	quantity := c.Query("quantity")
	if name == "" || quantity == "" {
		util.RespParamErr(c)
		return
	}
	//name不重复
	exist, err := service.SearchIfNameExist(userId, name)
	if err != nil {
		fmt.Printf("search name exist err:%v", err)
		util.RespInternalErr(c)
		return
	}
	if exist {
		util.NormErr(c, 400, "repeated name")
		return
	}
	//数据转化
	unitPrice, err := service.SearchPriceByName(name)
	if err != nil {
		fmt.Printf("search product err:%v", err)
		util.RespInternalErr(c)
		return
	}
	Quantity, err := strconv.ParseFloat(quantity, 64)
	if err != nil {
		util.NormErr(c, 400, "convert err")
		return
	}
	UnitPrice, err := strconv.ParseFloat(unitPrice, 64)
	if err != nil {
		util.NormErr(c, 400, "convert err")
		return
	}
	Price := math.Round(Quantity * UnitPrice)
	price := strconv.FormatFloat(Price, 'f', 2, 64)
	//插入数据
	err = service.InsertCart(model.Cart{
		UserId:    userId,
		Name:      name,
		UnitPrice: unitPrice,
		Quantity:  quantity,
		Price:     price,
	})
	util.RespOK(c, "add to cart success")
}

// ViewCart 实现查看购物车接口
func ViewCart(c *gin.Context) {
	//获取用户id
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	//查看购物车
	carts, err := service.SearchCartByUserId(userId)
	if err != nil {
		fmt.Printf("search cart err:%v", err)
		util.RespInternalErr(c)
		return
	}
	c.JSON(http.StatusOK, model.RespCart{
		Status: 200,
		Info:   "view cart success",
		Data:   carts,
	})
}

// DeleteCart 实现删除购物车接口
func DeleteCart(c *gin.Context) {
	//获取用户id
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	//获取商品
	name := c.Query("name")
	if name == "" {
		util.RespParamErr(c)
		return
	}
	//name存在
	exist, err := service.SearchIfNameExist(userId, name)
	if err != nil {
		fmt.Printf("search name exist err:%v", err)
		util.RespInternalErr(c)
		return
	}
	if !exist {
		util.NormErr(c, 400, "not exist name")
		return
	}
	//删除
	err = service.DeleteCart(userId, name)
	if err != nil {
		fmt.Printf("delete cart err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "delete cart success")
}

// AddToCollection 实现添加收藏接口
func AddToCollection(c *gin.Context) {
	//获取用户id
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	//获取商品名称
	name := c.Query("name")
	if name == "" {
		util.RespParamErr(c)
		return
	}
	//collection不重复
	exist, err := service.SearchIfCollectionExist(userId, name)
	if err != nil {
		fmt.Printf("search collection exist err:%v", err)
		util.RespInternalErr(c)
		return
	}
	if exist {
		util.NormErr(c, 400, "repeated name")
		return
	}
	//添加至我的收藏
	err = service.InsertCollection(userId, name)
	if err != nil {
		fmt.Printf("insert collection err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "add collection success")
}

// ViewCollection 实现查看收藏接口
func ViewCollection(c *gin.Context) {
	//获取用户id
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	//查看收藏
	collections, err := service.SearchCollectionByUserId(userId)
	if err != nil {
		fmt.Printf("view collection err:%v", err)
		util.RespInternalErr(c)
		return
	}
	c.JSON(http.StatusOK, model.RespCollection{
		Status: 200,
		Info:   "view collection success",
		Data:   collections,
	})
}

// DeleteCollection 实现删除收藏接口
func DeleteCollection(c *gin.Context) {
	//获取用户id
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	//获取name
	name := c.Query("name")
	if name == "" {
		util.RespParamErr(c)
		return
	}
	//collection存在
	exist, err := service.SearchIfCollectionExist(userId, name)
	if err != nil {
		fmt.Printf("search collection exist err:%v", err)
		util.RespInternalErr(c)
		return
	}
	if !exist {
		util.NormErr(c, 400, "not exist name")
		return
	}
	//删除收藏
	err = service.DeleteCollection(userId, name)
	if err != nil {
		fmt.Printf("delete collection exist err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "delete collection success")
}

// ViewProductDetail 实现查看商品详情接口
func ViewProductDetail(c *gin.Context) {
	//获取商品id
	productId := c.Param("product_id")
	if productId == "" {
		util.RespParamErr(c)
		return
	}
	//查看商品详情
	detail, err := service.SearchDetailByProductId(productId)
	if err != nil {
		fmt.Printf("search detail err:%v", err)
		util.RespInternalErr(c)
		return
	}
	c.JSON(http.StatusOK, model.RespDetail{
		Status: 200,
		Info:   "view product detail success",
		Data:   detail,
	})
}

// GiveReview 实现评价接口
func GiveReview(c *gin.Context) {
	//获取用户id
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	//获取评论信息
	review := model.Review{
		UserId:    userId,
		ProductId: c.PostForm("product_id"),
		Content:   c.PostForm("content"),
		Rating:    c.PostForm("rating"),
	}
	if review.ProductId == "" || review.Content == "" || review.Rating == "" {
		util.RespParamErr(c)
		return
	}
	name, err := service.SearchNameByProductId(review.ProductId)
	if err != nil {
		fmt.Printf("search name err:%v", err)
		util.RespInternalErr(c)
		return
	}
	review.Name = name
	//插入数据
	err = service.InsertReview(review)
	if err != nil {
		fmt.Printf("insert view err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "give review success")
}

// ViewReview 实现查看评价接口
func ViewReview(c *gin.Context) {
	//获取商品id
	productId := c.Param("product_id")
	if productId == "" {
		util.RespParamErr(c)
		return
	}
	//查看评价
	reviews, err := service.SearchReviewByProductId(productId)
	if err != nil {
		fmt.Printf("search review err:%v", err)
		util.RespInternalErr(c)
		return
	}
	c.JSON(http.StatusOK, model.RespReview{
		Status: 200,
		Info:   "view review success",
		Data:   reviews,
	})
}
