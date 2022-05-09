package model

import "time"

const (
	dayInHours = 24
)

type PlanetStorageModel struct {
	ID               string `bson:"_id"`
	Name             string
	Climate          []string
	Terrain          []string
	MovieAppearances int
	CreationDate     time.Time
}

func (m PlanetStorageModel) ToDomain() Planet {
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

func MongoToDomainList(storageModelList []PlanetStorageModel) []Planet {
	domainList := make([]Planet, 0, len(storageModelList))
	for _, storageModel := range storageModelList {
		domainList = append(domainList, storageModel.ToDomain())
	}

	return domainList
}
