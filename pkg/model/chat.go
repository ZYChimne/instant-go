package model

import "gorm.io/gorm"

type Conversation struct {
	gorm.Model
	ConversationName string `gorm:"not null"`
	ConversationType int    `gorm:"not null"`
	LastMessageID    uint
	Users            []User `gorm:"many2many:conversation_users;"`
}

type Message struct {
	gorm.Model
	ConversationID uint
	SenderID       uint
	MessageType    int
	Content        string
}

type InboxMessage struct {
	Message
	MessageID uint
	UserID    uint
	IsRead    bool
}

type RecentConversation struct {
	ID               uint   `json:"conversationID"`
	ConversationName string `json:"conversationName"`
	ConversationType int    `json:"conversationType"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
	Content          string `json:"content"`
	SenderID         uint   `json:"senderID"`
	Users            string `json:"-"`
}
