package constants

const (
	InputFilePath         = "resource/paths.csv"
	OutputFilePath         = "resource/rideReport.csv"
	NoOfConcurrentRequest = 6
	BufferSize            = 1000
	EarthRadiusKm          = 6371
	MinFare                = 3.47
	DayNightChangingHour   = 5
	DayNightChangingMinute = DayNightChangingHour * 60
	FlagFare               = 1.3
	DayTimeFarePerKm       = 0.74
	NightTimeFarePerKm     = 1.3
	IdleFarePerSec         = 11.90 / 3600
	MaxValidSpeedKmH       = 100
	MaxIdleSpeedKmH        = 10
	KmhToKmSec             = 0.00027778
	MaxValidSpeedKmSec     = MaxValidSpeedKmH * KmhToKmSec
	MaxIdleSpeedKmSec      = MaxIdleSpeedKmH * KmhToKmSec
	MinLat                 = -90
	MaxLat                 = 90
	MinLon                 = -180
	MaxLon                 = 180
	ZoneId                 = "Europe/Athens"
)
