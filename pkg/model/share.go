package model

import "time"

type Share struct {
	ShareID         string    `json:"commentID"       bson:"_id"`
	Created         time.Time `json:"created"`
	LastModified    time.Time `json:"lastModified"`
	InsID           string    `json:"insID"`
	ForwardedFromID string    `json:"ForwardedFromID"`
	UserID          string    `json:"userID"`
	Username        string    `json:"username"`
	Avatar          int       `json:"avatar"`
	Content         string    `json:"content"`
	Direct          bool      `json:"direct"`
}
