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
	"log"
	"github.com/jimlawless/whereami"
	"github.com/pkg/errors"
)

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

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

func textAsInt(inputText string) int64 {
	if inputText == "" {
		return 0
	}
	value, err := strconv.ParseInt(inputText, 10, 64)
	if err != nil {
		panic(err)
	}
	return value
}

func getApiRateLimits(resp *http.Response) ApiLimits {
	results := ApiLimits{}
	if resp.StatusCode != 429 {
		return results
	}
	results.Limit = textAsInt(resp.Header.Get("x-rate-limit"))
	results.Remaining = textAsInt(resp.Header.Get("x-rate-limit-remaining"))
	parsedTime := textAsInt(resp.Header.Get("x-rate-limit-reset"))
	if parsedTime > 0 {
		results.ResetTime = time.Unix(parsedTime, 0)
	}
	return results
}

func getData(url string) ([]byte, error) {

	netClient := &http.Client{
		Timeout: time.Minute * 5,
	}

	resp, err := netClient.Get(url)
	if err != nil {
		log.Printf("Unable to create HTTP client", whereami.WhereAmI())
		panic(err)
	}
	if (resp.StatusCode >= 500 && resp.StatusCode < 600) {
		log.Printf("Solcast API error, post to GitHub or here https://forums.solcast.com.au/ please", whereami.WhereAmI())
		panic(err)
	}
	if resp.StatusCode >= 400 && resp.StatusCode < 500{
		if resp.StatusCode == 429 {
			limits := getApiRateLimits(resp)
			log.Printf("Request rate limit exceeded please wait and try again %v", limits, whereami.WhereAmI())
			return []byte{}, errors.New(fmt.Sprintf("Retry request at %v", limits.ResetTime))
		}
		log.Printf("Bad request, check your inputs", whereami.WhereAmI())
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failure to read the HTTP body", whereami.WhereAmI())
		panic(err)
	}
	return body, nil
}

func powerEstimatedActuals(location datatypes.PowerLatLng, config Config) datatypes.PowerEstimatedActualsResponse {
	results := datatypes.PowerEstimatedActualsResponse{}
	queryParams := &datatypes.PowerQueryParams{
		Format: "json",
		Latitude: toString(location.Latitude, 6),
		Longitude: toString(location.Longitude, 6),
		APIKey: config.APIKey,
		Capacity: location.Capacity,
	}
	v, _ := query.Values(queryParams)
	url := fmt.Sprintf("%v/pv_power/estimated_actuals?%v", config.Url, v.Encode())

	data, err := getData(url)
	if err != nil {
		log.Printf("HTTP request failed to %v %v", err, whereami.WhereAmI())
		panic(err)
	}
	if err := json.Unmarshal(data, &results); err != nil {
		log.Printf("Failure to parse HTTP response body to %v", whereami.WhereAmI())
		panic(err)
	}
	return results
}

func PowerEstimatedActualsWithKey(location datatypes.PowerLatLng, apiKey string) datatypes.PowerEstimatedActualsResponse {
	config := Read()
	config.APIKey = apiKey
	return powerEstimatedActuals(location, config)
}

func PowerEstimatedActuals(location datatypes.PowerLatLng, ) datatypes.PowerEstimatedActualsResponse {
	return powerEstimatedActuals(location, Read())
}

func asyncPowerEstimatedActuals(location datatypes.PowerLatLng, config Config) <- chan datatypes.PowerEstimatedActualsResponse {
	ch := make(chan datatypes.PowerEstimatedActualsResponse, 1) // buffered
	go func(location datatypes.PowerLatLng) {
		ch <- powerEstimatedActuals(location, config)
	}(location)
	return ch
}

func AsyncPowerEstimatedActuals(location datatypes.PowerLatLng) <- chan datatypes.PowerEstimatedActualsResponse {
	return asyncPowerEstimatedActuals(location, Read())
}

func AsyncPowerEstimatedActualsWithKey(location datatypes.PowerLatLng, apiKey string) <- chan datatypes.PowerEstimatedActualsResponse {
	config := Read()
	config.APIKey = apiKey
	return asyncPowerEstimatedActuals(location, config)
}

