package main

type EncapParkingSpot struct{}

func (e *EncapParkingSpot) Occupy(vehicle *EncapVehicle) {
	_ = vehicle
}

type EncapVehicle struct {
	Type string
}

// Exposes mutable state publicly
type ParkingLotBad struct {
	Spots []*EncapParkingSpot
}
