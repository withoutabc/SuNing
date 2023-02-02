package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"suning/model"
	"suning/service"
	"suning/util"
)

// Register 实现了用户注册接口
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")
	//判断是否有效输入
	if username == "" || password == "" || confirmPassword == "" {
		util.RespParamErr(c)
		return
	}
	//检索数据库
	u, err := service.SearchUserByUsername(username)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RespInternalErr(c)
		return
	}
	//用户是否存在
	if u.Username != "" {
		util.NormErr(c, 300, "user has existed")
		return
	}
	//两次密码是否一致
	if confirmPassword != password {
		util.NormErr(c, 300, "different password")
		return
	}
	//用户信息写入数据库
	err = service.CreateUser(model.User{
		Username: username,
		Password: password,
	})
	if err != nil {
		fmt.Printf("create user err:%v", err)
		util.RespInternalErr(c)
		return
	}
	//查找用户
	u, err = service.SearchUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "user has exist")
		} else {
			log.Printf("search user error:%v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}
	//创建账户
	err = service.CreateAccount(model.Account{
		UserId:   u.UserId,
		Username: username,
		Balance:  0,
	})
	if err != nil {
		fmt.Printf("create account err:%v", err)
		util.RespInternalErr(c)
		return
	}
	//创建信息表
	err = service.CreateInformation(username, int(u.UserId))
	if err != nil {
		log.Printf("create information:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "register success")
}

// Login 实现了用户登录接口
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	//有效输入
	if username == "" || password == "" {
		util.RespParamErr(c)
		return
	}
	//查找用户
	u, err := service.SearchUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "user don't exist")
		} else {
			log.Printf("search user error:%v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}
	//密码错误
	if u.Password != password {
		util.NormErr(c, 300, "wrong password")
		return
	}
	//密码正确
	aToken, rToken, err := service.GenToken(strconv.Itoa(u.UserId), "user")
	if err != nil {
		fmt.Printf("refresh err:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"info":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.RespLogin{
		Status: 200,
		Info:   "login success",
		Data: model.Login{
			Uid:          u.UserId,
			Token:        aToken,
			RefreshToken: rToken,
		},
	})
}

// Refresh 实现了刷新token接口
func Refresh(c *gin.Context) {
	//refresh_token
	rToken := c.PostForm("refresh_token")
	if rToken == "" {
		util.RespParamErr(c)
		return
	}
	_, err := service.ParseToken(rToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 2005,
			"info":   "无效的token",
		})
		return
	}
	//生成新的token
	newAToken, newRToken, err := service.RefreshToken(rToken)
	if err != nil {
		fmt.Printf("refresh err:%v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"info":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.RespToken{
		Status: 200,
		Info:   "refresh token success",
		Data: model.Token{
			Token:        newAToken,
			RefreshToken: newRToken,
		},
	})
}
