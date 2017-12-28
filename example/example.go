package main

import (
	solcast "../solcast"
	datatypes "../solcast/types"
	"log"
	"fmt"
	"errors"
)

/*
func Init() {
	if currentKey != "" && len(currentKey) == 32 {
		return
	}
	var commandArgKey string
	flag.StringVar(&commandArgKey, "key", "", "API key for Solcast library")
	flag.Parse()

	if len(commandArgKey) == 32 {
		currentKey = commandArgKey
	} else {
		currentKey = os.Getenv(Solcast_API_KeyName)
	}
}
*/

func testRadiationForecast(location datatypes.LatLng) {
	result := *solcast.RadiationForecast(location)

	if len(result.Forecasts) != 336 {
		errors.New("Unexpected amount of forecasts")
	}

	log.Printf("Forecast %v", result.Forecasts)
	fmt.Println(result.Forecasts)
}

func testRadiationEstimatedActuals(location datatypes.LatLng) datatypes.RadiationEstimatedActuals {
	result := solcast.RadiationEstimatedActuals(location)
	return *result
}


func main() {
	//solcast.Init()
	testLocation := datatypes.LatLng{ Longitude: -97, Latitude: 32}
	testRadiationForecast(testLocation)


	//items := len(radiationForecast.Forecasts)

	//radiation_estimated_actuals := testRadiationEstimatedActuals(testLocation)
	//log.Printf("EstimatedActuals %v", radiation_estimated_actuals)

	/*
	radiation := struct {
		forecast datatypes.RadiationForecast
		actuals	datatypes.RadiationEstimatedActuals
	}{
		solcast.RadiationForecast(testLocation),
		solcast.RadiationEstimatedActuals(testLocation),
	}
	*/
}