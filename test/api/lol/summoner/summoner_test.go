package test

import (
	"riotapi/api/lol/summoner"
	"riotapi/client"
	"testing"
)

var testClient *client.RiotApiClient
var summonerApi *summoner.SummonerApi

func TestSummonerApi(t *testing.T) {
	testClient = client.GetTestRiotApiClient()
	summonerApi = summoner.NewSummonerApi(testClient)
}

func TestSummonerApi_ByName(t *testing.T) {
	data, err := summonerApi.ByName("ClownChu")

	if err != nil {
		t.Errorf("TestSummonerApi_ByName error = %v", err)
		return
	}

	expectedId := "KBmay8aQ2HaS_bZkDqWHRVWj9hn9Aobt3NjdIKpDhojTgKA"
	if data.Id != expectedId {
		t.Errorf("TestSummonerApi_ByName Id = %v, want %v", data.Id, expectedId)
	}
}
