package main

import (
	"github.com/labstack/echo/v4"
	"starwars-api-go/app/commons"
	"starwars-api-go/app/planets"
	"starwars-api-go/conf/storage/mongo"
)

func main() {
	client, _ := mongo.StartDB()
	e := echo.New()

	planets.APIRouter(e, client)

	port := commons.GetDefaultAPIPort()
	e.Logger.Fatal(e.Start(":" + port))
}
