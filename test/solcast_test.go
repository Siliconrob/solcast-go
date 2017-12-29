package solcast

import (
	solcast "../solcast"
	datatypes "../solcast/types"
	"github.com/jimlawless/whereami"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var radiationLocation = datatypes.LatLng{Longitude: -97, Latitude: 32}
var powerLocation = datatypes.PowerLatLng{Capacity: 1000, LatLng: radiationLocation}

const forecastCount = 336
const actualsCount = 317

func TestRadiationForecast(t *testing.T) {
	result := solcast.RadiationForecast(radiationLocation)
	assert.Equal(t, len(result.Forecasts), forecastCount, "Radiation forecast count should be", forecastCount)
	log.Println("%v passed", whereami.WhereAmI())
}

func TestRadiationEstimatedActuals(t *testing.T) {
	result := solcast.RadiationEstimatedActuals(radiationLocation)
	assert.Equal(t, len(result.EstimatedActuals), actualsCount, " Radiation estimated actuals count should be", actualsCount)
	log.Println("%v passed", whereami.WhereAmI())
}

func TestPowerForecast(t *testing.T) {
	result := solcast.PowerForecast(powerLocation)
	assert.Equal(t, len(result.Forecasts), forecastCount, " Power forecast count should be", forecastCount)
	log.Println("%v passed", whereami.WhereAmI())
}

func TestPowerEstimatedActuals(t *testing.T) {
	result := solcast.PowerEstimatedActuals(powerLocation)
	assert.Equal(t, len(result.EstimatedActuals), actualsCount, " Power estimated actuals count should be", actualsCount)
	log.Println("%v passed", whereami.WhereAmI())
}
