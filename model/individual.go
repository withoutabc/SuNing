package model

type Information struct {
	Uid      int    `json:"uid" form:"uid"`
	Username string `json:"username" form:"uid"`
	Nickname string `json:"nickname" form:"uid"`
	Gender   string `json:"gender" form:"uid"`
	PhoneNum string `json:"phone_num" form:"uid"`
	Email    string `json:"email" form:"uid"`
	Year     string `json:"year" form:"uid"`
	Month    string `json:"month" form:"uid"`
	Day      string `json:"day" form:"uid"`
	Avatar   string `json:"avatar" form:"uid"`
}
