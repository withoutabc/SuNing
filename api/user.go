package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"suning/model"
	"suning/service"
	"suning/util"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//判断是否有效输入
	if username == "" || password == "" {
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
	//用户信息写入数据库
	err = service.CreateUser(model.User{
		Username: username,
		Password: password,
	})
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//有效输入
	if username == " " || password == "" {
		util.RespParamErr(c)
		return
	}
	//检索用户处理
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
	util.RespOK(c)
	//设置cookie
	c.SetCookie("if_login", "true", 3600, "/", "localhost", false, true)
}
