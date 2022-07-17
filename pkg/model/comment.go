package model

import "time"

type Comment struct {
	CommentID    string    `json:"commentID"    bson:"_id"`
	Created      time.Time `json:"created"`
	LastModified time.Time `json:"lastModified"`
	InsID        string    `json:"insID"`
	ReplyToID    string    `json:"replyToID"`
	UserID       string    `json:"userID"`
	Content      string    `json:"content"`
	Direct       bool      `json:"direct"`
}
