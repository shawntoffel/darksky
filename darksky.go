package darksky

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// DarkSky API endpoint
var (
	BaseUrl = "https://api.darksky.net/forecast"
)

// DarkSky Api client
type DarkSky interface {
	Forecast(request ForecastRequest) (ForecastResponse, error)
}

type darkSky struct {
	ApiKey string
	Client *http.Client
}

// New creates a new DarkSky client
func New(apiKey string) DarkSky {
	return &darkSky{apiKey, &http.Client{}}
}

// NewWithClient creates a new DarkSky client with custom http.Client
func NewWithClient(apiKey string, client *http.Client) DarkSky {
	return &darkSky{apiKey, client}
}

// Forecast request a forecast
func (d *darkSky) Forecast(request ForecastRequest) (ForecastResponse, error) {
	response := ForecastResponse{}

	url := d.buildRequestUrl(request)

	err := get(d.Client, url, &response)

	return response, err
}

func (d *darkSky) buildRequestUrl(request ForecastRequest) string {
	url := fmt.Sprintf("%s/%s/%f,%f", BaseUrl, d.ApiKey, request.Latitude, request.Longitude)

	if request.Time > 0 {
		url = url + fmt.Sprintf(",%d", request.Time)
	}

	values, _ := query.Values(request.Options)
	queryString := values.Encode()

	if len(queryString) > 0 {
		url = url + "?" + queryString
	}

	return url
}
