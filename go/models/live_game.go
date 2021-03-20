package models

// Events happening during a live game. @see #LiveGameEvent @see #LiveGameUnit
type LiveGameEventType int

const (
	// This is the 'main' event, nothing in particular is happening; one of the teams is in possession of the ball.
	LIVE_GAME_EVENT_GENERAL         LiveGameEventType = iota // 0
	LIVE_GAME_EVENT_START_MATCH                              // 1
	LIVE_GAME_EVENT_HALF_TIME                                // 2
	LIVE_GAME_EVENT_EXTRA_TIME                               // 3
	LIVE_GAME_EVENT_END_MATCH                                // 4
	LIVE_GAME_EVENT_LOST_POSSESSION                          // 5
	LIVE_GAME_EVENT_SCORING_CHANCE                           // 6
	LIVE_GAME_EVENT_HEADER                                   // 7
	LIVE_GAME_EVENT_PENALTY                                  // 8
	LIVE_GAME_EVENT_FREE_KICK                                // 9
	LIVE_GAME_EVENT_GOAL                                     // 10
	LIVE_GAME_EVENT_OWN_GOAL                                 // 11
	LIVE_GAME_EVENT_POST                                     // 12
	LIVE_GAME_EVENT_MISS                                     // 13
	LIVE_GAME_EVENT_SAVE                                     // 14
	LIVE_GAME_EVENT_CROSS_BAR                                // 15
	LIVE_GAME_EVENT_FOUL                                     // 16
	LIVE_GAME_EVENT_FOUL_YELLOW                              // 17
	LIVE_GAME_EVENT_FOUL_RED                                 // 18
	LIVE_GAME_EVENT_FOUL_RED_INJURY                          // 19
	LIVE_GAME_EVENT_SEND_OFF                                 // 20
	LIVE_GAME_EVENT_INJURY                                   // 21
	// An injury that permits the player to continue after some brief time.
	LIVE_GAME_EVENT_TEMP_INJURY                 // 22
	LIVE_GAME_EVENT_PENALTIES                   // 23
	LIVE_GAME_EVENT_STADIUM                     // 24
	LIVE_GAME_EVENT_STADIUM_BREAKDOWN           // 25
	LIVE_GAME_EVENT_STADIUM_RIOTS               // 26
	LIVE_GAME_EVENT_STADIUM_FIRE                // 27
	LIVE_GAME_EVENT_SUBSTITUTION                // 28
	LIVE_GAME_EVENT_STRUCTURE_CHANGE            // 29
	LIVE_GAME_EVENT_STYLE_CHANGE_ALL_OUT_DEFEND // 30
	LIVE_GAME_EVENT_STYLE_CHANGE_DEFEND         // 31
	LIVE_GAME_EVENT_STYLE_CHANGE_BALANCED       // 32
	LIVE_GAME_EVENT_STYLE_CHANGE_ATTACK         // 33
	LIVE_GAME_EVENT_STYLE_CHANGE_ALL_OUT_ATTACK // 34
	LIVE_GAME_EVENT_BOOST_CHANGE_ANTI           // 35
	LIVE_GAME_EVENT_BOOST_CHANGE_OFF            // 36
	LIVE_GAME_EVENT_BOOST_CHANGE_ON             // 37
	LIVE_GAME_EVENT_END
)

type LiveGameUnitArea int

const (
	LIVE_GAME_UNIT_AREA_DEFEND LiveGameUnitArea = iota
	LIVE_GAME_UNIT_AREA_MIDFIELD
	LIVE_GAME_UNIT_AREA_ATTACK
	LIVE_GAME_UNIT_AREA_END
)

// Indices for the time variable of th #LiveGameUnit struct.
type LiveGameUnitTime int

const (
	LIVE_GAME_UNIT_TIME_FIRST_HALF LiveGameUnitTime = iota
	LIVE_GAME_UNIT_TIME_SECOND_HALF
	LIVE_GAME_UNIT_TIME_EXTRA_TIME
	LIVE_GAME_UNIT_TIME_PENALTIES
	LIVE_GAME_UNIT_TIME_END
)

