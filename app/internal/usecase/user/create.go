package user

import (
	"catchbook/internal/model"
	"github.com/go-playground/validator/v10"
)

type CreateUserUseCaseInterface interface {
	CreateUser(dto model.UserDto) (*model.User, error)
}

type CreateUserServiceInterface interface {
	Create(dto model.UserDto) (*model.User, error)
}

type CreateUserUseCase struct {
	service CreateUserServiceInterface
}

func NewUseCaseCreateUser(s CreateUserServiceInterface) CreateUserUseCaseInterface {
	return &CreateUserUseCase{
		service: s,
	}
}

var validate *validator.Validate

func (c *CreateUserUseCase) CreateUser(dto model.UserDto) (user *model.User, err error) {
	//validate = validator.New(validator.WithRequiredStructEnabled())
	//err := validate.Struct(dto)
	user, err = c.service.Create(dto)
	if err != nil {
		return
	}
	return
}
