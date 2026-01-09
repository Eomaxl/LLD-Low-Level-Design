package pricing

import "math"

type FlatHourlyPricing struct {
	HourlyRateCents int64
	MinHours        int64
}

func NewFlatHourlyPricing(hourlyRateCents int64) *FlatHourlyPricing {
	return &FlatHourlyPricing{
		HourlyRateCents: hourlyRateCents,
		MinHours:        1,
	}
}

func (p *FlatHourlyPricing) ChargeCents(durationMs int64) int64 {
	if durationMs < 0 {
		durationMs = 0
	}

	const hourMs = int64(60 * 60 * 1000)

	hours := int64(math.Ceil(float64(durationMs) / float64(hourMs)))
	if hours < p.MinHours {
		hours = p.MinHours
	}

	return hours * p.HourlyRateCents
}
