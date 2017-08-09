package darksky

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseForecastResponse(t *testing.T) {
	ds := New("api key")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bytes, _ := ioutil.ReadFile("testdata/newyork.json")
		fmt.Fprintln(w, string(bytes))
	}))

	defer server.Close()

	ds.SetBaseUrl(server.URL)
	forecast, err := ds.Forecast(ForecastRequest{})

	if err != nil {
		t.Error(err.Error())
	}

	actual := forecast.Timezone
	expected := "America/New_York"

	assert.Equal(t, actual, expected, "Timezone should match.")
}
