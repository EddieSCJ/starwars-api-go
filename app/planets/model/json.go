package model

type PlanetJSON struct {
	ID               string   `json:"id"`
	Name             string   `json:"Name"`
	Climate          []string `json:"climate"`
	Terrain          []string `json:"terrain"`
	MovieAppearances int      `json:"movie_appearances"`
	CacheInDays      int      `json:"cache_in_days"`
}

func FromDomain(planet Planet) PlanetJSON {
	return PlanetJSON(planet)
}

func FromDomainList(planets []Planet) []PlanetJSON {
	planetJSONList := make([]PlanetJSON, 0, len(planets))
	for _, planet := range planets {
		planetJSONList = append(planetJSONList, FromDomain(planet))
	}
	return planetJSONList
}
