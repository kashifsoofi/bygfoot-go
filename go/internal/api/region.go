package api

import (
	"context"

	"github.com/kashifsoofi/bygfoot-go/pkg/api"
	log "github.com/sirupsen/logrus"
)

// GetBook logs Book from client and returns new Book
func (s *Server) GetRegions(ctx context.Context, input *api.GetRegionsRequest) (*api.GetRegionsResponse, error) {
	log.Info("GetRegions requested by client")

	regions := getRegions()
	return &api.GetRegionsResponse{Regions: regions}, nil
}

func getRegions() []*api.Region {
	return []*api.Region{
		&api.Region{
			Id:     1,
			Sid:    "africa",
			Name:   "Africa",
			Rating: 1,
		},
		&api.Region{
			Id:     2,
			Sid:    "asia",
			Name:   "Asia",
			Rating: 1,
		},
		&api.Region{
			Id:     3,
			Sid:    "australia",
			Name:   "Australia",
			Rating: 1,
		},
		&api.Region{
			Id:     4,
			Sid:    "europe",
			Name:   "Europe",
			Rating: 1,
		},
		&api.Region{
			Id:     5,
			Sid:    "north_america",
			Name:   "North America",
			Rating: 1,
		},
		&api.Region{
			Id:     5,
			Sid:    "south_america",
			Name:   "South America",
			Rating: 1,
		},
	}
}
