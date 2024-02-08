package service

import (
	"chat/internal/model"
	"chat/internal/repository"
	"chat/pkg/db/scylla"
	"context"
)

type Chat interface {
	Get(ctx context.Context, chatID string) (*model.Chat, error)
	GetAllChats(ctx context.Context, userID string) ([]model.UserChatRole, error)
	//Create(ctx context.Context, data schema.ChatCreate) (model.Chat, error)
}

type ChatService struct {
	chatRepo         repository.Chat
	userChatRoleRepo scylla.QueryBuilder[model.UserChatRole]
}

func NewChatService(chatRepo repository.Chat, userChatRoleRepo scylla.QueryBuilder[model.UserChatRole]) *ChatService {
	return &ChatService{
		chatRepo:         chatRepo,
		userChatRoleRepo: userChatRoleRepo,
	}
}

func (s *ChatService) Get(ctx context.Context, chatID string) (*model.Chat, error) {
	chat, err := s.chatRepo.Select(ctx, &model.Chat{
		ChatID: chatID,
	})
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (s *ChatService) GetAllChats(ctx context.Context, userID string) ([]model.UserChatRole, error) {
	userChats, err := s.userChatRoleRepo.Select(ctx, &model.UserChatRole{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	return userChats, nil
}

//func (s *ChatService) Create(ctx context.Context, data schema.ChatCreate) (model.Chat, error) {
//	chat, err := s.chatRepo.Create(ctx, &model.Chat{
//		ChatID:        data.C,
//		Description: data.Description,
//		TypeCode:    data.TypeCode,
//	})
//	if err != nil {
//		return model.Chat{}, err
//	}
//
//	return chat, nil
//}
