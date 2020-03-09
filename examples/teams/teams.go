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

	seasons, err := client.GetTeamsByTeamID("1")

	if err != nil {
		panic(err)
	}

	pretty.MakePrettyJSON(seasons)
}

func GetTeamsByLeague() {
	client := api.NewNBAapiClient(os.Getenv("BASEURL"), os.Getenv("HOST"), os.Getenv("KEY"))

	seasons, err := client.GetTeamsByLeague("standard")

	if err != nil {
		panic(err)
	}

	pretty.MakePrettyJSON(seasons)
}

func GetTeamsByCity() {
	client := api.NewNBAapiClient(os.Getenv("BASEURL"), os.Getenv("HOST"), os.Getenv("KEY"))

	seasons, err := client.GetTeamsByCity("Atlanta")

	if err != nil {
		panic(err)
	}

	pretty.MakePrettyJSON(seasons)
}

func GetTeamsByShortName() {
	client := api.NewNBAapiClient(os.Getenv("BASEURL"), os.Getenv("HOST"), os.Getenv("KEY"))

	seasons, err := client.GetTeamsByShortName("ATL")

	if err != nil {
		panic(err)
	}

	pretty.MakePrettyJSON(seasons)
}

func GetTeamsByNickName() {
	client := api.NewNBAapiClient(os.Getenv("BASEURL"), os.Getenv("HOST"), os.Getenv("KEY"))

	seasons, err := client.GetTeamsByNickName("Hawks")

	if err != nil {
		panic(err)
	}

	pretty.MakePrettyJSON(seasons)
}

func GetTeamsByConfName() {
	client := api.NewNBAapiClient(os.Getenv("BASEURL"), os.Getenv("HOST"), os.Getenv("KEY"))

	seasons, err := client.GetTeamsByConfName("East")

	if err != nil {
		panic(err)
	}

	pretty.MakePrettyJSON(seasons)
}

func GetTeamsByDivName() {
	client := api.NewNBAapiClient(os.Getenv("BASEURL"), os.Getenv("HOST"), os.Getenv("KEY"))

	seasons, err := client.GetTeamsByDivName("Southeast")

	if err != nil {
		panic(err)
	}

	pretty.MakePrettyJSON(seasons)
}
