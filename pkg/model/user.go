package model

import "time"

type User struct {
	UserID        int       `json:"userid"`
	MailBox       string    `json:"mailbox"`
	Phone         string    `json:"phone"`
	Password      string    `json:"password"`
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
	Mymode        int       `json:"mymode"`
	Introduction  string    `json:"introduction"`
	ProfileImage  int       `json:"profileimage"`
	Tag           []string  `json:"tag"`
	Following     []string  `json:"following"`
	Followed      []string  `json:"followed"`
	BlackList     []string  `json:"blacklist"`
	InsIDList     []int     `json:"insidlist"`
	LikeIDList    []int     `json:"likeidlist"`
	CommentIDList []int     `json:"commentidlist"`
}
