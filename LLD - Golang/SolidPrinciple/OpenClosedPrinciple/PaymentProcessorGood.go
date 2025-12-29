package openclosedprinciple

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
