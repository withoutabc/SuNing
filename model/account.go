package model

type Account struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	Balance  int    `json:"balance"`
}
