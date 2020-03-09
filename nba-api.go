package nba

import (
	"nba-api-go/leagues"
	"nba-api-go/requests"
	"nba-api-go/seasons"
	"nba-api-go/teams"

	"net/http"
)

type NBAapi struct {
	Requests *requests.Requests
}

func NewNBAapiClient(url string, host string, key string) *NBAapi {

	return &NBAapi{
		Requests: &requests.Requests{
			BaseURL: url,
			Host:    host,
			Key:     key,
			Client:  &http.Client{},
		},
	}
}

func (n *NBAapi) GetSeasons() (seasons.Seasons, error) {
	return seasons.GetSeasons(n.Requests)
}

func (n *NBAapi) GetLeagues() (leagues.Leagues, error) {
	return leagues.GetLeagues(n.Requests)
}

func (n *NBAapi) GetTeamsByLeague(league string) (teams.Teams, error) {
	return teams.GetTeamsByLeague(n.Requests, league)
}
