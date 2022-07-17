package riot

import "strings"

// Api struct describes a Riot Api
type RiotApiScope struct {
	RiotApiId  string   // e.g. lol_summoner
	Path       []string // e.g. ["challengerleagues", "by-queue"]
	Parameters []string // e.g. ["league"]
}

var RiotApiScopes = []RiotApiScope{
	{"lol_summoner", []string{"summoners", "by-account"}, []string{"encryptedAccountId"}},
	{"lol_summoner", []string{"summoners", "by-name"}, []string{"summonerName"}},
	{"lol_summoner", []string{"summoners", "by-puuid"}, []string{"encryptedPUUID"}},
	{"lol_summoner", []string{"summoners"}, []string{"encryptedSummonerId"}},
	{"lol_summoner", []string{"summoners", "me"}, []string{}},
}

func (a *RiotApiScope) GetId() string {
	return strings.Join(a.Path, `/`)
}

func (a *RiotApiScope) GetPath(parameters []string) string {
	path := strings.Join(a.Path, `/`)
	if len(parameters) > 0 {
		path += "/" + strings.Join(parameters, `/`)
	}
	return path
}
