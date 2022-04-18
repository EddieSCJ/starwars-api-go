package nosql

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"starwars-api-go/app/commons"
	"time"
)

func StartDB() (*mongo.Client, context.CancelFunc) {
	log.Info().Msg("Connecting to MongoDB")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	mongoUri := buildMongoUri()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))

	if err != nil {
		defer cancel()
		log.Error().Msgf("Error creating options to connect to MongoDB: %s", err.Error())
		return nil, cancel
	}

	pingErr := client.Ping(ctx, nil)
	if pingErr != nil {
		defer cancel()
		log.Error().Msgf("Error connecting to MongoDB: %s", pingErr.Error())
		return nil, cancel
	}

	log.Error().Msg(mongoUri)
	log.Info().Msg("Connected to MongoDB")
	return client, cancel
}

func buildMongoUri() string {
	host := commons.GetEnv("MONGO_HOST", "mongoservice")
	port := commons.GetEnv("MONGO_PORT", "27017")
	username := commons.GetEnv("MONGO_USER", "")
	password := commons.GetEnv("MONGO_PASSWORD", "")
	database := commons.GetEnv("MONGO_DB", "starwars")

	if username != "" && password != "" {
		return "mongodb://" + username + ":" + password + "@" + host + ":" + port + "/" + database
	}
	return "mongodb://" + host + ":" + port + "/" + database
}
