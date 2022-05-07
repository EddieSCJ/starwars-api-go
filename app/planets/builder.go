package planets

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"starwars-api-go/app/planets/list"
)

func APIRouter(e *echo.Echo, client *mongo.Client) {
	handler := list.BuildListHandler(client)
	e.GET("/planets", handler.List)
}
