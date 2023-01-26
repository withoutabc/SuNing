package service

import (
	"suning/dao"
	"suning/model"
)

func SearchProduct(keyword string) (products []model.Product, err error) {
	products, err = dao.SearchByKeyword(keyword)
	return
}

func SearchAndSort(keyword, sortBy, order string) (products []model.Product, err error) {
	products, err = dao.SearchAndSort(keyword, sortBy, order)
	return
}

func Category(category string) (products []model.Product, err error) {
	products, err = dao.Category(category)
	return
}
