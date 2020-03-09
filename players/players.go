package players

import (
	"encoding/json"
	"nba-api-go/requests"
)

var (
	playersURL = "/players/"
)

type Players struct {
	API struct {
		Status  int      `json:"status"`
		Message string   `json:"message"`
		Results int      `json:"results"`
		Filters []string `json:"filters"`
		Players []Player `json:"players"`
	} `json:"api"`
}

type Player struct {
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	TeamID            string `json:"teamId"`
	YearsPro          string `json:"yearsPro"`
	CollegeName       string `json:"collegeName"`
	Country           string `json:"country"`
	PlayerID          string `json:"playerId"`
	DateOfBirth       string `json:"dateOfBirth"`
	Affiliation       string `json:"affiliation"`
	StartNba          string `json:"startNba"`
	HeightInMeters    string `json:"heightInMeters"`
	WeightInKilograms string `json:"weightInKilograms"`
	Leagues           struct {
		Standard struct {
			Jersey string `json:"jersey"`
			Active string `json:"active"`
			Pos    string `json:"pos"`
		} `json:"standard"`
	} `json:"leagues"`
}

func GetPlayers(requests *requests.Requests, filter string, filterValue string) (players Players, err error) {

	teamsResponse, err := requests.NewGetRequest(playersURL + filter + filterValue)

	if err != nil {
		return players, err
	}

	err = json.Unmarshal(teamsResponse, &players)

	if err != nil {
		return players, err
	}

	return players, err
}
