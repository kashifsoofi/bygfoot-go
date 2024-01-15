package domain

// A struct representing a betting possibility: a fixture plus odds.
type BetMatch struct {
	fixId int        // The match the bet is about.
	odds  [3]float64 // The odds for a win, draw or loss of the first team.
}

// A struct representing a bet by a user.
type BetUser struct {
	fixId   int // Match the user betted on.
	outcome int // Which outcome he picked.
	wager   int // How much he wagered.
}
