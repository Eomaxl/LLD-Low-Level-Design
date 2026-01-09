package internal

import (
	"math"
	"time"
)

// calculateFeeCents bills using ceil(hours) with minimum 1 hour (common parking policy).
func calculateFeeCents(entryMs, exitMs int64, hourlyRateCents int64) int64 {
	if exitMs < entryMs {
		exitMs = entryMs
	}
	durationMs := exitMs - entryMs

	hours := math.Ceil(float64(durationMs) / float64(time.Hour.Milliseconds()))
	if hours < 1 {
		hours = 1
	}
	return int64(hours) * hourlyRateCents
}
