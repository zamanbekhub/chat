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
	UserID      string       `db:"user_id" json:"user_id"`
	RoleID      string       `db:"role_id" json:"role_id"`
	Name        string       `db:"name" json:"name"`
	Description string       `db:"description" json:"description"`
	TypeCode    ChatTypeCode `db:"type_code" json:"type_code"`
}

func NewChatTable() *table.Table {
	m := table.Metadata{
		Name: "tracking_data",
		Columns: []string{
			"first_name", "last_name", "timestamp", "heat",
			"location", "speed", "telepathy_powers",
		},
		PartKey: []string{"first_name", "last_name"},
		SortKey: []string{"timestamp"},
	}
	return table.New(m)
}
