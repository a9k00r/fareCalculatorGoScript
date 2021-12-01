package computeAlgorithm

import (
	"Beat/constants"
	"Beat/datamodel"
	rideRule "Beat/rules"
	haversineUtil "Beat/utils"
	"log"
	"time"
)

var loc *time.Location = getLocation()

func CalculateRideFare(tupleBatch datamodel.TupleBatch) datamodel.RideReport {
	lastValidPosition := tupleBatch.Positions[0]
	RemoveIndex(tupleBatch.Positions, 0)

	totalFare := 0.0
	for _, position := range tupleBatch.Positions {
		segment := calculateSegment(lastValidPosition, position)
		if rideRule.IsValidSegment(segment) {
			segmentCost := rideRule.ApplyRulesOnSegment(segment)
			totalFare += segmentCost
			log.Printf("New segment calculated: start at '%v', stop at '%v', result '%v', totalFare '%v'\n", lastValidPosition, position, segment, segmentCost)
			lastValidPosition = position
		}
	}

	report := datamodel.RideReport{Id: tupleBatch.RiderId, TotalFare: totalFare}
	return rideRule.ApplyRulesOnRide(report)
}

func calculateSegment(start datamodel.Position, stop datamodel.Position) datamodel.Segment {
	segmentDistance := haversineUtil.CalculateDistance(start.LatLong, stop.LatLong)
	segmentTime := stop.TimeStamp - start.TimeStamp
	segmentSpeed := segmentDistance / float64(segmentTime)
	localMinutesOfDay := getLocalMinutesOfDay(stop.TimeStamp)

	return datamodel.Segment{Distance: segmentDistance, Duration: segmentTime, Speed: segmentSpeed, LocalMinutesOfDay: localMinutesOfDay}
}

func getLocalMinutesOfDay(timestamp int64) int64 {
	unixTimeUTC := time.Unix(timestamp, 0).In(loc)
	return int64(unixTimeUTC.Hour()*60 + unixTimeUTC.Minute())
}

func getLocation() *time.Location {
	loc, err := time.LoadLocation(constants.ZoneId)
	if err != nil {
		log.Fatalf("unable to load location from timezone %s", err)
	}
	return loc
}

func RemoveIndex(s []datamodel.Position, index int) []datamodel.Position {
	return append(s[:index], s[index+1:]...)
}
