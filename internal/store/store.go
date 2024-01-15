package store

import (
	"github.com/kashifsoofi/bygfoot-go/internal/domain"
)

type Store interface {
	Close()
	Region() RegionStore
	League() LeagueStore
}

type RegionStore interface {
	GetCountries() ([]*domain.Region, error)
}

type LeagueStore interface {
	GetLeaguesByCountryId(countryId int) ([]*domain.League, error)
}
