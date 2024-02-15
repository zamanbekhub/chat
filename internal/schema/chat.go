package schema

import "chat/internal/model"

type CreateChat struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	TypeCode    model.ChatTypeCode `json:"type_code" binding:"required"`
}

type GetUserChat struct {
	Name               string             `json:"name"`
	Avatar             string             `json:"avatar"`
	ChatType           model.ChatTypeCode `json:"chat_type"`
	LastUserName       string             `json:"last_user_name"`
	LastMessageText    string             `json:"last_message_text"`
	UnreadMessageCount uint               `json:"unread_message_count"`
}
