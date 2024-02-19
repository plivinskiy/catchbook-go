package user

import (
	"gorm.io/gorm"
	"sync"
)

type Repository struct {
	mu sync.RWMutex
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *Repository {
	return &Repository{
		db: conn,
	}
}
