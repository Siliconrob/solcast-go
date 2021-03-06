package main

import (
	solcast "github.com/Siliconrob/solcast-go/solcast"
	datatypes "github.com/Siliconrob/solcast-go/solcast/types"
	"errors"
	"fmt"
	"log"
	"os"
)

var YOUR_API_KEY = "<API KEY HERE>"

func testRadiationForecast(location datatypes.LatLng) error {
	result := solcast.RadiationForecast(location)
	if len(result.Forecasts) != 336 {
		return errors.New("Unexpected amount of forecasts")
	}
	log.Printf("Forecast %v", len(result.Forecasts))
	return nil
}

func testRadiationEstimatedActuals(location datatypes.LatLng) error {
	result := solcast.RadiationEstimatedActuals(location)
	if len(result.EstimatedActuals) != 317 {
		return errors.New("Unexpected amount of estimated actuals")
	}
	log.Printf("Estimated Actuals %v", len(result.EstimatedActuals))
	return nil
}

func asyncTestRadiationForecast(locations[] datatypes.LatLng) error {

	results := solcast.BatchRadiationForecast(locations)

	for current := 0; current < len(results); current++ {
	}

	return nil
}

func testPowerForecast(location datatypes.PowerLatLng) error {
	result := solcast.PowerForecastWithKey(location, YOUR_API_KEY)
	if len(result.Forecasts) != 336 {
		return errors.New("Unexpected amount of forecasts")
	}
	log.Printf("Forecast %v", len(result.Forecasts))
	return nil
}

func testPowerEstimatedActuals(location datatypes.PowerLatLng) error {
	result := solcast.PowerEstimatedActuals(location)
	if len(result.EstimatedActuals) != 317 {
		return errors.New("Unexpected amount of estimated actuals")
	}
	log.Printf("Estimated Actuals %v", len(result.EstimatedActuals))
	return nil
}

func main() {
	testRadiationLocation := datatypes.LatLng{Longitude: -97, Latitude: 32}
	/*if err := testRadiationForecast(testRadiationLocation); err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}
	if err := testRadiationEstimatedActuals(testRadiationLocation); err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}*/
	var testRadiationLocations []datatypes.LatLng
	testRadiationLocations = append(testRadiationLocations, testRadiationLocation)
	if err := asyncTestRadiationForecast(testRadiationLocations); err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}


	testPowerLocation := datatypes.PowerLatLng{Capacity: 1000, LatLng: testRadiationLocation}
	var testPowerLocations []datatypes.PowerLatLng
	testPowerLocations = append(testPowerLocations, testPowerLocation)
	if err := testPowerEstimatedActuals(testPowerLocation); err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}
}
