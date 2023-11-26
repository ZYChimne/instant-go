package database

import (
	"zychimne/instant/config"
	"zychimne/instant/pkg/model"
	"zychimne/instant/pkg/schema"
)

func GetRecentConversations(userID uint, offset int, limit int, conversations *[]model.RecentConversation) error {
	query := `
        SELECT 
            conversations.id,
            conversations.conversation_name,
            conversations.conversation_type,
            conversations.created_at,
            conversations.updated_at,
            messages.content,
            messages.sender_id,
            COALESCE(json_agg(users), '[]'::json) as users
        FROM conversations
        LEFT JOIN messages ON conversations.last_message_id = messages.id
        LEFT JOIN (
            SELECT
                u.id,
                u.username,
                u.nickname,
                u.avatar,
                u.user_type,
                cu.conversation_id
            FROM users u
            JOIN conversation_users cu ON u.id = cu.user_id
        ) users ON users.conversation_id = conversations.id AND users.id != ?
        WHERE conversations.id IN (SELECT conversation_id FROM conversation_users WHERE user_id = ?)
        GROUP BY 
            conversations.id,
            conversations.conversation_name,
            conversations.conversation_type,
            conversations.created_at,
            conversations.updated_at,
            messages.content,
            messages.sender_id
        ORDER BY conversations.updated_at DESC
        OFFSET ? LIMIT ?
    `
	return PostgresDB.Raw(query, userID, userID, offset, limit).Scan(conversations).Error
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
		inboxMessages[i].Message=*message
		inboxMessages[i].ID = 0
	}
	if err := tx.CreateInBatches(&inboxMessages, config.Conf.Database.App.CreateInstantBatchSize).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
