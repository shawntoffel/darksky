package darksky

import (
	"fmt"
	"github.com/google/go-querystring/query"
)

var baseUrl = "https://api.darksky.net/forecast"

type DarkSky interface {
	Forecast(request ForecastRequest) (ForecastResponse, error)
}

type darkSky struct {
	ApiKey string
}

func New(apiKey string) DarkSky {
	return &darkSky{apiKey}
}

func (d *darkSky) Forecast(request ForecastRequest) (ForecastResponse, error) {
	response := ForecastResponse{}

	url := buildRequestUrl(baseUrl, request)

	err := get(url, nil, &response)

	return response, err
}

func buildRequestUrl(baseUrl string, request ForecastRequest) string {
	url := fmt.Sprintf("%s/%s/%f,%f", baseUrl, request.ApiKey, request.Latitude, request.Longitude)

	if request.Time != nil {
		url = url + fmt.Sprintf(",%d", request.Time)
	}

	values, _ := query.Values(request.Options)
	queryString := values.Encode()

	if len(queryString) > 0 {
		url = url + "?" + queryString
	}

	return url
}
