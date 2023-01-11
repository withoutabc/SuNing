package service

import (
	"suning/dao"
	"suning/model"
)

func SearchSellerByName(sellerName string) (s model.Seller, err error) {
	s, err = dao.SearchSellerByName(sellerName)
	return
}

func CreateSeller(s model.Seller) error {
	err := dao.InsertSeller(s)
	return err
}

func AddProduct(p model.Product) (err error) {
	err = dao.InsertProduct(p)
	return
}

func SearchNameBySeller(seller string) (products []model.Product, err error) {
	products, err = dao.SearchNameBySeller(seller)
	return
}

func UpdateProduct(p model.Product) (err error) {
	err = dao.UpdateProduct(p)
	return
}

func DeleteProduct(seller string, name string) (err error) {
	err = dao.DeleteProduct(seller, name)
	return
}
