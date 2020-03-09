package main

import (
	api "nba-api-go"
	pretty "nba-api-go/examples"
	"os"
)

func main() {
	client := api.NewNBAapiClient(os.Getenv("BASEURL"), os.Getenv("HOST"), os.Getenv("KEY"))

	leagues, err := client.Leagues().GetLeagues()

	if err != nil {
		panic(err)
	}

	pretty.MakePrettyJSON(leagues)
}
