package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	BasicUser
	InstantID string
	ReplyToID string
	Content   string
}
