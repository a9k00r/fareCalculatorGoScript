package utils

import (
	"Beat/constants"
	"Beat/datamodel"
	"math"
)

func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

func CalculateDistance(p, q datamodel.Coordinate) float64 {
	lat1 := degreesToRadians(p.Lat)
	lon1 := degreesToRadians(p.Lon)
	lat2 := degreesToRadians(q.Lat)
	lon2 := degreesToRadians(q.Lon)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(diffLon/2), 2)

	return 2 * math.Asin(math.Sqrt(a)) * constants.EarthRadiusKm
}
