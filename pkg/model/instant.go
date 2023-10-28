package model

import (
	"gorm.io/gorm"
)

type Instant struct {
	gorm.Model
	BasicUser
	InstantType  int
	Content      string
	RefOriginID  uint
	LikeCount    int
	CommentCount int
	ShareCount   int
	Likes        []Like    `gorm:"foreignKey:InstantID"`
	Comments     []Comment `gorm:"foreignKey:InstantID"`
}

type Feed struct {
	gorm.Model
	InstantID   uint
	UserID      uint
	IsCommented bool
	IsShared    bool
	Attitude    int
}
