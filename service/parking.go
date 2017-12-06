package service

import (
	"context"
	"strconv"

	"github.com/bluehawk27/ridecell/store"
)

// Service : object for the service layer
type Service struct {
	Store store.StoreType
}

// New Service : calls on the store interface
func NewService() *Service {
	store, err := store.NewStore()
	if err != nil {
		return nil
	}

	service := &Service{
		Store: store,
	}

	return service
}

// List : Service layer list method
func (s *Service) List(ctx context.Context, lat, long, radius string) (*[]store.ParkingSpot, error) {
	point, perr := qStringToPoint(lat, long, radius)
	if perr != nil {
		return nil, perr
	}
	spots, err := s.Store.GetAllSpotsInRadius(ctx, *point)
	if err != nil {
		return nil, err
	}
	return spots, nil
}

func qStringToPoint(lat, long, rad string) (*store.Point, error) {
	latitude, err := strconv.ParseFloat(lat, 64)
	longitude, lerr := strconv.ParseFloat(long, 64)
	radius, rerr := strconv.ParseFloat(rad, 64)
	if err != nil || lerr != nil || rerr != nil {
		return nil, err
	}

	p := &store.Point{
		Latitude:  latitude,
		Longitude: longitude,
		Radius:    radius,
	}

	return p, nil
}
