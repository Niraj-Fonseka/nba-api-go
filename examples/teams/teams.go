package main

import (
	api "nba-api-go"
	pretty "nba-api-go/examples"
	"os"
)

func main() {
	GetTeamsTeamID()
	GetTeamsByLeague()
	GetTeamsByCity()
	GetTeamsByShortName()
	GetTeamsByNickName()
	GetTeamsByConfName()
	GetTeamsByDivName()
}

func GetTeamsTeamID() {
	client := api.NewNBAapiClient(os.Getenv("BASEURL"), os.Getenv("HOST"), os.Getenv("KEY"))

	seasons, err := client.Teams().GetTeamsByTeamID("1")

	if err != nil {
		panic(err)
	}

	pretty.MakePrettyJSON(seasons)
}

func GetTeamsByLeague() {
	client := api.NewNBAapiClient(os.Getenv("BASEURL"), os.Getenv("HOST"), os.Getenv("KEY"))

	seasons, err := client.Teams().GetTeamsByLeague("standard")

	if err != nil {
		panic(err)
	}

	pretty.MakePrettyJSON(seasons)
}

func GetTeamsByCity() {
	client := api.NewNBAapiClient(os.Getenv("BASEURL"), os.Getenv("HOST"), os.Getenv("KEY"))

	seasons, err := client.Teams().GetTeamsByCity("Atlanta")

	if err != nil {
		panic(err)
	}

	pretty.MakePrettyJSON(seasons)
}

func GetTeamsByShortName() {
	client := api.NewNBAapiClient(os.Getenv("BASEURL"), os.Getenv("HOST"), os.Getenv("KEY"))

	seasons, err := client.Teams().GetTeamsByShortName("ATL")

	if err != nil {
		panic(err)
	}

	pretty.MakePrettyJSON(seasons)
}

func GetTeamsByNickName() {
	client := api.NewNBAapiClient(os.Getenv("BASEURL"), os.Getenv("HOST"), os.Getenv("KEY"))

	seasons, err := client.Teams().GetTeamsByNickName("Hornets")

	if err != nil {
		panic(err)
	}

	pretty.MakePrettyJSON(seasons)
}

func GetTeamsByConfName() {
	client := api.NewNBAapiClient(os.Getenv("BASEURL"), os.Getenv("HOST"), os.Getenv("KEY"))

	seasons, err := client.Teams().GetTeamsByConfName("East")

	if err != nil {
		panic(err)
	}

	pretty.MakePrettyJSON(seasons)
}

func GetTeamsByDivName() {
	client := api.NewNBAapiClient(os.Getenv("BASEURL"), os.Getenv("HOST"), os.Getenv("KEY"))

	seasons, err := client.Teams().GetTeamsByDivName("Southeast")

	if err != nil {
		panic(err)
	}

	pretty.MakePrettyJSON(seasons)
}
