package list

import (
	"context"
	"starwars-api-go/app/planets/list/storage"
	"starwars-api-go/app/planets/model"
)

type PlanetStore interface {
	GetAll(ctx context.Context, options storage.MongoOptions) (interface{}, error)
}

type Service struct {
	planetStore PlanetStore
}

func NewService(planetStore PlanetStore) *Service {
	return &Service{
		planetStore: planetStore,
	}
}

func (s *Service) GetAll(ctx context.Context, filter Filter) (interface{}, error) {
	options := storage.NewMongoOptions(filter.offset, filter.limit)
	result, err := s.planetStore.GetAll(ctx, options)
	if err != nil {
		return nil, err
	}

	return model.ToDomainList(result.([]model.PlanetMongo)), nil
}
