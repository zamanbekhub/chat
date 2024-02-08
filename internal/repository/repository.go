package repository

import (
	"chat/internal/model"
	"chat/pkg/db/scylla"
	"github.com/gocql/gocql"
)

type Repository struct {
	Chat             Chat
	UserChatRoleRepo scylla.QueryBuilder[model.UserChatRole]
}

func NewRepositories(
	session *gocql.Session,
) *Repository {
	userChatRoleRepo := scylla.NewQueryBuider[model.UserChatRole](model.NewUserChatRoleTable(), session)

	return &Repository{
		Chat:             NewChatDB(session),
		UserChatRoleRepo: userChatRoleRepo,
	}
}
