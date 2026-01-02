package main

type ParkingLotGood struct {
	spots []*EncapParkingSpot
}

func NewParking() *ParkingLotGood {
	return &ParkingLotGood{
		spots: []*EncapParkingSpot{}}
}

func (p *ParkingLotGood) ParkVehicle(vehicle *EncapVehicle) bool {
	spot := p.findAvailableSpot(vehicle)
	if spot == nil {
		return false
	}
	spot.Occupy(vehicle)
	return true
}

func (p *ParkingLotGood) findAvailableSpot(vehicle *EncapVehicle) *EncapParkingSpot {
	_ = vehicle
	if len(p.spots) == 0 {
		return nil
	}
	return p.spots[0]
}

func (p *ParkingLotGood) GetSpots() []*EncapParkingSpot {
	copySpots := make([]*EncapParkingSpot, len(p.spots))
	copy(copySpots, p.spots)
	return copySpots
}
