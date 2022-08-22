# DarkSky
[![GoDoc](https://godoc.org/github.com/shawntoffel/darksky?status.svg)](https://godoc.org/github.com/shawntoffel/darksky) [![Go Report Card](https://goreportcard.com/badge/github.com/shawntoffel/darksky)](https://goreportcard.com/report/github.com/shawntoffel/darksky) [![CircleCI](https://circleci.com/gh/shawntoffel/darksky.svg?style=svg)](https://circleci.com/gh/shawntoffel/darksky)

Dark Sky API client in Go https://darksky.net/dev/docs

### A note on Apple Weather
DarkSky's technology has been integrated into Apple Weather. 
I've started a separate project [go-weatherkit](https://github.com/shawntoffel/go-weatherkit) for Go client development against Apple's new [WeatherKit](https://developer.apple.com/weatherkit/) REST API. 

### Installing

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

```sh
go get github.com/shawntoffel/darksky
```

### Usage

Import the package into your project.

```go
import "github.com/shawntoffel/darksky"
```

Construct a new DarkSky client

```go
client := darksky.New("api key")
```

Build a request

```go
request := darksky.ForecastRequest{
    Latitude:  40.7128,
    Longitude: -74.0059,
    Options: darksky.ForecastRequestOptions{
        Exclude: "hourly,minutely",
    },
}
```

Get the forecast
```go
forecast, err := client.Forecast(request)
```

### Configuration

A different API URL may be targeted using the `BaseUrl` package scoped variable
```go
darksky.BaseUrl = "https://api.darksky.net/forecast"
```
