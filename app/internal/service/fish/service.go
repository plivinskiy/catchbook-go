package fish

import (
	"catchbook/internal/model"
	"context"
	"github.com/google/uuid"
)

type ServiceInterface interface {
	GetFish(id string) (*model.Fish, error)
	Create(dto model.FishDto) (*model.Fish, error)
}

type RepositoryInterface interface {
	FindOneById(ctx context.Context, id string) (*model.Fish, error)
	Create(ctx context.Context, id string, dto model.FishDto) (*model.Fish, error)
}

type Service struct {
	fr RepositoryInterface
}

func NewService(r RepositoryInterface) ServiceInterface {
	return &Service{
		fr: r,
	}
}

func (s *Service) GetFish(id string) (*model.Fish, error) {
	return nil, nil
}

func (s *Service) Create(dto model.FishDto) (*model.Fish, error) {
	id, _ := uuid.NewUUID()
	ctx := context.Background()
	f, err := s.fr.Create(ctx, id.String(), dto)
	if err != nil {
		return nil, err
	}
	return f, nil
}
