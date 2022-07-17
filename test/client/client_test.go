package test

import (
	"testing"

	"github.com/ClownChu/riotapi/api/riot"
	"github.com/ClownChu/riotapi/client"
)

var testClient *client.RiotApiClient

func TestClient(t *testing.T) {
	testClient = client.GetTestRiotApiClient()
}

func TestClient_GetRegion(t *testing.T) {
	testClientRegion := testClient.GetRegion()
	if testClientRegion.Name != riot.RiotRegions[0].Name {
		t.Errorf("TestClient_GetRegion Expected = %v, got %v", testClientRegion.Name, riot.RiotRegions[0].Name)
	}
}

func TestClient_Get(t *testing.T) {
	var data interface{}
	resp, err := testClient.Get(`/`, data)

	if resp.StatusCode != 401 && err != nil {
		t.Errorf("TestClient_Get expected HTTP 401 error = %v", err)
		return
	}
}
