package v1

import (
	"chat/internal/common/middleware"
	"chat/internal/schema"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initMessage(v1 *gin.RouterGroup) {
	group := v1.Group("/message")
	group.GET("/list", middleware.GinErrorHandle(h.PushMessage))
	group.POST("/push", middleware.GinErrorHandle(h.PushMessage))
}

// GetMessageList
// WhoAmi godoc
// @Summary Получить сообщение чата
// @Accept json
// @Produce json
// @Param chat_id query string true "Chat ID"
// @Success 200 {object} schema.Response[[]model.Message]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags message
// @Router /api/v1/message/list [get]
func (h *Handler) GetMessageList(c *gin.Context) error {
	messages, err := h.services.Message.GetMessageList(c.Request.Context(), c.Query("chat_id"))
	if err != nil {
		return err
	}
	return schema.Respond(messages, c)
}

// PushMessage
// WhoAmi godoc
// @Summary Новое сообщение
// @Accept json
// @Produce json
// @Param data body schema.MessagePush true "Новое сообщение"
// @Success 200 {object} schema.Response[model.Message]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags message
// @Router /api/v1/message/push [post]
func (h *Handler) PushMessage(c *gin.Context) error {
	var data schema.MessagePush
	if err := c.BindJSON(&data); err != nil {
		return err
	}

	message, err := h.services.Message.Push(c.Request.Context(), data)
	if err != nil {
		return err
	}
	return schema.Respond(message, c)
}
