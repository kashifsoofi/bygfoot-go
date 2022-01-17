package models

// Enumeration describing the type of a job.
type JobType int

const (
	// Job offer by a team from the country the user's playing.
	JOB_TYPE_NATIONAL JobType = iota
	// Job offer by a team from a different country than user's playing.
	JOB_TYPE_INTERNATIONAL
	// Job offer by a team from a different country than user's playing; the team participates in an international cup.
	JOB_TYPE_CUP
	JOB_TYPE_END
)

// A structure representing a job in the job exchange.
type Job struct {
	Type                                 JobType // Whether it's an international job or a national one.
	time                                 int     // How many weeks remaining the offer is on the list.
	countryFile, countryName, leagueName string  // Only relevant for international jobs.
	leagueLayer                          int     // Only relevant for international jobs.
	countryRating                        int     // Only relevant for international jobs.
	talentPercent                        int     // Average talent compared to the league average in percent.
	teamId                               int     // The id of the team the job describes.
}
