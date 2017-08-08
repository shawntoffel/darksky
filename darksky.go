package darksky

import (
	"fmt"
	"github.com/google/go-querystring/query"
)

const (
	baseUrl = "https://api.darksky.net/forecast"
)

type DarkSky struct {
	apiKey     string
	BaseUrl    string
	RestClient *RestClient
}

func New(apiKey string) *DarkSky {
	restClient := NewRestClient()
	return &DarkSky{apiKey, baseUrl, restClient}
}

func (d *DarkSky) Forecast(request ForecastRequest) (ForecastResponse, error) {
	response := ForecastResponse{}

	values, _ := query.Values(request.Options)

	queryString := values.Encode()

	url := fmt.Sprintf("%s/%s/%f,%f", d.BaseUrl, d.apiKey, request.Latitude, request.Longitude)

	if len(queryString) > 0 {
		url = url + "?" + queryString
	}

	err := d.RestClient.Get(url, nil, &response)

	return response, err
}
