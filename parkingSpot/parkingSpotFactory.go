package parkingSpot

import (
	"errors"
	"parkingLot/constant"
)

type ParkingSpotFactory struct {
}

func NewParkingSpot(spotType string, floorNumber int) (ParkingSpot, error) {
	switch spotType {
	case constant.Bike:
		return NewBikeSpot(floorNumber), nil
	case constant.Car:
		return NewCarSpot(floorNumber), nil
	case constant.Truck:
		return NewTruckSpot(floorNumber), nil
	default:
		return nil, errors.New("not a valid vahicle")
	}
}
