package service

import (
	"errors"
	"suning/dao"
	"suning/model"
)

// UserService 是一个处理用户业务逻辑的结构体
type UserService struct {
	dao dao.UserDao
}

// Register 使用 dao 层的 Register 方法将用户注册信息保存到数据库
func (s *UserService) Register(username, password string) error {
	if len(username) == 0 || len(password) == 0 {
		return errors.New("username and password are required")
	}
	return s.dao.Register(username, password)
}

// Login 使用 dao 层的 Login 方法检查用户名和密码是否匹配
func (s *UserService) Login(username, password string) (bool, error) {
	if len(username) == 0 || len(password) == 0 {
		return false, errors.New("username and password are required")
	}
	return s.dao.Login(username, password)
}

func SearchUserByUsername(username string) (u model.User, err error) {
	u, err = dao.SearchUserByUsername(username)
	return
}

func CreateUser(u model.User) error {
	err := dao.InsertUser(u)
	return err
}
