package fish

import (
	"catchbook/internal/model"
	"context"
)

func (r *Repository) Create(ctx context.Context, id string, dto model.FishDto) (*model.Fish, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	q := `
		INSERT INTO fish (id,name,created_at) VALUES (?,?,NOW())
	`
	insert, err := r.db.Query(q, id, dto.Name)
	if err != nil {
		return nil, err
	}
	_ = insert.Close()
	return &model.Fish{
		Id:   id,
		Name: dto.Name,
	}, nil
}
