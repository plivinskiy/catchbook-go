package user

import (
	"catchbook/internal/model"
	"context"
	"fmt"
)

func (r *Repository) Create(ctx context.Context, id string, dto model.UserDto) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	stmt, err := r.db.Prepare("INSERT INTO user (id,status,email,username,password,firstname,lastname,created_at) VALUES (?,?,?,?,?,?,?,NOW())")
	if err != nil {
		return nil, err
	}
	insert, err := stmt.Query(id, dto.Status, dto.Email, dto.Username, dto.Password, dto.Firstname, dto.Lastname)
	defer insert.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}
	return &model.User{
		Id:       id,
		Email:    dto.Email,
		Username: dto.Username,
	}, nil
}
