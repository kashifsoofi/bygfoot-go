package models

type RegionType int

const (
	RegionTypeWorld RegionType = iota
	RegionTypeContinent
	RegionTypeCountry
)

type Region struct {
	Id           int
	Sid          string
	Name         string
	Symbol       string
	Rating       int
	RegionType   RegionType
	ParentId     int
	childRegions []*Region
}
