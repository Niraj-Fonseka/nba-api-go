package teams

import (
	"encoding/json"
	"nba-api-go/requests"
)

var (
	teamsURL = "/teams/"
)

type Teams struct {
	API struct {
		Status  int      `json:"status"`
		Message string   `json:"message"`
		Results int      `json:"results"`
		Filters []string `json:"filters"`
		Teams   []Team   `json:"teams"`
	} `json:"api"`
}

type Team struct {
	City         string `json:"city"`
	FullName     string `json:"fullName"`
	TeamID       string `json:"teamId"`
	Nickname     string `json:"nickname"`
	Logo         string `json:"logo"`
	ShortName    string `json:"shortName"`
	AllStar      string `json:"allStar"`
	NbaFranchise string `json:"nbaFranchise"`
	Leagues      struct {
		Standard struct {
			ConfName string `json:"confName"`
			DivName  string `json:"divName"`
		} `json:"standard"`
	} `json:"leagues"`
}

func GetTeamsByLeague(requests *requests.Requests, league string) (teams Teams, err error) {
	teamResponse, err := requests.NewGetRequest(teamsURL + "league/" + league)

	if err != nil {
		return teams, err
	}

	err = json.Unmarshal(teamResponse, &teams)

	if err != nil {
		return teams, err
	}

	return teams, err
}
