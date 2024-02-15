package user

import (
	"catchbook/internal/model"
	"context"
	"fmt"
)

func (r *Repository) FetchOneByUsernameAndPassword(ctx context.Context, username, password string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	stmt, err := r.db.Prepare("SELECT * FROM user WHERE username=? AND password=?")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}

	result, err := stmt.Query(username, password)
	defer result.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}
	var u model.User
	if !result.Next() {
		return nil, fmt.Errorf("user by username %s not found", username)
	}
	err = result.Scan(&u.Id, &u.Status, &u.Email, &u.Username, &u.Password, &u.Firstname, &u.Lastname, &u.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to scan: %w", err)
	}
	return &u, nil
}
