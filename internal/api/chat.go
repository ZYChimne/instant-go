package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
	"zychimne/instant/config"
	database "zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"
	"zychimne/instant/pkg/schema"

	"github.com/gin-gonic/gin"
)

func GetRecentConversations(c *gin.Context) {
	userID := c.MustGet("UserID").(uint)
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetRecentConversationsError))
		return
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetRecentConversationsError))
		return
	}
	var conversations []schema.ConversationResponse
	if err := database.GetRecentConversations(userID, offset, limit, &conversations); err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(GetRecentConversationsError))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": conversations})
}

func CreateConversation(c *gin.Context) {
	userID := c.MustGet("UserID").(uint)
	var conversationSchema schema.UpsertConversationRequest
	if err := c.Bind(&conversationSchema); err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(CreateConversationError))
		return
	}
	users := make([]model.User, len(conversationSchema.Users)+1)
	for i, userID := range conversationSchema.Users {
		users[i].ID = userID
	}
	users[len(users)-1].ID = userID
	conversationModel := model.Conversation{
		ConversationName: conversationSchema.ConversationName,
		ConversationType: conversationSchema.ConversationType,
		Users:            users,
	}
	if err := database.AddConversation(&conversationModel); err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(CreateConversationError))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": conversationModel.ID})
}

func Listen(c *gin.Context) {
	userID := c.MustGet("UserID").(uint)
	mID, err := strconv.ParseUint(c.Query("mID"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(ListenMessagesError))
		return
	}
	curMID := uint(mID)
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(ListenMessagesError))
		return
	}
	c.Stream(func(w io.Writer) bool {
		var messages []schema.MessageResponse
		if err := database.GetMessagesNewerThan(userID, curMID, limit, &messages); err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, errors.New(ListenMessagesError))
			return false
		}
		if len(messages) > 0 {
			payload, err := json.Marshal(messages)
			if err != nil {
				log.Println(err)
				c.AbortWithError(http.StatusInternalServerError, errors.New(ListenMessagesError))
				return false
			}
			c.SSEvent("message", string(payload))
			curMID = messages[len(messages)-1].ID
		}
		time.Sleep(time.Duration(config.Conf.Instant.Chat.RetrieveInterval) * time.Second)
		return true
	})
}

func Send(c *gin.Context) {
	userID := c.MustGet("UserID").(uint)
	var messageSchema schema.UpsertMessageRequest
	if err := c.Bind(&messageSchema); err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(SendMessagesError))
		return
	}
	messageModel := model.Message{
		SenderID:    userID,
		ConversationID: messageSchema.ConversationID,
		MessageType: messageSchema.MessageType,
		Content:     messageSchema.Content,
	}
	if err := database.AddMessage(&messageModel); err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(SendMessagesError))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": messageModel.ID})
}
