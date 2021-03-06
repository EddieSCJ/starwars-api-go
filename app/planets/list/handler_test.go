package list

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"starwars-api-go/app/commons"
	"starwars-api-go/app/planets/list/mocks"
	"starwars-api-go/app/planets/model"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	defaultEndpoint = "/planets?offset=5&limit=10"
)

func TestListEmpty(t *testing.T) {
	mockService := mocks.PlanetService{}
	mockService.On("List", mock.Anything, mock.Anything).Return([]model.Planet{}, nil)

	handler := NewHandler(&mockService)
	request := httptest.NewRequest(http.MethodGet, defaultEndpoint, nil)
	recorder := httptest.NewRecorder()

	ctx := echo.New().NewContext(request, recorder)

	if assert.NoError(t, handler.List(ctx)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "[]\n", recorder.Body.String())
	}
}

func TestListSuccessfully(t *testing.T) {
	mockService := mocks.PlanetService{}
	mockService.On("List", mock.Anything, mock.Anything).Return([]model.Planet{
		{
			ID:               time.Now().String(),
			Name:             "Alderaan",
			Climate:          []string{"temperate", "murky"},
			Terrain:          []string{"grasslands", "mountains"},
			MovieAppearances: 21,
			CacheInDays:      1,
		},
		{
			ID:               time.Now().String(),
			Name:             "Maricota",
			Climate:          []string{"temperate", "murky"},
			Terrain:          []string{"grasslands", "mountains"},
			MovieAppearances: 21,
			CacheInDays:      1,
		},
	}, nil)

	handler := NewHandler(&mockService)
	request := httptest.NewRequest(http.MethodGet, defaultEndpoint, nil)
	recorder := httptest.NewRecorder()

	ctx := echo.New().NewContext(request, recorder)

	if assert.NoError(t, handler.List(ctx)) {
		var planets []model.PlanetJSON
		assert.Equal(t, http.StatusOK, recorder.Code)

		err := json.Unmarshal(recorder.Body.Bytes(), &planets)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(planets))
		assert.Equal(t, "Alderaan", planets[0].Name)
		assert.Equal(t, "Maricota", planets[1].Name)
	}
}

func TestListFailBadRequest(t *testing.T) {
	mockService := mocks.PlanetService{}
	mockService.On("List", mock.Anything, mock.Anything).Return([]model.Planet{}, nil)

	handler := NewHandler(&mockService)
	request := httptest.NewRequest(http.MethodGet, defaultEndpoint+"f", nil)
	recorder := httptest.NewRecorder()

	ctx := echo.New().NewContext(request, recorder)

	if assert.NoError(t, handler.List(ctx)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		var httpError commons.HTTPError
		err := json.Unmarshal(recorder.Body.Bytes(), &httpError)

		assert.NoError(t, err)
		assert.Equal(t, "invalid request", httpError.Message)
	}
}

func TestListFailInternalServerError(t *testing.T) {
	mockService := mocks.PlanetService{}
	mockService.On("List", mock.Anything, mock.Anything).Return(nil, errors.New("error"))

	handler := NewHandler(&mockService)
	request := httptest.NewRequest(http.MethodGet, defaultEndpoint, nil)
	recorder := httptest.NewRecorder()

	ctx := echo.New().NewContext(request, recorder)

	if assert.NoError(t, handler.List(ctx)) {
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

		var httpError commons.HTTPError
		err := json.Unmarshal(recorder.Body.Bytes(), &httpError)
		assert.NoError(t, err)
		assert.Equal(t, "an unknown error occurred", httpError.Message)
	}
}
