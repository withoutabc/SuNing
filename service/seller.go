package service

import (
	"suning/dao"
	"suning/model"
)

// SearchSellerByName 通过商品名称查询对应商品信息
func SearchSellerByName(seller string) (s model.Seller, err error) {
	s, err = dao.SearchSellerByName(seller)
	return
}

// CreateSeller 创建商家信息
func CreateSeller(s model.Seller) error {
	err := dao.InsertSeller(s)
	return err
}

// AddProduct 添加商品
func AddProduct(p model.Product) (err error) {
	err = dao.InsertProduct(p)
	return
}

// SearchNameBySellerId 根据商家id查询其商品
func SearchNameBySellerId(sellerId string) (products []model.Product, err error) {
	products, err = dao.SearchNameBySellerId(sellerId)
	return
}

// UpdateProduct 更新商品信息
func UpdateProduct(p model.Product) (err error) {
	err = dao.UpdateProduct(p)
	return
}

// DeleteProduct 删除商品
func DeleteProduct(sellerId string, name string) (err error) {
	err = dao.DeleteProduct(sellerId, name)
	return
}

// SearchSellerBySellerId 根据商家id查询商家名称
func SearchSellerBySellerId(sellerId string) (s model.Seller, err error) {
	s, err = dao.SearchSellerBySellerId(sellerId)
	return
}
