package main

import (
	"suning/api"
	"suning/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