func radiationEstimatedActuals(location datatypes.LatLng, config Config) datatypes.RadiationEstimatedActualsResponse {
	results := datatypes.RadiationEstimatedActualsResponse{}
	queryParams := &datatypes.RadiationQueryParams{
		Format: "json",
		Latitude: toString(location.Latitude, 6),
		Longitude: toString(location.Longitude, 6),
		APIKey: config.APIKey,
	}
	v, _ := query.Values(queryParams)
	url := fmt.Sprintf("%v/radiation/estimated_actuals?%v", config.Url, v.Encode())
	data, err := getData(url)
	if err != nil {
		log.Printf("HTTP request failed to %v %v", err, whereami.WhereAmI())
		panic(err)
	}
	if err := json.Unmarshal(data, &results); err != nil {
		log.Printf("Failure to parse HTTP response body to %v", whereami.WhereAmI())
		panic(err)
	}
	return results
}

func RadiationEstimatedActualsWithKey(location datatypes.LatLng, apiKey string) datatypes.RadiationEstimatedActualsResponse {
	config := Read()
	config.APIKey = apiKey
	return radiationEstimatedActuals(location, config)
}

func RadiationEstimatedActuals(location datatypes.LatLng) datatypes.RadiationEstimatedActualsResponse {
	return radiationEstimatedActuals(location, Read())
}

func asyncRadiationEstimatedActuals(location datatypes.LatLng, config Config) <- chan datatypes.RadiationEstimatedActualsResponse {
	ch := make(chan datatypes.RadiationEstimatedActualsResponse, 1) // buffered
	go func(location datatypes.LatLng) {
		ch <- radiationEstimatedActuals(location, config)
	}(location)
	return ch
}

func AsyncRadiationEstimatedActuals(location datatypes.LatLng) <- chan datatypes.RadiationEstimatedActualsResponse {
	return asyncRadiationEstimatedActuals(location, Read())
}

func AsyncRadiationEstimatedActualsWithKey(location datatypes.LatLng, apiKey string) <- chan datatypes.RadiationEstimatedActualsResponse {
	config := Read()
	config.APIKey = apiKey
	return asyncRadiationEstimatedActuals(location, config)
}

func powerForecast(location datatypes.PowerLatLng, config Config) datatypes.PowerForecastsResponse {
	results := datatypes.PowerForecastsResponse{}
	queryParams := &datatypes.PowerQueryParams{
		Format: "json",
		Latitude: toString(location.Latitude, 6),
		Longitude: toString(location.Longitude, 6),
		APIKey: config.APIKey,
		Capacity: location.Capacity,
	}
	v, _ := query.Values(queryParams)
	url := fmt.Sprintf("%v/pv_power/forecasts?%v", config.Url, v.Encode())
	data, err := getData(url)
	if err != nil {
		log.Printf("HTTP request failed to %v %v", err, whereami.WhereAmI())
		panic(err)
	}
	if err := json.Unmarshal(data, &results); err != nil {
		log.Printf("Failure to parse HTTP response body to %v", whereami.WhereAmI())
		panic(err)
	}
	return results
}

func PowerForecast(location datatypes.PowerLatLng) datatypes.PowerForecastsResponse {
	return powerForecast(location, Read())
}

func PowerForecastWithKey(location datatypes.PowerLatLng, apiKey string) datatypes.PowerForecastsResponse {
	config := Read()
	config.APIKey = apiKey
	return powerForecast(location, config)
}

func asyncPowerForecast(location datatypes.PowerLatLng, config Config) <- chan datatypes.PowerForecastsResponse {
	ch := make(chan datatypes.PowerForecastsResponse, 1) // buffered
	go func(location datatypes.PowerLatLng) {
		ch <- powerForecast(location, config)
	}(location)
	return ch
}

func AsyncPowerForecast(location datatypes.PowerLatLng) <- chan datatypes.PowerForecastsResponse {
	return asyncPowerForecast(location, Read())
}

func AsyncPowerForecastWithKey(location datatypes.PowerLatLng, apiKey string) <- chan datatypes.PowerForecastsResponse {
	config := Read()
	config.APIKey = apiKey
	return asyncPowerForecast(location, config)
}

func radiationForecast(location datatypes.LatLng, config Config) datatypes.RadiationForecastsResponse {
	results := datatypes.RadiationForecastsResponse{}
	queryParams := &datatypes.RadiationQueryParams{
		Format: "json",
		Latitude: toString(location.Latitude, 6),
		Longitude: toString(location.Longitude, 6),
		APIKey: config.APIKey,
	}
	v, _ := query.Values(queryParams)
	url := fmt.Sprintf("%v/radiation/forecasts?%v", config.Url, v.Encode())
	data, err := getData(url)
	if err != nil {
		log.Printf("HTTP request failed to %v %v", err, whereami.WhereAmI())
		panic(err)
	}
	if err := json.Unmarshal(data, &results); err != nil {
		log.Printf("Failure to parse HTTP response body to %v", whereami.WhereAmI())
		panic(err)
	}
	return results
}

