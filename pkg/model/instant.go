package model

import "time"

type Instant struct {
	InsID       int       `json:"insID"`
	UserID      int       `json:"userID"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
	Content     string    `json:"content"`
	RefOriginId int       `json:"refOriginInsId"`
}
