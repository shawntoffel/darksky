package darksky

import (
	"fmt"
	"github.com/google/go-querystring/query"
)

const (
	baseUrl = "https://api.darksky.net/forecast"
)

type DarkSky struct {
	BaseUrl string
}

func New() *DarkSky {
	return &DarkSky{baseUrl}
}

func (d *DarkSky) Forecast(request ForecastRequest) (ForecastResponse, error) {
	response := ForecastResponse{}

	url := d.buildRequestUrl(request)

	err := get(url, nil, &response)

	return response, err
}

func (d *DarkSky) buildRequestUrl(request ForecastRequest) string {
	url := fmt.Sprintf("%s/%s/%f,%f", d.BaseUrl, request.ApiKey, request.Latitude, request.Longitude)

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
