package main

import (
	"fmt"

	"github.com/shawntoffel/darksky"
)

// print current temp in new york
func main() {

	client := darksky.New("api key")

	request := darksky.ForecastRequest{
		Latitude:  40.7128,
		Longitude: -74.0059,
		Options: darksky.ForecastRequestOptions{
			Exclude: "hourly,minutely",
		},
	}

	response, err := client.Forecast(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(response.Currently.Temperature)
	fmt.Println(response.Currently.Icon)
}
