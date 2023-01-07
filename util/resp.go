package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type respTemplate struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
	Data   struct {
	} `json:"data"`
}

var OK = respTemplate{
	Status: 200,
	Info:   "success",
	Data:   struct{}{},
}

func RespOK(c *gin.Context) {
	c.JSON(http.StatusOK, OK)
}

var ParamError = respTemplate{
	Status: 400,
	Info:   "params error",
	Data:   struct{}{},
}

func RespParamErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, ParamError)
}

var InternalErr = respTemplate{
	Status: 500,
	Info:   "internal error",
	Data:   struct{}{},
}

func RespInternalErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, InternalErr)
}

func NormErr(c *gin.Context, status int, info string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"info":   info,
	})
}
