package service

import (
	"suning/dao"
	"suning/model"
)

func SearchStyleByProductId(productId string) (Styles []model.Style, err error) {
	Styles, err = dao.SearchStyleByProductId(productId)
	return
}

func SearchPriceByName(productId string) (price string, err error) {
	price, err = dao.SearchPriceByName(productId)
	return
}

func InsertCart(c model.Cart) (err error) {
	err = dao.InsertCart(c)
	return
}

func SearchCartByUserId(userId string) (Carts []model.Cart, err error) {
	Carts, err = dao.SearchCartByUserId(userId)
	return
}

func InsertCollection(userId, name string) (err error) {
	err = dao.InsertCollection(userId, name)
	return
}

func SearchIfNameExist(userId, name string) (exist bool, err error) {
	exist, err = dao.SearchIfNameExist(userId, name)
	return
}

func DeleteCart(userId, name string) (err error) {
	err = dao.DeleteCart(userId, name)
	return
}

func SearchCollectionByUserId(userId string) (Collections []model.Collection, err error) {
	Collections, err = dao.SearchCollectionByUserId(userId)
	return
}

func SearchIfCollectionExist(userId, name string) (exist bool, err error) {
	exist, err = dao.SearchIfCollectionExist(userId, name)
	return
}
