package model

type Information struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username" `
	Nickname string `json:"nickname"`
	Gender   string `json:"gender" `
	PhoneNum string `json:"phone_num"`
	Email    string `json:"email" `
	Year     string `json:"year" `
	Month    string `json:"month"`
	Day      string `json:"day"`
	Avatar   string `json:"avatar"`
}
