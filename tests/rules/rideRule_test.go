package rules

import (
	"Beat/constants"
	"Beat/datamodel"
	"Beat/rules"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplyRulesOnRide(t *testing.T) {
	report1 := rules.ApplyRulesOnRide(datamodel.RideReport{Id: 1, TotalFare: 0})
	assert.Equal(t, constants.MinFare, report1.TotalFare)

	report2 := rules.ApplyRulesOnRide(datamodel.RideReport{Id: 1, TotalFare: 10})
	assert.Equal(t, 10+constants.FlagFare, report2.TotalFare)
}

func TestApplyRulesOnSegment(t *testing.T) {
	assert.Equal(t, 0.074, rules.ApplyRulesOnSegment(datamodel.Segment{Distance: 0.1, Duration: 60, Speed: constants.DayNightChangingMinute - 10}))
	assert.Equal(t, 0.5416666666666666, constants.DayTimeFarePerKm-rules.ApplyRulesOnSegment(datamodel.Segment{Distance: 1, Duration: 60, Speed: 1 / 60, LocalMinutesOfDay: constants.DayNightChangingMinute + 10}))
}

func TestIsValidSegment(t *testing.T) {
	assert.True(t, rules.IsValidSegment(datamodel.Segment{Distance: 1, Duration: 60, Speed: 1 / 60}))
	assert.False(t, rules.IsValidSegment(datamodel.Segment{Distance: 1, Duration: 10, Speed: 500}))
}
