package model

import "time"

type Order struct {
	OrderId          string    `json:"order_id"`
	OrderNumber      string    `json:"order_number"`
	OrderTime        time.Time `json:"order_time"`
	Status           string    `json:"status"`
	PaymentMethod    string    `json:"payment_method"`
	PaymentAmount    float64   `json:"payment_amount"`
	PaymentTime      string    `json:"payment_time"`
	RecipientName    string    `json:"recipient_name"`
	RecipientAddress string    `json:"recipient_address"`
	RecipientPhone   string    `json:"recipient_phone"`
	UserId           string    `json:"user_id"`
}

type RespOrder struct {
	Status      int             `json:"status"`
	Info        string          `json:"info"`
	Order       []Order         `json:"order"`
	OrderDetail [][]OrderDetail `json:"order_detail"`
}

type OrderDetail struct {
	OrderDetailId string `json:"order_detail_id"`
	OrderId       string `json:"order_id"`
	Name          string `json:"name"`
	Quantity      string `json:"quantity"`
	Price         string `json:"price"`
}
