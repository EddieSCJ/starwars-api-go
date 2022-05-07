package list

import (
	listStorage "starwars-api-go/app/planets/list/storage"

	"go.mongodb.org/mongo-driver/mongo"
)

func BuildListHandler(client *mongo.Client) *Handler {
	mongoStore := listStorage.NewMongoStore(client)
	return NewHandler(NewService(mongoStore))
}
