package types

import "time"

type PowerForecast struct {
	Forecasts []struct {
		PeriodEnd time.Time `json:"period_end"`
		Period string `json:"period"`
		PvEstimate float64 `json:"pv_estimate"`
	} `json:"forecasts"`
}