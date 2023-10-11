package model

import "gorm.io/gorm"

type Following struct {
	gorm.Model
	UserID     uint `gorm:"uniqueIndex:compositeIndex"`
	TargetID   uint `gorm:"uniqueIndex:compositeIndex"`
	IsFriend   bool
	TargetType int
}

type JointFollowing struct {
	ID          uint   `json:"followingID"       gorm:"primaryKey"`
	UserID      uint   `json:"userID"`
	TargetID    uint   `json:"targetID"`
	IsFriend    bool   `json:"isFriend"`
	AccountType int    `json:"accountType"`
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
}
