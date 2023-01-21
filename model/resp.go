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

type Products struct {
}
