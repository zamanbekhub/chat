package model

import "github.com/scylladb/gocqlx/table"

type ChatTypeCode = string

const (
	CHAT_TYPE_CODE_PERSONAL ChatTypeCode = "personal"
	CHAT_TYPE_CODE_GROUP    ChatTypeCode = "group"
	CHAT_TYPE_CODE_CHANNEL  ChatTypeCode = "channel"
)

type Chat struct {
	TimestampMixin
	ChatID      string       `db:"chat_id" json:"chat_id"`
	Name        string       `db:"name" json:"name"`
	Description string       `db:"description" json:"description"`
	TypeCode    ChatTypeCode `db:"type_code" json:"type_code"`
}

func NewChatTable() table.Table {
	m := table.Metadata{
		Name: "chat",
		Columns: []string{
			"chat_id", "name",
			"description", "type_code",
			"created_at", "updated_at",
		},
		PartKey: []string{"user_id"},
	}
	return *table.New(m)
}
