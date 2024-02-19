package user

import (
	"catchbook/internal/model"
	"context"
)

func (r *Repository) FetchAll(ctx context.Context) ([]*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return nil, nil
}
