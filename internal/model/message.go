package model

import (
	"github.com/scylladb/gocqlx/table"
)

type Message struct {
	TimestampMixin
	ChatID            string `db:"chat_id" json:"chat_id"`
	TimePart          string `db:"time_part" json:"time_part"`
	UserID            string `db:"user_id" json:"user_id"`
	MessageID         string `db:"message" json:"message_id"`
	Text              string `db:"text" json:"text"`
	ContentType       string `db:"content_type" json:"content_type"`
	CentrifugoChannel string `db:"centrifugo_channel" json:"centrifugo_channel"`
}

func NewMessageTable() table.Table {
	m := table.Metadata{
		Name: "message",
		Columns: []string{
			"chat_id", "time_part", "user_id", "message_id",
			"text", "content_type", "created_at", "updated_at",
		},
		PartKey: []string{"chat_id", "time_part"},
		SortKey: []string{"created_at"},
	}
	return *table.New(m)
}
