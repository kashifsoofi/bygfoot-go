package domain

// A list of first names and last names from a file.
type NameList struct {
	sid                   string   // The file id (the part between 'player_names_' and '.xml').
	firstNames, lastNames []string // Arrays of strings holding the names.
}
