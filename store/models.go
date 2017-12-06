package store

// ParkingSpot object
type ParkingSpot struct {
	ID         int64    `json:"id" db:"id"`
	Available  bool     `json:"available" db:"available"`
	Latitude   *float64 `json:"latitude" db:"lat"`
	Longitude  *float64 `json:"longitude" db:"long"`
	ReservedAt *string  `json:"reserved_at" db:"reservedAt"`
	Price      *float64 `json:"price" db:"price"`
	StartTime  *string  `json:"start_time" db:"startTime"`
	EndTime    *string  `json:"end_time" db:"endTime"`
}

// Point : point where we are searching from
type Point struct {
	Latitude  float64
	Longitude float64
	Radius    float64
}
