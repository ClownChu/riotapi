package test

import (
	"os"
	"riotapi/client"
	"testing"
)

var testClient *client.RiotApiClient

func TestMain(m *testing.M) {
	// Make sure client connection is possible
	// before running anything
	testClient = client.GetTestRiotApiClient()
	os.Exit(m.Run())
}
