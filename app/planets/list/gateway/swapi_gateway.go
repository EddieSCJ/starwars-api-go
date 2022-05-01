package gateway

import (
	"net/http"
	"starwars-api-go/app/commons"
)

type SWAPIGateway struct {
	client  *http.Client
	baseURL string
}

func NewSWAPIGateway() *SWAPIGateway {
	gateway := new(SWAPIGateway)
	gateway.baseURL = commons.GetSWAPIURL()
	gateway.client = &http.Client{
		Timeout: commons.GetDefaultTimeout(),
	}

	return gateway
}

func (s SWAPIGateway) GetPlanets(filter Filter) (*http.Response, error) {
	url := s.baseURL + "/planets"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	if filter.name != "" {
		query.Add("name", filter.name)
	}
	if filter.page != "" {
		query.Add("page", filter.page)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
