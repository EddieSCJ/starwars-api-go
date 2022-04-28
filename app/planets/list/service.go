package list

import (
	"starwars-api-go/app/planets/list/gateway"
	"starwars-api-go/app/planets/list/storage"
	"starwars-api-go/app/planets/model"
)

type Service interface {
	GetAll(filter Filter) ([]model.Planet, error)
}

type service struct {
	gateway    gateway.SWAPIGateway
	repository storage.Repository
}

func NewService(repository storage.Repository, gateway gateway.SWAPIGateway) Service {
	return &service{
		gateway:    gateway,
		repository: repository,
	}
}

func (s service) GetAll(filter Filter) ([]model.Planet, error) {
	options := storage.NewMongoOptions(filter.offset, filter.limit)
	mongoResultPlanets, err := s.repository.GetAll(options)
	if err != nil {
		return nil, err
	}

	return model.ToDomainList(mongoResultPlanets), nil
}
