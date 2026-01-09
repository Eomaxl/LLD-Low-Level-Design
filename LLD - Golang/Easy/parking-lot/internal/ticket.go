package internal

type Ticket struct {
	id          string
	spotID      string
	vehicleType VehicleType
	entryTime   int64
}

func NewTicket(id, spotID string, vehicleType VehicleType, entryTime int64) *Ticket {
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

func (t *Ticket) GetEntryTime() int64 {
	return t.entryTime
}
