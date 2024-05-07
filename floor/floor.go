package floor

import (
	"errors"
	"parkingLot/constant"
	parkingspot "parkingLot/parkingSpot"
)

type Floor struct {
	ID                            string
	parkingLotID                  string
	NumberOfParkingSpots          int
	ParkingSpots                  []parkingspot.ParkingSpot
	NumberOfFreeCarParkingSpots   int
	NumberOfCarOccupiedSpots      int
	NumberOfFreeBikeParkingSpots  int
	NumberOfBikeOccupiedSpots     int
	NumberOfFreeTruckParkingSpots int
	NumberOfTruckOccupiedSpots    int
	FloorNumber                   string
}

func NewFloor(floorNumber string, parkingLotID string, numberOfParkingSpots int) Floor {
	parkingSpots := make([]parkingspot.ParkingSpot, numberOfParkingSpots)
	numberOfTruckOccupiedSpots := 0
	numberOfFreeTruckParkingSpots := constant.NumberOfTruckParkingSpots
	numberOfBikeOccupiedSpots := 0
	numberOfFreeBikeParkingSpots := constant.NumberOfBikeParkingSpots
	numberOfCarOccupiedSpots := 0
	numberOfFreeCarParkingSpots := numberOfParkingSpots - (numberOfFreeBikeParkingSpots + numberOfFreeTruckParkingSpots)
	return Floor{
		ID:                            parkingLotID + "_" + floorNumber,
		parkingLotID:                  parkingLotID,
		NumberOfParkingSpots:          numberOfParkingSpots,
		ParkingSpots:                  parkingSpots,
		NumberOfFreeCarParkingSpots:   numberOfFreeCarParkingSpots,
		NumberOfCarOccupiedSpots:      numberOfCarOccupiedSpots,
		NumberOfFreeBikeParkingSpots:  numberOfFreeBikeParkingSpots,
		NumberOfBikeOccupiedSpots:     numberOfBikeOccupiedSpots,
		NumberOfFreeTruckParkingSpots: numberOfFreeTruckParkingSpots,
		NumberOfTruckOccupiedSpots:    numberOfTruckOccupiedSpots,
		FloorNumber:                   floorNumber,
	}
}

func (f *Floor) AddSpot(spotType string) error {
	if !isSpotvalid(spotType) {
		return errors.New("invalid spot type")
	}

	switch spotType {
	case constant.Car:
		f.NumberOfFreeCarParkingSpots++
	case constant.Bike:
		f.NumberOfFreeBikeParkingSpots++
	case constant.Truck:
		f.NumberOfFreeTruckParkingSpots++
	}

	f.NumberOfParkingSpots++

	return nil
}

func isSpotvalid(spotType string) bool {
	for _, v := range constant.ValidSpotTypes {
		if v == spotType {
			return true
		}
	}

	return false
}

func (f *Floor) FindNearestAvailableSpot(vehicleType string) (parkingspot.ParkingSpot, error) {
	switch vehicleType {
	case constant.Bike:
		if f.NumberOfBikeOccupiedSpots == constant.NumberOfBikeParkingSpots {
			return nil, errors.New("no spot available")
		} else {
			index := 1
			for i := 0; i < constant.NumberOfBikeParkingSpots; i++ {
				if f.ParkingSpots[index].IsSpotFree() {
					return f.ParkingSpots[index], nil
				}
				index++
			}
		}
	case constant.Truck:
		if f.NumberOfTruckOccupiedSpots == constant.NumberOfTruckParkingSpots {
			return nil, errors.New("no spot available")
		} else {
			index := 0
			for i := 0; i < constant.NumberOfTruckParkingSpots; i++ {
				if f.ParkingSpots[index].IsSpotFree() {
					return f.ParkingSpots[index], nil
				}
				index++
			}
		}
	case constant.Car:
		if f.NumberOfFreeCarParkingSpots == 0 {
			return nil, errors.New("no spot available")
		} else {
			index := 3
			for i := 0; i < constant.NumberOfTruckParkingSpots; i++ {
				if f.ParkingSpots[index].IsSpotFree() {
					return f.ParkingSpots[index], nil
				}
				index++
			}
		}
	}

	return nil, errors.New("not a valid vehicle type")
}

func (f *Floor) BookASlot() {

}
