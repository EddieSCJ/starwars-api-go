package gateway

import (
	"context"
	"net/http"
	"starwars-api-go/app/commons"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
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

func (s SWAPIGateway) GetPlanets(ctx context.Context, filter Filter) (*http.Response, error) {
	logger := log.Ctx(ctx)
	logger.Info().Msgf("Starting getting planets from Star Wars API client. name: %s. page: %s", filter.name, filter.page)

	req, err := s.mountRequest(ctx)
	if err != nil {
		return nil, err
	}

	s.setQueryParams(req, filter)
	return s.performRequest(ctx, req, filter)
}

func (s *SWAPIGateway) mountRequest(ctx context.Context) (*http.Request, error) {
	logger := log.Ctx(ctx)
	url := s.baseURL + "/planets/"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		message := "error while building request to get planets in swapi client"
		logger.Err(err).Msg(message)
		return nil, errors.Wrap(err, message)
	}

	return req, nil
}

func (s *SWAPIGateway) setQueryParams(req *http.Request, filter Filter) {
	query := req.URL.Query()
	if filter.name != "" {
		query.Add("search", filter.name)
	}
	if filter.page != "" {
		query.Add("page", filter.page)
	}
	req.URL.RawQuery = query.Encode()
}

func (s *SWAPIGateway) performRequest(ctx context.Context, req *http.Request, filter Filter) (*http.Response, error) {
	logger := log.Ctx(ctx)
	resp, err := s.client.Do(req)
	if err != nil {
		message := "error while performing request to swapi client"
		logger.Err(err).Msg(message)
		return nil, errors.Wrap(err, message)
	}

	logger.Info().Msgf("Planets retrieved with success. name: %s. page: %s", filter.name, filter.page)
	return resp, nil
}
