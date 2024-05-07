package parkingSpot

import (
	"parkingLot/constant"
)

type TruckSpot struct {
	FloorNumber int
	spotType    string
	isFree      bool
}

func NewTruckSpot(floorNumber int) TruckSpot {
	return TruckSpot{
		FloorNumber: floorNumber,
		spotType:    constant.Truck,
		isFree:      false,
	}
}

func (t TruckSpot) IsSpotFree() bool {
	return t.isFree
}
