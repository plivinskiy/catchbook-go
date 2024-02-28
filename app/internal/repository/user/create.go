package user

import (
	"catchbook/internal/model"
	"context"
)

func (r *Repository) Create(ctx context.Context, dto model.UserDto) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	user := model.User{
		Status:    dto.Status,
		Email:     dto.Email,
		Password:  dto.Password,
		Username:  dto.Username,
		Firstname: dto.Firstname,
		Lastname:  dto.Lastname,
		Address: model.Address{
			City:     dto.Address.City,
			Postcode: dto.Address.Postcode,
			Country:  dto.Address.Country,
			Street:   dto.Address.Street,
			Number:   dto.Address.Number,
		},
	}
	result := r.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
