package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"suning/model"
)

func ViewBalance(c *gin.Context, status int, info string, a model.Account) {
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"info":   info,
		"data":   a,
	})
}

func ViewInformation(c *gin.Context, status int, info string, i model.Information) {
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"info":   info,
		"data":   i,
	})
}
