package internal

func canFit(v VehicleType, s SpotType) bool {
	return vehicleRank(v) <= spotRank(s)
}

func vehicleRank(v VehicleType) int {
	switch v {
	case VehicleMotorcyle:
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
