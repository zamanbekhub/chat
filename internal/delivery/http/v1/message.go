package v1

import (
	"chat/internal/common/middleware"
	"chat/internal/schema"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initMessage(v1 *gin.RouterGroup) {
	group := v1.Group("/message")
	group.POST("/push", middleware.GinErrorHandle(h.PushMessage))
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
