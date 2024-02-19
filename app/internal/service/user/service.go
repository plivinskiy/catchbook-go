package user

import (
	"catchbook/internal/model"
	"context"
)

type ServiceInterface interface {
	GetUser(id uint) (*model.User, error)
	GetOneByUsernameAndEmail(username, password string) (*model.User, error)
	Create(dto model.UserDto) (*model.User, error)
	List() ([]*model.User, error)
}

type RepositoryInterface interface {
	FindOneById(ctx context.Context, id uint) (*model.User, error)
	Create(ctx context.Context, dto model.UserDto) (*model.User, error)
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

func (s *Service) GetUser(id uint) (*model.User, error) {
	ctx := context.Background()
	user, err := s.ur.FindOneById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) Create(dto model.UserDto) (*model.User, error) {
	ctx := context.Background()
	user, err := s.ur.Create(ctx, dto)
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
