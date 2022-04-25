package list

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"starwars-api-go/app/planets/model"
)

type Repository interface {
	GetAll() ([]model.PlanetMongo, error)
}

type repository struct {
	collection *mongo.Collection
}

func NewRepository(client *mongo.Database) Repository {
	return &repository{
		collection: client.Collection("planets"),
	}
}

func (r repository) GetAll() ([]model.PlanetMongo, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var planets []model.PlanetMongo
	if err := cursor.All(nil, &planets); err != nil {
		return nil, err
	}

	return planets, nil
}
