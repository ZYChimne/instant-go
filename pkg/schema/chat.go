package schema

import "time"

type ChatRequest struct {
	Token       string    `json:"token"`
	LocalMsgSeq int       `json:"localMsgSeq"`
	SendTime    time.Time `json:"sendTime"`
	From        int       `json:"from"`
	Group       int       `json:"group"`
	Content     string    `json:"content"`
}

type ChatResponse struct {
	ChatID       string    `json:"chatID"       bson:"_id"`
	Created      time.Time `json:"created"`
	LastModified time.Time `json:"lastModified"`
	From         int       `json:"from"`
	Group        int       `json:"group"`
	Type         int       `json:"type"`
	LocalMsgSeq  int       `json:"localMsgSeq"`
	Content      string    `json:"content"`
}

type OpenAIChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIChatRequest struct {
	Model    string              `json:"model"`
	Messages []OpenAIChatMessage `json:"messages"`
	Stream  bool                `json:"stream"`
}

type OpenAIChatStreamResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Delta struct{
			Content string `json:"content"`
		} `json:"delta"`
		Index int `json:"index"`
		Logprobs      float64 `json:"logprobs"`
		FinishReason  string  `json:"finish_reason"`
	} `json:"choices"`
}


type OpenAIChatResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text          string  `json:"text"`
		Logprobs      float64 `json:"logprobs"`
		FinishReason  string  `json:"finish_reason"`
		Index         int     `json:"index"`
	} `json:"choices"`
}