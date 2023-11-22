package model

import (
	"gorm.io/gorm"
)

type Instant struct {
	gorm.Model
	BasicUser
	InstantType  int `gorm:"default:0"`
	Content      string
	Visibility   int `gorm:"default:3"` // 0: public, 1: followers, 2: friends, 3: private
	RefOriginID  uint
	LikeCount    int           `gorm:"default:0"`
	CommentCount int           `gorm:"default:0"`
	ShareCount   int           `gorm:"default:0"`
	Likes        []InstantLike `gorm:"foreignKey:InstantID"`
	Comments     []Comment     `gorm:"foreignKey:InstantID"`
}

// A combination of speed and space efficiency
// Feed Table only stores the instantID rather than the whole content
type Feed struct {
	gorm.Model
	InstantID   uint
	UserID      uint
	IsCommented bool `gorm:"default:false"`
	IsShared    bool `gorm:"default:false"`
	Attitude    int  `gorm:"default:0"`
}

type UpsertInstant struct {
	ID          uint
	UserID      uint
	InstantType int
	Content     string
	Visibility  int
	RefOriginID uint
}
