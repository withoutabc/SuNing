package dao

import (
	"fmt"
	"strings"
	"suning/model"
)

func SearchSellerByName(seller string) (s model.Seller, err error) {
	row := DB.QueryRow("select * from seller where seller=?", seller)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&s.SellerId, &s.Seller, &s.Password)
	return
}

func InsertSeller(s model.Seller) (err error) {
	_, err = DB.Exec("insert into seller (seller,password) values (?,?)", s.Seller, s.Password)
	return
}

func InsertProduct(p model.Product) (err error) {
	_, err = DB.Exec("insert into product (seller_id,seller,name,price,sales,rating,category,image) values (?,?,?,?,?,?,?,?)", p.SellerId, p.Seller, p.Name, p.Price, p.Sales, p.Rating, p.Category, p.Image)
	return
}

func SearchNameBySellerId(sellerId string) (products []model.Product, err error) {
	rows, err := DB.Query("select * from product where seller_id=?", sellerId)
	if err != nil {
		return nil, err
	}
	//处理查询结果
	for rows.Next() {
		var p model.Product
		if err = rows.Scan(&p.ProductId, &p.SellerId, &p.Name, &p.Price, &p.Sales, &p.Rating, &p.Category, &p.Image, &p.Seller); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

// UpdateProduct 数据库修改商品信息
func UpdateProduct(p model.Product) (err error) {
	var sql strings.Builder
	var arg []interface{}
	sql.WriteString("update product set")
	if p.Price != "" {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" price=?")
		arg = append(arg, p.Price)
	}
	if p.Sales != "" {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" sales=?")
		arg = append(arg, p.Sales)
	}
	if p.Rating != "" {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" rating=?")
		arg = append(arg, p.Rating)
	}
	if p.Category != "" {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" category=?")
		arg = append(arg, p.Category)
	}
	if p.Image != "" {
		if len(arg) > 0 {
			sql.WriteString(",")
		}
		sql.WriteString(" image=?")
		arg = append(arg, p.Image)
	}
	sql.WriteString(" where seller_id=? and name=?")
	arg = append(arg, p.SellerId)
	arg = append(arg, p.Name)
	fmt.Println(sql.String())
	_, err = DB.Exec(sql.String(), arg...)
	return
}

// DeleteProduct 数据库删除商品
func DeleteProduct(sellerId string, name string) (err error) {
	_, err = DB.Exec("delete from product where seller_id=? and name=?", sellerId, name)
	return
}

// SearchSellerBySellerId 数据库查找
func SearchSellerBySellerId(sellerId string) (s model.Seller, err error) {
	row := DB.QueryRow("select * from seller where seller_id=?", sellerId)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&s.SellerId, &s.Seller, &s.Password)
	return
}
