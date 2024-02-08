package service

import (
	"chat/internal/integration"
	"chat/internal/repository"
)

type Services struct {
	Chat    Chat
	Message Message
}

func NewServices(repos *repository.Repository, integrations *integration.Integrations) *Services {
	return &Services{
		Chat:    NewChatService(repos.Chat, repos.UserChatRoleRepo),
		Message: NewMessageService(integrations.CentrifugoServer),
	}
}
