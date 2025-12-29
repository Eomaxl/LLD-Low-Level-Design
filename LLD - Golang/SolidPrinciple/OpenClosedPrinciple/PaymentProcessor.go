package openclosedprinciple

type PaymentProcessorBad struct{}

func (PaymentProcessorBad) Process(paymentType string, amount float64) {
	switch paymentType {
	case "credit":
	// credit card logic
	case "paypal":
		// paypal logic
	}
	_ = amount
}
