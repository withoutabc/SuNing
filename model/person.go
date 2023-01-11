package model

type Information struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Gender   string `json:"gender"`
	PhoneNum string `json:"phone_num"`
	Email    string `json:"email"`
	Year     string `json:"year"`
	Month    string `json:"month"`
	Day      string `json:"day"`
	Avatar   string `json:"avatar"`
}