func RadiationForecastWithKey(location datatypes.LatLng, apiKey string) datatypes.RadiationForecastsResponse {
	config := Read()
	config.APIKey = apiKey
	return radiationForecast(location, config)
}

func RadiationForecast(location datatypes.LatLng) datatypes.RadiationForecastsResponse {
	return radiationForecast(location, Read())
}

func asyncRadiationForecast(location datatypes.LatLng, config Config) <- chan datatypes.RadiationForecastsResponse {
	ch := make(chan datatypes.RadiationForecastsResponse, 1) // buffered
	go func(location datatypes.LatLng) {
		ch <- radiationForecast(location, config)
	}(location)
	return ch
}

func AsyncRadiationForecast(location datatypes.LatLng) <- chan datatypes.RadiationForecastsResponse {
	return asyncRadiationForecast(location, Read())
}

func AsyncRadiationForecastWithKey(location datatypes.LatLng, apiKey string) <- chan datatypes.RadiationForecastsResponse {
	config := Read()
	config.APIKey = apiKey
	return asyncRadiationForecast(location, config)
}

func batchPowerForecast(locations []datatypes.PowerLatLng, config Config) <- chan datatypes.PowerForecastsResponse {
	ch := make(chan datatypes.PowerForecastsResponse, len(locations)) // buffered
	for _, location := range locations {
		go func(location datatypes.PowerLatLng) {
			ch <- powerForecast(location, config)
		}(location)
	}
	return ch
}

func BatchPowerForecast(locations []datatypes.PowerLatLng) <- chan datatypes.PowerForecastsResponse {
	return batchPowerForecast(locations, Read())
}

func BatchPowerForecastWithKey(locations []datatypes.PowerLatLng, apiKey string) <- chan datatypes.PowerForecastsResponse {
	config := Read()
	config.APIKey = apiKey
	return batchPowerForecast(locations, Read())
}

func batchPowerEstimatedActuals(locations []datatypes.PowerLatLng, config Config) <- chan datatypes.PowerEstimatedActualsResponse {
	ch := make(chan datatypes.PowerEstimatedActualsResponse, len(locations)) // buffered
	for _, location := range locations {
		go func(location datatypes.PowerLatLng) {
			ch <- powerEstimatedActuals(location, config)
		}(location)
	}
	return ch
}

func BatchPowerEstimatedActuals(locations []datatypes.PowerLatLng) <- chan datatypes.PowerEstimatedActualsResponse {
	return batchPowerEstimatedActuals(locations, Read())
}

func BatchPowerEstimatedActualsWithKey(locations []datatypes.PowerLatLng, apiKey string) <- chan datatypes.PowerEstimatedActualsResponse {
	config := Read()
	config.APIKey = apiKey
	return batchPowerEstimatedActuals(locations, Read())
}

func batchRadiationForecast(locations []datatypes.LatLng, config Config) <- chan datatypes.RadiationForecastsResponse {
	ch := make(chan datatypes.RadiationForecastsResponse, len(locations)) // buffered
	for _, location := range locations {
		go func(location datatypes.LatLng) {
			ch <- radiationForecast(location, config)
		}(location)
	}
	return ch
}

func BatchRadiationForecast(locations []datatypes.LatLng) <- chan datatypes.RadiationForecastsResponse {
	return batchRadiationForecast(locations, Read())
}

func BatchRadiationForecastWithKey(locations []datatypes.LatLng, apiKey string) <- chan datatypes.RadiationForecastsResponse {
	config := Read()
	config.APIKey = apiKey
	return batchRadiationForecast(locations, Read())
}

func batchRadiationEstimatedActuals(locations []datatypes.LatLng, config Config) <- chan datatypes.RadiationEstimatedActualsResponse {
	ch := make(chan datatypes.RadiationEstimatedActualsResponse, len(locations)) // buffered
	for _, location := range locations {
		go func(location datatypes.LatLng) {
			ch <- radiationEstimatedActuals(location, config)
		}(location)
	}
	return ch
}

func BatchRadiationEstimatedActuals(locations []datatypes.LatLng) <- chan datatypes.RadiationEstimatedActualsResponse {
	return batchRadiationEstimatedActuals(locations, Read())
}

func BatchRadiationEstimatedActualsWithKey(locations []datatypes.LatLng, apiKey string) <- chan datatypes.RadiationEstimatedActualsResponse {
	config := Read()
	config.APIKey = apiKey
	return batchRadiationEstimatedActuals(locations, Read())
}