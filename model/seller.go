package model

type Seller struct {
	Sid      int    `json:"sid"`
	Seller   string `json:"seller"`
	Password string `json:"password"`
}

type Product struct {
	Pid      int    `json:"pid"`
	Sid      string `json:"sid"`
	Seller   string `json:"seller"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	Sales    string `json:"sales"`
	Rating   string `json:"rating"`
	Category string `json:"category"`
	Image    string `json:"image"`
}
