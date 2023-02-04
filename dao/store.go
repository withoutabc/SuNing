package dao

import (
	"database/sql"
	"suning/model"
)

func InsertAnnouncement(sellerId string) (err error) {
	_, err = DB.Exec("insert into announcement (seller_id,content) values (?,?)", sellerId, "暂时还未发布公告哦。")
	return
}

func UpdateAnnouncement(sellerId, title, content string) (err error) {
	_, err = DB.Exec("update announcement set title=?,content=? where seller_id=?", title, content, sellerId)
	return
}

func ViewAnnouncement(sellerId string) (announcement model.Announcement, err error) {
	var row *sql.Row
	row = DB.QueryRow("select * from announcement where seller_id=?", sellerId)
	err = row.Scan(&announcement.AnnouncementId, &announcement.SellerId, &announcement.Title, &announcement.Content)
	return
}

func StoreSort(sellerId, sortBy, order string) (products []model.Product, err error) {
	var rows *sql.Rows
	rows, err = DB.Query("select * from product where seller_id=? "+" order by "+sortBy+" "+order, sellerId)
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

func StoreCategory(sellerId, category string) (products []model.Product, err error) {
	var rows *sql.Rows
	rows, err = DB.Query("select * from product where category=? and seller_id=?", category, sellerId)
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
