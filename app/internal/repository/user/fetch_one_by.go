package user

import (
	"catchbook/internal/model"
	"context"
)

func (r *Repository) FetchOneByUsernameAndPassword(ctx context.Context, username, password string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var user model.User
	result := r.db.First(&user, "username=? AND password=?", username, password)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, result.Error
}
