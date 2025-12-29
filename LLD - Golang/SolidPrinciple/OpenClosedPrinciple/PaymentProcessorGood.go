package openclosedprinciple

type PaymentMethod interface {
	Process(amount float64)
}

type CreditCardProcessor struct{}

func (CreditCardProcessor) Process(amount float64) {
	_ = amount
}

type PaypalProcessor struct{}

func (PaypalProcessor) Process(amount float64) {
	_ = amount
}

type CryptoProcessor struct{}

func (CryptoProcessor) Process(amount float64) {
	_ = amount
}

type PaymentProcessorGood struct{}

func (PaymentProcessorGood) Process(method PaymentMethod, amount float64) {
	method.Process(amount)
}
