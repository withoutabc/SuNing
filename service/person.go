package service

import (
	"suning/dao"
	"suning/model"
)

func SearchBalancerFromUsername(username string) (a model.Account, err error) {
	a, err = dao.SearchBalanceFromUsername(username)
	return
}
func CreateAccount(a model.Account) error {
	err := dao.InsertAccount(a)
	return err
}
func RechargeToAccount(username string, accounted int) (err error) {
	err = dao.UpdateAccount(username, accounted)
	return
}
