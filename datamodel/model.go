package datamodel

type Coordinate struct {
	Lat float64
	Lon float64
}

type Position struct {
	LatLong   Coordinate
	TimeStamp int64
}

type Tuple struct {
	RiderId  int64
	Position Position
}

type Segment struct {
	Distance          float64
	Duration          int64
	Speed             float64
	LocalMinutesOfDay int64
}

type TupleBatch struct {
	RiderId   int64
	Positions []Position
}

type RideReport struct {
	Id        int64   `csv:"id_ride"`
	TotalFare float64 `csv:"fare_estimate"`
}
