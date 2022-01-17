package sql

import (
	"fmt"

	"github.com/kashifsoofi/bygfoot-go/models"
	"github.com/kashifsoofi/bygfoot-go/store"
)

type SqlLeagueStore struct {
	*SqlStore
}

func NewSqlLeagueStore(sqlStore *SqlStore) store.LeagueStore {
	return &SqlLeagueStore{sqlStore}
}

func (s SqlLeagueStore) GetLeaguesByCountryId(countryId int) ([]*models.League, error) {
	var regions []*models.League
	if err := s.dbx.Select(&regions,
		`SELECT
			*
		FROM Leagues
		WHERE RegionId = $1`, countryId); err != nil {
		return nil, fmt.Errorf("could not get leagues %w", err)
	}
	return regions, nil
}
