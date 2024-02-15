package service

import (
	"chat/internal/integration"
	"chat/internal/model"
	"chat/internal/schema"
	"chat/pkg/db/scylla"
	"context"
	"fmt"
)

type Message interface {
	GetMessageList(ctx context.Context, chatID string) ([]model.Message, error)
	Push(ctx context.Context, data schema.MessagePush) (model.Message, error)
}

type MessageService struct {
	centrifugServerClient integration.CentrifugoServer
	messageRepo           scylla.QueryBuilder[model.Message]
	userChatRoleRepo      scylla.QueryBuilder[model.UserChatRole]
}

func NewMessageService(
	centrifugServerClient integration.CentrifugoServer,
	messageRepo scylla.QueryBuilder[model.Message],
	userChatRoleRepo scylla.QueryBuilder[model.UserChatRole],
) *MessageService {
	return &MessageService{
		centrifugServerClient: centrifugServerClient,
		messageRepo:           messageRepo,
		userChatRoleRepo:      userChatRoleRepo,
	}
}

func (s *MessageService) GetMessageList(ctx context.Context, chatID string) ([]model.Message, error) {
	messages, err := s.messageRepo.Select(ctx, &model.Message{
		ChatID: chatID,
	})
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (s *MessageService) Push(ctx context.Context, data schema.MessagePush) (model.Message, error) {
	// TODO Validate User and Chat Relation
	userChatRole, err := s.userChatRoleRepo.Get(ctx, &model.UserChatRole{
		UserID: data.UserID,
		ChatID: data.ChatID,
	})
	if err != nil {
		return model.Message{}, err
	}
	if userChatRole == nil {
		return model.Message{}, fmt.Errorf("there is not chat to current user")
	}

	message := model.Message{
		ChatID:            data.ChatID,
		UserID:            data.UserID,
		MessageID:         "",
		Text:              data.Text,
		ContentType:       "Text",
		CentrifugoChannel: data.Channel,
	}
	err = s.messageRepo.Insert(ctx, &message)
	if err != nil {
		return message, err
	}

	err = s.centrifugServerClient.Push(ctx, data.Channel, data.Text)
	if err != nil {
		return message, err
	}

	return message, nil
}
