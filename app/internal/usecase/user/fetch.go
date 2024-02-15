package user

import "catchbook/internal/model"

type FetchUserUseCaseInterface interface {
	FetchUser(id string) (*model.User, error)
}

type ServiceInterface interface {
	GetUser(id string) (*model.User, error)
}

type FetchUserUseCase struct {
	service ServiceInterface
}

func NewUseCaseFetchUser(service ServiceInterface) FetchUserUseCaseInterface {
	return &FetchUserUseCase{
		service: service,
	}
}

func (r FetchUserUseCase) FetchUser(id string) (*model.User, error) {
	u, err := r.service.GetUser(id)
	if err != nil {
		return nil, err
	}
	return u, nil
}
