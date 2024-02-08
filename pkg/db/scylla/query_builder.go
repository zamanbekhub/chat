package scylla

//import (
//	"context"
//	"fmt"
//	"github.com/gocql/gocql"
//	"github.com/scylladb/gocqlx/table"
//)
//
//type T any
//
//type QueryBuilder[t T] interface {
//	Insert(ctx context.Context, insertData *t) error
//	Delete(ctx context.Context, dataToBeDeleted *t) error
//	DeleteAllFromPartitioningKey(ctx context.Context, dataToBeDeleted *t) error
//	Select(ctx context.Context, dataToGet *t) ([]t, error)
//	Get(ctx context.Context, dataToGet *t) (*t, error)
//	SelectAll(ctx context.Context) ([]t, error)
//}
//
//type queryBuilder[T any] struct {
//	model   table.Table
//	session gocql.Session
//}
//
///*
//	It will insert data into table.
//
//INSERT INTO table VALUES {};
//*/
//func (queryBuilder *queryBuilder[T]) Insert(ctx context.Context, insertData *T) error {
//	insertStatement, insertNames := queryBuilder.model.Insert()
//	insertQuery := gocql.Query()queryBuilder.session.Query(insertStatement).WithContext(ctx)
//
//	err := insertQuery.BindStruct(insertData).ExecRelease()
//	if err != nil {
//		queryBuilder.logger.Error(fmt.Sprintf("Insert() error %s", err.Error()))
//		return err
//	}
//
//	return nil
//}
//
///*
//	It will delete from table based on the Primary Key (Partition Key + Clustering Key (if exists))
//
//DELETE FROM table WHERE PK = {};
//*/
//func (queryBuilder *queryBuilder[T]) Delete(ctx context.Context, dataToBeDeleted *T) error {
//	deleteStatement, deleteNames := queryBuilder.model.Delete()
//	deleteQuery := queryBuilder.session.Query(deleteStatement, deleteNames).WithContext(ctx)
//
//	err := deleteQuery.BindStruct(dataToBeDeleted).WithContext(ctx).ExecRelease()
//	if err != nil {
//		queryBuilder.logger.Error(fmt.Sprintf("Delete by Primary Key error: %s", err.Error()))
//		return err
//	}
//
//	return nil
//}
//
//func (queryBuilder *queryBuilder[T]) DeleteAllFromPartitioningKey(ctx context.Context, dataToBeDeleted *T) error {
//	query := fmt.Sprintf("DELETE FROM %s WHERE ", queryBuilder.model.Metadata().M.Name)
//
//	for index, value := range queryBuilder.model.Metadata().M.PartKey {
//		if index == 0 {
//			query += fmt.Sprintf("%s=? ", value)
//			continue
//		}
//
//		query += fmt.Sprintf("AND %s=?", value)
//	}
//
//	deleteQuery := queryBuilder.session.Query(query, queryBuilder.model.Metadata().M.PartKey).WithContext(ctx)
//
//	err := deleteQuery.BindStruct(dataToBeDeleted).WithContext(ctx).ExecRelease()
//	if err != nil {
//		queryBuilder.logger.Error(fmt.Sprintf("Delete by Partition Key error: %s", err.Error()))
//		return err
//	}
//
//	return nil
//}
//
///*
//	It will return data based on the Partition Key
//
//SELECT * FROM table WHERE {partition key = {}};
//*/
//func (queryBuilder *queryBuilder[T]) Select(ctx context.Context, dataToGet *T) ([]T, error) {
//	selectStatement, selectNames := queryBuilder.model.Select()
//	selectQuery := queryBuilder.session.Query(selectStatement, selectNames).WithContext(ctx)
//
//	var results []T
//	err := selectQuery.BindStruct(dataToGet).WithContext(ctx).SelectRelease(&results)
//	if err != nil {
//		queryBuilder.logger.Error(fmt.Sprintf("Select error: %s", err.Error()))
//		return nil, err
//	}
//
//	return results, nil
//}
//
///*
//	It will return data based on the Primary Key (Partition + Clustering key)
//
//SELECT * FROM table WHERE {primary key = {}};
//*/
//func (queryBuilder *queryBuilder[T]) Get(ctx context.Context, dataToGet *T) (*T, error) {
//	selectStatement, selectNames := queryBuilder.model.Get()
//	selectQuery := queryBuilder.session.Query(selectStatement, selectNames).WithContext(ctx)
//
//	var result []T
//	err := selectQuery.BindStruct(dataToGet).WithContext(ctx).SelectRelease(&result)
//	if err != nil {
//		queryBuilder.logger.Error(fmt.Sprintf("Get error: %s", err.Error()))
//		return nil, err
//	}
//
//	if len(result) > 0 {
//		return &result[0], nil
//	}
//
//	return nil, nil
//}
//
///*
//	It will everything from table.
//
//SELECT * FROM table;
//*/
//func (queryBuilder *queryBuilder[T]) SelectAll(ctx context.Context) ([]T, error) {
//	selectAllStatement, selectAllNames := queryBuilder.model.SelectAll()
//	selectAllQuery := queryBuilder.session.Query(selectAllStatement, selectAllNames).WithContext(ctx)
//
//	var results []T
//	err := selectAllQuery.WithContext(ctx).SelectRelease(&results)
//	if err != nil {
//		queryBuilder.logger.Error(fmt.Sprintf("SelectAll error: %s", err.Error()))
//		return nil, err
//	}
//
//	return results, nil
//}
//
//func NewQueryBuider[T any](model igocqlxtable.ITable, session igocqlx.ISessionx, logger interfaces.ILogger) *queryBuilder[T] {
//	return &queryBuilder[T]{
//		model,
//		session,
//		logger,
//	}
//}
