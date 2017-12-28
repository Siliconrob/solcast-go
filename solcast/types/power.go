package solcast

import "time"

type PowerForecast struct {
	Forecasts []struct {
		PeriodEnd time.Time `json:"period_end"`
		Period string `json:"period"`
		PvEstimate float64 `json:"pv_estimate"`
	} `json:"forecasts"`
}

type PowerEstimatedActuals struct {
	EstimatedActuals []struct {
		PeriodEnd  time.Time `json:"period_end"`
		Period     string    `json:"period"`
		PvEstimate int       `json:"pv_estimate"`
	} `json:"estimated_actuals"`
}

type GetPowerForecast struct {
	Capacity int
	Tilt int
	Azimuth int
	InstallDate time.Time
	LossFactor float64
}