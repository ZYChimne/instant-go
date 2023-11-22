package model

import (
	"gorm.io/gorm"
)

type InstantLike struct {
	gorm.Model
	BasicUser
	InstantID string
	Attitude  int `gorm:"default:0"`
}

type CommentLike struct {
	gorm.Model
	BasicUser
	CommentID string
	Attitude  int `gorm:"default:0"`
}
