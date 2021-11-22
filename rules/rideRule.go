package rules

import (
	rideConstant "Beat/constants"
	"Beat/datamodel"
	"log"
)

func ApplyRulesOnRide(report datamodel.RideReport) datamodel.RideReport {
	report.TotalFare = report.TotalFare + rideConstant.FlagFare
	if report.TotalFare < rideConstant.MinFare {
		report.TotalFare = rideConstant.MinFare
	}
	return report
}

func ApplyRulesOnSegment(segment datamodel.Segment) float64 {
	if segment.Distance == float64(0) || segment.Speed < rideConstant.MaxIdleSpeedKmSec {
		return rideConstant.IdleFarePerSec * float64(segment.Duration)
	} else {
		if segment.LocalMinutesOfDay == 0 || segment.LocalMinutesOfDay > rideConstant.DayNightChangingMinute {
			return rideConstant.DayTimeFarePerKm * segment.Distance
		} else {
			return rideConstant.NightTimeFarePerKm * segment.Distance
		}
	}
}

func IsValidSegment(s datamodel.Segment) bool {
	if s.Speed > rideConstant.MaxValidSpeedKmSec {
		log.Printf("Max allowed speed breach: %v km/h", s.Speed/rideConstant.KmhToKmSec)
		return false
	}
	return true
}
