package gateway

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"starwars-api-go/app/commons"
	"testing"

	"github.com/stretchr/testify/assert"
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
	t.Parallel()
	g := NewSWAPIGateway()
	req, err := g.mountRequest(context.TODO(), "planets")
	if err != nil {
		t.Error(err)
	}

	g.setQueryParams(req, Filter{name: "Tatooine"})
	assert.Equal(t, "Tatooine", req.URL.Query().Get("search"))
}

func TestSetPageParam(t *testing.T) {
	t.Parallel()
	g := NewSWAPIGateway()
	req, err := g.mountRequest(context.TODO(), "planets")
	if err != nil {
		t.Error(err)
	}

	g.setQueryParams(req, Filter{page: "2"})
	assert.Equal(t, "2", req.URL.Query().Get("page"))
}

func TestSetAllQueryParams(t *testing.T) {
	t.Parallel()
	g := NewSWAPIGateway()
	req, err := g.mountRequest(context.TODO(), "planets")
	if err != nil {
		t.Error(err)
	}

	g.setQueryParams(req, Filter{name: "Tatooine", page: "2"})
	assert.Equal(t, "Tatooine", req.URL.Query().Get("search"))
	assert.Equal(t, "2", req.URL.Query().Get("page"))
}

func TestGetPlanets(t *testing.T) {
	expected, _ := ioutil.ReadFile("mock_data/planets_page_1_200.json")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expected)
	}))
	defer server.Close()

	gateway := NewSWAPIGateway()
	response, err := gateway.GetPlanets(context.TODO(), Filter{page: "1"})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, response.Request.Method, http.MethodGet)
	assert.Equal(t, response.Request.URL.String(), commons.GetSWAPIURL()+"/planets/?page=1")
	assert.Equal(t, "1", response.Request.URL.Query().Get("page"))
}

func TestGetPlanetsByName(t *testing.T) {
	expected, _ := ioutil.ReadFile("mock_data/planets_by_name_tatooine_200.json")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expected)
	}))
	defer server.Close()

	gateway := NewSWAPIGateway()
	response, err := gateway.GetPlanets(context.TODO(), Filter{name: "Tatooine"})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, response.Request.Method, http.MethodGet)
	assert.Equal(t, response.Request.URL.String(), commons.GetSWAPIURL()+"/planets/?search=Tatooine")
	assert.Equal(t, "Tatooine", response.Request.URL.Query().Get("search"))
}
