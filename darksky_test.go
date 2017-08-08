package darksky

import (
	"fmt"
	"testing"
)

func TestGetForecast(t *testing.T) {
	request := ForecastRequest{}
	request.Latitude = 0
	request.Longitude = 0
	request.Options = ForecastRequestOptions{}
	request.Options.Exclude = "minutely,hourly,daily"

	ds := New("invalid token")

	forecast, err := ds.Forecast(request)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(forecast.Currently.Icon)
}
