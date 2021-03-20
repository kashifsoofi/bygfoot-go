package models

// Structure representing a fixture, or, in other words, a match.
type Fixture struct {
	clid                               int       // The cup or league the fixture belongs to.
	id                                 int       // The unique id of the fixture.
	round                              int       // The round (in a cup) the fixture belongs to.
	replayNumber                       int       // The replay number (ie. how often the match was repeated because of a draw).
	weekNumber, weekRoundNumber        int       // When it takes place.
	teams                              [2]*Team  // The teams involved.
	teamIds                            [2]int    // Ids of the teams. Needed when the team pointers get invalid (e.g. after promotion/relegation).
	result                             [2][3]int // The number of goals for each team in regulation, extra time and penalty shoot-out.
	homeAdvantage, secondLeg, decisive bool      // Whether there's home advantage, this is second leg, or the game has to be decided.
	attendance                         int       // How many people attended and whether there were special events.
}
