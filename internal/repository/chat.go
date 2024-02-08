package repository

import (
	"chat/internal/model"
	"context"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/table"
)

type Chat interface {
	Get(ctx context.Context, data *model.Chat) (*model.Chat, error)
	Select(ctx context.Context, data *model.Chat) (*model.Chat, error)
	Create(ctx context.Context, data *model.Chat) error
}

type ChatDB struct {
	session *gocql.Session
	model   table.Table
}

func NewChatDB(session *gocql.Session) *ChatDB {
	return &ChatDB{
		session: session,
		model:   model.NewChatTable(),
	}
}

func (r *ChatDB) Create(ctx context.Context, data *model.Chat) error {
	insertStatement, insertNames := r.model.Insert()
	err := gocqlx.Query(r.session.Query(insertStatement), insertNames).
		WithContext(ctx).
		BindStruct(data).
		ExecRelease()

	if err != nil {
		return err
	}

	return nil
}

func (r *ChatDB) Get(ctx context.Context, data *model.Chat) (*model.Chat, error) {
	var result []model.Chat

	selectStatement, selectNames := r.model.Get()
	err := gocqlx.Query(r.session.Query(selectStatement), selectNames).
		WithContext(ctx).
		BindStruct(data).
		SelectRelease(&result)

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return &result[0], nil
	}

	return nil, nil
}

func (r *ChatDB) Select(ctx context.Context, data *model.Chat) (*model.Chat, error) {
	var result []model.Chat

	selectStatement, selectNames := r.model.Select()
	err := gocqlx.Query(r.session.Query(selectStatement), selectNames).
		WithContext(ctx).
		BindStruct(data).
		SelectRelease(&result)

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return &result[0], nil
	}

	return nil, nil
}
