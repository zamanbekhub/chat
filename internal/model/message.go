package model

type Message struct {
	TimestampMixin
	ChatID      uint   `db:"chat_id" json:"chat_id"`
	TimePart    string `db:"time_part" json:"time_part"`
	Time        string `db:"time" json:"time"`
	MessageID   string `db:"message" json:"message_id"`
	Text        string `db:"text" json:"text"`
	ContentType string `db:"content_type" json:"content_type"`
}
