package model

type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RespToken struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
	Data   Token  `json:"data"`
}

type RespLogin struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
	Data   Login  `json:"data"`
}

type Login struct {
	Uid          int    `json:"uid"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RespProducts struct {
	Status int       `json:"status"`
	Info   string    `json:"info"`
	Data   []Product `json:"product"`
}

type RespBalance struct {
	Status int     `json:"status"`
	Info   string  `json:"info"`
	Data   Account `json:"data"`
}

type RespInformation struct {
	Status int         `json:"status"`
	Info   string      `json:"info"`
	Data   Information `json:"data"`
}
