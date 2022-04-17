package databases

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"starwars-api-go/commons"
	"time"
)

func StartMongoDB() (*mongo.Client, context.CancelFunc) {
	log.Info().Msg("Connecting to MongoDB")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	mongoUri := buildMongoUri()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))

	if err != nil {
		defer cancel()
		log.Error().Msgf("Error connecting to MongoDB: %s", err.Error())
		return nil, cancel
	}

	log.Info().Msg("Connected to MongoDB")
	return client, cancel
}

func buildMongoUri() string {
	host := commons.GetEnvVar("MONGO_HOST", "localhost")
	port := commons.GetEnvVar("MONGO_PORT", "27017")
	username := commons.GetEnvVar("MONGO_USERNAME", "")
	password := commons.GetEnvVar("MONGO_PASSWORD", "")
	database := commons.GetEnvVar("MONGO_DB", "starwars")

	if username != "" && password != "" {
		return "mongodb://" + username + ":" + password + "@" + host + ":" + port + "/" + database
	}
	return "mongodb://" + host + ":" + port + "/" + database
}
