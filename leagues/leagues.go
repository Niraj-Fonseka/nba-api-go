package leagues

import (
	"encoding/json"
	"nba-api-go/requests"
)

var (
	leaguesURL = "/leagues/"
)

type Leagues struct {
	Requests *requests.Requests
	API      struct {
		Status  int      `json:"status"`
		Message string   `json:"message"`
		Results int      `json:"results"`
		Filters []string `json:"filters"`
		Leagues []string `json:"leagues"`
	} `json:"api"`
}

func (n *Leagues) GetLeagues() (Leagues, error) {
	return getLeagues(n.Requests)
}

func getLeagues(requests *requests.Requests) (leagues Leagues, err error) {

	leaguesResponse, err := requests.NewGetRequest(leaguesURL)

	if err != nil {
		return leagues, err
	}

	err = json.Unmarshal(leaguesResponse, &leagues)

	if err != nil {
		return leagues, err
	}

	return leagues, err
}
