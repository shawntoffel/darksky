package darksky

type timestamp int64
type measurement float64

type baseRequest struct {
	ApiKey    string
	Latitude  measurement
	Longitude measurement
}

// ForecaseRequest contains all available options for requesting a forecast
type ForecastRequest struct {
	baseRequest
	Time    *timestamp
	Options ForecastRequestOptions
}

// ForecastRequestOptions are optional and passed as query parameters
type ForecastRequestOptions struct {
	Exclude string `url:"exclude,omitempty"`
	Extend  string `url:"extend,omitempty"`
	Lang    string `url:"lang,omitempty"`
	Units   string `url:"units,omitempty"`
}

// ForecastResponse is the response containing all requested properties
type ForecastResponse struct {
	Latitude  measurement `json:"latitude,omitempty"`
	Longitude measurement `json:"longitude,omitempty"`
	Timezone  string      `json:"timezone,omitempty"`
	Currently DataPoint   `json:"currently,omitempty"`
	Minutely  DataBlock   `json:"minutely,omitempty"`
	Hourly    DataBlock   `json:"hourly,omitempty"`
	Daily     DataBlock   `json:"daily,omitempty"`
	Alerts    []Alert     `json:"alerts,omitempty"`
	Flags     Flags       `json:"flags,omitempty"`
}

// DataPoint contains various properties, each representing the average (unless otherwise specified) of a particular weather phenomenon occurring during a period of time.
type DataPoint struct {
	ApparentTemperature        measurement `json:"apparentTemperature,omitempty"`
	ApparentTemperatureMax     measurement `json:"apparentTemperatureMax,omitempty"`
	ApparentTemperatureMaxTime timestamp   `json:"apparentTemperatureMaxTime,omitempty"`
	ApparentTemperatureMin     measurement `json:"apparentTemperatureMin,omitempty"`
	ApparentTemperatureMinTime timestamp   `json:"apparentTemperatureMinTime,omitempty"`
	CloudCover                 measurement `json:"cloudCover,omitempty"`
	DewPoint                   measurement `json:"dewPoint,omitempty"`
	Humidity                   measurement `json:"humidity,omitempty"`
	Icon                       string      `json:"icon,omitempty"`
	MoonPhase                  measurement `json:"moonPhase,omitempty"`
	NearestStormBearing        measurement `json:"nearestStormBearing,omitempty"`
	NearestStormDistance       measurement `json:"nearestStormDistance,omitempty"`
	Ozone                      measurement `json:"ozone,omitempty"`
	PrecipAccumulation         measurement `json:"precipAccumulation,omitempty"`
	PrecipIntensity            measurement `json:"precipIntensity,omitempty"`
	PrecipIntensityMax         measurement `json:"precipIntensityMax,omitempty"`
	PrecipIntensityMaxTime     timestamp   `json:"precipIntensityMaxTime,omitempty"`
	PrecipProbability          measurement `json:"precipProbability,omitempty"`
	PrecipType                 string      `json:"precipType,omitempty"`
	Pressure                   measurement `json:"pressure,omitempty"`
	Summary                    string      `json:"summary,omitempty"`
	SunriseTime                timestamp   `json:"sunriseTime,omitempty"`
	SunsetTime                 timestamp   `json:"sunsetTime,omitempty"`
	Temperature                measurement `json:"temperature,omitempty"`
	TemperatureMax             measurement `json:"temperatureMax,omitempty"`
	TemperatureMaxTime         timestamp   `json:"temperatureMaxTime,omitempty"`
	TemperatureMin             measurement `json:"temperatureMin,omitempty"`
	TemperatureMinTime         timestamp   `json:"temperatureMinTime,omitempty"`
	Time                       timestamp   `json:"time,omitempty"`
	UvIndex                    int64       `json:"uvIndex,omitempty"`
	UvIndexTime                timestamp   `json:"uvIndexTime,omitempty"`
	Visibility                 measurement `json:"visibility,omitempty"`
	WindBearing                measurement `json:"windBearing,omitempty"`
	WindGust                   measurement `json:"windGust,omitempty"`
	WindGustTime               timestamp   `json:"windGustTime,omitempty"`
	WindSpeed                  measurement `json:"windSpeed,omitempty"`
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
	Expires     timestamp `json:"expires,omitempty"`
	Regions     []string  `json:"regions,omitempty"`
	Severity    string    `json:"severity,omitempty"`
	Time        timestamp `json:"time,omitempty"`
	Title       string    `json:"title,omitempty"`
	Uri         string    `json:"uri,omitempty"`
}

// Flags contains various metadata information related to the request
type Flags struct {
	DarkSkyUnavailable string   `json:"darksky-unavailable,omitempty"`
	Sources            []string `json:"sources,omitempty"`
	Units              string   `json:"units,omitempty"`
}
