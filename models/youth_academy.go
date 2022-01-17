package models

// The data about a user's youth academy is bundled here.
type YouthAcademy struct {
	team               *Team   // Pointer to the team the academy belongs to.
	positionPreference int     // Position preference of the user.
	coach              int     // Coach quality, like scout or physio.
	percentage         int     // Percentage of income the user devotes to the youth academy.
	averageCoach       float64 // Average coach and percentage values; these are used when a new youth comes into the academy to determine the quality of the youth.
	averagePercentage  float64
	counterYouth       float64   // When this counter is <= 0, a new youth appears.
	players            []*Player // The youths are simply young players.
}
