package gamedetails

import (
	"encoding/json"
	"nba-api-go/requests"
	"time"
)

type GameDetails struct {
	Requests *requests.Requests
	API      struct {
		Status  int      `json:"status"`
		Message string   `json:"message"`
		Results int      `json:"results"`
		Filters []string `json:"filters"`
		Game    []Game   `json:"game"`
	} `json:"api"`
}

type Game struct {
	SeasonYear      string       `json:"seasonYear"`
	League          string       `json:"league"`
	GameID          string       `json:"gameId"`
	StartTimeUTC    time.Time    `json:"startTimeUTC"`
	EndTimeUTC      time.Time    `json:"endTimeUTC"`
	Arena           string       `json:"arena"`
	City            string       `json:"city"`
	Country         string       `json:"country"`
	Clock           string       `json:"clock"`
	GameDuration    string       `json:"gameDuration"`
	TimesTied       string       `json:"timesTied"`
	LeadChanges     string       `json:"leadChanges"`
	CurrentPeriod   string       `json:"currentPeriod"`
	Halftime        string       `json:"halftime"`
	EndOfPeriod     string       `json:"EndOfPeriod"`
	SeasonStage     string       `json:"seasonStage"`
	StatusShortGame string       `json:"statusShortGame"`
	StatusGame      string       `json:"statusGame"`
	VTeam           PlayingTeams `json:"vTeam"`
	HTeam           PlayingTeams `json:"hTeam"`
	Officials       []Official   `json:"officials"`
}

type Official struct {
	Name string `json:"name"`
}

type PlayingTeams struct {
	FullName     string `json:"fullName"`
	TeamID       string `json:"teamId"`
	Nickname     string `json:"nickname"`
	Logo         string `json:"logo"`
	ShortName    string `json:"shortName"`
	AllStar      string `json:"allStar"`
	NbaFranchise string `json:"nbaFranchise"`
	Score        struct {
		Win        string   `json:"win"`
		Loss       string   `json:"loss"`
		SeriesWin  string   `json:"seriesWin"`
		SeriesLoss string   `json:"seriesLoss"`
		Linescore  []string `json:"linescore"`
		Points     string   `json:"points"`
	} `json:"score"`
	Leaders []struct {
		Points   string `json:"points,omitempty"`
		PlayerID string `json:"playerId"`
		Name     string `json:"name"`
		Rebounds string `json:"rebounds,omitempty"`
		Assists  string `json:"assists,omitempty"`
	} `json:"leaders"`
}

var (
	gameDetailsURL = "/gameDetails/"
)

func (n *GameDetails) GetGameDetailsByID(gameID string) (GameDetails, error) {
	return getGameDetails(n.Requests, gameID)
}

func getGameDetails(requests *requests.Requests, gameID string) (gamedetails GameDetails, err error) {

	gameDetailsResponse, err := requests.NewGetRequest(gameDetailsURL + gameID)

	if err != nil {
		return gamedetails, err
	}

	err = json.Unmarshal(gameDetailsResponse, &gamedetails)

	if err != nil {
		return gamedetails, err
	}

	return gamedetails, err
}
