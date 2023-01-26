package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"suning/model"
	"suning/service"
	"suning/util"
)

// SearchAndSort 实现搜索和规则排序接口
func SearchAndSort(c *gin.Context) {
	//获取关键字
	keyword := c.Query("keyword")
	if keyword == "" {
		util.RespParamErr(c)
		return
	}
	//获取排序规则
	sortBy := c.DefaultQuery("sort_by", "sales")
	order := c.DefaultQuery("order", "desc")
	//排序
	products, err := service.SearchAndSort(keyword, sortBy, order)
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

// Category 实现分类展示接口
func Category(c *gin.Context) {
	//获取类别
	category := c.Query("category")
	if category == "" {
		util.RespParamErr(c)
		return
	}
	//分类
	products, err := service.Category(category)
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
