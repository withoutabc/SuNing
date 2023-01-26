package model

type Announcement struct {
	AnnouncementId string `json:"announcement_id"`
	SellerId       string `json:"seller_id"`
	Title          string `json:"title"`
	Content        string `json:"content"`
}

type RespAnnouncement struct {
	Status int          `json:"status"`
	Info   string       `json:"info"`
	Data   Announcement `json:"data"`
}
