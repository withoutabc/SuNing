package model

type Order struct {
	OrderId          string `json:"order_id"`
	OrderNumber      string `json:"order_number"`
	OrderTime        string `json:"order_time"`
	Status           string `json:"status"`
	PaymentMethod    string `json:"payment_method"`
	PaymentAmount    string `json:"payment_amount"`
	PaymentTime      string `json:"payment_time"`
	RecipientName    string `json:"recipient_name"`
	RecipientAddress string `json:"recipient_address"`
	RecipientPhone   string `json:"recipient_phone"`
}
