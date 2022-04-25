package model

type PlanetJson struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Climate          []string `json:"climate"`
	Terrain          []string `json:"terrain"`
	MovieAppearances int      `json:"movie_appearances"`
	CacheInDays      int      `json:"cache_in_days"`
}
