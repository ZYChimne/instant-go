package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"zychimne/instant/pkg/model"

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
	var chatRequest model.ChatRequest
	var chatResponse model.ChatResponse
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
		chatResponse = model.ChatResponse{ChatID: "abc", Created: time.Now(), From: chatRequest.From, Group: chatRequest.Group, Type: 0, LocalMsgSeq: chatRequest.LocalMsgSeq, Content: ""}
		msg, err = json.Marshal(&chatResponse)
		if err != nil {
			log.Println(err.Error())
		}
		// Write message back to browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			log.Println(err.Error())
		}
		chatResponse = model.ChatResponse{ChatID: "abc", Created: time.Now(), From: chatRequest.From, Group: chatRequest.Group, Type: 1, LocalMsgSeq: chatRequest.LocalMsgSeq, Content: chatRequest.Content}
		msg, err = json.Marshal(&chatResponse)
		if err != nil {
			log.Println(err.Error())
		}
		if err = conn.WriteMessage(msgType, msg); err != nil {
			log.Println(err.Error())
		}
	}
}
