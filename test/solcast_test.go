package solcast

import (
	solcast "github.com/Siliconrob/solcast-go/solcast"
	datatypes "github.com/Siliconrob/solcast-go/solcast/types"
	"github.com/jimlawless/whereami"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"fmt"
)

var radiationLocation = datatypes.LatLng{Longitude: -97, Latitude: 32}
var powerLocation = datatypes.PowerLatLng{Capacity: 1000, LatLng: radiationLocation}

const forecastCount = 336
const actualsCount = 317

func TestRadiationForecast(t *testing.T) {
	result := solcast.RadiationForecast(radiationLocation)
	log.Println(fmt.Sprintf("%v", whereami.WhereAmI()))
	assert.Equal(t, len(result.Forecasts), forecastCount, "Radiation forecast count should be %v", forecastCount)
}

func TestRadiationEstimatedActuals(t *testing.T) {
	result := solcast.RadiationEstimatedActuals(radiationLocation)
	log.Println(fmt.Sprintf("%v", whereami.WhereAmI()))
	assert.Equal(t, len(result.EstimatedActuals), actualsCount, " Radiation estimated actuals count should be %v", actualsCount)
}

func TestPowerForecast(t *testing.T) {
	result := solcast.PowerForecast(powerLocation)
	log.Println(fmt.Sprintf("%v", whereami.WhereAmI()))
	assert.Equal(t, len(result.Forecasts), forecastCount, " Power forecast count should be %v", forecastCount)
}

func TestPowerEstimatedActuals(t *testing.T) {
	result := solcast.PowerEstimatedActuals(powerLocation)
	log.Println(fmt.Sprintf("%v", whereami.WhereAmI()))
	assert.Equal(t, len(result.EstimatedActuals), actualsCount, " Power estimated actuals count should be %v", actualsCount)
}
