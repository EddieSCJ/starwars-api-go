package main

import (
	"fmt"
	"starwars-api-go/app/planets/storage"
	"starwars-api-go/conf/storage/nosql"
)

func main() {
	nosql.StartDB()
	swapiClient := storage.NewSWAPIClient()
	fmt.Println(swapiClient.GetPlanets(1))
}
