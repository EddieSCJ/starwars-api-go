package main

import (
	"starwars-api-go/app/commons"
	"starwars-api-go/app/commons/middleware"
	"starwars-api-go/app/planets"
	"starwars-api-go/conf/storage/mongo"

	"github.com/labstack/echo/v4"
)

func main() {
	client, _ := mongo.StartDB()
	e := echo.New()

	e.Use(
		middleware.SetLoggerInContext(commons.NewZeroLogger()),
		middleware.AddMetadata(),
	)

	planets.APIRouter(e, client)

	port := commons.GetDefaultAPIPort()
	e.Logger.Fatal(e.Start(":" + port))
}
