package auth

import (
	"catchbook/internal/config"
	"catchbook/internal/model"
	"catchbook/pkg/jwt"
	"fmt"
)

type AuthorizeUseCaseInterface interface {
	Authorize(username, password string) (*model.User, error)
	Token(user *model.User) (interface{}, error)
}

type AuthorizeUseCase struct {
	us  UserServiceInterface
	js  jwt.ServiceInterface
	cfg *config.Config
}

type UserServiceInterface interface {
	GetOneByUsernameAndEmail(username, password string) (*model.User, error)
}

func NewAuthorizeUseCase(us UserServiceInterface, js jwt.ServiceInterface) AuthorizeUseCaseInterface {
	return &AuthorizeUseCase{
		us: us,
		js: js,
	}
}

func (c AuthorizeUseCase) Authorize(username, password string) (*model.User, error) {
	user, err := c.us.GetOneByUsernameAndEmail(username, password)
	if err != nil {
		return nil, fmt.Errorf("user not found by username %s: %v", username, err)
	}
	return user, nil
}

func (c AuthorizeUseCase) Token(user *model.User) (interface{}, error) {
	token, err := c.js.GenerateAccessToken(user, c.cfg.GetSecret())
	if err != nil {
		return nil, err
	}
	return token, nil
}
