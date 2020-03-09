package leagues

import (
	"encoding/json"
	"nba-api-go/requests"
)

var (
	leaguesURL = "/leagues/"
)

type Leagues struct {
	API struct {
		Status  int      `json:"status"`
		Message string   `json:"message"`
		Results int      `json:"results"`
		Filters []string `json:"filters"`
		Leagues []string `json:"leagues"`
	} `json:"api"`
}

func GetLeagues(requests *requests.Requests) (leagues Leagues, err error) {

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
