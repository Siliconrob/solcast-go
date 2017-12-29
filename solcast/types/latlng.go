package solcast

// LatLng location on the Earth, expected projection 4326
type LatLng struct {
	Latitude  float64
	Longitude float64
}

// Expanded LatLng location on the Earth, expected projection 4326 with a Capacity property
type PowerLatLng struct {
	LatLng
	Capacity int
}
