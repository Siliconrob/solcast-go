package solcast

import "time"

// JSON API response from /pv_power/forecasts
type PowerForecastsResponse struct {
	Forecasts []struct {
		PeriodEnd  time.Time `json:"period_end"`
		Period     string    `json:"period"`
		PvEstimate float64   `json:"pv_estimate"`
	} `json:"forecasts"`
}

// JSON API response from /pv_power/estimated_actuals
type PowerEstimatedActualsResponse struct {
	EstimatedActuals []struct {
		PeriodEnd  time.Time `json:"period_end"`
		Period     string    `json:"period"`
		PvEstimate float64   `json:"pv_estimate"`
	} `json:"estimated_actuals"`
}

// Possible parameters that the PvPower Forecasts/EstimatedActuals services can work with
type PowerQueryParams struct {
	Format      string    `url:"format,omitempty"`
	Latitude    string    `url:"latitude,omitempty"`
	Longitude   string    `url:"longitude,omitempty"`
	APIKey      string    `url:"api_key,omitempty"`
	Capacity    int       `url:"capacity,omitempty"`
	Tilt        int       `url:"tilt,omitempty"`
	Azimuth     int       `url:"azimuth,omitempty"`
	InstallDate time.Time `url:"install_date,omitempty"`
	LossFactor  float64   `url:"loss_factor,omitempty"`
}
