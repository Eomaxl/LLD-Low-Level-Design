package internal

func (p *ParkingLot) findNearestAvailable(order []string, v VehicleType) (ParkingSpot, bool) {
	for _, spotID := range order {
		if _, occupied := p.occupiedSpotIds[spotID]; occupied {
			continue
		}
		spot := p.spotByID[spotID]
		if canFit(v, spot.spotType) {
			return spot, true
		}
	}
	return ParkingSpot{}, false
}
