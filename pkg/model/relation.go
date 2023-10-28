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
	ID          uint
	UserID      uint
	TargetID    uint
	IsFriend    bool
	UserType int
	Username    string
	Nickname    string
	Avatar      string
}
