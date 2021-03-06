package list

import (
	"context"
	"errors"
	"starwars-api-go/app/planets/list/mocks"
	"starwars-api-go/app/planets/model"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	planetStore := new(mocks.PlanetStore)
	service := NewService(planetStore)

	assert.NotNil(t, service)
	assert.NotNil(t, service.planetStore)
}

func TestGetAll(t *testing.T) {
	testTable := []struct {
		name             string
		mockMethodName   string
		mockMethodParams func() (interface{}, interface{})
		mockReturnValue  interface{}
		expected         interface{}
	}{
		{
			name:           "Empty Result",
			mockMethodName: "FindAll",
			mockMethodParams: func() (interface{}, interface{}) {
				return context.TODO(), mock.Anything
			},
			mockReturnValue: []model.PlanetStorageModel{},
			expected:        []model.Planet{},
		},
		{
			name:           "Not Empty Result",
			mockMethodName: "FindAll",
			mockMethodParams: func() (interface{}, interface{}) {
				return context.TODO(), mock.Anything
			},
			mockReturnValue: []model.PlanetStorageModel{
				{
					Name:         "Alderaan",
					CreationDate: time.Now(),
				},
			},
			expected: []model.Planet{
				{
					Name: "Alderaan",
				},
			},
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			planetStore := new(mocks.PlanetStore)
			param1, param2 := test.mockMethodParams()
			planetStore.On(test.mockMethodName, param1, param2).Return(test.mockReturnValue, nil)

			service := NewService(planetStore)
			result, err := service.List(context.TODO(), model.Filter{Offset: 2, Limit: 10})

			assert.Nil(t, err)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestGetAllFail(t *testing.T) {
	planetStore := new(mocks.PlanetStore)
	service := NewService(planetStore)

	planetStore.On("FindAll", mock.Anything, mock.Anything).Return(nil, errors.New("error"))
	planets, err := service.List(context.TODO(), model.Filter{Offset: 2, Limit: 10})

	assert.Error(t, err)
	assert.Nil(t, planets)
}
