package sql

import (
	"fmt"

	"github.com/kashifsoofi/bygfoot-go/models"
	"github.com/kashifsoofi/bygfoot-go/store"
)

type SqlRegionStore struct {
	*SqlStore
}

func NewSqlRegionStore(sqlStore *SqlStore) store.RegionStore {
	return &SqlRegionStore{sqlStore}
}

func (s SqlRegionStore) GetCountries() ([]*models.Region, error) {
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
