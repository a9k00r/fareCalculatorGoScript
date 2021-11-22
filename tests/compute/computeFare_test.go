package compute

import (
	"Beat/computeAlgorithm"
	"Beat/datamodel"
	"github.com/stretchr/testify/assert"
	"testing"
)

var positions = []datamodel.Position{
	{LatLong: datamodel.Coordinate{Lat: 38, Lon: 24}, TimeStamp: 1405595000},
	{LatLong: datamodel.Coordinate{Lat: 38, Lon: 24}, TimeStamp: 1405595000},
	{LatLong: datamodel.Coordinate{Lat: 38.02, Lon: 24.02}, TimeStamp: 1405595200},
	{LatLong: datamodel.Coordinate{Lat: 38.04, Lon: 24.04}, TimeStamp: 1405595400},
	{LatLong: datamodel.Coordinate{Lat: 38.06, Lon: 24.02}, TimeStamp: 1405595600},
	{LatLong: datamodel.Coordinate{Lat: 38.08, Lon: 24}, TimeStamp: 1405595800},
}

func TestCalculateRide(t *testing.T) {
	actualRideFareReport := computeAlgorithm.CalculateRideFare(datamodel.TupleBatch{RiderId: 1, Positions: positions})
	assert.Equal(t, int64(1), actualRideFareReport.Id, "invalid RiderId")
	assert.Equal(t, 9.679192443960003, actualRideFareReport.TotalFare, "invalid Fare")
}
