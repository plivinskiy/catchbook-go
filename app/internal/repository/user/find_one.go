package user

import (
	"catchbook/internal/model"
	"context"
)

func (r *Repository) FindOneById(ctx context.Context, id uint) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var user model.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
