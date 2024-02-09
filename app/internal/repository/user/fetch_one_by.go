package user

import (
	"catchbook/internal/model"
	"context"
)

func (r *Repository) FetchOneByUsernameAndPassword(ctx context.Context, username, password string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	q := `
	SELECT * FROM user WHERE username=? and password=?
	`
	result, err := r.db.Query(q, username, password)
	if err != nil {
		return nil, err
	}
	var u model.User
	result.Next()
	err = result.Scan(&u.Id, &u.Status, &u.Email, &u.Username, &u.Password, &u.Firstname, &u.Lastname, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
