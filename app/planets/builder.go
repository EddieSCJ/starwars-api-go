package planets

import (
	"starwars-api-go/app/planets/list"
	listStorage "starwars-api-go/app/planets/list/storage"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func APIRouter(e *echo.Echo, client *mongo.Client) {
	buildListEndpoints(e, client)
}

func buildListEndpoints(e *echo.Echo, client *mongo.Client) {
	mongoStore := listStorage.NewMongoStore(client)
	listHandler := list.NewHandler(list.NewService(mongoStore))
	e.GET("/planets", listHandler.List)
}
