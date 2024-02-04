package db

import (
	"context"
	"database/sql"
	"fmt"
)

type MysqlClient interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) error
	Query(ctx context.Context, sql string, args ...interface{}) (sql.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) sql.Row
	Begin(ctx context.Context) error
}

func NewMysqlClient(ctx context.Context, dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("cannot connect to database: %s", err.Error()))
	}
	return db
}
