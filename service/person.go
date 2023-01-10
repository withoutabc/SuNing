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
func CreateInformation(username string) (err error) {
	err = dao.InsertInformation(username)
	return
}
func SearchInformationByUsername(username string) (i model.Information, err error) {
	i, err = dao.SearchInformationByUsername(username)
	return
}
func ChangeInformation(i model.Information) (err error) {
	err = dao.UpdateInformation(i)
	return
}
