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
		if err = rows.Scan(&product.Pid, &product.Sid, &product.Name, &product.Price, &product.Sales, &product.Rating, &product.Category, &product.Image, &product.Seller); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return
}

func Sort(sortBy, order string) (products []model.Product, err error) {
	var rows *sql.Rows
	rows, err = DB.Query("select * from product order by " + sortBy + " " + order)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var product model.Product
		if err = rows.Scan(&product.Pid, &product.Sid, &product.Name, &product.Price, &product.Sales, &product.Rating, &product.Category, &product.Image, &product.Seller); err != nil {
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
		if err = rows.Scan(&product.Pid, &product.Sid, &product.Name, &product.Price, &product.Sales, &product.Rating, &product.Category, &product.Image, &product.Seller); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return
}
