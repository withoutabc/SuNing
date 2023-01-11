package service

import (
	"suning/dao"
	"suning/model"
)

func SearchSellerByName(seller string) (s model.Seller, err error) {
	s, err = dao.SearchSellerByName(seller)
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

func SearchNameBySid(sid string) (products []model.Product, err error) {
	products, err = dao.SearchNameBySid(sid)
	return
}

func UpdateProduct(p model.Product) (err error) {
	err = dao.UpdateProduct(p)
	return
}

func DeleteProduct(sid string, name string) (err error) {
	err = dao.DeleteProduct(sid, name)
	return
}

func SearchSellerBySid(sid string) (s model.Seller, err error) {
	s, err = dao.SearchSellerBySid(sid)
	return
}
