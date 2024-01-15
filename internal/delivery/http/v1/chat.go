package v1

import (
	"chat/internal/common/middleware"
	"chat/internal/schema"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) initChat(v1 *gin.RouterGroup) {
	group := v1.Group("/chat")
	group.GET("", middleware.GinErrorHandle(h.GetChat))
	group.POST("", middleware.GinErrorHandle(h.CreateChat))
}

// GetChat
// WhoAmi godoc
// @Summary Получить всех пользователей
// @Accept json
// @Produce json
// @Param chat_id query int true "Chat ID"
// @Success 200 {object} schema.Response[model.Chat]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags chat
// @Router /api/v1/chat [get]
func (h *Handler) GetChat(c *gin.Context) error {
	rawID := c.Query("chat_id")
	chatID, err := strconv.ParseUint(rawID, 10, 64)
	if err != nil {
		return err
	}
	chat, err := h.services.Chat.GetByID(c.Request.Context(), uint(chatID))
	if err != nil {
		return err
	}
	return schema.Respond(chat, c)
}

// CreateChat
// WhoAmi godoc
// @Summary Создание пользователя
// @Accept json
// @Produce json
// @Param data body schema.ChatCreate true "Создание пользователыя"
// @Success 200 {object} schema.Response[model.Chat]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags chat
// @Router /api/v1/chat [post]
func (h *Handler) CreateChat(c *gin.Context) error {
	var data schema.ChatCreate
	if err := c.BindJSON(&data); err != nil {
		return err
	}

	chat, err := h.services.Chat.Create(c.Request.Context(), data)
	if err != nil {
		return err
	}
	return schema.Respond(chat, c)
}
