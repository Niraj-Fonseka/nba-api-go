package standings

import (
	"encoding/json"
	"nba-api-go/requests"
)

type Standings struct {
	Requests *requests.Requests
	API      struct {
		Status    int        `json:"status"`
		Message   string     `json:"message"`
		Results   int        `json:"results"`
		Filters   []string   `json:"filters"`
		Standings []Standing `json:"standings"`
	} `json:"api"`
}

type Standing struct {
	League           string     `json:"league"`
	TeamID           string     `json:"teamId"`
	Win              string     `json:"win"`
	Loss             string     `json:"loss"`
	GamesBehind      string     `json:"gamesBehind"`
	LastTenWin       string     `json:"lastTenWin"`
	LastTenLoss      string     `json:"lastTenLoss"`
	Streak           string     `json:"streak"`
	SeasonYear       string     `json:"seasonYear"`
	Conference       Conference `json:"conference"`
	Division         Division   `json:"division"`
	WinPercentage    string     `json:"winPercentage"`
	LossPercentage   string     `json:"lossPercentage"`
	Home             Location   `json:"home"`
	Away             Location   `json:"away"`
	WinStreak        string     `json:"winStreak"`
	TieBreakerPoints string     `json:"tieBreakerPoints"`
}

type Conference struct {
	Name string `json:"name"`
	Rank string `json:"rank"`
	Win  string `json:"win"`
	Loss string `json:"loss"`
}

type Location struct {
	Win  string `json:"win"`
	Loss string `json:"loss"`
}

type Division struct {
	Name        string `json:"name"`
	Rank        string `json:"rank"`
	Win         string `json:"win"`
	Loss        string `json:"loss"`
	GamesBehind string `json:"GamesBehind"`
}

var (
	standingsURL = "/standings/"
)

func (n *Standings) GetStandingsByConference(league string, seasonYear string, conference string) (Standings, error) {
	return getStandings(n.Requests, league, seasonYear, "conference/", conference)
}

func (n *Standings) GetStandingsByDivision(league string, seasonYear string, division string) (Standings, error) {
	return getStandings(n.Requests, league, seasonYear, "division/", division)
}

func (n *Standings) GetStandingsByTeamID(league string, seasonYear string, teamID string) (Standings, error) {
	return getStandings(n.Requests, league, seasonYear, "teamId/", teamID)
}

func (n *Standings) GetStandings(league string, seasonYear string) (Standings, error) {
	return getStandings(n.Requests, league, seasonYear, "", "")
}

func getStandings(requests *requests.Requests, league string, year string, filter string, filterValue string) (standings Standings, err error) {

	teamsResponse, err := requests.NewGetRequest(standingsURL + league + "/" + year + "/" + filter + filterValue)

	if err != nil {
		return standings, err
	}

	err = json.Unmarshal(teamsResponse, &standings)

	if err != nil {
		return standings, err
	}

	return standings, err
}
