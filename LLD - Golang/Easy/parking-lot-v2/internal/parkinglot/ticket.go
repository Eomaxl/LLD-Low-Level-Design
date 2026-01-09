package parkinglot

type Ticket struct {
	id        string
	vehicle   VehicleType
	spotType  string
	entryTime int64
}

func NewTicket(id, spotType string, vehicle VehicleType, entryTime int64) *Ticket {
	return &Ticket{
		id:        id,
		vehicle:   vehicle,
		spotType:  spotType,
		entryTime: entryTime,
	}
}

func (t *Ticket) GetID() string {
	return t.id
}

func (t *Ticket) GetVehicle() VehicleType {
	return t.vehicle
}

func (t *Ticket) GetSpotType() string {
	return t.spotType
}

func (t *Ticket) GetEntrytime() int64 {
	return t.entryTime
}
