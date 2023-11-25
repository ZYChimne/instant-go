package database

import (
	"zychimne/instant/config"
	"zychimne/instant/pkg/model"
	"zychimne/instant/pkg/schema"
)

func GetRecentConversations(userID uint, offset int, limit int, conversations *[]schema.ConversationResponse) error {
	if err := PostgresDB.Table("conversations").Select("conversations.id", "conversations.conversation_name", "conversations.conversation_type", "conversations.created_at", "conversations.updated_at, messages.content, messages.sender_id").
		Joins("left join messages on conversations.last_message_id = messages.id").
		Where("conversations.id in (select conversation_id from conversation_users where user_id = ?)", userID).
		Order("updated_at desc").Offset(offset).Limit(limit).Scan(&conversations).Error; err != nil {
		return err
	}
	for i, conversation := range *conversations {
		if err := PostgresDB.Table("users").Select("id", "username", "nickname", "avatar", "user_type").Where("id in (select user_id from conversation_users where conversation_id = ? and user_id != ?)", conversation.ID, userID).Scan(&(*conversations)[i].Users).Error; err != nil {
			return err
		}
	}
	return nil
}

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
	if err := tx.Table("conversations").Where("id = ?", message.ConversationID).Update("last_message_id", message.ID).Error; err != nil {
		tx.Rollback()
		return err
	}
	var users []uint
	if err := tx.Table("conversation_users").Select("user_id").Where("conversation_id = ?", message.ConversationID).Scan(&users).Error; err != nil {
		tx.Rollback()
		return err
	}
	inboxMessages := make([]model.InboxMessage, len(users))
	for i, user := range users {
		inboxMessages[i] = model.InboxMessage{
			UserID: user,
		}
		inboxMessages[i].MessageID = message.ID
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
