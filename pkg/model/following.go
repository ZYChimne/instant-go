package model

import "time"

type Following struct {
	UserID       string    `json:"userID"`
	FollowingID  string    `json:"followingID"`
	Created      time.Time `json:"created"`
	LastModified time.Time `json:"lastModified"`
}
