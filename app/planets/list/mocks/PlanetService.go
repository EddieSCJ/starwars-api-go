// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "starwars-api-go/app/planets/model"

	testing "testing"
)

// PlanetService is an autogenerated mock type for the PlanetService type
type PlanetService struct {
	mock.Mock
}

// List provides a mock function with given fields: ctx, filter
func (_m *PlanetService) List(ctx context.Context, filter model.Filter) ([]model.Planet, error) {
	ret := _m.Called(ctx, filter)

	var r0 []model.Planet
	if rf, ok := ret.Get(0).(func(context.Context, model.Filter) []model.Planet); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Planet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPlanetService creates a new instance of PlanetService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewPlanetService(t testing.TB) *PlanetService {
	mock := &PlanetService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
