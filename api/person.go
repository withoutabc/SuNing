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
	util.ViewBalance(c, "view balance successfully", a)
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
		log.Printf("search account error:%v", err)
		util.RespInternalErr(c)
		return
	}
	//充值，写入数据库
	accounted := a.Balance + account
	err = service.RechargeToAccount(username, accounted)
	if err != nil {
		log.Printf("recharge error:%v", err)
		util.RespInternalErr(c)
		return
	}
	a = model.Account{
		Username: username,
		Balance:  accounted,
	}
	util.ViewBalance(c, "recharge successfully", a)
}

func ViewInformation(c *gin.Context) {
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
	//查询数据库
	var i model.Information
	i, err = service.SearchInformationByUsername(username)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search information error:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.ViewInformation(c, "view information successfully", i)
}

func ChangeInformation(c *gin.Context) {
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
	//传入要修改的信息
	i := model.Information{
		Username: username,
		Nickname: c.PostForm("nickname"),    //最多15个字符
		Gender:   c.PostForm("gender"),      //0保密 1男 2女
		PhoneNum: c.PostForm("phoneNumber"), //11位
		Email:    c.PostForm("email"),       //
		Year:     c.PostForm("year"),        //4位
		Month:    c.PostForm("month"),       //1-12
		Day:      c.PostForm("day"),         //1-31
		Avatar:   c.PostForm("avatar"),
	}
	//分析是否符合格式
	//if len(i.Nickname) > 15 {
	//	util.NormErr(c, 400, "nickname over 15")
	//	return
	//}
	//if i.Gender != "" {
	//	if i.Gender != "0" && i.Gender != "1" && i.Gender != "2" {
	//		util.NormErr(c, 400, "invalid gender")
	//		return
	//	}
	//}
	//
	//if i.PhoneNum != "" {
	//	if len(i.PhoneNum) != 11 {
	//		util.NormErr(c, 400, "invalid phone number")
	//		return
	//	}
	//}
	//if i.Email != "" {
	//	// 定义一个正则表达式，用于匹配邮箱地址
	//	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	//	// 使用MatchString函数判断字符串是否符合正则表达式
	//	var match bool
	//	match, err = regexp.MatchString(pattern, i.Email)
	//	if err != nil {
	//		util.NormErr(c, 400, "invalid email")
	//		return
	//	}
	//	if match != true {
	//		util.NormErr(c, 400, "invalid email")
	//		return
	//	}
	//}
	//插入到数据库
	err = service.ChangeInformation(i)
	if err != nil {
		log.Printf("change information err:%v", err)
		util.RespInternalErr(c)
		return
	}
	////查询数据库
	//i, err = service.SearchInformationByUsername(username)
	//if err != nil && err != sql.ErrNoRows {
	//	log.Printf("search information error:%v", err)
	//	util.RespInternalErr(c)
	//	return
	//}
	util.RespOK(c)
}
