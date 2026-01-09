package parkinglot

type SpotAllocator interface {
	FindSpot(spots []ParkingSpot, occupied map[string]struct{}, vehicleType VehicleType) (ParkingSpot, bool)
}

type FirstFitAllocator struct{}

func (a FirstFitAllocator) FindSpot(spots []ParkingSpot, occupied map[string]struct{}, vehicleType VehicleType) (ParkingSpot, bool) {
	for _, s := range spots {
		if _, used := occupied[s.GetID()]; used {
			continue
		}
		if canFit(vehicleType, s.GetSpotType()) {
			return s, true
		}
	}
	return ParkingSpot{}, false
}

func canFit(v VehicleType, s SpotType) bool {
	return vehicleRank(v) <= spotRank(s)
}

func vehicleRank(v VehicleType) int {
	switch v {
	case VehicleMotorcycle:
		return 1
	case VehicleCar:
		return 2
	case VehicleLarge:
		return 3
	default:
		return 999
	}
}

func spotRank(s SpotType) int {
	switch s {
	case SpotMotorcycle:
		return 1
	case SpotCar:
		return 2
	case SpotLarge:
		return 3
	default:
		return 0
	}
}
