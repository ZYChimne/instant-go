package model

import "time"

type Instant struct {
	InsID        string    `json:"insID" bson:"_id"`
	UserID       string    `json:"userID"`
	Created      time.Time `json:"created"`
	LastModified time.Time `json:"lastModified"`
	Content      string    `json:"content"`
	RefOriginID  string    `json:"refOriginID"`
	Likes        int       `json:"likes"`
	Shares       int       `json:"shares"`
}
