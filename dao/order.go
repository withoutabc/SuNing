package dao

import (
	"database/sql"
	"suning/model"
	"time"
)

func Payment(userId string, payProducts []string) (totalPrice float64, err error) {
	var pay float64
	for i := 0; i < len(payProducts); i++ {
		var row *sql.Row
		row = DB.QueryRow("select price from cart where user_id=? and name=?", userId, payProducts[i])
		err = row.Scan(&pay)
		if err != nil {
			return 0, err
		}
		totalPrice += pay
	}
	return
}

func DeleteCartName(userId string, payProducts []string) (err error) {
	for i := 0; i < len(payProducts); i++ {
		_, err = DB.Exec("delete * from cart where user_id=? and name=? ", userId, payProducts[i])
		if err != nil {
			return err
		}
	}
	return
}

func SearchOrder(userId, status string) (orders []model.Order, err error) {
	var rows *sql.Rows
	rows, err = DB.Query("select * from order where user_id=? and status=?", userId, status)
	var order model.Order
	for rows.Next() {
		err = rows.Scan(&order.OrderId, &order.OrderNumber, &order.OrderTime, &order.Status, &order.PaymentMethod, &order.PaymentAmount, &order.PaymentTime, &order.RecipientName, &order.RecipientAddress, &order.RecipientPhone)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return
}

func ViewOrder(userId string) (orders []model.Order, err error) {
	var rows *sql.Rows
	rows, err = DB.Query("select * from order where user_id=?", userId)
	var order model.Order
	for rows.Next() {
		err = rows.Scan(&order.OrderId, &order.OrderNumber, &order.OrderTime, &order.Status, &order.PaymentMethod, &order.PaymentAmount, &order.PaymentTime, &order.RecipientName, &order.RecipientAddress, &order.RecipientPhone)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return
}

func UpdateOrderStatus(orderId, status string) (err error) {
	_, err = DB.Exec("update order set status=? where order_id=?", status, orderId)
	return
}

func UpdatePaymentMethodTime(orderId, paymentMethod string, paymentTime time.Time) (err error) {
	_, err = DB.Exec("update order set payment_method=?,payment_time=? where order_id=?", paymentMethod, paymentTime, orderId)
	return
}

func SearchCartInOrder(userId, name string) (cart model.Cart, err error) {
	row := DB.QueryRow("select * from cart where user_id=? and name=?", userId, name)
	if err = row.Err(); err != nil {
		return model.Cart{}, err
	}
	err = row.Scan(&cart.CartId, &cart.UserId, &cart.Name, &cart.UnitPrice, &cart.Quantity, &cart.Price, &cart.Image)
	return
}

func InsertOrder(o model.Order) (err error) {
	_, err = DB.Exec("insert into order (order_number,order_time,status,payment_method,payment_amount,payment_time,recipient_name,recipient_address,recipient_phone,user_id) values (?,?,?,?,?,?,?,?,?,?)", o.OrderNumber, o.OrderTime, o.Status, o.PaymentMethod, o.PaymentAmount, o.PaymentTime, o.RecipientName, o.RecipientAddress, o.RecipientPhone, o.UserId)
	return
}

func SearchOrderIdByNumber(orderNumber string) (orderId string, err error) {
	row := DB.QueryRow("select order_id from order where order_number=?", orderNumber)
	if err = row.Err(); err != nil {
		return "", err
	}
	err = row.Scan(&orderId)
	return
}

func InsertOrderDetail(orderId string, c model.Cart) (err error) {
	_, err = DB.Exec("insert into order_detail (order_id,name,quantity,price) values (?,?,?,?)", orderId, c.Name, c.Quantity, c.Price)
	return
}

func SearchOrderDetailByOrderId(orderId string) (orderDetails []model.OrderDetail, err error) {
	var rows *sql.Rows
	rows, err = DB.Query("select * from order_detail where order_id=?", orderId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var orderDetail model.OrderDetail
		err = rows.Scan(&orderDetail.OrderDetailId, &orderDetail.OrderId, &orderDetail.Name, &orderDetail.Quantity, &orderDetail.Price)
		if err != nil {
			return nil, err
		}
		orderDetails = append(orderDetails, orderDetail)
	}
	return orderDetails, nil
}

func SearchTotalPriceById(orderId string) (totalPrice string, err error) {
	row := DB.QueryRow("select payment_amount from order where order_id=?", orderId)
	if err = row.Err(); err != nil {
		return "", err
	}
	err = row.Scan(&totalPrice)
	return totalPrice, err
}

func DeleteOrder(orderId string) (err error) {
	_, err = DB.Exec("delete * from order where order_id=?", orderId)
	return
}

func DeleteOrderDetail(orderId string) (err error) {
	_, err = DB.Exec("delete * from order_detail where order_id=?", orderId)
	return
}
