package service

import (
	"suning/dao"
	"suning/model"
)

func SearchBalancerFromUserId(uid string) (a model.Account, err error) {
	a, err = dao.SearchBalanceFromUserId(uid)
	return
}
func CreateAccount(a model.Account) error {
	err := dao.InsertAccount(a)
	return err
}
func RechargeToAccount(username string, accounted float64) (err error) {
	err = dao.UpdateAccount(username, accounted)
	return
}
func CreateInformation(username string, uid int) (err error) {
	err = dao.InsertInformation(username, uid)
	return
}
func SearchInformationByUserId(uid string) (i model.Information, err error) {
	i, err = dao.SearchInformationByUserId(uid)
	return
}
func ChangeInformation(i model.Information) (err error) {
	err = dao.UpdateInformation(i)
	return
}
