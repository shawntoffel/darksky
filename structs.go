package darksky

import "net/url"

// Timestamp is an int64 timestamp
type Timestamp int64

// Measurement is a float64 measurement
type Measurement float64

// ForecastRequest contains all available options for requesting a forecast
type ForecastRequest struct {
	Latitude  Measurement
	Longitude Measurement
	Time      Timestamp
	Options   ForecastRequestOptions
}

// ForecastRequestOptions are optional and passed as query parameters
type ForecastRequestOptions struct {
	Exclude string
	Extend  string
	Lang    string
	Units   string
}

// Encode into URL encoded query string parameters (exclude=hourly&units=si)
func (o ForecastRequestOptions) Encode() string {
	q := url.Values{}

	if o.Exclude != "" {
		q.Add("exclude", o.Exclude)
	}
	if o.Extend != "" {
		q.Add("extend", o.Extend)
	}
	if o.Lang != "" {
		q.Add("lang", o.Lang)
	}
	if o.Units != "" {
		q.Add("units", o.Units)
	}

	return q.Encode()
}

// ForecastResponse is the response containing all requested properties
type ForecastResponse struct {
	Latitude  Measurement `json:"latitude,omitempty"`
	Longitude Measurement `json:"longitude,omitempty"`
	Timezone  string      `json:"timezone,omitempty"`
	Currently *DataPoint  `json:"currently,omitempty"`
	Minutely  *DataBlock  `json:"minutely,omitempty"`
	Hourly    *DataBlock  `json:"hourly,omitempty"`
	Daily     *DataBlock  `json:"daily,omitempty"`
	Alerts    []*Alert    `json:"alerts,omitempty"`
	Flags     *Flags      `json:"flags,omitempty"`
}

// DataPoint contains various properties, each representing the average (unless otherwise specified) of a particular weather phenomenon occurring during a period of time.
type DataPoint struct {
	ApparentTemperature         Measurement `json:"apparentTemperature,omitempty"`
	ApparentTemperatureHigh     Measurement `json:"apparentTemperatureHigh,omitempty"`
	ApparentTemperatureHighTime Timestamp   `json:"apparentTemperatureHighTime,omitempty"`
	ApparentTemperatureLow      Measurement `json:"apparentTemperatureLow,omitempty"`
	ApparentTemperatureLowTime  Timestamp   `json:"apparentTemperatureLowTime,omitempty"`
	CloudCover                  Measurement `json:"cloudCover,omitempty"`
	DewPoint                    Measurement `json:"dewPoint,omitempty"`
	Humidity                    Measurement `json:"humidity,omitempty"`
	Icon                        string      `json:"icon,omitempty"`
	MoonPhase                   Measurement `json:"moonPhase,omitempty"`
	NearestStormBearing         Measurement `json:"nearestStormBearing,omitempty"`
	NearestStormDistance        Measurement `json:"nearestStormDistance,omitempty"`
	Ozone                       Measurement `json:"ozone,omitempty"`
	PrecipAccumulation          Measurement `json:"precipAccumulation,omitempty"`
	PrecipIntensity             Measurement `json:"precipIntensity,omitempty"`
	PrecipIntensityError        Measurement `json:"precipIntensityError,omitempty"`
	PrecipIntensityMax          Measurement `json:"precipIntensityMax,omitempty"`
	PrecipIntensityMaxTime      Timestamp   `json:"precipIntensityMaxTime,omitempty"`
	PrecipProbability           Measurement `json:"precipProbability,omitempty"`
	PrecipType                  string      `json:"precipType,omitempty"`
	Pressure                    Measurement `json:"pressure,omitempty"`
	Summary                     string      `json:"summary,omitempty"`
	SunriseTime                 Timestamp   `json:"sunriseTime,omitempty"`
	SunsetTime                  Timestamp   `json:"sunsetTime,omitempty"`
	Temperature                 Measurement `json:"temperature,omitempty"`
	TemperatureHigh             Measurement `json:"temperatureHigh,omitempty"`
	TemperatureHighTime         Timestamp   `json:"temperatureHighTime,omitempty"`
	TemperatureLow              Measurement `json:"temperatureLow,omitempty"`
	TemperatureLowTime          Timestamp   `json:"temperatureLowTime,omitempty"`
	Time                        Timestamp   `json:"time,omitempty"`
	UvIndex                     int64       `json:"uvIndex,omitempty"`
	UvIndexTime                 Timestamp   `json:"uvIndexTime,omitempty"`
	Visibility                  Measurement `json:"visibility,omitempty"`
	WindBearing                 Measurement `json:"windBearing,omitempty"`
	WindGust                    Measurement `json:"windGust,omitempty"`
	WindGustTime                Timestamp   `json:"windGustTime,omitempty"`
	WindSpeed                   Measurement `json:"windSpeed,omitempty"`
}

// DataBlock represents the various weather phenomena occurring over a period of time
type DataBlock struct {
	Data    []DataPoint `json:"data,omitempty"`
	Summary string      `json:"summary,omitempty"`
	Icon    string      `json:"icon,omitempty"`
}

// Alert contains objects representing the severe weather warnings issued for the requested location by a governmental authority
type Alert struct {
	Description string    `json:"description,omitempty"`
	Expires     Timestamp `json:"expires,omitempty"`
	Regions     []string  `json:"regions,omitempty"`
	Severity    string    `json:"severity,omitempty"`
	Time        Timestamp `json:"time,omitempty"`
	Title       string    `json:"title,omitempty"`
	Uri         string    `json:"uri,omitempty"`
}

// Flags contains various metadata information related to the request
type Flags struct {
	DarkSkyUnavailable string      `json:"darksky-unavailable,omitempty"`
	NearestStation     Measurement `json:"nearest-station"`
	Sources            []string    `json:"sources,omitempty"`
	Units              string      `json:"units,omitempty"`
}
