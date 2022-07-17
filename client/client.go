package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"riotapi/api/riot"
	"time"

	"github.com/throttled/throttled/store/memstore"
	"github.com/throttled/throttled/v2"
)

var (
	ErrAPIKeyRequired    = errors.New("required APIKey is missing")
	ErrClientNotAttached = errors.New("required client is missing")
	ErrUnknownRegion     = errors.New("unknown region")
	ErrUnknownRiotApi    = errors.New("unknown api")
)

const riotApiScheme = "https"
const riotApiBaseUrl = "api.riotgames.com"

type RiotApiClient struct {
	httpClient  *http.Client
	apiKey      string
	region      *riot.RiotRegion
	rateLimiter *throttled.GCRARateLimiter
}

type IRiotApiService interface {
	*RiotApiClient
	AttachClient(client *RiotApiClient)
	riot.RiotApi
}

func GetTestRiotApiClient() *RiotApiClient {
	apiKey := os.Getenv("RIOTAPI_KEY")

	testClient, err := NewRiotApiClient(apiKey, riot.RiotRegions[0].Id)
	if err != nil {
		println("Unable to initate RIOT client, err: ", err.Error())
		os.Exit(1)
	}
	return testClient
}

// NewClient initializes and returns a riotapi Client struct
func NewRiotApiClient(apiKey string, regionId string) (*RiotApiClient, error) {
	if len(apiKey) == 0 {
		return nil, ErrAPIKeyRequired
	}

	// Data store. Since we only vary keys by regionId, use that to calculate size
	var rateLimiter *throttled.GCRARateLimiter
	store, err := memstore.New(1024)
	if err == nil {
		// Set up rate limiting to 50/minute to fit the 500/10 minutes developer
		// limit
		quota := throttled.RateQuota{throttled.PerMin(50), 0}
		rateLimiter, _ = throttled.NewGCRARateLimiter(store, quota)
	}

	region := riot.GetRiotRegion(regionId)
	if region == nil {
		return nil, ErrUnknownRegion
	}

	c := &RiotApiClient{
		httpClient:  new(http.Client),
		apiKey:      apiKey,
		region:      region,
		rateLimiter: rateLimiter,
	}

	return c, nil
}

func (c *RiotApiClient) GetRegion() riot.RiotRegion {
	return *c.region
}

func (c *RiotApiClient) Get(path string, decoded interface{}) (*http.Response, error) {
	requestUrl := url.URL{
		Scheme:   riotApiScheme,
		Host:     c.region.Id + "." + riotApiBaseUrl,
		RawQuery: fmt.Sprintf("api_key=%s", url.QueryEscape(c.apiKey)),
		Path:     path,
	}

	// Check rate limiting and wait if necessary
	if c.rateLimiter != nil {
		// rate limits are per region, so use the region in the key
		for {
			limited, context, err := c.rateLimiter.RateLimit("riotapi-"+c.region.Id, 1)
			if err != nil {
				return nil, err
			}
			if limited {
				time.Sleep(context.RetryAfter)
			} else {
				break
			}
		}
	}

	req, err := http.NewRequest("GET", requestUrl.String(), nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotModified {
		errBody := new(bytes.Buffer)
		resp.Write(errBody)
	} else {
		if decoded != nil {
			err = json.NewDecoder(resp.Body).Decode(decoded)
		}
	}

	return resp, err
}
