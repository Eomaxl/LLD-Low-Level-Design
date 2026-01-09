package parkinglot

import (
	"crypto/rand"
	"encoding/hex"
	"parking-lot/internal/parkinglot/pricing"
	"time"
)

type ParkingLot struct {
	spots []ParkingSpot

	occupiedSpotIds map[string]struct{}

	activeTickets map[string]Ticket

	allocator SpotAllocator

	pricer pricing.PricingStrategy
}

func NewParkingLot(spots []ParkingSpot, allocator SpotAllocator, pricer pricing.PricingStrategy) *ParkingLot {
	if allocator == nil {
		allocator = FirstFitAllocator{}
	}

	return &ParkingLot{
		spots:           spots,
		occupiedSpotIds: make(map[string]struct{}),
		activeTickets:   make(map[string]Ticket),
		allocator:       allocator,
		pricer:          pricer,
	}
}

// Entry(vehicleType) -> Ticket
func (p *ParkingLot) Enter(vehicleType VehicleType) (Ticket, error) {
	spot, ok := p.allocator.FindSpot(p.spots, p.occupiedSpotIds, vehicleType)
	if !ok {
		return Ticket{}, ErrNoSpotAvailable
	}

	p.occupiedSpotIds[spot.GetID()] = struct{}{}

	ticketID := newID()
	entryMs := time.Now().UnixMilli()
	t := Ticket{
		id:        ticketID,
		spotType:  spot.GetID(),
		vehicle:   vehicleType,
		entryTime: entryMs,
	}

	p.activeTickets[ticketID] = t

	return t, nil
}

// Exit(ticketID) -> long(cent)
func (p *ParkingLot) Exit(ticketID string) (int64, error) {
	t, ok := p.activeTickets[ticketID]
	if !ok {
		return 0, ErrInvalidTicket
	}

	exitMs := time.Now().UnixMilli()
	duration := exitMs - t.GetEntrytime()
	fee := p.pricer.ChargeCents(duration)

	// free resources
	delete(p.activeTickets, ticketID)
	delete(p.occupiedSpotIds, t.GetID())

	return fee, nil
}

func newID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
