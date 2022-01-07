package utils

import (
	"Beat/datamodel"
	haversineUtil "Beat/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	p                datamodel.Coordinate
	q                datamodel.Coordinate
	expectedDistance float64
}{
	{
		datamodel.Coordinate{Lat: 51.45, Lon: 1.15},  // Oxford, United Kingdom
		datamodel.Coordinate{Lat: 41.54, Lon: 12.27}, // Vatican, City Vatican City
		1389.1793118293067,
	},
	{
		datamodel.Coordinate{Lat: 22.34, Lon: 17.05}, // Windhoek, Namibia
		datamodel.Coordinate{Lat: 51.56, Lon: 4.29},  // Rotterdam, Netherlands
		3429.89310043882,
	},
	{
		datamodel.Coordinate{Lat: 63.24, Lon: 56.59}, // Esperanza, Argentina
		datamodel.Coordinate{Lat: 8.50, Lon: 13.14},  // Luanda, Angola
		6996.18595539861,
	},
	{
		datamodel.Coordinate{Lat: 90.00, Lon: 0.00}, // North/South Poles
		datamodel.Coordinate{Lat: 48.51, Lon: 2.21}, // Paris,  France
		4613.477506482742,
	},
	{
		datamodel.Coordinate{Lat: 45.04, Lon: 7.42},  // Turin, Italy
		datamodel.Coordinate{Lat: 3.09, Lon: 101.42}, // Kuala Lumpur, Malaysia
		10078.111954385415,
	},
}

func TestHaversineDistance(t *testing.T) {
	for _, input := range tests {
		actualDistance := haversineUtil.CalculateDistance(input.p, input.q)
		assert.Equal(t, actualDistance, input.expectedDistance, "fail: want %v %v -> %v got %v", input.p, input.q, input.expectedDistance, actualDistance)

	}
}
