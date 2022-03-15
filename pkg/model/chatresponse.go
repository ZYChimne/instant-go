package model

import "time"

type ChatResponse struct {
	ChatID      int       `json:"chatid"`
	Time        time.Time `json:"time"`
	From        int       `json:"from"`
	Group       int       `json:"group"`
	Type        int       `json:"type"`
	LocalMsgSeq int       `json:"localmsgseq"`
	Content     string    `json:"content"`
}
