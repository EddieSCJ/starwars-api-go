package model

import "strings"

type swapiPlanet struct {
	Name    string   `json:"Name"`
	Climate string   `json:"climate"`
	Terrain string   `json:"terrain"`
	Films   []string `json:"films"`
}

type SWAPIResponseBody struct {
	Count    int           `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []swapiPlanet `json:"results"`
}

func (r SWAPIResponseBody) ToDomainList() []Planet {
	planets := make([]Planet, 0, len(r.Results))
	for _, planet := range r.Results {
		planets = append(planets, planet.ToDomain())
	}
	return planets
}

func (s *swapiPlanet) ToDomain() Planet {
	climates := strings.Split(s.Climate, ",")
	terrains := strings.Split(s.Terrain, ",")

	return Planet{
		Name:             s.Name,
		Climate:          climates,
		Terrain:          terrains,
		MovieAppearances: len(s.Films),
	}
}
