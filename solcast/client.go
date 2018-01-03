package solcast

import (
	datatypes "github.com/Siliconrob/solcast-go/solcast/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
	"github.com/jimlawless/whereami"
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"github.com/vardius/worker-pool"
	"runtime"
	"sync"
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
	return float64(round(num*output)) / output
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
		log.Printf("Unable to create HTTP client %v", whereami.WhereAmI())
		panic(err)
	}
	if resp.StatusCode >= 500 && resp.StatusCode < 600 {
		log.Printf("Solcast API error, post to GitHub or here https://forums.solcast.com.au/ please %v", whereami.WhereAmI())
		panic(err)
	}
	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		if resp.StatusCode == 429 {
			limits := getApiRateLimits(resp)
			log.Printf("Request rate limit exceeded please wait and try again %v %v", limits, whereami.WhereAmI())
			return []byte{}, errors.New(fmt.Sprintf("Retry request at %v", limits.ResetTime))
		}
		log.Printf("Bad request, check your inputs %v", whereami.WhereAmI())
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failure to read the HTTP body %v", whereami.WhereAmI())
		panic(err)
	}
	return body, nil
}

func powerEstimatedActuals(location datatypes.PowerLatLng, config Config) datatypes.PowerEstimatedActualsResponse {
	results := datatypes.PowerEstimatedActualsResponse{}
	queryParams := &datatypes.PowerQueryParams{
		Format:    "json",
		Latitude:  toString(location.Latitude, 6),
		Longitude: toString(location.Longitude, 6),
		APIKey:    config.APIKey,
		Capacity:  location.Capacity,
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

func PowerEstimatedActuals(location datatypes.PowerLatLng) datatypes.PowerEstimatedActualsResponse {
	return powerEstimatedActuals(location, Read())
}

func radiationEstimatedActuals(location datatypes.LatLng, config Config) datatypes.RadiationEstimatedActualsResponse {
	results := datatypes.RadiationEstimatedActualsResponse{}
	queryParams := &datatypes.RadiationQueryParams{
		Format:    "json",
		Latitude:  toString(location.Latitude, 6),
		Longitude: toString(location.Longitude, 6),
		APIKey:    config.APIKey,
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

func powerForecast(location datatypes.PowerLatLng, config Config) datatypes.PowerForecastsResponse {
	results := datatypes.PowerForecastsResponse{}
	queryParams := &datatypes.PowerQueryParams{
		Format:    "json",
		Latitude:  toString(location.Latitude, 6),
		Longitude: toString(location.Longitude, 6),
		APIKey:    config.APIKey,
		Capacity:  location.Capacity,
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

func radiationForecast(location datatypes.LatLng, config Config) datatypes.RadiationForecastsResponse {
	results := datatypes.RadiationForecastsResponse{}
	queryParams := &datatypes.RadiationQueryParams{
		Format:    "json",
		Latitude:  toString(location.Latitude, 6),
		Longitude: toString(location.Longitude, 6),
		APIKey:    config.APIKey,
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

func batchRadiationForecast(locations[] datatypes.LatLng, config Config) []datatypes.RadiationForecastsResponse {
	var wg sync.WaitGroup

	toDo := len(locations)
	workersAmount := runtime.NumCPU()

	// create new pool
	pool := workerpool.New(runtime.NumCPU())
	out := make(chan datatypes.LatLng, toDo)

	pool.Start(workersAmount, func(location datatypes.LatLng) {
		defer wg.Done()
		out <- location
	})

	wg.Add(workersAmount)

	for i := 0; i < toDo; i++ {
		pool.Delegate(i)
	}

	go func() {
		// stop all workers after jobs are done
		wg.Wait()
		close(out)
		pool.Stop()
	}()

	var results []datatypes.RadiationForecastsResponse
	for location := range out {
		results = append(results, radiationForecast(location, config))
	}
	return results
}

func BatchRadiationForecast(locations[] datatypes.LatLng) []datatypes.RadiationForecastsResponse {
	return batchRadiationForecast(locations, Read())
}

func BatchRadiationForecastWithKey(locations[] datatypes.LatLng, apiKey string) []datatypes.RadiationForecastsResponse {
	config := Read()
	config.APIKey = apiKey
	return batchRadiationForecast(locations, config)
}

func batchPowerForecast(locations[] datatypes.PowerLatLng, config Config) []datatypes.PowerForecastsResponse {
	return nil
}

func BatchPowerForecast(locations[] datatypes.PowerLatLng) []datatypes.PowerForecastsResponse {
	return batchPowerForecast(locations, Read())
}

func BatchPowerForecastWithKey(locations[] datatypes.PowerLatLng, apiKey string) []datatypes.PowerForecastsResponse {
	config := Read()
	config.APIKey = apiKey
	return batchPowerForecast(locations, config)
}