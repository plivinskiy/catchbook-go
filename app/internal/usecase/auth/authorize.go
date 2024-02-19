package auth

import (
	"catchbook/internal/config"
	"catchbook/internal/model"
	"catchbook/pkg/jwt"
)

type AuthorizeUseCaseInterface interface {
	Authorize(username, password string) (*model.User, error)
	Token(u *model.User) (interface{}, error)
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
	u, err := c.us.GetOneByUsernameAndEmail(username, password)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (c AuthorizeUseCase) Token(u *model.User) (interface{}, error) {
	t, err := c.js.GenerateAccessToken(u, c.cfg.GetSecret())
	if err != nil {
		return nil, err
	}
	return t, nil
}
