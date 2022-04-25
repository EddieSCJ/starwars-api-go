package main

import (
	"encoding/json"
	"fmt"
	"starwars-api-go/app/planets/list"
	"starwars-api-go/conf/storage/nosql"
)

func main() {
	client, _ := nosql.StartDB()
	starwarsDB := client.Database("starwars")

	rep := list.NewRepository(starwarsDB)
	a, _ := rep.GetAll()
	documents, _ := json.Marshal(a)
	fmt.Println(documents)
}
