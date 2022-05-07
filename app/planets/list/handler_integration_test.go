//go:build integration

package list

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"net/http/httptest"
	"starwars-api-go/app/commons"
	"starwars-api-go/app/planets/model"
	"starwars-api-go/conf/storage/mongo"
	"testing"
	"time"
)

var data = []model.PlanetStorageModel{
	{
		ID:               time.Now().String(),
		Name:             "Alderaan",
		Climate:          []string{"temperate", "murky"},
		Terrain:          []string{"grasslands", "mountains"},
		MovieAppearances: 21,
		CreationDate:     time.Now(),
	},
	{
		ID:               time.Now().String(),
		Name:             "Maricota",
		Climate:          []string{"temperate", "murky"},
		Terrain:          []string{"grasslands", "mountains"},
		MovieAppearances: 21,
		CreationDate:     time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
}

func TestAPIListEmpty(t *testing.T) {
	client, cancel := mongo.StartDB()
	cleanUp(client, t)
	defer cancel()

	handler := BuildListHandler(client)
	request := httptest.NewRequest(http.MethodGet, "/planets", nil)
	recorder := httptest.NewRecorder()

	ctx := echo.New().NewContext(request, recorder)
	if assert.NoError(t, handler.List(ctx)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "[]\n", recorder.Body.String())
	}

	cleanUp(client, t)
	err := client.Disconnect(context.TODO())
	assert.NoError(t, err)
}

func TestAPIListSuccessfully(t *testing.T) {
	client, _ := mongo.StartDB()
	cleanUp(client, t)
	prepareData(client, t)

	handler := BuildListHandler(client)
	request := httptest.NewRequest(http.MethodGet, "/planets", nil)
	recorder := httptest.NewRecorder()

	ctx := echo.New().NewContext(request, recorder)

	if assert.NoError(t, handler.List(ctx)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		planetsJson := make([]model.PlanetJson, 0, len(data))
		err := json.Unmarshal(recorder.Body.Bytes(), &planetsJson)

		assert.NoError(t, err)
		assert.Equal(t, len(data), len(planetsJson))
		for i, planet := range planetsJson {
			cacheInDays := int(time.Now().Sub(data[i].CreationDate).Hours() / 24)
			assert.Equal(t, data[i].ID, planet.ID)
			assert.Equal(t, data[i].Name, planet.Name)
			assert.Equal(t, data[i].Climate, planet.Climate)
			assert.Equal(t, data[i].Terrain, planet.Terrain)
			assert.Equal(t, data[i].MovieAppearances, planet.MovieAppearances)
			assert.Equal(t, cacheInDays, planet.CacheInDays)
		}
	}

	cleanUp(client, t)
	err := client.Disconnect(context.TODO())
	assert.NoError(t, err)
}

func TestAPIListBadRequest(t *testing.T) {
	client, _ := mongo.StartDB()
	cleanUp(client, t)

	handler := BuildListHandler(client)
	request := httptest.NewRequest(http.MethodGet, "/planets?offset=5f", nil)
	recorder := httptest.NewRecorder()

	ctx := echo.New().NewContext(request, recorder)

	if assert.NoError(t, handler.List(ctx)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		var httpError commons.HttpError
		err := json.Unmarshal(recorder.Body.Bytes(), &httpError)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, httpError.Code)
		assert.Equal(t, "invalid request", httpError.Message)
		assert.Equal(t, "failed while binding filter. review your query params", httpError.AdditionalInfo[0])
	}

	err := client.Disconnect(context.TODO())
	assert.NoError(t, err)
}

func TestAPIListInternalServerError(t *testing.T) {

	handler := BuildListHandler(&mongodriver.Client{})
	request := httptest.NewRequest(http.MethodGet, "/planets?offset=5", nil)
	recorder := httptest.NewRecorder()

	ctx := echo.New().NewContext(request, recorder)

	if assert.NoError(t, handler.List(ctx)) {
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

		var httpError commons.HttpError
		err := json.Unmarshal(recorder.Body.Bytes(), &httpError)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, httpError.Code)
		assert.Equal(t, "an unknown error occurred", httpError.Message)
	}

}

func prepareData(client *mongodriver.Client, t *testing.T) {
	collection := client.Database(commons.GetMongoDBName()).Collection("planets")
	documents := make([]interface{}, 0, len(data))
	for _, d := range data {
		documents = append(documents, d)
	}

	result, err := collection.InsertMany(context.TODO(), documents)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.InsertedIDs))
}

func cleanUp(client *mongodriver.Client, t *testing.T) {
	collection := client.Database(commons.GetMongoDBName()).Collection("planets")
	_, err := collection.DeleteMany(context.TODO(), options.Delete())
	assert.NoError(t, err)
}
