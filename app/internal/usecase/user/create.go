package user

import "catchbook/internal/model"

type CreateUserUseCaseInterface interface {
	Create(dto model.UserDto) (*model.User, error)
}

type CreateUserServiceInterface interface {
	Create(dto model.UserDto) (*model.User, error)
}

type CreateUserUseCase struct {
	service CreateUserServiceInterface
}

func NewUseCaseCreateUser(service CreateUserServiceInterface) CreateUserServiceInterface {
	return &CreateUserUseCase{
		service: service,
	}
}

func (c *CreateUserUseCase) Create(dto model.UserDto) (*model.User, error) {
	// todo add validation of user
	user, err := c.service.Create(dto)
	if err != nil {
		return nil, err
	}
	return user, nil
}