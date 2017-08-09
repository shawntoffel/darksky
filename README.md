# DarkSky


[![GoDoc](https://godoc.org/github.com/shawntoffel/darksky?status.svg)](https://godoc.org/github.com/shawntoffel/darksky) [![Go Report Card](https://goreportcard.com/badge/github.com/shawntoffel/darksky)](https://goreportcard.com/report/github.com/shawntoffel/darksky) [![Build Status](https://travis-ci.org/shawntoffel/darksky.svg?branch=master)](https://travis-ci.org/shawntoffel/darksky)

Dark Sky API client https://darksky.net/dev/docs

### Installing

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

`go get` *will always pull the latest released version from the master branch.*

```sh
go get github.com/shawntoffel/darksky
```

If you want to use the dev branch, follow these steps next.

```sh
cd $GOPATH/src/github.com/shawntoffel/darksky
git checkout dev
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
request := darksky.ForecastRequest{}
request.Latitude = 40.7128
request.Longitude = -74.0059
```

Get the forecast
```go
forecast, err := client.Forecast(request)
```
