<h1 align="center">clownchu/riotapi - RIOT API client</h1>
<h3 align="center">RIOT API client implementation written in GO</h3>

<hr>

<h2 align="center" id="requirements">Requirements</h2>

* [GO 1.18+](https://go.dev/doc/install)
* RIOT account and API key - [Get here](https://developer.riotgames.com/)

<hr>

<h2 id="get-in-go" align="center">Get in go</h2>

```bash
go get github.com/ClownChu/riotapi
```

<hr>

<h2 id="making-api-calls" align="center">Making API calls</h2>

```go
apiKey := `RGAPI-f90d223b-9f91-48d0-b9e2-2ff146c6b7e4`
testClient, _ := NewRiotApiClient(apiKey, `NA1`)
summonerApi := summoner.NewSummonerApi(testClient)
summonerData, _ := summonerApi.ByName("ClownChu")
```

<hr>

<h2 id="apis-implementated" align="center">APIs implementated</h2>

* SUMMONER-V4 - [Get summoners information][summoner-api](https://developer.riotgames.com/apis#summoner-v4/)
    * [summoners/by-name](https://developer.riotgames.com/apis#summoner-v4/GET_getBySummonerName)

<h2 id="api-keys-limitations" align="center">API keys limitations</h2>

Riot Games has restrictions for API usage. [read more here](https://developer.riotgames.com/docs/portal#web-apis_personal-api-keys)

<h2 id="testing" align="center">Testing</h2>

You will need to set the environment variable `RIOTAPI_KEY` to you Riot Developer API Key in order to run tests
with `go test ./test/...`. You can obtain your API Key from Riot at [Riot's Developer Website](https://developer.riotgames.com/).

As mentioned above, the tests specify a rate limiter using the developer limits API rate limits.

<hr>

<h2 align="center" id="license">License</h2>
<div align="center">
    <a href="https://github.com/ClownChu/riotapi" target="_blank">clownchu/riotapi source code</a> is made available under the <a href="https://www.gnu.org/licenses/agpl-3.0.en.html" target="_blank">GNU Affero General Public License v3.0</a> license. (<a href="https://choosealicense.com/licenses/agpl-3.0/" target="_blank">Read more</a>)
</div>