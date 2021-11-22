package utils

import (
	rideConstant "Beat/constants"
)

func IsValidLat(lat float64) bool {
	if lat < rideConstant.MinLat || lat > rideConstant.MaxLat {
		return false
	}
	return true
}

func IsValidLon(lon float64) bool {
	if lon < rideConstant.MinLon || lon > rideConstant.MaxLon {
		return false
	}
	return true
}
