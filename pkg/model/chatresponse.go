package model

import "time"

type ChatResponse struct {
	ChatID       string    `json:"chatID" bson:"_id"`
	Created      time.Time `json:"created"`
	LastModified time.Time `json:"lastModified"`
	From         int       `json:"from"`
	Group        int       `json:"group"`
	Type         int       `json:"type"`
	LocalMsgSeq  int       `json:"localMsgSeq"`
	Content      string    `json:"content"`
}
