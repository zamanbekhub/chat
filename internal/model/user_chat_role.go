package model

import "github.com/scylladb/gocqlx/table"

type UserChatRole struct {
	TimestampMixin
	ChatID string `db:"chat_id" json:"chat_id"`
	UserID string `db:"user_id" json:"user_id"`
	RoleID string `db:"role_id" json:"role_id"`
}

func NewUserChatRoleTable() table.Table {
	m := table.Metadata{
		Name: "user_chat_role",
		Columns: []string{
			"user_id", "chat_id", "role_id",
			"created_at", "updated_at",
		},
		PartKey: []string{"user_id", "chat_id"},
	}
	return *table.New(m)
}
