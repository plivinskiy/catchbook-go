package user

import "catchbook/internal/model"

type ListUserUseCaseInterface interface {
	ListUsers() ([]*model.User, error)
}

type ListUserServiceInterface interface {
	List() ([]*model.User, error)
}

type ListUserUseCase struct {
	service ListUserServiceInterface
}

func NewUseCaseListUser(service ListUserServiceInterface) ListUserUseCaseInterface {
	return &ListUserUseCase{
		service: service,
	}
}

func (c ListUserUseCase) ListUsers() ([]*model.User, error) {
	list, err := c.service.List()
	if err != nil {
		return nil, err
	}
	return list, nil
}
