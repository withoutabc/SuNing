package model

type Account struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Balance  int    `json:"balance"`
}
