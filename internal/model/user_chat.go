package model

import "time"

type UserChat struct {
	TimestampMixin
	ChatID         uint      `gorm:"chat_id" json:"chat_id"`
	UserID         uint      `gorm:"user_id" json:"user_id"`
	Blocked        bool      `gorm:"blocked" json:"blocked"`
	BlockedEndTime time.Time `gorm:"blocked_end_time" json:"blocked_end_time"`
}

func (UserChat) TableName() string {
	return "chat.user_chat"
}
