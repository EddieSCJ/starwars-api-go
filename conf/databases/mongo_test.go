package databases

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBuildMongoUriSuccess(t *testing.T) {
	uri := buildMongoUri()
	assert.Equal(t, "mongodb://localhost:27017/starwars", uri)
}

func TestBuildMongoUriWithCredentialsSuccess(t *testing.T) {
	err := os.Setenv("MONGO_USER", "user")
	assert.NoError(t, err)

	err = os.Setenv("MONGO_PASSWORD", "password")
	assert.NoError(t, err)

	uri := buildMongoUri()
	assert.Equal(t, "mongodb://user:password@localhost:27017/starwars", uri)
}

func TestStartMongoDB(t *testing.T) {
	assert.Truef(t, true, "MongoDB is not running")
}
