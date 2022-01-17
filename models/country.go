package models

type Country struct {
	name    string   // Name of the country.
	symbol  string   // Symbol of the country, eg a flag pixmap.
	sid     string   // Id of the country, eg 'england'.
	rating  int      // A rating point from 0-10 telling us how good the first league of the country is. Spain, for instance, has rating 10, whereas Ireland has only 5.
	leagues []League // Leagues arrays.
	cups    []Cup    // Cups arrays.
	allcups []Cup    // Pointer array holding all cups.
}
