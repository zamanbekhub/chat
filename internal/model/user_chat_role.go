package model

type UserChatRole struct {
	TimestampMixin
	DeleteMixin
	ChatID uint `gorm:"chat_id" json:"chat_id"`
	UserID uint `gorm:"user_id" json:"user_id"`
	RoleID uint `gorm:"role_id" json:"role_id"`
}

func (UserChatRole) TableName() string {
	return "chat.user_chat_role"
}
