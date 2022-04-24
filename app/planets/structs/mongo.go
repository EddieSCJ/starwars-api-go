package structs

import "time"

type PlanetMongo struct {
	ID               string
	Name             string
	Climate          []string
	Terrain          []string
	MovieAppearances int
	CreationDate     time.Time
}
