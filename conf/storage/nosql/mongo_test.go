package nosql

import (
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBuildMongoUriSuccess(t *testing.T) {
	uri := buildMongoUri()
	assert.Equal(t, "mongodb://localhost:27017/starwars", uri)
}

func TestBuildMongoUriWithCredentialsSuccess(t *testing.T) {
	changeCredentials("user", "password")
	uri := buildMongoUri()
	assert.Equal(t, "mongodb://user:password@localhost:27017/starwars", uri)
	changeCredentials("", "")
}

func changeCredentials(user, password string) {
	err := os.Setenv("MONGO_USER", user)
	if err != nil {
		log.Error().Err(err).Msg("Error setting env variable MONGO_USER")
	}

	err = os.Setenv("MONGO_PASSWORD", password)
	if err != nil {
		log.Error().Err(err).Msg("Error setting env variable MONGO_PASSWORD")
	}
}
