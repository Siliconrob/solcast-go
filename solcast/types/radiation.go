package types

import "time"

type RadiationForecast struct {
	Forecasts []struct {
		Ghi          int       `json:"ghi"`
		Ghi90        int       `json:"ghi90"`
		Ghi10        int       `json:"ghi10"`
		Ebh          int       `json:"ebh"`
		Dni          int       `json:"dni"`
		Dni10        int       `json:"dni10"`
		Dni90        int       `json:"dni90"`
		Dhi          int       `json:"dhi"`
		AirTemp      int       `json:"air_temp"`
		Zenith       int       `json:"zenith"`
		Azimuth      int       `json:"azimuth"`
		CloudOpacity int       `json:"cloud_opacity"`
		PeriodEnd    time.Time `json:"period_end"`
		Period       string    `json:"period"`
	} `json:"forecasts"`
}