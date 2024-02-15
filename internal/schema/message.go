package schema

type MessagePush struct {
	ChatID  string `json:"chatID"`
	UserID  string `json:"userID"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}
