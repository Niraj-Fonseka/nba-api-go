package seasons

import (
	"encoding/json"
	"nba-api-go/requests"
)

var (
	seasonsURL = "/seasons/"
)

type Seasons struct {
	API struct {
		Status  int      `json:"status"`
		Message string   `json:"message"`
		Results int      `json:"results"`
		Filters []string `json:"filters"`
		Seasons []string `json:"seasons"`
	} `json:"api"`
}

func GetSeasons(requests *requests.Requests) (seasons Seasons, err error) {

	seasonsResponse, err := requests.NewGetRequest(seasonsURL)

	if err != nil {
		return seasons, err
	}

	err = json.Unmarshal(seasonsResponse, &seasons)

	if err != nil {
		return seasons, err
	}
	return seasons, err
}
