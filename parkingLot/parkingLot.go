package parkinglot

import (
	"errors"
	"parkingLot/constant"
	"parkingLot/floor"
	parkingspot "parkingLot/parkingSpot"
	"parkingLot/ticket"
)

type ParkingLot struct {
	ID             string
	NumberOfFloors int
	floors         []floor.Floor
}

func NewParkingLot(ID string, numberOfFloors int) ParkingLot {
	return ParkingLot{
		ID:             ID,
		NumberOfFloors: numberOfFloors,
		floors:         []floor.Floor{},
	}
}

func (p *ParkingLot) AddFloor(floor floor.Floor) {
	p.NumberOfFloors += 1
	p.floors = append(p.floors, floor)
}

func (p *ParkingLot) FindNearestAvailableSpot(vehicleType string) (parkingspot.ParkingSpot, *floor.Floor, error) {
	if !isVehicleValid(vehicleType) {
		return nil, nil, errors.New("not a valid vehicle")
	}
	for _, floor := range p.floors {
		spot, err := floor.FindNearestAvailableSpot(vehicleType)
		if err != nil {
			return nil, nil, err
		} else {
			return spot, &floor, nil
		}
	}

	return nil, nil, errors.New("parkking lot is full")

}

func isVehicleValid(vehicleType string) bool {
	for _, v := range constant.ValidSpotTypes {
		if v == vehicleType {
			return true
		}
	}

	return false
}

func (p *ParkingLot) BookASpot(vehicleType string, licenseNumber string, color string) (*ticket.Ticket, error) {
	spot, floor, err := p.FindNearestAvailableSpot(vehicleType)
	if err != nil {
		return nil, err
	}

	// ticket := ticket.NewTicket(p.ID, spot.IsSpotFree())
	// return nil, nil

}
