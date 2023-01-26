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

type RespInformation struct {
	Status int         `json:"status"`
	Info   string      `json:"info"`
	Data   Information `json:"data"`
}

type Address struct {
	AddressId        string `json:"address_id"`
	UserId           string `json:"user_id"`
	RecipientName    string `json:"recipient_name"`
	RecipientPhone   string `json:"recipient_phone"`
	Province         string `json:"province"`
	City             string `json:"city"`
	StateOrCommunity string `json:"state_or_community"`
}

type RespAddress struct {
	Status int       `json:"status"`
	Info   string    `json:"info"`
	Data   []Address `json:"address"`
}
