package mongo

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

var (
	host     = "localhost"
	port     = "27017"
	database = "development"
)

func TestBuildMongoUriSuccess(t *testing.T) {
	uri := buildMongoURI()
	assert.Equal(t, fmt.Sprintf("mongodb://%s:%s/%s", host, port, database), uri)
}

func TestBuildMongoUriWithCredentialsSuccess(t *testing.T) {
	changeCredentials("user", "password")
	uri := buildMongoURI()
	assert.Equal(t, fmt.Sprintf("mongodb://user:password@%s:%s/%s", host, port, database), uri)
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

func BenchmarkBuildMongoUri(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildMongoURI()
	}
}
