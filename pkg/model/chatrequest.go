package model

import "time"

type ChatRequest struct {
	Token       string    `json:"token"`
	LocalMsgSeq int       `json:"localmsgseq"`
	SendTime    time.Time `json:"sendtime"`
	From        int       `json:"from"`
	Group       int       `json:"group"`
	Content     string    `json:"content"`
}
