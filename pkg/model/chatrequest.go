package model

import "time"

type ChatRequest struct {
	Token       string    `json:"token"`
	LocalMsgSeq int       `json:"localMsgSeq"`
	SendTime    time.Time `json:"sendTime"`
	From        int       `json:"from"`
	Group       int       `json:"group"`
	Content     string    `json:"content"`
}
