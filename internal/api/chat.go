package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
	"zychimne/instant/config"
	"zychimne/instant/pkg/schema"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// allow cros
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Echo(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	var chatRequest schema.ChatRequest
	var chatResponse schema.ChatResponse
	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
		}
		err = json.Unmarshal(msg, &chatRequest)
		if err != nil {
			log.Println(err.Error())
		}
		chatResponse = schema.ChatResponse{
			ChatID:      "abc",
			Created:     time.Now(),
			From:        chatRequest.From,
			Group:       chatRequest.Group,
			Type:        0,
			LocalMsgSeq: chatRequest.LocalMsgSeq,
			Content:     "",
		}
		msg, err = json.Marshal(&chatResponse)
		if err != nil {
			log.Println(err.Error())
		}
		// Write message back to browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			log.Println(err.Error())
		}
		chatResponse = schema.ChatResponse{
			ChatID:      "abc",
			Created:     time.Now(),
			From:        chatRequest.From,
			Group:       chatRequest.Group,
			Type:        1,
			LocalMsgSeq: chatRequest.LocalMsgSeq,
			Content:     chatRequest.Content,
		}
		msg, err = json.Marshal(&chatResponse)
		if err != nil {
			log.Println(err.Error())
		}
		if err = conn.WriteMessage(msgType, msg); err != nil {
			log.Println(err.Error())
		}
	}
}

func Chat(c *gin.Context) {
	// _ = c.MustGet("UserID")
	errMsg := "Chat error"
	var chatRequest schema.ChatRequest
	if err := c.Bind(&chatRequest); err != nil {
		handleError(c, err, errMsg, ParameterError)
		return
	}
	resp := _OpenAIChat(chatRequest.Content)
	c.Stream(func(w io.Writer) bool {
		defer resp.Body.Close()
		bytes := make([]byte, 1024)
		for {
			clear(bytes)
			n, err := resp.Body.Read(bytes)
			if err != nil {
				log.Println(err.Error())
				return false
			}
			if n == 0 {
				break
			}
			messages := strings.Split(string(bytes[:n]), "\r")
			for _, message := range messages {
				message = strings.Trim(message, "\n")
				if len(message) > 0 {
					c.SSEvent("", message)
				}
			}
		}
		return true
	})
}

func _OpenAIChat(query string) *http.Response {
	chatRequest := schema.OpenAIChatRequest{
		Model: "Llama-2",
		Messages: []schema.OpenAIChatMessage{
			{
				Role:    "user",
				Content: query,
			},
		},
		Stream: true,
	}
	body, err := json.Marshal(chatRequest)
	if err != nil {
		log.Println(err.Error())
	}
	response, err := http.Post(strings.Join([]string{config.Conf.OpenAI.URL, "v1/chat/completions"}, "/"), "application/json", bytes.NewReader(body))
	if err != nil {
		log.Println(err.Error())
	}
	return response
}
