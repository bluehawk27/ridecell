package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/bluehawk27/ridecell/store"
)

var s = NewService()

//TODO use the storeType interface signature to create a mockStore for testing
func TestService_List(t *testing.T) {
	type fields struct {
		Store store.StoreType
	}
	type args struct {
		ctx    context.Context
		lat    string
		long   string
		radius string
	}
	f := fields{
		Store: s.Store,
	}
	ctx := context.Background()
	a := args{
		ctx:    ctx,
		lat:    "37.486714",
		long:   "-122.226306",
		radius: ".1",
	}
	lat := 37.48654
	long := -122.226354
	lat1 := 37.487281
	long1 := -122.226418

	p := store.ParkingSpot{
		ID:        6,
		Available: false,
		Latitude:  &lat,
		Longitude: &long,
	}
	p1 := store.ParkingSpot{
		ID:        7,
		Available: false,
		Latitude:  &lat1,
		Longitude: &long1,
	}
	ps := []store.ParkingSpot{p, p1}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]store.ParkingSpot
		wantErr bool
	}{
		{"Successful-list", f, a, &ps, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Store: tt.fields.Store,
			}
			got, err := s.List(tt.args.ctx, tt.args.lat, tt.args.long, tt.args.radius)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_qStringToPoint(t *testing.T) {
	type args struct {
		lat  string
		long string
		rad  string
	}
	a := args{
		lat:  "37.486714",
		long: "-122.226306",
		rad:  ".5",
	}

	p := &store.Point{
		Latitude:  37.486714,
		Longitude: -122.226306,
		Radius:    .5,
	}
	tests := []struct {
		name    string
		args    args
		want    *store.Point
		wantErr bool
	}{
		{"string-to-point", a, p, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := qStringToPoint(tt.args.lat, tt.args.long, tt.args.rad)
			if (err != nil) != tt.wantErr {
				t.Errorf("qStringToPoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("qStringToPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
