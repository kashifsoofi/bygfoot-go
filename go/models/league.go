package models

type PromRelType int

const (
	PROM_REL_PROMOTION PromRelType = iota
	PROM_REL_RELEGATION
	PROM_REL_NONE
)

// An element representing a promotion or relegation rule.
// This means, a PromRelElement specifies a range of teams
// that get promoted or relegated to a given league.
// @see PromRel
type PromRelElement struct {
	ranks  [2]int      // The range of teams; default 0 and 0
	estSid string      // The id of the destination league. Default ""
	Type   PromRelType // PromRelType. Promotion or relegation or none.
}

// This structure specifies how promotion and relegation is handled in a league.
// It contains promotion and relegation rules in an array and possibly also
// a rule about promotion games to be played.
// @see PromRelElement
type PromRel struct {
	promGamesDestSid         string            // The id of the league the promotion games winner gets promoted to. Default ""
	promGamesLoserSid        string            // The id of the league the promotion games losers get moved to. Default ""
	promGamesNumberOfAdvance int               // Number of teams that advance from the promotion games. Default: 1.
	elements                 []*PromRelElement // Array with promotion/relegation rules. @see PromRelElement
	promGamesCupSid          string            // The cup determining how the promotion games are handled.
}

// Representation of a league. @see PromRel @see Table
type League struct {
	Id                           int        // Numerical id, as opposed to the string id 'sid'.
	Name, ShortName, Sid, Symbol string     // Default value ""
	NamesFile                    string     // The sid of the player names file the teams in the league take their names from. Default: 'general', meaning the 'player_names_general.xml' file.
	PromRel                      PromRel    // @see PromRel
	Layer                        int        // Layer of the league; this specifies which leagues are parallel.
	FirstWeek                    int        // The first week games are played. Default 1.
	WeekGap                      int        // Weeks between two matchdays. Default 1.
	TwoMatchWeeks                [2]int     // Here we store intervals of fixtures during which there should be two matches in a week instead of one.
	RoundRobins                  int        // How many round robins are played. Important for small leagues with 10 teams or so. Default: 1.
	RrBreak                      int        // Number of weeks between the parts of a round robin.
	YellowRed                    int        // Number of yellow cards until a player gets banned. Default 1000 (which means 'off', basically).
	AverageTalent                float64    // Average talent for the first season. Default: -1.
	Team                         []*Team    // Array of teams in the league. @see Team
	Table                        Table      // League table. @see Table
	Fixtures                     []*Fixture // The fixtures of a season for the league.
	Stats                        LeagueStat // The current league statistics.
	Active                       bool       // Whether the league is only a container for teams or really a league with fixtures and all.
}
