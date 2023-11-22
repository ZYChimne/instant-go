package database

import (
	"zychimne/instant/config"
	"zychimne/instant/pkg/model"
	"zychimne/instant/pkg/schema"
)

func AddConversation(conversation *model.Conversation) error {
	return PostgresDB.Create(&conversation).Error
}

func GetMessagesNewerThan(userID uint, mID uint, limit int, messages *[]schema.MessageResponse) error {
	return PostgresDB.Table("inbox_messages").Where("user_id = ?", userID).Where("id > ?", mID).Order("created_at asc").Limit(limit).Scan(&messages).Error
}

func AddMessage(message *model.Message) error {
	tx := PostgresDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Create(&message).Error; err != nil {
		tx.Rollback()
		return err
	}
	var users []uint
	if err := tx.Table("conversations").Select("user_id").Where("conversation_id = ?", message.ConversationID).Scan(&users).Error; err != nil {
		tx.Rollback()
		return err
	}
	users = append(users, message.SenderID)
	inboxMessages := make([]model.InboxMessage, len(users))
	for i, user := range users {
		inboxMessages[i] = model.InboxMessage{
			UserID: user,
		}
		inboxMessages[i].ID = message.ID
		inboxMessages[i].ConversationID = message.ConversationID
		inboxMessages[i].SenderID = message.SenderID
		inboxMessages[i].MessageType = message.MessageType
		inboxMessages[i].Content = message.Content
		inboxMessages[i].CreatedAt = message.CreatedAt
		inboxMessages[i].UpdatedAt = message.UpdatedAt
	}
	if err := tx.CreateInBatches(&inboxMessages, config.Conf.Database.App.CreateInstantBatchSize).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
