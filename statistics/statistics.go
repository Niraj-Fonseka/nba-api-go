package statistics

import (
	"encoding/json"
	"nba-api-go/requests"
)

type Statistics struct {
	Requests *requests.Requests
	API      struct {
		Status     int         `json:"status"`
		Message    string      `json:"message"`
		Results    int         `json:"results"`
		Filters    []string    `json:"filters"`
		Statistics []Statistic `json:"statistics"`
	} `json:"api"`
}

type Statistic struct {
	GameID             string `json:"gameId"`
	TeamID             string `json:"teamId"`
	FastBreakPoints    string `json:"fastBreakPoints"`
	PointsInPaint      string `json:"pointsInPaint"`
	BiggestLead        string `json:"biggestLead"`
	SecondChancePoints string `json:"secondChancePoints"`
	PointsOffTurnovers string `json:"pointsOffTurnovers"`
	LongestRun         string `json:"longestRun"`
	Points             string `json:"points"`
	Fgm                string `json:"fgm"`
	Fga                string `json:"fga"`
	Fgp                string `json:"fgp"`
	Ftm                string `json:"ftm"`
	Fta                string `json:"fta"`
	Ftp                string `json:"ftp"`
	Tpm                string `json:"tpm"`
	Tpa                string `json:"tpa"`
	Tpp                string `json:"tpp"`
	OffReb             string `json:"offReb"`
	DefReb             string `json:"defReb"`
	TotReb             string `json:"totReb"`
	Assists            string `json:"assists"`
	PFouls             string `json:"pFouls"`
	Steals             string `json:"steals"`
	Turnovers          string `json:"turnovers"`
	Blocks             string `json:"blocks"`
	PlusMinus          string `json:"plusMinus"`
	Min                string `json:"min"`
}

var (
	statisticsURL = "/statistics/"
)

func (n *Statistics) GetGameStatisticsByGameID(gameID string) (Statistics, error) {
	return getStastics(n.Requests, "games/gameId/", gameID)
}

func (n *Statistics) GetPlayerStatisticsByGameID(gameID string) (Statistics, error) {
	return getStastics(n.Requests, "players/gameId/", gameID)
}

func (n *Statistics) GetPlayerStatisticsPlayerID(playerID string) (Statistics, error) {
	return getStastics(n.Requests, "players/playerId/", playerID)
}

func getStastics(requests *requests.Requests, filter string, filterValue string) (stats Statistics, err error) {

	teamsResponse, err := requests.NewGetRequest(statisticsURL + filter + filterValue)

	if err != nil {
		return stats, err
	}

	err = json.Unmarshal(teamsResponse, &stats)

	if err != nil {
		return stats, err
	}

	return stats, err
}
