package repository

import (
	"chat/internal/model"
	"chat/pkg/db/scylla"
	"github.com/gocql/gocql"
)

type Repository struct {
	Chat             Chat
	UserChatRoleRepo scylla.QueryBuilder[model.UserChatRole]
	MessageRepo      scylla.QueryBuilder[model.Message]
}

func NewRepositories(
	session *gocql.Session,
) *Repository {
	userChatRoleRepo := scylla.NewQueryBuilder[model.UserChatRole](model.NewUserChatRoleTable(), session)
	messageRepo := scylla.NewQueryBuilder[model.Message](model.NewMessageTable(), session)

	return &Repository{
		Chat:             NewChatDB(session),
		UserChatRoleRepo: userChatRoleRepo,
		MessageRepo:      messageRepo,
	}
}
