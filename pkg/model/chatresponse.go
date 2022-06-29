package model

import "time"

type ChatResponse struct {
	ChatID      int       `json:"chatID"`
	Time        time.Time `json:"time"`
	From        int       `json:"from"`
	Group       int       `json:"group"`
	Type        int       `json:"type"`
	LocalMsgSeq int       `json:"localMsgSeq"`
	Content     string    `json:"content"`
}
