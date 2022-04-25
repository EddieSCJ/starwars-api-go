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
