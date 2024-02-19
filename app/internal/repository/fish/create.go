package fish

import (
	"catchbook/internal/model"
	"context"
)

func (r *Repository) Create(ctx context.Context, id string, dto model.FishDto) (*model.Fish, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return nil, nil
}
