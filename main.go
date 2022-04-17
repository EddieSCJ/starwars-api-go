package main

import "starwars-api-go/conf/storage"

func main() {
	storage.StartMongoDB()
}
