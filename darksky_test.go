package darksky

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseForecastFullResponse(t *testing.T) {
	ds := New("api key")

	server := getMockServerWithFileData("full.json")

	defer server.Close()

	BaseUrl = server.URL

	forecast, err := ds.Forecast(ForecastRequest{})

	if err != nil {
		t.Error(err.Error())
	}

	want := "America/New_York"
	have := forecast.Timezone

	if want != have {
		t.Errorf("want %s, have %s", want, have)
	}
}

func TestParseForecastExcludedResponse(t *testing.T) {
	ds := New("api key")

	server := getMockServerWithFileData("allexcluded.json")

	defer server.Close()

	BaseUrl = server.URL

	forecast, err := ds.Forecast(ForecastRequest{})

	if err != nil {
		t.Error(err.Error())
	}

	want := "America/New_York"
	have := forecast.Timezone

	if want != have {
		t.Errorf("want %s, have %s", want, have)
	}
}

func TestUrlConstructionLatitudeLongitude(t *testing.T) {
	ds := darkSky{
		ApiKey: "apikey",
	}

	r := ForecastRequest{
		Latitude:  40.712800,
		Longitude: -74.005900,
	}

	BaseUrl = "http://localhost"

	want := "http://localhost/apikey/40.712800,-74.005900"
	have := ds.buildRequestUrl(r)

	if want != have {
		t.Errorf("want %s, have %s", want, have)
	}
}

func TestUrlConstructionTimeMachine(t *testing.T) {
	ds := darkSky{
		ApiKey: "apikey",
	}

	r := ForecastRequest{
		Latitude:  40.712800,
		Longitude: -74.005900,
		Time:      1547889618,
	}

	BaseUrl = "http://localhost"

	want := "http://localhost/apikey/40.712800,-74.005900,1547889618"
	have := ds.buildRequestUrl(r)

	if want != have {
		t.Errorf("want %s, have %s", want, have)
	}
}

func TestUrlConstructionPartialForecastRequestOptions(t *testing.T) {
	ds := darkSky{
		ApiKey: "apikey",
	}

	r := ForecastRequest{
		Latitude:  40.712800,
		Longitude: -74.005900,
		Options: ForecastRequestOptions{
			Exclude: "hourly,minutely",
		},
	}

	BaseUrl = "http://localhost"

	want := "http://localhost/apikey/40.712800,-74.005900?exclude=hourly%2Cminutely"
	have := ds.buildRequestUrl(r)

	if want != have {
		t.Errorf("want %s, have %s", want, have)
	}
}

func TestUrlConstructionFull(t *testing.T) {
	ds := darkSky{
		ApiKey: "apikey",
	}

	r := ForecastRequest{
		Latitude:  40.712800,
		Longitude: -74.005900,
		Time:      1547889618,
		Options: ForecastRequestOptions{
			Exclude: "minutely,daily",
			Extend:  "hourly",
			Lang:    "en",
			Units:   "si",
		},
	}

	BaseUrl = "http://localhost"

	want := "http://localhost/apikey/40.712800,-74.005900,1547889618?exclude=minutely%2Cdaily&extend=hourly&lang=en&units=si"
	have := ds.buildRequestUrl(r)

	if want != have {
		t.Errorf("want %s, have %s", want, have)
	}
}

func getMockServerWithFileData(filename string) *httptest.Server {
	bytes, _ := ioutil.ReadFile("testdata/" + filename)

	return getMockServer(string(bytes))
}

func getMockServer(data string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, data)
	}))

	return server
}
