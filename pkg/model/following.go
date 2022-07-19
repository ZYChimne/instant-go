package model

import "time"

type Following struct {
	UserID       string    `json:"userID"`
	FollowingID  string    `json:"followingID"`
	IsFriend     bool      `json:"isFriend"`
	Created      time.Time `json:"created"`
	LastModified time.Time `json:"lastModified"`
}
