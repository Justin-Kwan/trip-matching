package geo

import (
	"math"

	// "googlemaps.github.io/maps"
)

type Trip struct {
	orig *Location
	dest *Location
}

type Location struct {
	lon float64
	lat float64
}

const (
	_earthRadiusKm = float64(6371)
	_urlBase       = "https://maps.googleapis.com/maps/api/distancematrix/json?units="
	_queryUrl      = "metric&origins=40.6655101,-73.89188969999998&destinations=40.6905615%2C-73.9976592&key=AIzaSyAdIxVqZL2LgJu84IlAsLIgiBDpK52kloI"
)

type GeoCalculator interface {
	TravelTime() float64
	Distance() float64
}

// NewTrip() *Trip {

// }

// distance by remote api
// func (t *Trip) TravelTime() float64 {
//
// 	return time
// }

// distance by calling maps api
// func (t *Trip) Distance() float64 {
// 	return dist
// }

// distance by calculation
// interface for different uses (calculation or google api)
func (t *Trip) Distance() float64 {
	deltaLat := (t.dest.lat - t.orig.lat) * (math.Pi / 180)
	deltaLon := (t.dest.lon - t.orig.lon) * (math.Pi / 180)

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(t.orig.lat*(math.Pi/180))*math.Cos(t.dest.lat*(math.Pi/180))*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	dist := _earthRadiusKm * c
	return dist
}
