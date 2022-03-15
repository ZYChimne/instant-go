package model

import "time"

type Comment struct {
	CommentID  int       `json:"commentid"`
	CreateTime time.Time `json:"createtime"`
	UpdateTime time.Time `json:"updatetime"`
	InsID      int       `json:"insid"`
	UserID     int       `json:"userid"`
	Content    string    `json:"content"`
}
