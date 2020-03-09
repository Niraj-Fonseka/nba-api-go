package nba

import (
	"nba-api-go/gamedetails"
	"nba-api-go/games"
	"nba-api-go/leagues"
	"nba-api-go/players"
	"nba-api-go/requests"
	"nba-api-go/seasons"
	"nba-api-go/standings"
	"nba-api-go/statistics"

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

func (n *NBAapi) Leagues() *leagues.Leagues {
	return &leagues.Leagues{
		Requests: n.Requests,
	}
}

func (n *NBAapi) Games() *games.Games {
	return &games.Games{
		Requests: n.Requests,
	}
}

func (n *NBAapi) Players() *players.Players {
	return &players.Players{
		Requests: n.Requests,
	}
}

func (n *NBAapi) Seasons() *seasons.Seasons {
	return &seasons.Seasons{
		Requests: n.Requests,
	}
}

func (n *NBAapi) Standings() *standings.Standings {
	return &standings.Standings{
		Requests: n.Requests,
	}
}

func (n *NBAapi) Statistics() *statistics.Statistics {
	return &statistics.Statistics{
		Requests: n.Requests,
	}
}

func (n *NBAapi) GameDetails() *gamedetails.GameDetails {
	return &gamedetails.GameDetails{
		Requests: n.Requests,
	}
}
