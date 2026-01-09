package internal

type SpotType int

const (
	SpotMotorcycle SpotType = iota
	SpotCar
	SpotLarge
)

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

func (s ParkingSpot) GetSpotType() SpotType {
	return s.spotType
}

func (s ParkingSpot) GetID() string {
	return s.id
}
