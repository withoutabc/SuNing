package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"suning/model"
)

func ViewBalance(c *gin.Context, info string, a model.Account) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   info,
		"data":   a,
	})
}

func ViewInformation(c *gin.Context, info string, i model.Information) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   info,
		"data":   i,
	})
}

func ViewProducts(c *gin.Context, info string, p []model.Product) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   info,
		"data":   p,
	})
}
