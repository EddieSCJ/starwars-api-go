package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"starwars-api-go/app/commons"
	"starwars-api-go/app/planets/structs"
)

type SWAPIClient interface {
	GetPlanets(page int) ([]structs.Planet, error)
}

type swapiClient struct {
	client  *http.Client
	baseUri string
}

func NewSWAPIClient() SWAPIClient {
	client := new(swapiClient)
	client.baseUri = commons.GetSWAPIURL()
	client.client = &http.Client{
		Timeout: commons.GetDefaultTimeout(),
	}

	return client
}

func (s *swapiClient) GetPlanets(page int) ([]structs.Planet, error) {
	planetsUri := fmt.Sprintf("%s/planets/?page=%d", s.baseUri, page)
	response, err := s.client.Get(planetsUri)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	var swapiPlanetsResponse structs.SWAPIPlanetResponse
	unmarshalError := json.Unmarshal(body, &swapiPlanetsResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return swapiPlanetsResponse.ToPlanetList(), nil

}
