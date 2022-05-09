//go:build integration

package storage

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/options"
	"starwars-api-go/conf/storage/mongo"
	"strconv"
	"testing"
)

var (
	client, _ = mongo.StartDB()
)

func insertData(mongoStore *MongoStore) {
	_, err := mongoStore.collection.InsertOne(context.TODO(), map[string]interface{}{
		"name": "Alderaan",
	})
	if err != nil {
		panic(err)
	}
}

func cleanData(store *MongoStore) {
	_, err := store.collection.DeleteMany(context.TODO(), options.Delete())
	if err != nil {
		panic(err)
	}
}

func TestNewMongoStore(t *testing.T) {
	mongoStore := NewMongoStore(client)
	assert.NotNil(t, mongoStore)
	assert.Equal(t, mongoStore.collection.Name(), "planets")
}

func TestCount(t *testing.T) {
	mongoStore := NewMongoStore(client)
	cleanData(mongoStore)
	insertData(mongoStore)

	count, err := mongoStore.Count(context.TODO())
	assert.Nil(t, err)
	assert.Equal(t, int64(1), count)
	cleanData(mongoStore)
}

func TestFindAll(t *testing.T) {
	testTable := []struct {
		name           string
		length         int
		insertDataFunc func(mongoStore *MongoStore)
		cleanDataFunc  func(mongoStore *MongoStore)
	}{
		{"Get All Empty", 0, func(mongoStore *MongoStore) { return }, func(mongoStore *MongoStore) { return }},
		{"Get All", 10, func(mongoStore *MongoStore) {
			for i := 0; i < 10; i++ {
				_, err := mongoStore.collection.InsertOne(context.TODO(), map[string]interface{}{
					"name": "Alderaan" + strconv.Itoa(i),
				})
				if err != nil {
					panic(err)
				}
			}
		}, cleanData},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			mongoStore := NewMongoStore(client)
			test.cleanDataFunc(mongoStore)
			test.insertDataFunc(mongoStore)

			planets, err := mongoStore.FindAll(context.TODO(), MongoOptions{offset: 0, limit: 10})
			assert.Nil(t, err)
			assert.Equal(t, test.length, len(planets))
			test.cleanDataFunc(mongoStore)
		})
	}
}
