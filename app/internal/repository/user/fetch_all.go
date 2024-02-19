package user

import (
	"catchbook/internal/model"
	"context"
	"fmt"
)

func (r *Repository) FetchAll(ctx context.Context) ([]*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	stmt, err := r.db.Prepare("SELECT * FROM user")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}
	result, err := stmt.Query()
	defer result.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}
	var users []*model.User
	for result.Next() {
		var u model.User
		err = result.Scan(&u.Id, &u.Status, &u.Email, &u.Username, &u.Password, &u.Firstname, &u.Lastname, &u.CreatedAt)
		users = append(users, &u)
	}
	return users, nil
}
