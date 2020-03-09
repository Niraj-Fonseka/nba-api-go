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

/*
	Seasons
*/

func (n *NBAapi) GetSeasons() (seasons.Seasons, error) {
	return seasons.GetSeasons(n.Requests)
}

/*
	Leagues
*/

func (n *NBAapi) GetLeagues() (leagues.Leagues, error) {
	return leagues.GetLeagues(n.Requests)
}

/*
	Teams
*/

func (n *NBAapi) GetTeamsByTeamID(teamId string) (teams.Teams, error) {
	return teams.GetTeams(n.Requests, "teamId/", teamId)
}

func (n *NBAapi) GetTeamsByLeague(league string) (teams.Teams, error) {
	return teams.GetTeams(n.Requests, "league/", league)
}

func (n *NBAapi) GetTeamsByCity(city string) (teams.Teams, error) {
	return teams.GetTeams(n.Requests, "city/", city)
}

func (n *NBAapi) GetTeamsByShortName(shortName string) (teams.Teams, error) {
	return teams.GetTeams(n.Requests, "shortName/", shortName)
}

func (n *NBAapi) GetTeamsByNickName(nickName string) (teams.Teams, error) {
	return teams.GetTeams(n.Requests, "league/", nickName)
}

func (n *NBAapi) GetTeamsByConfName(confName string) (teams.Teams, error) {
	return teams.GetTeams(n.Requests, "confName/", confName)
}

func (n *NBAapi) GetTeamsByDivName(divName string) (teams.Teams, error) {
	return teams.GetTeams(n.Requests, "divName/", divName)
}

/*
	Players
*/

func (n *NBAapi) GetPlayersByPlayerID(playerID string) (players.Players, error) {
	return players.GetPlayers(n.Requests, "playerId/", playerID)
}

func (n *NBAapi) GetPlayersByTeamID(teamID string) (players.Players, error) {
	return players.GetPlayers(n.Requests, "teamId/", teamID)
}

func (n *NBAapi) GetPlayersByLeague(league string) (players.Players, error) {
	return players.GetPlayers(n.Requests, "league/", league)
}

func (n *NBAapi) GetPlayersByCountry(country string) (players.Players, error) {
	return players.GetPlayers(n.Requests, "country/", country)
}

func (n *NBAapi) GetPlayersByFirstName(firstName string) (players.Players, error) {
	return players.GetPlayers(n.Requests, "firstName/", firstName)
}

func (n *NBAapi) GetPlayersByLastName(lastName string) (players.Players, error) {
	return players.GetPlayers(n.Requests, "lastName/", lastName)
}

/*
	Games
*/

func (n *NBAapi) GetGamesBySeasonYear(year string) (games.Games, error) {
	return games.GetGames(n.Requests, "seasonYear/", year)
}

func (n *NBAapi) GetGamesByLeague(year string) (games.Games, error) {
	return games.GetGames(n.Requests, "seasonYear/", year)
}

func (n *NBAapi) GetGamesByGameID(gameId string) (games.Games, error) {
	return games.GetGames(n.Requests, "gameId/", gameId)
}

func (n *NBAapi) GetGamesByTeamID(teamId string) (games.Games, error) {
	return games.GetGames(n.Requests, "teamId/", teamId)
}

func (n *NBAapi) GetGamesByDate(date string) (games.Games, error) {
	return games.GetGames(n.Requests, "date/", date)
}

func (n *NBAapi) GetGamesByLive(year string) (games.Games, error) {
	return games.GetGames(n.Requests, "live/", "")
}

/*
	GameDetails
*/

func (n *NBAapi) GetGameDetailsByID(gameID string) (gamedetails.GameDetails, error) {
	return gamedetails.GetGameDetails(n.Requests, gameID)
}

/*
	Standings
*/

func (n *NBAapi) GetStandingsByConference(league string, seasonYear string, conference string) (standings.Standings, error) {
	return standings.GetStandings(n.Requests, league, seasonYear, "conference/", conference)
}

func (n *NBAapi) GetStandingsByDivision(league string, seasonYear string, division string) (standings.Standings, error) {
	return standings.GetStandings(n.Requests, league, seasonYear, "division/", division)
}

func (n *NBAapi) GetStandingsByTeamID(league string, seasonYear string, teamID string) (standings.Standings, error) {
	return standings.GetStandings(n.Requests, league, seasonYear, "teamId/", teamID)
}

func (n *NBAapi) GetStandings(league string, seasonYear string) (standings.Standings, error) {
	return standings.GetStandings(n.Requests, league, seasonYear, "", "")
}

/*
	Statistics
*/

func (n *NBAapi) GetGameStatisticsByGameID(gameID string) (statistics.Statistics, error) {
	return statistics.GetStastics(n.Requests, "games/gameId/", gameID)
}

func (n *NBAapi) GetPlayerStatisticsByGameID(gameID string) (statistics.Statistics, error) {
	return statistics.GetStastics(n.Requests, "players/gameId/", gameID)
}

func (n *NBAapi) GetPlayerStatisticsPlayerID(playerID string) (statistics.Statistics, error) {
	return statistics.GetStastics(n.Requests, "players/playerId/", playerID)
}
