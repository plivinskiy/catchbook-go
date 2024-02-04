package user

import (
	"catchbook/internal/model"
	"context"
	"github.com/google/uuid"
)

type ServiceInterface interface {
	GetUser(id string) (*model.User, error)
	GetOneByUsernameAndEmail(username, password string) (*model.User, error)
	Create(dto model.UserDto) (*model.User, error)
	List() ([]*model.User, error)
}

type RepositoryInterface interface {
	FindOneById(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, id string, dto model.UserDto) (*model.User, error)
	FetchAll(ctx context.Context) ([]*model.User, error)
	FetchOneByUsernameAndPassword(ctx context.Context, username, password string) (*model.User, error)
}

type Service struct {
	ur RepositoryInterface
}

func NewService(r RepositoryInterface) ServiceInterface {
	return &Service{
		ur: r,
	}
}

func (s *Service) List() ([]*model.User, error) {
	ctx := context.Background()
	user, err := s.ur.FetchAll(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) GetUser(id string) (*model.User, error) {
	ctx := context.Background()
	user, err := s.ur.FindOneById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) Create(dto model.UserDto) (*model.User, error) {
	id, _ := uuid.NewUUID()
	ctx := context.Background()
	user, err := s.ur.Create(ctx, id.String(), dto)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) GetOneByUsernameAndEmail(username, password string) (*model.User, error) {
	ctx := context.Background()
	user, err := s.ur.FetchOneByUsernameAndPassword(ctx, username, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
