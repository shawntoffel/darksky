package darksky

import (
	"fmt"
	"net/http"
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

	requestUrl := d.buildRequestUrl(request)

	err := get(d.Client, requestUrl, &response)

	return response, err
}

func (d *darkSky) buildRequestUrl(request ForecastRequest) string {
	requestUrl := fmt.Sprintf("%s/%s/%g,%g", BaseUrl, d.ApiKey, request.Latitude, request.Longitude)

	if request.Time > 0 {
		requestUrl = requestUrl + fmt.Sprintf(",%d", request.Time)
	}

	queryString := request.Options.Encode()
	if queryString != "" {
		requestUrl = requestUrl + "?" + queryString
	}

	return requestUrl
}
