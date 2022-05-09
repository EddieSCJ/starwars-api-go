package planets

import (
	"starwars-api-go/app/planets/list"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func APIRouter(e *echo.Echo, client *mongo.Client) {
	handler := list.BuildListHandler(client)
	e.GET("/planets", handler.List)
}
