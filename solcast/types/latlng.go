package solcast

type LatLng struct {
	Latitude float64
	Longitude float64
}

type PowerLatLng  struct {
	LatLng
	Capacity int
}