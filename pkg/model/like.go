package model

import "time"

type Like struct {
	LikeID       string    `json:"likeID" bson:"_id"`
	Created      time.Time `json:"created"`
	LastModified time.Time `json:"lastModified"`
	InsID        string    `json:"insID"`
	UserID       string    `json:"userID"`
	Attitude     int       `json:"attitude"`
}
