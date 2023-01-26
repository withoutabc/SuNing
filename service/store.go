package service

import (
	"suning/dao"
	"suning/model"
)

func InsertAnnouncement(sellerId string) (err error) {
	err = dao.InsertAnnouncement(sellerId)
	return
}

func UpdateAnnouncement(sellerId, title, content string) (err error) {
	err = dao.UpdateAnnouncement(sellerId, title, content)
	return
}

func ViewAnnouncement(sellerId string) (announcement model.Announcement, err error) {
	announcement, err = dao.ViewAnnouncement(sellerId)
	return
}

func Sort(sellerId, sortBy, order string) (products []model.Product, err error) {
	products, err = dao.StoreSort(sellerId, sortBy, order)
	return
}

func StoreCategory(sellerId, category string) (products []model.Product, err error) {
	products, err = dao.StoreCategory(sellerId, category)
	return
}
