package model

import "time"

type Instant struct {
	InsID       int       `json:"insid"`
	UserID      int       `json:"userid"`
	CreateTime  time.Time `json:"createtime"`
	UpdateTime  time.Time `json:"updatetime"`
	Content     string    `json:"content"`
	RefOriginId int       `json:"reforigininsid"`
}
