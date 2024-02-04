package user

import (
	"catchbook/internal/model"
	"context"
)

func (r *Repository) Create(ctx context.Context, id string, dto model.UserDto) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	q := `
		INSERT INTO user (id,status,email,username,firstname,lastname,created_at) VALUES (?,?,?,?,?,?,NOW())
	`
	insert, err := r.db.Query(q, id, dto.Status, dto.Email, dto.Username, dto.Firstname, dto.Lastname)
	if err != nil {
		return nil, err
	}
	_ = insert.Close()
	return &model.User{
		Id:       id,
		Email:    dto.Email,
		Username: dto.Username,
	}, nil
}
