package service

import (
	"suning/dao"
	"suning/model"
	"time"
)

func SearchOrder(userId, status string) (orders []model.Order, err error) {
	orders, err = dao.SearchOrder(userId, status)
	return
}

func UpdateOrderStatus(orderId, status string) (err error) {
	err = dao.UpdateOrderStatus(orderId, status)
	return
}

func UpdatePaymentMethodTime(orderId, paymentMethod string, paymentTime time.Time) (err error) {
	err = dao.UpdatePaymentMethodTime(orderId, paymentMethod, paymentTime)
	return
}

func GenOrder(userId string, payProducts []string) (totalPrice float64, err error) {
	totalPrice, err = dao.Payment(userId, payProducts)
	if err != nil {
		return 0, err
	}
	err = dao.DeleteCartName(userId, payProducts)
	return totalPrice, err
}

func SearchCartInOrder(userId string, payProducts []string) (carts []model.Cart, err error) {
	for k := range payProducts {
		var cart model.Cart
		cart, err = dao.SearchCartInOrder(userId, payProducts[k])
		if err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}
	return carts, nil
}

func InsertOrder(o model.Order) (err error) {
	err = dao.InsertOrder(o)
	return
}

func SearchOrderIdByNumber(orderNumber string) (orderId string, err error) {
	orderId, err = dao.SearchOrderIdByNumber(orderNumber)
	return
}

func InsertOrderDetail(orderId string, c model.Cart) (err error) {
	err = dao.InsertOrderDetail(orderId, c)
	return
}

func SearchOrderDetailByOrderId(orders []model.Order) (orderDetailsS [][]model.OrderDetail, err error) {
	for _, order := range orders {
		var orderDetails []model.OrderDetail
		orderDetails, err = dao.SearchOrderDetailByOrderId(order.OrderId)
		if err != nil {
			return nil, err
		}
		orderDetailsS = append(orderDetailsS, orderDetails)
	}
	return orderDetailsS, nil
}

func SearchTotalPriceById(orderId string) (totalPrice string, err error) {
	totalPrice, err = dao.SearchTotalPriceById(orderId)
	return
}

func ViewOrder(userId string) (orders []model.Order, err error) {
	orders, err = dao.ViewOrder(userId)
	return
}
