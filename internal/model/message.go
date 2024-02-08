package model

import "github.com/scylladb/gocqlx/table"

type Message struct {
	TimestampMixin
	ChatID      uint   `db:"chat_id" json:"chat_id"`
	TimePart    string `db:"time_part" json:"time_part"`
	MessageID   string `db:"message" json:"message_id"`
	Text        string `db:"text" json:"text"`
	ContentType string `db:"content_type" json:"content_type"`
}

func NewMessageTable() table.Table {
	m := table.Metadata{
		Name: "message",
		Columns: []string{
			"chat_id", "time_part", "message_id",
			"text", "content_type",
		},
		PartKey: []string{"chat_id", "time_part"},
		SortKey: []string{"created_at"},
	}
	return *table.New(m)
}
