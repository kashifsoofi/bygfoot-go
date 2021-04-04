package store

import "github.com/kashifsoofi/bygfoot-go/models"

type Store interface {
	Close()
	Region() RegionStore
}

type RegionStore interface {
	List() ([]*models.Region, error)
}
