package main

import (
	"starwars-api-go/conf/storage/mongo"
)

func main() {
	mongo.StartDB()
}
