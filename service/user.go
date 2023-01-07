package service

import (
	"suning/dao"
	"suning/model"
)

func SearchUserByUsername(username string) (u model.User, err error) {
	u, err = dao.SearchUserByUsername(username)
	return
}

func CreateUser(u model.User) error {
	err := dao.InsertUser(u)
	return err
}
