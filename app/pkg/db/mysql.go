package db

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlClient(ctx context.Context, dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("cannot connect to database: %s", err.Error()))
	}
	return db
}
