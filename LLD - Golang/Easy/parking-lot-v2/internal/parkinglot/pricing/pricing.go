package pricing

type PricingStrategy interface {
	ChargeCents(durationMs int64) int64
}
