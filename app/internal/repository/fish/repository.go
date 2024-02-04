package fish

import (
	"database/sql"
	"sync"
)

type Repository struct {
	mu sync.RWMutex
	db *sql.DB
}

func NewFishRepository(conn *sql.DB) *Repository {
	return &Repository{
		db: conn,
	}
}
