package list

import (
	"starwars-api-go/app/planets/list/storage"
	"starwars-api-go/app/planets/model"
)

type PlanetStore interface {
	GetAll(options storage.MongoOptions) (interface{}, error)
}

type Service struct {
	planetStore PlanetStore
}

func NewService(planetStore PlanetStore) *Service {
	return &Service{
		planetStore: planetStore,
	}
}

func (s *Service) GetAll(filter Filter) (interface{}, error) {
	options := storage.NewMongoOptions(filter.offset, filter.limit)
	result, err := s.planetStore.GetAll(options)
	if err != nil {
		return nil, err
	}

	return model.ToDomainList(result.([]model.PlanetMongo)), nil
}
