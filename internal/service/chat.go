package service

import (
	"chat/internal/model"
	"chat/internal/repository"
	"chat/internal/schema"
	"context"
)

type Chat interface {
	GetByID(ctx context.Context, chatID uint) (model.Chat, error)
	Create(ctx context.Context, data schema.ChatCreate) (model.Chat, error)
}

type ChatService struct {
	chatRepo repository.Chat
}

func NewChatService(chatRepo repository.Chat) *ChatService {
	return &ChatService{
		chatRepo: chatRepo,
	}
}

func (s *ChatService) GetByID(ctx context.Context, chatID uint) (model.Chat, error) {
	user, err := s.chatRepo.Get(ctx, repository.GetChatParams{ChatID: &chatID})
	if err != nil {
		return model.Chat{}, err
	}

	return user, nil
}

func (s *ChatService) Create(ctx context.Context, data schema.ChatCreate) (model.Chat, error) {
	chat, err := s.chatRepo.Create(ctx, model.Chat{
		Name:        data.Name,
		Description: data.Description,
		TypeCode:    data.TypeCode,
	})
	if err != nil {
		return model.Chat{}, err
	}

	return chat, nil
}
