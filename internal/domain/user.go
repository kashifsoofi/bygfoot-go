package domain

// Indices for the money_in array.
type MonIn int

const (
	MON_IN_PRIZE MonIn = iota
	MON_IN_TICKET
	MON_IN_SPONSOR
	MON_IN_BETS
	MON_IN_TRANSFERS
	MON_IN_END
)

// Indices for the money_out array.
type MonOut int

const (
	MON_OUT_WAGE MonOut = iota
	MON_OUT_PHYSIO
	MON_OUT_SCOUT
	MON_OUT_YC
	MON_OUT_YA
	MON_OUT_JOURNEY
	MON_OUT_COMPENSATIONS
	MON_OUT_BETS
	MON_OUT_BOOST
	MON_OUT_TRANSFERS
	MON_OUT_STADIUM_IMPROVEMENT
	MON_OUT_STADIUM_BILLS
	MON_OUT_TRAINING_CAMP
	MON_OUT_END
)

// Indices for the counters variable in #User.
type CounterValue int

const (
	COUNT_USER_LOAN             CounterValue = iota // How many weeks until user has to pay back his loan.
	COUNT_USER_OVERDRAWN                            // How often the user overdrew his bank account.
	COUNT_USER_POSITIVE                             // How many weeks until the bank account has to be positive or at least not overdrawn).
	COUNT_USER_SUCCESS                              // How successful the user is.
	COUNT_USER_WARNING                              // Whether there was already a warning about rumours (new coach).
	COUNT_USER_INC_CAP                              // How many weeks until the stadium capacity is increased.
	COUNT_USER_INC_SAF                              // How often the stadium safety was increased (in a week).
	COUNT_USER_STADIUM_CAPACITY                     // Counter for building stadium seats.
	COUNT_USER_STADIUM_SAFETY                       // Counter for increasing stadium safety.
	COUNT_USER_SHOW_RES                             // Whether the latest result is shown when the main window gets refreshed.
	COUNT_USER_TOOK_TURN                            // Whether the user took his turn in a week round.
	COUNT_USER_NEW_SPONSOR                          // A new sponsor offer has to be shown.
	COUNT_USER_END
)

// User-related things that get recorded.
type UserHistoryType int

const (
	USER_HISTORY_START_GAME UserHistoryType = iota
	USER_HISTORY_FIRE_FINANCE
	USER_HISTORY_FIRE_FAILURE
	USER_HISTORY_JOB_OFFER_ACCEPTED
	USER_HISTORY_END_SEASON
	USER_HISTORY_PROMOTED
	USER_HISTORY_RELEGATED
	USER_HISTORY_WIN_FINAL
	USER_HISTORY_LOSE_FINAL
	USER_HISTORY_REACH_CUP_ROUND
	USER_HISTORY_CHAMPION
	USER_HISTORY_END
)

// A memorable match (a recording of a live game) of a user.
type MemMatch struct {
	countryName     string   // Name of the country the user was playing with.
	competitionName string   // The name of the competition, including the cup round name.
	neutral         bool     // Whether the match was on neutral ground. Only relevant for MM list display.
	userTeam        int      // 0 or 1, indicating which team was the user's.
	liveGame        LiveGame // The recording.
}

// A structure holding an element of a user's history, e.g. the event of being fired. */
type UserHistory struct {
	season, week int // When the event happened.

	userHistorytype UserHistoryType // The type (see #UserHistoryType) of the history event.
	teamName        string          // The team of the user at the time.
	info            []string        // These can hold various information like team or league/cup ids.
	//? gchar *string[3];
}

// A user sponsor.
type UserSponsor struct {
	name     string
	benefit  int // Money per week.
	contract int // Contract length in weeks.
}

// A structure representing a human player.
type User struct {
	name          string                       // Username.
	team          *Team                        // The team the user manages.
	teamId        int                          // The team id (needed when the team pointer gets invalid)
	options       OptionList                   // User options.
	events        []*Event                     // Events shown each week.
	history       []*UserHistory               // User history.
	counters      [COUNT_USER_END]CounterValue // User counters (not changeable by the user), like number of weeks until debt has to be paid back.
	money, debt   int                          // The user's money, debt, income and expenses. We have double arrays to store information about the current and the past week.
	moneyIn       [2][MON_IN_END]int
	moneyOut      [2][MON_OUT_END]int
	scout, physio int          // The user's scout and physio qualities. @see #Quality
	liveGame      LiveGame     // The variable for the latest user live game. @see #Game
	sponsor       UserSponsor  // Sponsor of the user.
	youthAcademy  YouthAcademy // Youth academy of the user.
	mmatches_file string       // The currently used MM file.
	mmatches      []*MemMatch  // The array of MMs.
	bets          [2]*BetUser  // Array of current and recent bets.
}

type EventType int

const (
	EVENT_TYPE_WARNING EventType = iota
	EVENT_TYPE_PLAYER_LEFT
	EVENT_TYPE_PAYBACK
	EVENT_TYPE_OVERDRAW
	EVENT_TYPE_JOB_OFFER
	EVENT_TYPE_FIRE_FINANCE
	EVENT_TYPE_FIRE_FAILURE
	EVENT_TYPE_TRANSFER_OFFER_USER
	EVENT_TYPE_TRANSFER_OFFER_CPU
	EVENT_TYPE_TRANSFER_OFFER_REJECTED_BETTER_OFFER
	EVENT_TYPE_TRANSFER_OFFER_REJECTED_FEE_WAGE
	EVENT_TYPE_TRANSFER_OFFER_REJECTED_FEE
	EVENT_TYPE_TRANSFER_OFFER_REJECTED_WAGE
	EVENT_TYPE_TRANSFER_OFFER_MONEY
	EVENT_TYPE_TRANSFER_OFFER_ROSTER
	EVENT_TYPE_PLAYER_CAREER_STOP
	EVENT_TYPE_CHARITY
	EVENT_TYPE_END
)

// A structure representing an event for a user. This is used to show information like a successful transfer or a job offer.
type Event struct {
	user           *User     // Pointer to the user the event belongs to.
	eventType      EventType // Type of the event. See #EventType.
	value1, value2 int       // Some values that are used for different purposes.
	value_pointer  *int      // A pointer for different purposes.
	value_string   string    // A string for different purposes.
}
