package main

import (
	api "nba-api-go"
	pretty "nba-api-go/examples"
	"os"
)

func main() {
	client := api.NewNBAapiClient(os.Getenv("BASEURL"), os.Getenv("HOST"), os.Getenv("KEY"))

	seasons, err := client.GetTeamsByLeague("standard")

	if err != nil {
		panic(err)
	}

	pretty.MakePrettyJSON(seasons)
}
