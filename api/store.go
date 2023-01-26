package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"suning/model"
	"suning/service"
	"suning/util"
)

// UpdateAnnouncement 实现了更新公告接口
func UpdateAnnouncement(c *gin.Context) {
	//获取商家id
	sellerId := c.Param("seller_id")
	if sellerId == "" {
		util.RespParamErr(c)
		return
	}
	//获取标题、内容，可为空
	title := c.PostForm("title")
	content := c.PostForm("content")
	//插入
	err := service.UpdateAnnouncement(sellerId, title, content)
	if err != nil {
		fmt.Printf("update announcement err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "update announcement success")
}

// ViewAnnouncement 实现了查看公告接口
func ViewAnnouncement(c *gin.Context) {
	//获取商家id
	sellerId := c.Param("seller_id")
	if sellerId == "" {
		util.RespParamErr(c)
		return
	}
	//查看公告
	announcement, err := service.ViewAnnouncement(sellerId)
	if err != nil {
		fmt.Printf("view announcement err:%v", err)
		util.RespInternalErr(c)
		return
	}
	c.JSON(http.StatusOK, model.RespAnnouncement{
		Status: 200,
		Info:   "view announcement success",
		Data:   announcement,
	})
}

// StoreSort 实现规则排序接口
func StoreSort(c *gin.Context) {
	//获取商家id
	sellerId := c.Param("seller_id")
	if sellerId == "" {
		util.RespParamErr(c)
		return
	}
	//获取排序规则
	sortBy := c.DefaultQuery("sort_by", "sales")
	order := c.DefaultQuery("order", "desc")
	//排序
	products, err := service.Sort(sellerId, sortBy, order)
	if err != nil {
		fmt.Printf("sort products err:%v", err)
		util.RespInternalErr(c)
		return
	}
	c.JSON(http.StatusOK, model.RespProducts{
		Status: 200,
		Info:   "search products success",
		Data:   products,
	})
}

// StoreCategory 实现商品分类接口
func StoreCategory(c *gin.Context) {
	//获取商家id
	sellerId := c.Param("seller_id")
	if sellerId == "" {
		util.RespParamErr(c)
		return
	}
	//获取分类规则
	category := c.Query("category")
	if category == "" {
		util.RespParamErr(c)
		return
	}
	//分类
	products, err := service.StoreCategory(sellerId, category)
	if err != nil {
		fmt.Printf("sort products err:%v", err)
		util.RespInternalErr(c)
		return
	}
	c.JSON(http.StatusOK, model.RespProducts{
		Status: 200,
		Info:   "category products success",
		Data:   products,
	})
}
