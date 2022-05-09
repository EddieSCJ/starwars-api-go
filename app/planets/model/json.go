package model

type PlanetJson struct {
	ID               string   `json:"id"`
	Name             string   `json:"Name"`
	Climate          []string `json:"climate"`
	Terrain          []string `json:"terrain"`
	MovieAppearances int      `json:"movie_appearances"`
	CacheInDays      int      `json:"cache_in_days"`
}

func FromDomain(planet Planet) PlanetJson {
	return PlanetJson{
		ID:               planet.ID,
		Name:             planet.Name,
		Climate:          planet.Climate,
		Terrain:          planet.Terrain,
		MovieAppearances: planet.MovieAppearances,
		CacheInDays:      planet.CacheInDays,
	}
}

func FromDomainList(planets []Planet) []PlanetJson {
	planetJsonList := make([]PlanetJson, 0, len(planets))
	for _, planet := range planets {
		planetJsonList = append(planetJsonList, FromDomain(planet))
	}
	return planetJsonList
}
