package summoner

import (
	"github.com/ClownChu/riotapi/api/riot"
	"github.com/ClownChu/riotapi/client"
)

// Platform struct describes a Riot Platform
type SummonerInfo struct {
	// Summoner ID.
	Id string `json:"id"`

	// Account ID.
	AccountId string `json:"accountId"`

	// Puu ID.
	PuuId string `json:"puuid"`

	// Summoner name.
	Name string `json:"name"`

	// ID of the summoner icon associated with the summoner.
	ProfileIconId int `json:"profileIconId"`

	// Date summoner was last modified specified as epoch milliseconds.
	// The following events will update this timestamp: profile icon change,
	// playing the tutorial or advanced tutorial, finishing a game, summoner name change
	RevisionDate int64 `json:"revisionDate"`

	// Summoner level associated with the summoner.
	SummonerLevel int64 `json:"summonerLevel"`
}

type SummonerApi struct {
	riotApiClient *client.RiotApiClient
	riotApi       *riot.RiotApi
}

func NewSummonerApi(riotApiClient *client.RiotApiClient) *SummonerApi {
	return &SummonerApi{
		riotApiClient: riotApiClient,
		riotApi:       riot.GetRiotApi(`lol`, `summoner`),
	}
}

func (a *SummonerApi) ByName(name string) (*SummonerInfo, error) {
	if a.riotApiClient == nil {
		return nil, client.ErrClientNotAttached
	}

	if a.riotApi == nil {
		return nil, client.ErrUnknownRiotApi
	}

	apiScope := a.riotApi.GetScope("summoners/by-name")
	parameters := []string{
		name,
	}

	path := a.riotApi.GetBasePath() + "/" + apiScope.GetPath(parameters)
	data := new(SummonerInfo)
	if _, err := a.riotApiClient.Get(path, data); err != nil {
		return nil, err
	}
	return data, nil
}