// Indices for the values in #LiveGameStats.
type LiveGameStatValue int

const (
	LIVE_GAME_STAT_VALUE_GOALS_REGULAR LiveGameStatValue = iota
	LIVE_GAME_STAT_VALUE_SHOTS
	LIVE_GAME_STAT_VALUE_SHOT_PERCENTAGE
	LIVE_GAME_STAT_VALUE_POSSESSION
	LIVE_GAME_STAT_VALUE_PENALTIES
	LIVE_GAME_STAT_VALUE_FOULS
	LIVE_GAME_STAT_VALUE_CARDS
	LIVE_GAME_STAT_VALUE_REDS
	LIVE_GAME_STAT_VALUE_INJURIES
	LIVE_GAME_STAT_VALUE_END
)

type LiveGameStatArray int

const (
	LIVE_GAME_STAT_ARRAY_SCORERS LiveGameStatArray = iota
	LIVE_GAME_STAT_ARRAY_YELLOWS
	LIVE_GAME_STAT_ARRAY_REDS
	LIVE_GAME_STAT_ARRAY_INJURED
	LIVE_GAME_STAT_ARRAY_END
)

// Indices for the team_value array. @see game_get_values()
type GameTeamValue int

const (
	GAME_TEAM_VALUE_GOALIE GameTeamValue = iota
	GAME_TEAM_VALUE_DEFEND
	GAME_TEAM_VALUE_MIDFIELD
	GAME_TEAM_VALUE_ATTACK
	GAME_TEAM_VALUE_END
)

// Some stats for a live game like ball possession, shots on goal etc.
type LiveGameStats struct {
	possession float64
	values     [2][LIVE_GAME_STAT_VALUE_END]int
	players    [2][LIVE_GAME_STAT_ARRAY_END]*Player
}

// A struct telling us what's happening at a certain moment in a game.
type LiveGameEvent struct {
	eventType LiveGameEventType // @see #LiveGameEventType
	verbosity int               // Verbosity value. The lower the more important the event.
	// Information about a team and two players involved in the event.
	team, player, player2 int // The commentary for the event.
	commentary            string
	commentary_id         int // Id of the commentary.
}

// A struct representing a fraction of a live game.
type LiveGameUnit struct {
	possession   int           // Tells us which of the teams is in possession of the ball.
	area         int           // The area of the pitch the ball is currently in.
	minute, time int           // Which minute of the game and which part of the game. If 'minute' is -1 we have an event like a substitution that doesn't count as a  match time consuming event. @see #GameUnitTime
	result       [2]int        // The match result at the time of this unit.
	event        LiveGameEvent // The event belonging to the game unit.
}

// A structure storing team settings during a live game pause (so that we know what users have changed in pauses.
type LiveGameTeamState struct {
	structure, style int
	boost            bool
	player_ids       [11]int
}

type LiveGame struct {
	fix           *Fixture                              // The fixture that belongs to the game.
	fixId         int                                   // Integer determining the fixture (needed because fixture pointers can change when new fixtures are added to the fixtures array).
	teamNames     [2]string                             // The names of the teams stored for later use (when the fixture already got freed).
	attendance    int                                   // Attendance stored for later use (when the fixture already got freed).
	subLeft       [2]int                                // Substitutions left for the teams.
	startedGame   int                                   // The team that started the game, 0 or 1.
	stadiumEvent  int                                   // We keep track of the stadium events because there shouldn't be more than one of them in a game.
	teamValues    [2][GAME_TEAM_VALUE_END]GameTeamValue // Attacking, defending etc. values for both teams. @see #GameTeamValue
	homeAdvantage float64                               // The home advantage factor.
	units         []*LiveGameUnit                       // The array of units. @see #GameUnit
	stats         LiveGameStats                         // Match statistics. @see #LiveGameStats
	team_state    [2]LiveGameTeamState
	actionIds     [2]string // Ids of strategy actions applied. Actions only get applied once.
}
