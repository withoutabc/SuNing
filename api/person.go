package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"suning/model"
	"suning/service"
	"suning/util"
)

func ViewBalance(c *gin.Context) {
	//获取用户名
	username, err := c.Cookie("username")
	if err != nil {
		util.RespUnauthorizedErr(c)
		return
	}
	if username == "" {
		util.RespUnauthorizedErr(c)
		return
	}
	//查询数据，传入结构体
	var a model.Account
	a, err = service.SearchBalancerFromUsername(username)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.ViewBalance(c, 200, "view balance successfully", a.Username, a.Balance)
}

func Recharge(c *gin.Context) {
	//获取用户名
	username, err := c.Cookie("username")
	if err != nil {
		util.RespUnauthorizedErr(c)
		return
	}
	if username == "" {
		util.RespUnauthorizedErr(c)
		return
	}
	//获取充值金额
	var account int
	account, err = strconv.Atoi(c.Query("account"))
	if err != nil {
		util.NormErr(c, 400, "invalid recharge")
		return
	}
	if account <= 0 {
		util.NormErr(c, 400, "invalid recharge")
		return
	}
	//查询用户目前余额
	var a model.Account
	a, err = service.SearchBalancerFromUsername(username)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RespInternalErr(c)
		return
	}
	//充值，写入数据库
	accounted := a.Balance + account
	err = service.RechargeToAccount(username, accounted)
	if err != nil {
		log.Printf("update password error:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.ViewBalance(c, 200, "recharge successfully", username, accounted)
}
