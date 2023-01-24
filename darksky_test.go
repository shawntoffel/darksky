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

func TestParseForecastOffsetFloat(t *testing.T) {
	ds := New("api key")
	server := getMockServerWithFileData("forecast_mumbai.json")

	defer server.Close()

	BaseUrl = server.URL

	forecast, err := ds.Forecast(ForecastRequest{})

	if err != nil {
		t.Error(err.Error())
	}

	want := float64(5.5)
	have := float64(forecast.Offset)
	if want != have {
		t.Errorf("want %f, have %f", want, have)
	}
}

func TestParseForecastOffsetInteger(t *testing.T) {
	ds := New("api key")
	server := getMockServerWithFileData("full.json")

	defer server.Close()

	BaseUrl = server.URL

	forecast, err := ds.Forecast(ForecastRequest{})

	if err != nil {
		t.Error(err.Error())
	}

	want := float64(-5)
	have := float64(forecast.Offset)
	if want != have {
		t.Errorf("want %f, have %f", want, have)
	}
}

func TestParseFloatUvIndex(t *testing.T) {
	ds := New("api key")
	server := getMockServerWithFileData("pirate_weather_full.json")

	defer server.Close()

	BaseUrl = server.URL

	forecast, err := ds.Forecast(ForecastRequest{})

	if err != nil {
		t.Error(err.Error())
	}

	want := int64(1)
	have := forecast.Currently.UvIndex
	if want != have {
		t.Errorf("want %d, have %d", want, have)
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
		Latitude:  40.7128,
		Longitude: -74.0059,
	}

	BaseUrl = "http://localhost"

	want := "http://localhost/apikey/40.7128,-74.0059"
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
		Latitude:  40.7128,
		Longitude: -74.0059,
		Time:      1547889618,
	}

	BaseUrl = "http://localhost"

	want := "http://localhost/apikey/40.7128,-74.0059,1547889618"
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
		Latitude:  40.7128,
		Longitude: -74.0059,
		Options: ForecastRequestOptions{
			Exclude: "hourly,minutely",
		},
	}

	BaseUrl = "http://localhost"

	want := "http://localhost/apikey/40.7128,-74.0059?exclude=hourly%2Cminutely"
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
		Latitude:  40.7128,
		Longitude: -74.0059,
		Time:      1547889618,
		Options: ForecastRequestOptions{
			Exclude: "minutely,daily",
			Extend:  "hourly",
			Lang:    "en",
			Units:   "si",
		},
	}

	BaseUrl = "http://localhost"

	want := "http://localhost/apikey/40.7128,-74.0059,1547889618?exclude=minutely%2Cdaily&extend=hourly&lang=en&units=si"
	have := ds.buildRequestUrl(r)

	if want != have {
		t.Errorf("want %s, have %s", want, have)
	}
}

func TestUrlConstructionMaintainsFullFloat64LatLongPrecision(t *testing.T) {
	ds := darkSky{
		ApiKey: "apikey",
	}

	r := ForecastRequest{
		Latitude:  40.71281234567891,
		Longitude: -74.0059123456123,
		Time:      1547889618,
		Options: ForecastRequestOptions{
			Exclude: "minutely,daily",
			Extend:  "hourly",
			Lang:    "en",
			Units:   "si",
		},
	}

	BaseUrl = "http://localhost"

	want := "http://localhost/apikey/40.71281234567891,-74.0059123456123,1547889618?exclude=minutely%2Cdaily&extend=hourly&lang=en&units=si"
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
