package models

// Lineup types for a CPU team (ie. which players are preferred when putting together the first 11).
type StratLineupType int

const (
	STRAT_LINEUP_BEST StratLineupType = iota
	STRAT_LINEUP_WEAKEST
	STRAT_LINEUP_FITTEST
	STRAT_LINEUP_UNFITTEST
	STRAT_LINEUP_END
)

// A struct describing the pre-match strategy settings of a CPU team.
type StrategyPrematch struct {
	condition            string   // A condition describing when the strategy should be applied.
	formations           []string // Array of possible formations, sorted by preference.
	boost, style, lineup int      // Boost, style values and lineup type.
	minFitness           float64  // The fitness value below which a player gets substituted if there's a sub with better fitness.
}

type StrategyMatchAction struct {
	condition, subCondition                    string // A condition describing when the action should be taken.
	boost, style                               int    // New boost and style values.
	subInPos, subInProp, subOutPos, subOutProp int    // Substitution specifiers (position and property). Property is taken from #StratLineupType.
	id                                         int    // An id to prevent actions from being applied again and again during a match.
}

// A CPU strategy.
type Strategy struct {
	sid, desc   string                 // String id and description of the strategy.
	priority    int                    // How often this strategy gets picked, relative to the other strategies.
	prematch    []*StrategyPrematch    // Array with prematch settings.
	matchAction []*StrategyMatchAction // Array with match settings.
}
