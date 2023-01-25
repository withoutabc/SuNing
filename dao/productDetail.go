package dao

import (
	"database/sql"
	"suning/model"
)

func SearchStyleByProductId(productId string) (Styles []model.Style, err error) {
	var rows *sql.Rows
	rows, err = DB.Query("select * from product_style where product_id=?", productId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var Style model.Style
		if err = rows.Scan(&Style.StyleId, &Style.ProductId, &Style.Product, &Style.Style); err != nil {
			return nil, err
		}
		Styles = append(Styles, Style)
	}
	return
}

func SearchPriceByName(Name string) (price string, err error) {
	row := DB.QueryRow("select price from product where name=?", Name)
	if err = row.Err(); err != nil {
		return
	}
	err = row.Scan(&price)
	return
}

func InsertCart(c model.Cart) (err error) {
	_, err = DB.Exec("insert into cart (user_id,name,unit_price,quantity,price) values (?,?,?,?,?)", c.UserId, c.Name, c.UnitPrice, c.Quantity, c.Price)
	return
}

func SearchCartByUserId(userId string) (Carts []model.Cart, err error) {
	var rows *sql.Rows
	rows, err = DB.Query("select * from cart where user_id=?", userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var Cart model.Cart
		if err = rows.Scan(&Cart.CartId, &Cart.UserId, &Cart.Name, &Cart.UnitPrice, &Cart.Quantity, &Cart.Price, &Cart.Image); err != nil {
			return nil, err
		}
		Carts = append(Carts, Cart)
	}
	return
}

func InsertCollection(userId, name string) (err error) {
	_, err = DB.Exec("insert into collection (user_id,name) values (?,?)", userId, name)
	return
}

func SearchIfNameExist(userId, name string) (exist bool, err error) {
	row := DB.QueryRow("select count(*) from cart where user_id=? and name=?", userId, name)
	if err = row.Err(); err != nil {
		return
	}
	var count int
	err = row.Scan(&count)
	if err != nil {
		return
	}
	return count > 0, nil
}

func DeleteCart(userId, name string) (err error) {
	_, err = DB.Exec("delete * from cart where user_id=? and name=?", userId, name)
	return
}

func SearchCollectionByUserId(userId string) (Collections []model.Collection, err error) {
	var rows *sql.Rows
	rows, err = DB.Query("select * from collection where user_id=?", userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var Collection model.Collection
		if err = rows.Scan(&Collection.CollectionId, &Collection.UserId, &Collection.Name); err != nil {
			return nil, err
		}
		Collections = append(Collections, Collection)
	}
	return
}

func SearchIfCollectionExist(userId, name string) (exist bool, err error) {
	row := DB.QueryRow("select count(*) from collection where user_id=? and name=?", userId, name)
	if err = row.Err(); err != nil {
		return
	}
	var count int
	err = row.Scan(&count)
	if err != nil {
		return
	}
	return count > 0, nil
}
