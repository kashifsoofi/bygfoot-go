package store

import "github.com/kashifsoofi/bygfoot-go/models"

type Store interface {
	Close()
	Region() RegionStore
	League() LeagueStore
}

type RegionStore interface {
	GetCountries() ([]*models.Region, error)
}

type LeagueStore interface {
	GetLeaguesByCountryId(countryId int) ([]*models.League, error)
}
