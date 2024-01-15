package model

type ChatTypeCode = string

const (
	CHAT_TYPE_CODE_PERSONAL ChatTypeCode = "personal"
	CHAT_TYPE_CODE_GROUP    ChatTypeCode = "group"
	CHAT_TYPE_CODE_CHANNEL  ChatTypeCode = "channel"
)

type Chat struct {
	ChatID      uint         `gorm:"primaryKey;chat_id" json:"chat_id"`
	Name        string       `gorm:"name" json:"name"`
	Description string       `gorm:"description" json:"description"`
	TypeCode    ChatTypeCode `gorm:"type_code" json:"type_code"`
}

func (Chat) TableName() string {
	return "chat.chat"
}
