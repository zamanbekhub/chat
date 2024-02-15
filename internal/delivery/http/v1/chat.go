package v1

import (
	"chat/internal/common/middleware"
	"chat/internal/schema"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initChat(v1 *gin.RouterGroup) {
	group := v1.Group("/chat")
	group.GET("/all", middleware.GinErrorHandle(h.GetUserChats))
	group.GET("", middleware.GinErrorHandle(h.GetChat))
	group.POST("", middleware.GinErrorHandle(h.CreateChat))
}

// GetUserChats
// WhoAmi godoc
// @Summary Получить всех пользователей
// @Accept json
// @Produce json
// @Param user_id query string true "User ID"
// @Success 200 {object} schema.Response[model.Chat]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags chat
// @Router /api/v1/chat/all [get]
func (h *Handler) GetUserChats(c *gin.Context) error {
	chat, err := h.services.Chat.GetAllChats(c.Request.Context(), c.Query("user_id"))
	if err != nil {
		return err
	}
	return schema.Respond(chat, c)
}

// GetChat
// WhoAmi godoc
// @Summary Получить всех пользователей
// @Accept json
// @Produce json
// @Param chat_id query string true "Chat ID"
// @Success 200 {object} schema.Response[model.Chat]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags chat
// @Router /api/v1/chat [get]
func (h *Handler) GetChat(c *gin.Context) error {
	chat, err := h.services.Chat.Get(c.Request.Context(), c.Query("chat_id"))
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
// @Param data body schema.ChatCreate true "Создание чата"
// @Success 200 {object} schema.Response[model.Chat]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags chat
// @Router /api/v1/chat [post]
func (h *Handler) CreateChat(c *gin.Context) error {
	var data schema.CreateChat
	if err := c.BindJSON(&data); err != nil {
		return err
	}

	chat, err := h.services.Chat.Create(c.Request.Context(), data)
	if err != nil {
		return err
	}
	return schema.Respond(chat, c)
}
