package schema

type UpsertMessageRequest struct {
	ConversationID uint   `json:"conversationID"`
	MessageType    int    `json:"messageType"`
	Content        string `json:"content"`
}

type MessageResponse struct {
	ID             uint   `json:"id"`
	ConversationID uint   `json:"conversationID"`
	SenderID       uint   `json:"senderID"`
	MessageType    int    `json:"messageType"`
	Content        string `json:"content"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
}

type UpsertConversationRequest struct {
	ConversationName string `json:"conversationName"`
	ConversationType int    `json:"conversationType"`
	Users            []uint `json:"users"`
}
