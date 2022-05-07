package storage

import (
	"context"
	"starwars-api-go/app/commons"
	"starwars-api-go/app/planets/model"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStore struct {
	collection *mongo.Collection
}

func NewMongoStore(client *mongo.Client) *MongoStore {
	return &MongoStore{
		collection: client.Database(commons.GetMongoDBName()).Collection("planets"),
	}
}

func (r *MongoStore) Count(ctx context.Context) (int64, error) {
	logger := log.Ctx(ctx)
	logger.Info().Msgf("Starting count documents in database.")

	result, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		message := "error while counting documents in database"
		log.Err(err).Msg(message)
		return 0, errors.Wrap(err, message)
	}

	logger.Info().Msgf("%d documents returned successfully", result)
	return result, nil
}

func (r *MongoStore) FindAll(ctx context.Context, mongoOptions MongoOptions) ([]model.PlanetStorageModel, error) {
	logger := log.Ctx(ctx)

	findOptions := mongoOptions.Build()
	cursor, err := r.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		message := "error setting up cursor of planets in database"
		logger.Err(err).Msg(message)
		return nil, errors.Wrap(err, message)
	}

	size := mongoOptions.limit - mongoOptions.offset
	return bindAll(ctx, size, cursor)
}

func bindAll(ctx context.Context, size int64, cursor *mongo.Cursor) ([]model.PlanetStorageModel, error) {
	logger := log.Ctx(ctx)

	planets := make([]model.PlanetStorageModel, 0, size)
	if err := cursor.All(ctx, &planets); err != nil {
		message := "error while binding cursor data to planet mongo type"
		logger.Err(err).Msg(message)
		return nil, errors.Wrap(err, message)
	}

	return planets, nil
}
