package parkingSpot

import (
	"parkingLot/constant"
)

type CarSpot struct {
	FloorNumber int
	spotType    string
	IsFree      bool
}

func NewCarSpot(floorNumber int) CarSpot {
	return CarSpot{
		FloorNumber: floorNumber,
		spotType:    constant.Bike,
		IsFree:      false,
	}
}

func (c CarSpot) IsSpotFree() bool {
	return c.IsFree
}
