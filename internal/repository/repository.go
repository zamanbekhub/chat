package repository

import "gorm.io/gorm"

type Repository struct {
	Chat Chat
}

func NewRepositories(
	db *gorm.DB,
) *Repository {
	return &Repository{
		Chat: NewChatDB(db),
	}
}
