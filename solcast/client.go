package solcast

import (
	datatypes "./types"
	"github.com/go-resty/resty"
	"strconv"
	"math"
	"log"
	"github.com/jimlawless/whereami"
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

func RadiationEstimatedActuals(location datatypes.LatLng) *datatypes.RadiationEstimatedActuals {
	currentConfig := Read()
	resp, err := resty.R().SetQueryParams(map[string]string{
		"longitude": toString(location.Longitude, 6),
		"latitude": toString(location.Latitude, 6),
		"api_key": currentConfig.APIKey,
	}).SetResult(&datatypes.RadiationEstimatedActuals{}).Get(currentConfig.Url + "/radiation/estimated_actuals")

	if err != nil {
		log.Printf("%v failed %v", whereami.WhereAmI(), err)
	}
	result := resp.Result().(*datatypes.RadiationEstimatedActuals)
	return result
}

func RadiationForecast(location datatypes.LatLng) interface{} {
	currentConfig := Read()
	resp, err := resty.R().SetQueryParams(map[string]string{
		"longitude": toString(location.Longitude, 6),
		"latitude": toString(location.Latitude, 6),
		"api_key": currentConfig.APIKey,
	}).SetResult(&datatypes.RadiationForecast{}).Get(currentConfig.Url + "/radiation/forecasts")

	if err != nil {
		log.Printf("%v failed %v", whereami.WhereAmI(), err)
	}
	//result := resp.Result().(*datatypes.RadiationForecast)
	data := resp.Result()
	return data
}

