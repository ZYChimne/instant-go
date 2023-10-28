package model

import (
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	BasicUser
	InstantID string
	Attitude  int
}
