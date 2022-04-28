package storage

import (
	"context"
	"starwars-api-go/app/planets/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Count() (int64, error)
	GetAll(findOptions MongoOptions) ([]model.PlanetMongo, error)
}

type repository struct {
	collection *mongo.Collection
}

func NewRepository(client *mongo.Database) Repository {
	return &repository{
		collection: client.Collection("planets"),
	}
}

func (r repository) Count() (int64, error) {
	return r.collection.CountDocuments(context.Background(), bson.M{})
}

func (r repository) GetAll(mongoOptions MongoOptions) ([]model.PlanetMongo, error) {
	findOptions := options.Find().
		SetSkip(mongoOptions.offset).
		SetLimit(mongoOptions.limit)

	cursor, err := r.collection.Find(context.Background(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	var planets []model.PlanetMongo
	if err := cursor.All(nil, &planets); err != nil {
		return nil, err
	}

	return planets, nil
}
