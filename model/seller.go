package model

type Seller struct {
	ID       int64  `json:"id"`
	Seller   string `json:"seller"`
	Password string `json:"password"`
}

type Product struct {
	Id       int    `json:"id"`
	Seller   string `json:"seller"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	Sales    string `json:"sales"`
	Rating   string `json:"rating"`
	Category string `json:"category"`
	Image    string `json:"image"`
}
