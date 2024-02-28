package user

import (
	"catchbook/internal/model"
	"context"
)

func (r *Repository) FindOneById(ctx context.Context, id uint) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var user model.User
	result := r.db.Find(&user, &model.User{ID: id})
	if result.Error != nil {
		return nil, result.Error
	}
	err := r.db.Model(&user).Association("Address").Find(&user.Address)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
