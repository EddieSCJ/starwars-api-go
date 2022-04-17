package main

import "starwars-api-go/conf/databases"

func main() {
	databases.StartMongoDB()
}
