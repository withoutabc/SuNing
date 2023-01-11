package service

import (
	"suning/dao"
	"suning/model"
)

func SearchBalancerFromUid(uid string) (a model.Account, err error) {
	a, err = dao.SearchBalanceFromUid(uid)
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
func CreateInformation(username string, uid int) (err error) {
	err = dao.InsertInformation(username, uid)
	return
}
func SearchInformationByUid(uid string) (i model.Information, err error) {
	i, err = dao.SearchInformationByUid(uid)
	return
}
func ChangeInformation(i model.Information) (err error) {
	err = dao.UpdateInformation(i)
	return
}
