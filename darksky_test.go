package darksky

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
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

	actual := forecast.Timezone
	expected := "America/New_York"

	assert.Equal(t, actual, expected, "Timezone should match.")
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

	actual := forecast.Timezone
	expected := "America/New_York"

	assert.Equal(t, actual, expected, "Timezone should match.")
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
