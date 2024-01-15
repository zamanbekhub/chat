package service

import (
	"chat/internal/integration"
	"chat/internal/schema"
	"context"
)

type Message interface {
	Push(ctx context.Context, data schema.MessagePush) error
}

type MessageService struct {
	centrifugServerClient integration.CentrifugoServer
}

func NewMessageService(centrifugServerClient integration.CentrifugoServer) *MessageService {
	return &MessageService{
		centrifugServerClient: centrifugServerClient,
	}
}

func (s *MessageService) Push(ctx context.Context, data schema.MessagePush) error {
	err := s.centrifugServerClient.Push(ctx, data.Channel, data.Message)
	if err != nil {
		return err
	}
	return nil
}
