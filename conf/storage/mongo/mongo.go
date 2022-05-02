package mongo

import (
	"context"
	"time"

	"starwars-api-go/app/commons"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const timeout = 10 * time.Second

func StartDB() (*mongo.Client, context.CancelFunc) {
	log.Info().Msg("Connecting to MongoDB")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	mongoURI := buildMongoURI()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		defer cancel()
		log.Err(err).Msgf("Error creating options to connect to MongoDB: %s", err.Error())
		return nil, cancel
	}

	pingErr := client.Ping(ctx, nil)
	if pingErr != nil {
		defer cancel()
		log.Err(err).Msgf("Error connecting to MongoDB: %s", pingErr.Error())
		return nil, cancel
	}

	log.Info().Msg("Connected to MongoDB successfully")
	return client, cancel
}

func buildMongoURI() string {
	host := commons.GetMongoHost()
	port := commons.GetMongoPort()
	username := commons.GetMongoUsername()
	password := commons.GetMongoPassword()
	database := commons.GetMongoDBName()

	if username != "" && password != "" {
		return "mongodb://" + username + ":" + password + "@" + host + ":" + port + "/" + database
	}
	return "mongodb://" + host + ":" + port + "/" + database
}
