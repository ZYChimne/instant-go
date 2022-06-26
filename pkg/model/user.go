package model

import "time"

type User struct {
	UserID        int       `json:"userID"`
	MailBox       string    `json:"mailbox" form:"mailbox"`
	Phone         string    `json:"phone" form:"phone"`
	Password      string    `json:"password" form:"password"`
	Username      string    `json:"username"`
	CreateTime    time.Time `json:"createtime"`
	UpdateTime    time.Time `json:"updatetime"`
	Avatar        int       `json:"avatar"`
	Gender        int       `json:"gender"`
	Country       int       `json:"country"`
	Province      int       `json:"province"`
	City          int       `json:"city"`
	Birthday      time.Time `json:"birthday"`
	School        string    `json:"school"`
	Company       string    `json:"company"`
	Job           string    `json:"job"`
	MyMode        int       `json:"mymode"`
	Introduction  string    `json:"introduction"`
	CoverPhoto  int       `json:"coverphoto"`
	Tag          string  `json:"tag"`
}
