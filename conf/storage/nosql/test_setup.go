package nosql

import (
	"context"
	"fmt"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client

func StartDBContainer(pool *dockertest.Pool) (*dockertest.Resource, error) {

	resource, err := pullMongoImage(pool)
	if err != nil {
		log.Error().Msgf("Could not start Mongo Image: %s", err)
		return nil, err
	}

	err = makeReadyToAcceptConnections(pool, resource)
	if err != nil {
		log.Error().Msgf("Could not connect to docker: %s", err)
		return nil, err
	}
	return resource, nil
}

func pullMongoImage(pool *dockertest.Pool) (*dockertest.Resource, error) {
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Name:       "mongoservice",
		Tag:        "5.0",
		PortBindings: map[docker.Port][]docker.PortBinding{
			"27017/tcp": {{HostIP: "", HostPort: "27017"}},
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})

	return resource, err
}

func makeReadyToAcceptConnections(pool *dockertest.Pool, resource *dockertest.Resource) error {
	err := pool.Retry(func() error {
		var err error
		dbClient, err = mongo.Connect(
			context.TODO(),
			options.Client().ApplyURI(
				fmt.Sprintf("mongodb://localhost:%s", resource.GetPort("27017/tcp")),
			),
		)
		if err != nil {
			return err
		}
		return dbClient.Ping(context.TODO(), nil)
	})

	return err
}

func RemoveDBContainer(pool *dockertest.Pool) {
	if err := pool.RemoveContainerByName("mongoservice"); err != nil {
		log.Error().Msgf("Could not purge resource: %s", err)
	}
}

func DisconnectDB() {
	if err := dbClient.Disconnect(context.TODO()); err != nil {
		log.Error().Msgf("Could not disconnect from DB: %s", err)
	}
}
