package internal

import "time"

type Ticket struct {
	id          string
	spotID      string
	vehicleType VehicleType
	entryTime   time.Time
}

func NewTicket(id, spotID string, vehicleType VehicleType, entryTime time.Time) *Ticket {
	return &Ticket{
		id:          id,
		spotID:      spotID,
		vehicleType: vehicleType,
		entryTime:   entryTime,
	}
}

func (t *Ticket) GetID() string {
	return t.id
}

func (t *Ticket) GetSpotType() string {
	return t.spotID
}

func (t *Ticket) GetVehicleType() VehicleType {
	return t.vehicleType
}

func (t *Ticket) GetEntryTime() time.Time {
	return t.entryTime
}
