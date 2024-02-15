package user

import (
	"catchbook/internal/model"
	"context"
	"fmt"
)

func (r *Repository) FindOneById(ctx context.Context, id string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	q := `
	SELECT * FROM user WHERE id=?
	`
	result, err := r.db.Query(q, id)
	defer result.Close()
	if err != nil {
		return nil, err
	}
	var u model.User
	if !result.Next() {
		return nil, fmt.Errorf("user not found")
	}
	err = result.Scan(&u.Id, &u.Status, &u.Email, &u.Username, &u.Password, &u.Firstname, &u.Lastname, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
