package main

import (
	"errors"
	"sync"
)

type Vehicle struct {
	kind string // private: immutable from outside
}

func NewVehicle(kind string) Vehicle {
	return Vehicle{kind: kind}
}

func (v Vehicle) Kind() string {
	return v.kind
}

// Parkingspot is internal state owned by the parking lot.
// Nothing here is exported, so outside package can't mutate it.
type parkingSpot struct {
	id       int
	occupied bool
	vehicle  *Vehicle
}

func (s *parkingSpot) canFit(v Vehicle) bool {
	// placeholder for real logic (size/type constraints, etc)
	_ = v
	return true
}

func (s *parkingSpot) occupy(v Vehicle) error {
	if s.occupied {
		return errors.New("Spot already occupied")
	}

	s.occupied = true
	s.vehicle = &v
	return nil
}

func (s *parkingSpot) vacate() error {
	if !s.occupied {
		return errors.New("Spot is already empty")
	}
	s.occupied = false
	s.vehicle = nil
	return nil
}

// SpotView is a read-only representation exposed to callers.
// Note: it contains no pointers to internal mutable state.
type SpotView struct {
	ID       int
	Occupied bool
	Vehicle  string // vehicle kind (or empty)
}

type ParkingLot struct {
	mu    sync.RWMutex
	spots []*parkingSpot // private slice: callers can't append/ remove / reorder
}

func NewParkingLot(numSpots int) *ParkingLot {
	spots := make([]*parkingSpot, 0, numSpots)
	for i := 0; i < numSpots; i++ {
		spots = append(spots, &parkingSpot{id: i + 1})
	}
	return &ParkingLot{spots: spots}
}

// Park tries to park a vehicle. Only ParkingLot decides where/how.
func (p *ParkingLot) Park(v Vehicle) (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, s := range p.spots {
		if !s.occupied && s.canFit(v) {
			if err := s.occupy(v); err != nil {
				return 0, err
			}
			return s.id, nil
		}
	}
	return 0, errors.New(" No available spot")
}

func (p *ParkingLot) Leave(spotId int) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	s := p.findById(spotId)
	if s == nil {
		return errors.New("Invalid spot id")
	}
	return s.vacate()
}

func (p *ParkingLot) Snapshot() []SpotView {
	p.mu.RLock()
	defer p.mu.RUnlock()

	out := make([]SpotView, 0, len(p.spots))
	for _, s := range p.spots {
		vehicleKind := ""
		if s.vehicle != nil {
			vehicleKind = s.vehicle.Kind()
		}
		out = append(out, SpotView{
			ID:       s.id,
			Occupied: s.occupied,
			Vehicle:  vehicleKind,
		})
	}
	return out
}

// internal helper
func (p *ParkingLot) findById(id int) *parkingSpot {
	for _, s := range p.spots {
		if s.id == id {
			return s
		}
	}
	return nil
}
