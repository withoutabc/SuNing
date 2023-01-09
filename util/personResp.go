package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ViewBalance(c *gin.Context, status int, info string, username string, balance int) {
	type Data struct {
		Username string `json:"username"`
		Balance  int    `json:"balance"`
	}
	var data = Data{
		Username: username,
		Balance:  balance,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"info":   info,
		"data":   data,
	})
}
