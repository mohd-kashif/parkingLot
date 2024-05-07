package ticket

import (
	"time"
)

type Ticket struct {
	ID            string
	Amount        float64
	ParkingSpotID string
	Timestamp     time.Time
	VehicleID     string
	ParkingLotID  string
	FloorNumber   string
}

func NewTicket(parkingLotID string, floorNumber string, spotID string) Ticket {
	ID := parkingLotID + "_" + floorNumber + "_" + spotID
	return Ticket{
		ID:            ID,
		Amount:        0,
		ParkingSpotID: spotID,
		Timestamp:     time.Time{},
		VehicleID:     ID,
		ParkingLotID:  parkingLotID,
		FloorNumber:   floorNumber,
	}
}
