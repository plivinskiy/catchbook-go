package fish

import (
	"catchbook/internal/model"
	"context"
)

func (r *Repository) FindOneById(ctx context.Context, id string) (*model.Fish, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	q := `
	SELECT * FROM fish WHERE id=?
	`
	result, err := r.db.Query(q, id)
	if err != nil {
		return nil, err
	}
	var fish model.Fish
	result.Next()
	err = result.Scan(&fish.Id, &fish.Name, &fish.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &fish, nil
}
