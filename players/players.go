package players

import (
	"encoding/json"
	"nba-api-go/requests"
)

var (
	playersURL = "/players/"
)

type Players struct {
	Requests *requests.Requests
	API      struct {
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

func (n *Players) GetPlayersByPlayerID(playerID string) (Players, error) {
	return getPlayers(n.Requests, "playerId/", playerID)
}

func (n *Players) GetPlayersByTeamID(teamID string) (Players, error) {
	return getPlayers(n.Requests, "teamId/", teamID)
}

func (n *Players) GetPlayersByLeague(league string) (Players, error) {
	return getPlayers(n.Requests, "league/", league)
}

func (n *Players) GetPlayersByCountry(country string) (Players, error) {
	return getPlayers(n.Requests, "country/", country)
}

func (n *Players) GetPlayersByFirstName(firstName string) (Players, error) {
	return getPlayers(n.Requests, "firstName/", firstName)
}

func (n *Players) GetPlayersByLastName(lastName string) (Players, error) {
	return getPlayers(n.Requests, "lastName/", lastName)
}

func getPlayers(requests *requests.Requests, filter string, filterValue string) (players Players, err error) {

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
