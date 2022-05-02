// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	storage "starwars-api-go/app/planets/list/storage"

	testing "testing"
)

// PlanetStore is an autogenerated mock type for the PlanetStore type
type PlanetStore struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: ctx, options
func (_m *PlanetStore) GetAll(ctx context.Context, options storage.MongoOptions) (interface{}, error) {
	ret := _m.Called(ctx, options)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, storage.MongoOptions) interface{}); ok {
		r0 = rf(ctx, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, storage.MongoOptions) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPlanetStore creates a new instance of PlanetStore. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewPlanetStore(t testing.TB) *PlanetStore {
	mock := &PlanetStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
