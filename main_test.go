//go:build integration

package main

import (
	"github.com/ory/dockertest"
	"github.com/rs/zerolog/log"
	"os"
	"starwars-api-go/conf/storage/nosql"
	"testing"
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

	nosql.RemoveDBContainer(pool)
	_, err := nosql.StartDBContainer(pool)
	if err != nil {
		log.Error().Msgf("Could not start MongoDB: %s", err)
		os.Exit(1)
	}

	code := m.Run()

	nosql.DisconnectDB()
	os.Exit(code)
}
