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
	UserID uint
	IsRead bool
}
