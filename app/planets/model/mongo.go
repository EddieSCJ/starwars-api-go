package model

import "time"

const (
	dayInHours = 24
)

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
	cacheInDays := difference.Hours() / dayInHours
	return Planet{
		ID:               m.ID,
		Name:             m.Name,
		Climate:          m.Climate,
		Terrain:          m.Terrain,
		MovieAppearances: m.MovieAppearances,
		CacheInDays:      int(cacheInDays),
	}
}

func MongoToDomainList(mongoPlanets []PlanetMongo) []Planet {
	domainList := make([]Planet, 0, len(mongoPlanets))
	for _, mongoPlanet := range mongoPlanets {
		domainList = append(domainList, mongoPlanet.ToDomain())
	}

	return domainList
}
