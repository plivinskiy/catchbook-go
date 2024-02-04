package user

import (
	"database/sql"
	"sync"
)

type Repository struct {
	mu sync.RWMutex
	db *sql.DB
}

func NewUserRepository(conn *sql.DB) *Repository {
	return &Repository{
		db: conn,
	}
}
