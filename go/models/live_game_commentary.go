package models

// Structure describing a commentary for the live game.
type LiveGameCommentary struct {
	text      string // The commentary text (possibly containing tokens).
	condition string // A condition (if not fulfilled, the commentary doesn't get shown).
	priority  int    // Priority of the commentary (compared to the other ones for the same event). The higher the priority the higher the probability that the commentary gets picked.
	id        int    // An id to keep track of already used commentaries in the live game (so as not to use the same one too frequently).
}
