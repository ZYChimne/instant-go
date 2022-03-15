package model

import "time"

type Like struct {
	LikeID     int       `json:"likeid"`
	CreateTime time.Time `json:"createtime"`
	UpdateTime time.Time `json:"updatetime"`
	InsID      int       `json:"insid"`
	UserID     int       `json:"userid"`
	Attitude   int       `json:"attitude"`
}
