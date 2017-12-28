package solcast

import (
	datatypes "./types"
	"strconv"
	"math"
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"time"
	"github.com/google/go-querystring/query"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num * output)) / output
}

func toString(num float64, precision int) string {
	result := strconv.FormatFloat(toFixed(num, precision), 'f', precision, 64)
	return result
}

func getData(url string) []byte {

	netClient := &http.Client{
		Timeout: time.Minute * 5,
	}

	resp, err := netClient.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}

func PowerEstimatedActuals(location datatypes.PowerLatLng) datatypes.PowerEstimatedActualsResponse {
	results := datatypes.PowerEstimatedActualsResponse{}
	currentConfig := Read()
	queryParams := &datatypes.PowerQueryParams{
		Format: "json",
		Latitude: toString(location.Latitude, 6),
		Longitude: toString(location.Longitude, 6),
		APIKey: currentConfig.APIKey,
		Capacity: location.Capacity,
	}
	v, _ := query.Values(queryParams)
	url := fmt.Sprintf("%v/pv_power/estimated_actuals?%v", currentConfig.Url, v.Encode())
	if err := json.Unmarshal(getData(url), &results); err != nil {
		panic(err)
	}
	return results
}

func RadiationEstimatedActuals(location datatypes.LatLng) datatypes.RadiationEstimatedActuals {
	results := datatypes.RadiationEstimatedActuals{}
	currentConfig := Read()
	queryParams := &datatypes.RadiationQueryParams{
		Format: "json",
		Latitude: toString(location.Latitude, 6),
		Longitude: toString(location.Longitude, 6),
		APIKey: currentConfig.APIKey,
	}
	v, _ := query.Values(queryParams)
	url := fmt.Sprintf("%v/radiation/estimated_actuals?%v", currentConfig.Url, v.Encode())
	if err := json.Unmarshal(getData(url), &results); err != nil {
		panic(err)
	}
	return results
}

func PowerForecast(location datatypes.PowerLatLng) datatypes.PowerForecastsResponse {
	results := datatypes.PowerForecastsResponse{}
	currentConfig := Read()
	queryParams := &datatypes.PowerQueryParams{
		Format: "json",
		Latitude: toString(location.Latitude, 6),
		Longitude: toString(location.Longitude, 6),
		APIKey: currentConfig.APIKey,
		Capacity: location.Capacity,
	}
	v, _ := query.Values(queryParams)
	url := fmt.Sprintf("%v/pv_power/forecasts?%v", currentConfig.Url, v.Encode())
	if err := json.Unmarshal(getData(url), &results); err != nil {
		panic(err)
	}
	return results
}

func RadiationForecast(location datatypes.LatLng) datatypes.RadiationForecastsResponse {
	results := datatypes.RadiationForecastsResponse{}
	currentConfig := Read()
	queryParams := &datatypes.RadiationQueryParams{
		Format: "json",
		Latitude: toString(location.Latitude, 6),
		Longitude: toString(location.Longitude, 6),
		APIKey: currentConfig.APIKey,
	}
	v, _ := query.Values(queryParams)
	url := fmt.Sprintf("%v/radiation/forecasts?%v", currentConfig.Url, v.Encode())
	if err := json.Unmarshal(getData(url), &results); err != nil {
		panic(err)
	}
	return results
}

