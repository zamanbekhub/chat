package model

type UserChatRole struct {
	TimestampMixin
	ChatID uint `db:"chat_id" json:"chat_id"`
	UserID uint `db:"user_id" json:"user_id"`
	RoleID uint `db:"role_id" json:"role_id"`
}

func (UserChatRole) TableName() string {
	return "chat.user_chat_role"
}
