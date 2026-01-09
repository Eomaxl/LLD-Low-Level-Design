package internal

import (
	"fmt"
)

type ParkingLot struct {
	spots []ParkingSpot

	spotByID        map[string]ParkingSpot
	occupiedSpotIds map[string]struct{} // Set<String>
	activeTickets   map[string]Ticket   // Map<ticketId, Ticket>

	spotRatesCents map[SpotType]int64  // per spot-type rate
	gateSpotOrder  map[string][]string // gate -> nearest-first spot IDs

	clock Clock

	// Concurrency: multiple gates calling Enter/Exit concurrently
	// must not double-book spots or close tickets twice.
	muMutex
}

// muMutex wrapper keeps parking_lot.go clean.
type muMutex struct {
	mu SimpleMutex
}

// SimpleMutex interface allows tests to replace locking if needed.
// In production we use real mutex implementation below.
type SimpleMutex interface {
	Lock()
	Unlock()
}

type realMutex struct{ ch chan struct{} }

// A simple mutex implementation backed by a buffered channel (works fine for this project).
func newRealMutex() realMutex {
	m := realMutex{ch: make(chan struct{}, 1)}
	m.ch <- struct{}{}
	return m
}
func (m realMutex) Lock()   { <-m.ch }
func (m realMutex) Unlock() { m.ch <- struct{}{} }

// NewParkingLotExtended builds a lot that supports multi-gate nearest allocation + per-type rates.
func NewParkingLotExtended(
	spots []ParkingSpot,
	spotRatesCents map[SpotType]int64,
	distances []GateDistance,
	clock Clock,
) (*ParkingLot, error) {

	if len(spots) == 0 || len(spotRatesCents) == 0 || len(distances) == 0 {
		return nil, ErrInvalidConfig
	}
	if clock == nil {
		clock = RealClock{}
	}

	spotByID := make(map[string]ParkingSpot, len(spots))
	for _, s := range spots {
		if s.id == "" {
			return nil, fmt.Errorf("%w: spot id empty", ErrInvalidConfig)
		}
		if _, exists := spotByID[s.id]; exists {
			return nil, fmt.Errorf("%w: duplicate spot id %s", ErrInvalidConfig, s.id)
		}
		if _, ok := spotRatesCents[s.spotType]; !ok {
			return nil, fmt.Errorf("%w: missing rate for spotType", ErrInvalidConfig)
		}
		spotByID[s.id] = s
	}

	gateSpotOrder := buildGateSpotOrder(distances)
	if len(gateSpotOrder) == 0 {
		return nil, fmt.Errorf("%w: no gate ordering built", ErrInvalidConfig)
	}

	// validate gateSpotOrder references known spots
	for gate, list := range gateSpotOrder {
		if len(list) == 0 {
			return nil, fmt.Errorf("%w: gate %s has empty ordering", ErrInvalidConfig, gate)
		}
		for _, sid := range list {
			if _, ok := spotByID[sid]; !ok {
				return nil, fmt.Errorf("%w: gate %s references unknown spot %s", ErrInvalidConfig, gate, sid)
			}
		}
	}

	pl := &ParkingLot{
		spots:           spots,
		spotByID:        spotByID,
		occupiedSpotIds: make(map[string]struct{}),
		activeTickets:   make(map[string]Ticket),
		spotRatesCents:  spotRatesCents,
		gateSpotOrder:   gateSpotOrder,
		clock:           clock,
	}
	pl.muMutex.mu = newRealMutex()
	return pl, nil
}

// EnterAtGate allocates nearest compatible spot for that gate and issues a ticket.
func (p *ParkingLot) EnterAtGate(gateID string, vehicleType VehicleType) (Ticket, error) {
	p.muMutex.mu.Lock()
	defer p.muMutex.mu.Unlock()

	order, ok := p.gateSpotOrder[gateID]
	if !ok {
		return Ticket{}, fmt.Errorf("%w: %s", ErrUnknownGate, gateID)
	}

	spot, found := p.findNearestAvailable(order, vehicleType)
	if !found {
		return Ticket{}, ErrNoAvailableSpot
	}

	// controlled denormalization: mark occupied in set
	p.occupiedSpotIds[spot.id] = struct{}{}

	// issue ticket
	ticketID := newID()
	entryMs := p.clock.Now().UnixMilli()
	ticket := NewTicket(ticketID, spot.id, vehicleType, entryMs)

	// O(1) ticket lookup
	p.activeTickets[ticketID] = ticket

	return ticket, nil
}

// Exit closes ticket and returns fee in cents.
func (p *ParkingLot) Exit(ticketID string) (int64, error) {
	p.muMutex.mu.Lock()
	defer p.muMutex.mu.Unlock()

	ticket, ok := p.activeTickets[ticketID]
	if !ok {
		return 0, ErrInvalidTicket
	}

	spot, ok := p.spotByID[ticket.spotID]
	if !ok {
		return 0, fmt.Errorf("%w: ticket references unknown spot", ErrInvalidConfig)
	}

	exitMs := p.clock.Now().UnixMilli()
	rate := p.spotRatesCents[spot.spotType]
	fee := calculateFeeCents(ticket.entryTime, exitMs, rate)

	// free spot + remove ticket
	delete(p.activeTickets, ticketID)
	delete(p.occupiedSpotIds, ticket.spotID)

	return fee, nil
}
