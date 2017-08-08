package darksky

import (
	"fmt"
	"github.com/google/go-querystring/query"
)

const (
	BaseUrl = "https://api.darksky.net/forecast"
)

type DarkSky interface {
	Forecast(request ForecastRequest) (Forecast, error)
}

type darkSky struct {
	ApiKey     string
	RestClient RestClient
}

func New(apiKey string) DarkSky {
	restClient := NewRestClient()
	return &darkSky{apiKey, restClient}
}

func NewWithRestClient(apiKey string, restClient RestClient) DarkSky {
	return &darkSky{apiKey, restClient}
}

func (d *darkSky) Forecast(request ForecastRequest) (Forecast, error) {
	response := Forecast{}

	values, _ := query.Values(request.Options)

	queryString := values.Encode()

	url := fmt.Sprintf("%s/%s/%f,%f", BaseUrl, d.ApiKey, request.Latitude, request.Longitude)

	if len(queryString) > 0 {
		url = url + "?" + queryString
	}

	fmt.Println(url)

	err := d.RestClient.Get(url, nil, &response)

	return response, err
}
