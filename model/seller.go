package model

type Seller struct {
	SellerId int    `json:"seller_id"`
	Seller   string `json:"seller"`
	Password string `json:"password"`
}

type Product struct {
	ProductId int    `json:"product_id"`
	SellerId  string `json:"seller_id"`
	Seller    string `json:"seller"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	Sales     string `json:"sales"`
	Rating    string `json:"rating"`
	Category  string `json:"category"`
	Image     string `json:"image"`
}
