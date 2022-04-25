package list

import (
	"starwars-api-go/app/planets/model"
)

type Service interface {
	GetAll() ([]model.Planet, error)
}

type service struct {
	gateway    SWAPIGateway
	repository Repository
}

func NewService() Service {
	return &service{
		gateway:    NewSWAPIGateway(),
		repository: nil,
	}
}

func (s service) GetAll() ([]model.Planet, error) {
	return []model.Planet{}, nil
}
