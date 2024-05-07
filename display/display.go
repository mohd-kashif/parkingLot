package display

import (
	"parkingLot/floor"
	parkingspot "parkingLot/parkingSpot"
)

type Display struct {
	ParkingLotID string
}

func NewDisplay(parkingLotID string) Display {
	return Display{
		ParkingLotID: parkingLotID,
	}
}

func ShowFreeSlots(floor floor.Floor, vehicleType string) []parkingspot.ParkingSpot {
	return []parkingspot.ParkingSpot{}
}

func ShowNumberOfFreeSlots(floor floor.Floor, vehicleType string) int {
	return 0
}

func ShowNumberOfOccupiedSlots(floor floor.Floor, vehicleType string) int {
	return 0
}
