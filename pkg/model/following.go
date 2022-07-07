package model

import "time"

type Following struct {
	UserID       string    `json:"userID"`
	FollowingID  string    `json:"followingID"`
	LastModified time.Time `json:"lastModified"`
}
