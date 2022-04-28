package model

import "time"

type PlanetMongo struct {
	ID               string `bson:"_id"`
	Name             string
	Climate          []string
	Terrain          []string
	MovieAppearances int
	CreationDate     time.Time
}

func (m PlanetMongo) ToDomain() Planet {
	difference := time.Now().Sub(m.CreationDate)
	cacheInDays := difference.Hours() / 24
	return Planet{
		ID:               m.ID,
		Name:             m.Name,
		Climate:          m.Climate,
		Terrain:          m.Terrain,
		MovieAppearances: m.MovieAppearances,
		CacheInDays:      int(cacheInDays),
	}
}

func ToDomainList(mongoPlanets []PlanetMongo) []Planet {
	var domainList []Planet
	for _, mongoPlanet := range mongoPlanets {
		domainList = append(domainList, mongoPlanet.ToDomain())
	}

	return domainList
}
