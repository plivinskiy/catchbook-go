package fish

import (
	"catchbook/internal/model"
	"context"
)

func (r *Repository) FindOneById(ctx context.Context, id string) (*model.Fish, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return nil, nil
}
