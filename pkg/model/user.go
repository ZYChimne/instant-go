package model

import "time"

type User struct {
	UserID       string    `json:"userID" bson:"_id"`
	MailBox      string    `json:"mailbox" form:"mailbox"`
	Phone        string    `json:"phone" form:"phone"`
	Password     string    `json:"password" form:"password"`
	Username     string    `json:"username"`
	Created      time.Time `json:"created"`
	LastModified time.Time `json:"lastModified"`
	Avatar       int       `json:"avatar"`
	Gender       int       `json:"gender"`
	Country      int       `json:"country"`
	Province     int       `json:"province"`
	City         int       `json:"city"`
	Birthday     time.Time `json:"birthday"`
	School       string    `json:"school"`
	Company      string    `json:"company"`
	Job          string    `json:"job"`
	MyMode       int       `json:"myMode"`
	Introduction string    `json:"introduction		"`
	CoverPhoto   int       `json:"coverPhoto"`
	Tags         []string  `json:"tags"`
	Followings   int	   `json:"followings"`
	Followers    int	   `json:"followers"`
}
