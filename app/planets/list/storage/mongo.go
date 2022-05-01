package storage

import (
	"context"
	"starwars-api-go/app/planets/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	collection *mongo.Collection
}

func NewMongoRepository(client *mongo.Database) *MongoStore {
	return &MongoStore{
		collection: client.Collection("planets"),
	}
}

func (r *MongoStore) Count() (int64, error) {
	return r.collection.CountDocuments(context.Background(), bson.M{})
}

func (r *MongoStore) GetAll(mongoOptions MongoOptions) (interface{}, error) {
	findOptions := options.Find().
		SetSkip(mongoOptions.offset).
		SetLimit(mongoOptions.limit)

	cursor, err := r.collection.Find(context.Background(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	size := mongoOptions.limit - mongoOptions.offset
	planets := make([]model.PlanetMongo, size)
	if err := cursor.All(nil, &planets); err != nil {
		return nil, err
	}

	return planets, nil
}
