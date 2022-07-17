package riot

import (
	"strconv"
	"strings"
)

// Api struct describes a Riot Api
type RiotApi struct {
	Collection string // e.g. "lol"
	Name       string // e.g. "league"
	Version    int    // e.g. "4"
}

var RiotApis = []RiotApi{
	{"riot", "account", 1},
	{"lol", "champion-mastery", 4},
	{"lol", "platform", 3},
	{"lol", "clash", 1},
	{"lol", "league-exp", 4},
	{"lol", "league", 4},
	{"lol", "challenges", 1},
	{"lol", "status", 4},
	{"lol", "match", 5},
	{"lol", "spectator", 4},
	{"lol", "summoner", 4},
	{"lol", "tournament-stub", 4},
	{"lol", "tournament", 4},
	{"lor", "deck", 1},
	{"lor", "inventory", 1},
	{"lor", "match", 1},
	{"lor", "ranked", 1},
	{"lor", "status", 1},
	{"tft", "league", 1},
	{"tft", "match", 1},
	{"tft", "summoner", 1},
	{"val", "content", 1},
	{"val", "match", 1},
	{"val", "ranked", 1},
	{"val", "status", 1},
}

func (a *RiotApi) GetId() string {
	return a.Collection + "_" + a.Name
}

func (a *RiotApi) GetBasePath() string {
	path := []string{
		a.Collection,
		a.Name,
		"v" + strconv.Itoa(a.Version),
	}
	return strings.Join(path, `/`)
}

func (a *RiotApi) GetScopes() []RiotApiScope {
	var scopes []RiotApiScope
	for _, r := range RiotApiScopes {
		if a.GetId() == r.RiotApiId {
			scopes = append(scopes, r)
		}
	}
	return scopes
}

func (a *RiotApi) GetScope(scopeId string) *RiotApiScope {
	scopes := a.GetScopes()
	for _, r := range scopes {
		if r.GetId() == scopeId {
			return &r
		}
	}
	return nil
}

// GetApi returns the Api description for a given Api id (case-insensitive)
func GetRiotApi(collection string, name string) *RiotApi {
	for _, r := range RiotApis {
		if strings.ToLower(collection) == r.Collection && strings.ToLower(name) == r.Name {
			return &r
		}
	}
	return nil
}
