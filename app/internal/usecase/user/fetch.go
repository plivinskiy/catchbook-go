package user

import "catchbook/internal/model"

type FetchUserUseCaseInterface interface {
	FetchUser(id uint) (*model.User, error)
}

type ServiceInterface interface {
	GetUser(id uint) (*model.User, error)
}

type FetchUserUseCase struct {
	service ServiceInterface
}

func NewUseCaseFetchUser(service ServiceInterface) FetchUserUseCaseInterface {
	return &FetchUserUseCase{
		service: service,
	}
}

func (r FetchUserUseCase) FetchUser(id uint) (*model.User, error) {
	user, err := r.service.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
