package models

type RegionType int

const (
	RegionTypeWorld RegionType = iota
	RegionTypeContinent
	RegionTypeCountry
)

type Region struct {
	id           int
	sid          string
	name         string
	symbol       string
	rating       int
	childRegions []*Region
}
