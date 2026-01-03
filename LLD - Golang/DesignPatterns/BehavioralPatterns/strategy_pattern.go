package behavioralpatterns

type PaymentStrategy interface {
	Pay(amount float64)
}

type CreditCardPayment struct {
	cardNumber string
}

func NewCreditCardPayment(cardNumber string) *CreditCardPayment {
	return &CreditCardPayment{
		cardNumber: cardNumber,
	}
}

func (c *CreditCardPayment) Pay(amount float64) bool {
	// credit card logic here
	return true
}

type PayPalPayment struct {
	email string
}

func NewPayPalPayment(email string) *PayPalPayment {
	return &PayPalPayment{
		email: email,
	}
}

func (p *PayPalPayment) Pay(amount float64) bool {
	// Paypal payment logic here
	return true
}

type ShoppingCart struct {
	paymentStrategy PaymentStrategy
}

func NewShoppingCart(paymentMethod PaymentStrategy) *ShoppingCart {
	return &ShoppingCart{
		paymentStrategy: paymentMethod,
	}
}

func (c *ShoppingCart) SetPaymentStrategy(paymentStrategy PaymentStrategy) {
	c.paymentStrategy = paymentStrategy
}

func (c *ShoppingCart) Checkout(amount float64) {
	if c.paymentStrategy != nil {
		c.paymentStrategy.Pay(amount)
	}
}

// Usage
// cart := NewShoppingCart()
// cart.SetPaymentStrategy(NewCreditCardPayment("3434-3434-3434"))
// cart.Checkout(1400)
