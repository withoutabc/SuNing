package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"suning/model"
	"suning/service"
	"suning/util"
	"time"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	confPassword := c.PostForm("confirmPassword")
	//判断是否有效输入
	if username == "" || password == "" || confPassword == "" {
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
	if confPassword != password {
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
			util.NormErr(c, 300, "user don't exist")
		} else {
			log.Printf("search user error:%v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}
	//创建账户
	err = service.CreateAccount(model.Account{
		Username: username,
		Balance:  0,
		Uid:      int(u.Uid),
	})
	if err != nil {
		fmt.Printf("create account err:%v", err)
		util.RespInternalErr(c)
		return
	}
	//创建信息表
	err = service.CreateInformation(username, int(u.Uid))
	if err != nil {
		log.Printf("create information:%v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	//有效输入
	if username == " " || password == "" {
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
	util.ViewUser(c, "login success", u)
	//设置cookie
	c.SetCookie("uid", strconv.Itoa(u.Uid), 3600, "/", "localhost", false, true)
}

func Logout(c *gin.Context) {
	//检测是否登录
	uid, err := c.Cookie("uid")
	if err != nil {
		util.RespUnauthorizedErr(c)
		return
	}
	if uid == "" {
		util.RespUnauthorizedErr(c)
		return
	}
	//清除登陆状态cookie
	c.SetCookie("uid", "", -1, "/", "localhost", false, true)
	util.RespOK(c)
}

func UserRefresh(c *gin.Context) {
	//判断cookie过没过期
	uid, err := c.Cookie("uid")
	if err != nil {
		util.RespUnauthorizedErr(c)
		c.Abort()
		return
	}
	if uid == "" {
		util.RespUnauthorizedErr(c)
		c.Abort()
		return
	}
	//没过期
	c.Next()
	// 设置新的cookie
	expiration := time.Now().Add(time.Hour)
	c.SetCookie("uid", uid, int(expiration.Unix()), "/", "localhost", false, true)
	util.RespOK(c)
}
