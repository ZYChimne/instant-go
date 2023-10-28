package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `gorm:"not null;uniqueIndex"`
	Phone          string `gorm:"not null;uniqueIndex"`
	Password       string
	Username       string `gorm:"not null;uniqueIndex"`
	Nickname       string
	UserType       int
	Avatar         string
	Gender         int
	Country        string
	State          string
	City           string
	ZipCode        string
	Birthday       time.Time
	School         string
	Company        string
	Job            string
	MyMode         string
	Introduction   string
	CoverPhoto     string
	FollowingCount int
	FollowerCount  int
	Followings     []Following `gorm:"foreignKey:UserID"`
	Followers      []Following `gorm:"foreignKey:TargetID"`
}

type BasicUser struct {
	UserID   uint
	Username string
	Nickname string
	Avatar   string
	UserType int
}
