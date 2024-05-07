package parkingSpot

import "parkingLot/constant"

type BikeSpot struct {
	FloorNumber int
	spotType    string
	isFree      bool
}

func NewBikeSpot(floorNumber int) BikeSpot {
	return BikeSpot{
		FloorNumber: floorNumber,
		spotType:    constant.Bike,
		isFree:      false,
	}
}

func (b BikeSpot) IsSpotFree() bool {
	return b.isFree
}
