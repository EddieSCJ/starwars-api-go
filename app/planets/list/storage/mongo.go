package storage

import (
	"context"
	"github.com/pkg/errors"
	"starwars-api-go/app/planets/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStore struct {
	collection *mongo.Collection
}

func NewMongoRepository(client *mongo.Database) *MongoStore {
	return &MongoStore{
		collection: client.Collection("planets"),
	}
}

func (r *MongoStore) Count(ctx context.Context) (int64, error) {
	return r.collection.CountDocuments(ctx, bson.M{})
}

func (r *MongoStore) GetAll(ctx context.Context, mongoOptions MongoOptions) ([]model.PlanetMongo, error) {
	findOptions := mongoOptions.Build()

	cursor, err := r.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, errors.Wrap(err, "error setting up cursor of planets in database")
	}

	size := mongoOptions.limit - mongoOptions.offset
	return bindAll(ctx, size, cursor)
}

func bindAll(ctx context.Context, size int64, cursor *mongo.Cursor) ([]model.PlanetMongo, error) {
	planets := make([]model.PlanetMongo, size)
	if err := cursor.All(ctx, &planets); err != nil {
		return nil, errors.Wrap(err, "error while binding cursor data to planet mongo type")
	}

	return planets, nil
}
