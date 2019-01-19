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
