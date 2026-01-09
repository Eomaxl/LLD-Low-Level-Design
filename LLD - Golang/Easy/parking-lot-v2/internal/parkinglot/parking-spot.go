package parkinglot

type ParkingSpot struct {
	id       string
	spotType SpotType
}

func NewParkingSpot(id string, spotType SpotType) *ParkingSpot {
	return &ParkingSpot{
		id:       id,
		spotType: spotType,
	}
}

func (p *ParkingSpot) GetID() string {
	return p.id
}

func (p *ParkingSpot) GetSpotType() SpotType {
	return p.spotType
}
