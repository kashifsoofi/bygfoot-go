package models

import "github.com/google/uuid"

//  Table element values. @see TableElement @see Table
type TableElementValues int

const (
	TABLE_PLAYED TableElementValues = 0
	TABLE_WON
	TABLE_DRAW
	TABLE_LOST
	TABLE_GF
	TABLE_GA
	TABLE_GD
	TABLE_PTS
	TABLE_END
)

// An element representing a team in the tables. @see Table @see #TableElementValues
type TableElement struct {
	team   *Team
	teamId uuid.UUID
	// The rank of the element before the last update. Used to display an arrow if the rank changed.
	oldRank int
	values  [TABLE_END]int
}

// A table belonging to a league or a cup with round robin. @see TableElement
type Table struct {
	name     string
	clid     int
	round    int // The cup round (or -1 if it's a league).
	elements []*TableElement
}
