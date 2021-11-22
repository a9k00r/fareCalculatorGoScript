package utils

import (
	"Beat/datamodel"
	util "Beat/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var validCoordinates = []struct {
	coordinate datamodel.Coordinate
}{
	{
		datamodel.Coordinate{Lat: 51.45, Lon: 1.15},
	},
	{
		datamodel.Coordinate{Lat: 22.34, Lon: 17.05},
	},
	{
		datamodel.Coordinate{Lat: 63.24, Lon: 56.59},
	},
}

var invalidCoordinates = []struct {
	coordinate datamodel.Coordinate
}{
	{
		datamodel.Coordinate{Lat: -91.45, Lon: 181.15},
	},
	{
		datamodel.Coordinate{Lat: 91.34, Lon: -181.05},
	},
	{
		datamodel.Coordinate{Lat: -193.24, Lon: 196.59},
	},
	{
		datamodel.Coordinate{Lat: 92.00, Lon: 190.00},
	},
}

func TestIsValidLat(t *testing.T) {
	for _, input := range validCoordinates {
		lat := util.IsValidLat(input.coordinate.Lat)
		assert.True(t, lat)
	}
}

func TestNotValidLat(t *testing.T) {
	for _, input := range invalidCoordinates {
		lat := util.IsValidLat(input.coordinate.Lat)
		assert.False(t, lat)
	}
}

func TestIsValidLon(t *testing.T) {
	for _, input := range validCoordinates {
		lon := util.IsValidLat(input.coordinate.Lon)
		assert.True(t, lon)
	}
}

func TestNotValidLon(t *testing.T) {
	for _, input := range invalidCoordinates {
		lon := util.IsValidLat(input.coordinate.Lon)
		assert.False(t, lon)
	}
}
