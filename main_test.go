//go:build integration

package main

import (
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/rs/zerolog/log"

	"starwars-api-go/conf/storage/mongo"
)

var pool *dockertest.Pool

func StartDocker() {
	var err error
	pool, err = dockertest.NewPool("")
	if err != nil {
		log.Error().Msgf("Could not connect to docker: %s", err.Error())
		os.Exit(1)
	}
}

func TestMain(m *testing.M) {
	StartDocker()

	mongo.RemoveDBContainer(pool)
	_, dbErr := mongo.StartDBContainer(pool)
	if dbErr != nil {
		log.Err(dbErr).Msg("Could not start MongoDB")
		os.Exit(1)
	}

	code := m.Run()

	mongo.DisconnectDB()
	os.Exit(code)
}
