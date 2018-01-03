package solcast

import (
	solcast "github.com/Siliconrob/solcast-go/solcast"
	datatypes "github.com/Siliconrob/solcast-go/solcast/types"
	"github.com/jimlawless/whereami"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"math"
)

var radiationLocation = datatypes.LatLng{Longitude: -97, Latitude: 32}
var powerLocation = datatypes.PowerLatLng{Capacity: 1000, LatLng: radiationLocation}

const forecastCount = 7 * 24 * (60/30) // 7 days * 24 hours * default period of 30 minutes
const actualsCount = 6.583 * 24 * (60/30) // 6.583 * 24 hours * default period of 30 minutes

func TestRadiationForecast(t *testing.T) {
	result := solcast.RadiationForecast(radiationLocation)
	log.Printf("%v", whereami.WhereAmI())
	recordCount := len(result.Forecasts)
	assert.Equal(t, recordCount, forecastCount, "Radiation forecast count should be %v", forecastCount)
}

func TestRadiationEstimatedActuals(t *testing.T) {
	result := solcast.RadiationEstimatedActuals(radiationLocation)
	log.Printf("%v", whereami.WhereAmI())
	recordCount := len(result.EstimatedActuals)
	refCount := int(math.Ceil(actualsCount))
	assert.True(t, recordCount >= refCount, " Radiation estimated actuals count should be %v", actualsCount)
}

func TestPowerForecast(t *testing.T) {
	result := solcast.PowerForecast(powerLocation)
	log.Printf("%v", whereami.WhereAmI())
	recordCount := len(result.Forecasts)
	assert.Equal(t, recordCount, forecastCount, " Power forecast count should be %v", forecastCount)
}

func TestPowerEstimatedActuals(t *testing.T) {
	result := solcast.PowerEstimatedActuals(powerLocation)
	log.Printf("%v", whereami.WhereAmI())
	recordCount := len(result.EstimatedActuals)
	refCount := int(math.Ceil(actualsCount))
	assert.True(t, recordCount >= refCount, " Power estimated actuals count should be %v", actualsCount)
}
