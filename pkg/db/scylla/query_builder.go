package scylla

import (
	"context"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/table"
)

type T any

type QueryBuilder[t T] interface {
	Select(ctx context.Context, dataToGet *t) ([]t, error)
	Get(ctx context.Context, dataToGet *t) (*t, error)
	//SelectAll(ctx context.Context) ([]t, error)
	Insert(ctx context.Context, insertData *t) error
	Delete(ctx context.Context, dataToBeDeleted *t) error
}

type QueryBuilderImpl[T any] struct {
	model   table.Table
	session *gocql.Session
}

/*
	It will insert data into table.

INSERT INTO table VALUES {};
*/
func (queryBuilder *QueryBuilderImpl[T]) Insert(ctx context.Context, insertData *T) error {
	insertStatement, insertNames := queryBuilder.model.Insert()
	err := gocqlx.Query(queryBuilder.session.Query(insertStatement), insertNames).
		WithContext(ctx).
		BindStruct(insertData).
		ExecRelease()

	if err != nil {
		return err
	}

	return nil
}

/*
	It will delete from table based on the Primary Key (Partition Key + Clustering Key (if exists))

DELETE FROM table WHERE PK = {};
*/
func (queryBuilder *QueryBuilderImpl[T]) Delete(ctx context.Context, dataToBeDeleted *T) error {
	deleteStatement, deleteNames := queryBuilder.model.Delete()
	err := gocqlx.Query(queryBuilder.session.Query(deleteStatement), deleteNames).
		WithContext(ctx).
		BindStruct(dataToBeDeleted).
		ExecRelease()

	if err != nil {
		return err
	}

	return nil
}

/*
	It will return data based on the Partition Key

SELECT * FROM table WHERE {partition key = {}};
*/
func (queryBuilder *QueryBuilderImpl[T]) Select(ctx context.Context, dataToGet *T) ([]T, error) {
	selectStatement, selectNames := queryBuilder.model.Select()
	selectQuery := gocqlx.Query(queryBuilder.session.Query(selectStatement), selectNames).WithContext(ctx)

	var results []T
	err := selectQuery.BindStruct(dataToGet).WithContext(ctx).SelectRelease(&results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

/*
	It will return data based on the Primary Key (Partition + Clustering key)

SELECT * FROM table WHERE {primary key = {}};
*/
func (queryBuilder *QueryBuilderImpl[T]) Get(ctx context.Context, dataToGet *T) (*T, error) {
	selectStatement, selectNames := queryBuilder.model.Get()
	selectQuery := gocqlx.Query(queryBuilder.session.Query(selectStatement), selectNames).WithContext(ctx)

	var result []T
	err := selectQuery.BindStruct(dataToGet).WithContext(ctx).SelectRelease(&result)
	if err != nil {
		return nil, err
	}

	if len(result) > 0 {
		return &result[0], nil
	}

	return nil, nil
}

func NewQueryBuilder[T any](model table.Table, session *gocql.Session) *QueryBuilderImpl[T] {
	return &QueryBuilderImpl[T]{
		model,
		session,
	}
}
