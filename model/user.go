package model

type User struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	Password string `json:"password"`
}
