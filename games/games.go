package games

import (
	"encoding/json"
	"nba-api-go/requests"
	"time"
)

var (
	playersURL = "/players/"
)

type Games struct {
	Requests *requests.Requests
	API      struct {
		Status  int      `json:"status"`
		Message string   `json:"message"`
		Results int      `json:"results"`
		Filters []string `json:"filters"`
		Games   []Game   `json:"games"`
	} `json:"api"`
}

type Game struct {
	SeasonYear      string    `json:"seasonYear"`
	League          string    `json:"league"`
	GameID          string    `json:"gameId"`
	StartTimeUTC    time.Time `json:"startTimeUTC"`
	EndTimeUTC      time.Time `json:"endTimeUTC"`
	Arena           string    `json:"arena"`
	City            string    `json:"city"`
	Country         string    `json:"country"`
	Clock           string    `json:"clock"`
	GameDuration    string    `json:"gameDuration"`
	CurrentPeriod   string    `json:"currentPeriod"`
	Halftime        string    `json:"halftime"`
	EndOfPeriod     string    `json:"EndOfPeriod"`
	SeasonStage     string    `json:"seasonStage"`
	StatusShortGame string    `json:"statusShortGame"`
	StatusGame      string    `json:"statusGame"`
	VTeam           struct {
		TeamID    string `json:"teamId"`
		ShortName string `json:"shortName"`
		FullName  string `json:"fullName"`
		NickName  string `json:"nickName"`
		Logo      string `json:"logo"`
		Score     struct {
			Points string `json:"points"`
		} `json:"score"`
	} `json:"vTeam"`
	HTeam struct {
		TeamID    string `json:"teamId"`
		ShortName string `json:"shortName"`
		FullName  string `json:"fullName"`
		NickName  string `json:"nickName"`
		Logo      string `json:"logo"`
		Score     struct {
			Points string `json:"points"`
		} `json:"score"`
	} `json:"hTeam"`
}

func (n *Games) GetGamesBySeasonYear(year string) (Games, error) {
	return getGames(n.Requests, "seasonYear/", year)
}

func (n *Games) GetGamesByLeague(year string) (Games, error) {
	return getGames(n.Requests, "seasonYear/", year)
}

func (n *Games) GetGamesByGameID(gameId string) (Games, error) {
	return getGames(n.Requests, "gameId/", gameId)
}

func (n *Games) GetGamesByTeamID(teamId string) (Games, error) {
	return getGames(n.Requests, "teamId/", teamId)
}

func (n *Games) GetGamesByDate(date string) (Games, error) {
	return getGames(n.Requests, "date/", date)
}

func (n *Games) GetGamesByLive(year string) (Games, error) {
	return getGames(n.Requests, "live/", "")
}

func getGames(requests *requests.Requests, filter string, filterValue string) (players Games, err error) {

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
