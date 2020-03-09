package teams

import (
	"encoding/json"
	"nba-api-go/requests"
)

var (
	teamsURL = "/teams/"
)

type Teams struct {
	Requests *requests.Requests
	API      struct {
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

func (n *Teams) GetTeamsByTeamID(teamId string) (Teams, error) {
	return getTeams(n.Requests, "teamId/", teamId)
}

func (n *Teams) GetTeamsByLeague(league string) (Teams, error) {
	return getTeams(n.Requests, "league/", league)
}

func (n *Teams) GetTeamsByCity(city string) (Teams, error) {
	return getTeams(n.Requests, "city/", city)
}

func (n *Teams) GetTeamsByShortName(shortName string) (Teams, error) {
	return getTeams(n.Requests, "shortName/", shortName)
}

func (n *Teams) GetTeamsByNickName(nickName string) (Teams, error) {
	return getTeams(n.Requests, "nickName/", nickName)
}

func (n *Teams) GetTeamsByConfName(confName string) (Teams, error) {
	return getTeams(n.Requests, "confName/", confName)
}

func (n *Teams) GetTeamsByDivName(divName string) (Teams, error) {
	return getTeams(n.Requests, "divName/", divName)
}

func getTeams(requests *requests.Requests, filter string, filterValue string) (teams Teams, err error) {

	teamsResponse, err := requests.NewGetRequest(teamsURL + filter + filterValue)

	if err != nil {
		return teams, err
	}

	err = json.Unmarshal(teamsResponse, &teams)

	if err != nil {
		return teams, err
	}

	return teams, err
}
