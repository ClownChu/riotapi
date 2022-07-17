package riot

import (
	"strings"
)

// Region struct describes a Riot Region
type RiotRegion struct {
	Id   string // e.g. "NA1"
	Name string // e.g. "North America"
}

var RiotRegions = []RiotRegion{
	{"NA1", "North America"},
	{"BR1", "Brazil"},
	{"EUN1", "EU Nordic & East "},
	{"EUW1", "EU West"},
	{"JP1", "Japan"},
	{"KR", "Korea"},
	{"LA1", "Latin America North"},
	{"LA2", "Latin America South"},
	{"OC1", "Oceania"},
	{"PBE1", "PBE"},
	{"RU", "Russia"},
	{"TR1", "Turkey"},
}

// GetRegion returns the Region description for a given Region id (case-insensitive)
func GetRiotRegion(id string) *RiotRegion {
	upperId := strings.ToUpper(id)
	for _, r := range RiotRegions {
		if upperId == r.Id {
			return &r
		}
	}
	return nil
}
