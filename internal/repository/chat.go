package repository

import (
	"chat/internal/model"
	"context"
	"errors"
	"gorm.io/gorm"
)

type Chat interface {
	Get(ctx context.Context, params GetChatParams) (user model.Chat, err error)
	Create(ctx context.Context, user model.Chat) (model.Chat, error)
}

type ChatDB struct {
	db *gorm.DB
}

func NewChatDB(db *gorm.DB) *ChatDB {
	return &ChatDB{
		db: db,
	}
}

func (r *ChatDB) Create(ctx context.Context, chat model.Chat) (model.Chat, error) {
	err := r.db.WithContext(ctx).Create(&chat).Error
	if err != nil {
		return chat, err
	}

	return chat, nil
}

func (r *ChatDB) Get(
	ctx context.Context,
	params GetChatParams,
) (chat model.Chat, err error) {
	query := r.db.Model(&model.Chat{})

	if params.ChatID != nil {
		query = query.Where(`chat_id = ?`, *params.ChatID)
	}

	err = query.First(&chat).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Chat{}, err
		}

		return model.Chat{}, err
	}

	return chat, nil
}

type GetChatParams struct {
	ChatID *uint
}
