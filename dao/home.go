package dao

import (
	"database/sql"
	"suning/model"
)

func SearchByKeyword(keyword string) (products []model.Product, err error) {
	var rows *sql.Rows
	rows, err = DB.Query("select * from product where name like ?", "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var product model.Product
		if err = rows.Scan(&product.ProductId, &product.SellerId, &product.Name, &product.Price, &product.Sales, &product.Rating, &product.Category, &product.Image, &product.Seller); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return
}

func SearchAndSort(keyword, sortBy, order string) (products []model.Product, err error) {
	var rows *sql.Rows
	rows, err = DB.Query("select * from product where name like ?"+" order by "+sortBy+" "+order, "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var product model.Product
		if err = rows.Scan(&product.ProductId, &product.SellerId, &product.Name, &product.Price, &product.Sales, &product.Rating, &product.Category, &product.Image, &product.Seller); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return
}

func Category(category string) (products []model.Product, err error) {
	var rows *sql.Rows
	rows, err = DB.Query("select * from product where category=?", category)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var product model.Product
		if err = rows.Scan(&product.ProductId, &product.SellerId, &product.Name, &product.Price, &product.Sales, &product.Rating, &product.Category, &product.Image, &product.Seller); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return
}
