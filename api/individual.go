package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"suning/model"
	"suning/service"
	"suning/util"
)

// ViewBalance 实现了查看账户余额接口
func ViewBalance(c *gin.Context) {
	//获取uid
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	//查询数据，传入结构体
	a, err := service.SearchBalancerFromUserId(userId)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RespInternalErr(c)
		return
	}
	c.JSON(http.StatusOK, model.RespBalance{
		Status: 200,
		Info:   "view balance success",
		Data:   a,
	})
}

// Recharge 实现了充值接口
func Recharge(c *gin.Context) {
	//获取uid
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	//获取充值金额
	account, err := strconv.ParseFloat(c.Query("account"), 64)
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
	a, err = service.SearchBalancerFromUserId(userId)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search account error:%v", err)
		util.RespInternalErr(c)
		return
	}
	initialAccount := float64(a.Balance)
	//充值，写入数据库
	accounted := initialAccount + account
	err = service.RechargeToAccount(userId, accounted)
	if err != nil {
		log.Printf("recharge error:%v", err)
		util.RespInternalErr(c)
		return
	}
	//查询用户目前余额
	a, err = service.SearchBalancerFromUserId(userId)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search account error:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "recharge success")
}

// ViewInformation 实现了查看个人信息接口
func ViewInformation(c *gin.Context) {
	//获取uid
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	//查询数据库
	i, err := service.SearchInformationByUserId(userId)
	if err != nil {
		log.Printf("search information error:%v", err)
		util.RespInternalErr(c)
		return
	}
	c.JSON(http.StatusOK, model.RespInformation{
		Status: 200,
		Info:   "view information success",
		Data:   i,
	})
}

// ChangeInformation 实现了修改个人信息接口
func ChangeInformation(c *gin.Context) {
	//获取uid
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	UserId, err := strconv.Atoi(userId)
	//传入要修改的信息
	i := model.Information{
		Nickname: c.PostForm("nickname"),
		Gender:   c.PostForm("gender"),
		PhoneNum: c.PostForm("phone_num"),
		Email:    c.PostForm("email"),
		Year:     c.PostForm("year"),
		Month:    c.PostForm("month"),
		Day:      c.PostForm("day"),
		Avatar:   c.PostForm("avatar"),
		UserId:   UserId,
	}
	if err != nil {
		util.RespParamErr(c)
		return
	}
	//插入到数据库
	if i.Nickname == "" && i.Gender == "" && i.PhoneNum == "" && i.Email == "" && i.Year == "" && i.Day == "" {
		util.NormErr(c, 400, "fail to update")
		return
	}
	err = service.ChangeInformation(i)
	if err != nil {
		log.Printf("change information err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "change information success")
}

// AddAddress 实现新增地址接口
func AddAddress(c *gin.Context) {
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	//获取地址
	address := model.Address{
		UserId:           userId,
		RecipientName:    c.PostForm("name"),
		RecipientPhone:   c.PostForm("phone"),
		Province:         c.PostForm("province"),
		City:             c.PostForm("city"),
		StateOrCommunity: c.PostForm("street_or_community"),
	}
	//写入数据库
	if address.RecipientName == "" || address.RecipientPhone == "" || address.Province == "" || address.City == "" || address.StateOrCommunity == "" {
		util.RespParamErr(c)
		return
	}
	err := service.AddAddress(address)
	if err != nil {
		log.Printf("add address err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "add address success")
}

// ViewAddress 实现查看地址接口
func ViewAddress(c *gin.Context) {
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	//查找地址
	addresses, err := service.SearchAddress(userId)
	if err != nil {
		log.Printf("search address err:%v", err)
		util.RespInternalErr(c)
		return
	}
	c.JSON(http.StatusOK, model.RespAddress{
		Status: 200,
		Info:   "view address success",
		Data:   addresses,
	})
}

// UpdateAddress 实现修改地址接口
func UpdateAddress(c *gin.Context) {
	addressId := c.Param("address_id")
	if addressId == "" {
		util.RespParamErr(c)
		return
	}
	//
	//获取地址
	address := model.Address{
		AddressId:        addressId,
		RecipientName:    c.PostForm("name"),
		RecipientPhone:   c.PostForm("phone"),
		Province:         c.PostForm("province"),
		City:             c.PostForm("city"),
		StateOrCommunity: c.PostForm("street_or_community"),
	}
	//修改地址
	err := service.UpdateAddress(address)
	if err != nil {
		log.Printf("update address err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "update address success")
}

// DeleteAddress 实现删除地址接口
func DeleteAddress(c *gin.Context) {
	addressId := c.Param("address_id")
	if addressId == "" {
		util.RespParamErr(c)
		return
	}
	//删除地址
	err := service.DeleteAddress(addressId)
	if err != nil {
		log.Printf("delete address err:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "delete address success")
}
