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

func DecreaseBalance(userId string, price float64) (err error) {
	err = dao.DecreaseBalance(userId, price)
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

func AddAddress(a model.Address) (err error) {
	err = dao.AddAddress(a)
	return
}

func SearchAddress(userId string) (addresses []model.Address, err error) {
	addresses, err = dao.SearchAddress(userId)
	return
}

func SearchAddressById(addressId string) (address model.Address, err error) {
	address, err = dao.SearchAddressById(addressId)
	return
}

func UpdateAddress(a model.Address) (err error) {
	err = dao.UpdateAddress(a)
	return
}

func DeleteAddress(addressId string) (err error) {
	err = dao.DeleteAddress(addressId)
	return
}
