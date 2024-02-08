package repository

import (
	"github.com/gocql/gocql"
)

type Repository struct {
	Chat Chat
}

func NewRepositories(
	db *gocql.Session,
) *Repository {
	return &Repository{
		Chat: NewChatDB(db),
	}
}
