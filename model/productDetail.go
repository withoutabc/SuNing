package model

type Style struct {
	StyleId   string `json:"style_id"`
	ProductId string `json:"product_id"`
	Product   string `json:"product"`
	Style     string `json:"style"`
}

type RespProducts struct {
	Status int       `json:"status"`
	Info   string    `json:"info"`
	Data   []Product `json:"product"`
}

type RespStyles struct {
	Status int     `json:"status"`
	Info   string  `json:"info"`
	Data   []Style `json:"style"`
}

type Cart struct {
	CartId    string `json:"cart_id"`
	UserId    string `json:"user_id"`
	Name      string `json:"name"`
	UnitPrice string `json:"unit_price"`
	Quantity  string `json:"quantity"`
	Price     string `json:"price"`
	Image     string `json:"image"`
}

type RespCart struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
	Data   []Cart `json:"cart"`
}

type Collection struct {
	CollectionId string `json:"collection_id"`
	UserId       string `json:"user_id"`
	Name         string `json:"name"`
}

type RespCollection struct {
	Status int          `json:"status"`
	Info   string       `json:"info"`
	Data   []Collection `json:"data"`
}
