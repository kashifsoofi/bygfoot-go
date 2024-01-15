package domain

// A statistics element holding some string and integer values.
type Stat struct {
	teamName               string
	value1, value2, value3 int
	valueString            string
}

/** A structure holding some stat arrays about a league. */
type LeagueStat struct {
	leagueSymbol                string
	leagueName                  string
	teamsOff, teamsDef          []*Team   // Best offensive and defensive teams.
	playerScores, playerGoalies []*Player // Best goal getters and goalies.
}

// A team name and a competition name.
type ChampStat struct {
	team_name, cl_name string
}

// A season statistics structure.
type SeasonStat struct {
	seasonNumber           int           // Which season
	leagueChamps, cupChams []*Team       // League and cup winners.
	leagueStats            []*LeagueStat // The league stats at the end of the season.
}
