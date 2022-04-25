package model

import "strings"

type swapiPlanet struct {
	Name    string   `json:"name"`
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

func (r SWAPIResponseBody) ToPlanetList() []Planet {
	var planets []Planet
	for _, planet := range r.Results {
		planets = append(planets, planet.ToPlanet())
	}
	return planets
}

func (s *swapiPlanet) ToPlanet() Planet {
	climates := strings.Split(s.Climate, ",")
	terrains := strings.Split(s.Terrain, ",")

	return Planet{
		Name:    s.Name,
		Climate: climates,
		Terrain: terrains,
		Films:   s.Films,
	}
}
