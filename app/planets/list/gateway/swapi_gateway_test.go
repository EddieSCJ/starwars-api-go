package gateway

import (
	"context"
	"github.com/stretchr/testify/assert"
	"starwars-api-go/app/commons"
	"testing"
)

func TestNewSWAPIGateway(t *testing.T) {
	g := NewSWAPIGateway()
	assert.Equal(t, commons.GetSWAPIURL(), g.baseURL)
	assert.Equal(t, commons.GetDefaultTimeout(), g.client.Timeout)
}

func TestMountRequest(t *testing.T) {
	g := NewSWAPIGateway()
	req, err := g.mountRequest(context.TODO(), "planets")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, commons.GetSWAPIURL()+"/planets/", req.URL.String())
}

func TestSetNameQueryParam(t *testing.T) {
	g := NewSWAPIGateway()
	req, err := g.mountRequest(context.TODO(), "planets")
	if err != nil {
		t.Error(err)
	}

	g.setQueryParams(req, Filter{name: "Tatooine"})
	assert.Equal(t, "Tatooine", req.URL.Query().Get("search"))
}

func TestSetPageParam(t *testing.T) {
	g := NewSWAPIGateway()
	req, err := g.mountRequest(context.TODO(), "planets")
	if err != nil {
		t.Error(err)
	}

	g.setQueryParams(req, Filter{page: "2"})
	assert.Equal(t, "2", req.URL.Query().Get("page"))
}

func TestSetAllQueryParams(t *testing.T) {
	g := NewSWAPIGateway()
	req, err := g.mountRequest(context.TODO(), "planets")
	if err != nil {
		t.Error(err)
	}

	g.setQueryParams(req, Filter{name: "Tatooine", page: "2"})
	assert.Equal(t, "Tatooine", req.URL.Query().Get("search"))
	assert.Equal(t, "2", req.URL.Query().Get("page"))
}
