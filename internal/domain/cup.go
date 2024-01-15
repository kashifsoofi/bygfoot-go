package domain

// Rules for a round of a cup.
// Cups consist of rounds, e.g. the final counts as
// a round or the round robin games.
type CupRound struct {
	homeAway                      bool             // Whether there are home and away games or only one leg. Default: TRUE.
	replay                        int              // How many times a match gets repeated if there's a draw. Default: 0.
	neutral                       bool             // Whether the matches are on neutral ground. Default: FALSE.
	randomiseTeams                bool             // Whether the teams array gets randomised before writing the fixtures. Default: TRUE. FALSE makes only sense in the first cup round.
	roundRobinNumberOfGroups      int              // How many round robin groups there are. Default: 0 (ie. no round robin).
	roundRobinNumberOfAdvance     int              // How many teams advance from each group. Default: -1.
	roundRobinNumberOfBestAdvance int              // How many teams advance apart from those that are specified already (e.g. the first two advance and additionally the best 3 from all the groups. Default: 0.
	newTeams                      int              // Number of new teams participating in the cup round (ie. teams that get loaded and are not advancing from a previous round).
	byes                          int              // The number of byes to be awarded for this cup round. The default is enough to bring the next round to a power of two.
	delay                         int              // Number of weeks the cup round is delayed (or scheduled sooner if the value is negative) with respect to the previous cup round and the week gap. Default: 0.
	twoMatchWeeks                 [2]int           // Here we store intervals of fixtures during which there should be two matches in a week instead of one. This is only relevant in round robin rounds.
	twoMatchWeek                  bool             // Whether the two matches of a home/away round are played in one week.
	teams                         []*Team          // The teams that got loaded for this cup round. Mostly this only happens in the first round.
	teamPtrs                      []*Team          // Pointers to all teams loaded in the cup round; these teams get passed to the fixture generation function together with the teams advancing from the previous round.
	choose_teams                  []*CupChooseTeam // Which new teams come into the cup (@see #CupChooseTeam)
	tables                        []Table          // The round robin tables (in case there is a round robin).
}

// Rules for choosing teams from the league files.
// This could tell us to select the first three teams
// from the league 'Italy 1' to participate in the cup.
type CupChooseTeam struct {
	sid              string // The string id of the league we choose from. Default: "".
	numberOfTeams    int    // The number of teams chosen. Default: -1 (ie. all teams are chosen).
	startIdx, endIdx int    // The range from which the teams are chosen, e.g. 2 and 5 means the teams from 2 to 5 are chosen (given that 'number_of_teams' is 4). Defaults: -1 (ie. the teams are chosen from the whole range of teams in the league).
	randomly         bool   // Whether the teams are chosen randomly, ie. 3 random teams from the range 1-20. Default: FALSE.
	generate         bool   // Whether the team is generated and loaded from a league file or taken from one of the country's leagues or cups. Default: FALSE.
}

// Structure representing a cup.
type Cup struct {
	name, shortName, symbol, sid string  // Name and short name of the cup, a pixmap path, and the string id (e.g. england_fa or so). Default: "".
	id                           int     // Numerical id.
	group                        int     // An integer specifying which cups are mutually exclusive for league teams, e.g. the same team can't participate in the UEFA Cup and the Champions' League.
	lastWeek, weekGap            int     // Last week (typically the week the final takes place) and weeks between matchdays. Default: -1 and 1.
	addWeek                      int     // This determines when the cup gets added to the acps pointer array (and becomes visible for the user). Also determines when the cup fixtures for the first round get written. Default: 0 (ie. the cup is visible from the first week).
	yellowRed                    int     // Number of yellow cards that lead to a missed match. Default: 1000 (off).
	talentDiff                   float64 // Difference of the average talent for the cup teams compared to the league with highest average talent. Default: 0.
	nextFixtureUpdateWeek        int     // The week and week_round at the beginning of which the fixtures have to be updated.
	nextFixtureUpdateWeekRound   int
	properties                   []string    // A gchar pointer array of properties (like "national").
	rounds                       []*CupRound // The rounds of the cup. @see #CupRound
	bye                          []*Team     // Pointer array containing teams that got a bye for a round of the cup.
	teams                        []*Team     // The teams belonging to the cup (stored in the cup rounds, these are only pointers). Relevant only if it's an international one.
	teamNames                    []string    // Pointer array with the names of all the teams in the cup. Also the teams from the country's leagues.
	fixtures                     []*Fixture  // The fixtures of a season for the cup.
}
