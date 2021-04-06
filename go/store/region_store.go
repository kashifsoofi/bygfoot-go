package store

import (
	"fmt"

	"github.com/kashifsoofi/bygfoot-go/models"
)

type SqlRegionStore struct {
	*SqlStore
}

func NewSqlRegionStore(sqlStore *SqlStore) RegionStore {
	return &SqlRegionStore{sqlStore}
}

func (s SqlRegionStore) ListCountries() ([]*models.Region, error) {
	var regions []*models.Region
	if err := s.dbx.Select(&regions,
		`SELECT
			*
		FROM Regions
		WHERE RegionType = 2`); err != nil {
		return nil, fmt.Errorf("could not list countries %w", err)
	}
	return regions, nil
}
