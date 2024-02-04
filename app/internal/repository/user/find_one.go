package user

import (
	"catchbook/internal/model"
	"context"
)

func (r *Repository) FindOneById(ctx context.Context, id string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	q := `
	SELECT * FROM user WHERE id=?
	`
	result, err := r.db.Query(q, id)
	if err != nil {
		return nil, err
	}
	var u model.User
	result.Next()
	err = result.Scan(&u.Id, &u.Status, &u.Email, &u.Username, &u.Firstname, &u.Lastname, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
