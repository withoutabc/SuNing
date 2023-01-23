package api

import "github.com/gin-gonic/gin"

// SearchStyle 实现查看商品款式接口
func SearchStyle(c *gin.Context) {
	productId := c.Param("product_id")

}

// AddToCart 实现加入购物车接口
func AddToCart(c *gin.Context) {
	userId := c.Param("user_id")
}

// ViewCart 实现查看购物车接口
func ViewCart(c *gin.Context) {
	userId := c.Param("user_id")
}

// AddToCollection 实现添加我的收藏接口
func AddToCollection(c *gin.Context) {
	userId := c.Param("user_id")
}

// ViewCollection 实现查看我的收藏接口
func ViewCollection(c *gin.Context) {
	userId := c.Param("user_id")

}

// ViewProductDetail 实现查看商品详情接口
func ViewProductDetail(c *gin.Context) {
	productId := c.Param("product_id")
}

// GiveComment 实现评价接口
func GiveComment(c *gin.Context) {
	productId := c.Param("product_id")
}

// ViewComment 实现查看评价接口
func ViewComment(c *gin.Context) {
	productId := c.Param("product_id")
}
