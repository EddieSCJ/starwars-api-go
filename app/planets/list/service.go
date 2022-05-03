package list

import (
	"context"
	"starwars-api-go/app/planets/list/storage"
	"starwars-api-go/app/planets/model"
)

type PlanetStore interface {
	FindAll(ctx context.Context, options storage.MongoOptions) (interface{}, error)
}

type Service struct {
	planetStore PlanetStore
}

func NewService(planetStore PlanetStore) *Service {
	return &Service{
		planetStore: planetStore,
	}
}

func (s *Service) List(ctx context.Context, filter model.Filter) ([]model.Planet, error) {
	options := storage.NewMongoOptions(filter.Offset, filter.Limit)
	result, err := s.planetStore.FindAll(ctx, options)
	if err != nil {
		return nil, err
	}

	return model.MongoToDomainList(result.([]model.PlanetMongo)), nil
}
