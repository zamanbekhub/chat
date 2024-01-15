package schema

type MessagePush struct {
	Channel string `json:"channel"`
	Message string `json:"message"`
}
