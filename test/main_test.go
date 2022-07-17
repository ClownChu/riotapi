package test

import (
	"os"
	"testing"

	"github.com/ClownChu/riotapi/client"
)

var testClient *client.RiotApiClient

func TestMain(m *testing.M) {
	// Make sure client connection is possible
	// before running anything
	testClient = client.GetTestRiotApiClient()
	os.Exit(m.Run())
}
