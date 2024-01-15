package schema

import "chat/internal/model"

type ChatCreate struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	TypeCode    model.ChatTypeCode `json:"type_code" binding:"required"`
}
