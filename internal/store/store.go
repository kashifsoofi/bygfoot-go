package store

import (
	"github.com/kashifsoofi/bygfoot-go/internal/domain"
)

type Store interface {
	Close()
	Hints() HintsStore
	League() LeagueStore
	Region() RegionStore
}

type HintsStore interface {
	LoadHintNumber() int
	SaveHintNumber(int)
	Hint(int) string
	TotalHints() int
}

type RegionStore interface {
	GetCountries() ([]*domain.Region, error)
}

type LeagueStore interface {
	GetLeaguesByCountryId(countryId int) ([]*domain.League, error)
}
