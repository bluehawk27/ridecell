package store

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

//  StoreType : Interface exposes methods that can later be used for mocking calls
type StoreType interface {
	GetAllSpots(ctx context.Context) (*[]ParkingSpot, error)
	GetAllSpotsInRadius(ctx context.Context, p Point) (*[]ParkingSpot, error)
}

// Store : Represents the DB object could add a config here to read in driver and address
type Store struct {
	db *sqlx.DB
}

// NewStore : New DB Connection
func NewStore() (StoreType, error) {
	// Todo Move diver/connection string to config
	db, err := sqlx.Connect("mysql", "root:@/ridecell?parseTime=true")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	s := &Store{
		db: db,
	}
	return s, nil
}

// GetAllSpots : Get All ParkingSpots
func (s *Store) GetAllSpots(ctx context.Context) (*[]ParkingSpot, error) {
	ps := []ParkingSpot{}
	if err := s.db.Select(&ps, GetAllSpotsNoLimit); err != nil {
		return nil, err
	}

	return &ps, nil
}

// GetAllSpotsInRadius : Get all Parking Spots for a radius from a given point
func (s *Store) GetAllSpotsInRadius(ctx context.Context, p Point) (*[]ParkingSpot, error) {
	ps := []ParkingSpot{}
	// https://github.com/kellydunn/golang-geo/blob/master/sql_mapper.go#L33
	selectSpot := "SELECT * FROM parkingSpot p"
	lat1 := fmt.Sprintf("sin(radians(%f)) * sin(radians(p.lat))", p.Latitude)
	lng1 := fmt.Sprintf("cos(radians(%f)) * cos(radians(p.lat)) * cos(radians(p.long) - radians(%f))", p.Latitude, p.Longitude)
	whereStr := fmt.Sprintf("WHERE acos(%s + %s) * %f <= %f", lat1, lng1, float64(EarthRadius), p.Radius)
	query := fmt.Sprintf("%s %s", selectSpot, whereStr)

	if err := s.db.Select(&ps, query); err != nil {
		return nil, err
	}

	return &ps, nil
}
