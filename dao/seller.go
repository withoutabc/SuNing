package dao

import (
	"fmt"
	"strings"
	"suning/model"
)

func SearchSellerByName(sellName string) (s model.Seller, err error) {
	row := DB.QueryRow("select * from seller where sellerName=?", sellName)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&s.ID, &s.Seller, &s.Password)
	return
}

func InsertSeller(s model.Seller) (err error) {
	_, err = DB.Exec("insert into seller (seller,password) values (?,?)", s.Seller, s.Password)
	return
}

func InsertProduct(p model.Product) (err error) {
	_, err = DB.Exec("insert into product (seller,name,price,sales,rating,category,image) values (?,?,?,?,?,?,?)", p.Seller, p.Name, p.Price, p.Sales, p.Rating, p.Category, p.Image)
	return
}

func SearchNameBySeller(seller string) (products []model.Product, err error) {
	rows, err := DB.Query("select * from product where seller=?", seller)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//处理查询结果
	for rows.Next() {
		var p model.Product
		if err = rows.Scan(&p.Id, &p.Seller, &p.Name, &p.Price, &p.Sales, &p.Rating, &p.Category, &p.Image); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

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
	sql.WriteString(" where seller=? and name=?")
	arg = append(arg, p.Seller)
	arg = append(arg, p.Name)
	fmt.Println(sql.String())
	_, err = DB.Exec(sql.String(), arg...)
	return
}

func DeleteProduct(seller string, name string) (err error) {
	_, err = DB.Exec("delete from product where seller=? and name=?", seller, name)
	return
}
