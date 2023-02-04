package model

type Account struct {
	UserId   int     `json:"user_id"`
	Username string  `json:"username"`
	Balance  float64 `json:"balance"`
}

type RespBalance struct {
	Status int     `json:"status"`
	Info   string  `json:"info"`
	Data   Account `json:"data"`
}
