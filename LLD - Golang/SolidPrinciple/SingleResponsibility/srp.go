package main

import "fmt"

// 1. TaxCalculator Interface
type TaxCalculator interface {
	CalculateTax(amount float64) float64
}

// 2. Region Specific Implementations
type IndiaTaxCalculator struct{}

func (i IndiaTaxCalculator) CalculateTax(amount float64) float64 {
	return amount * 0.18 // GST
}

type USTaxCalculator struct{}

func (u USTaxCalculator) CalculateTax(amount float64) float64 {
	return amount * 0.08 // Sales Tax
}

type UKTaxCalculator struct{}

func (u UKTaxCalculator) CalculateTax(amount float64) float64 {
	return amount * 0.12 // VAT
}

// 3. Invoice struct using dependency injection
type Invoice struct {
	Amount        float64
	TaxCalculator TaxCalculator
}

func (inv Invoice) GetTotalAmount() float64 {
	return inv.Amount + inv.TaxCalculator.CalculateTax(inv.Amount)
}

// 4. Main function
func main() {
	amount := 1000.0
	indiaInvoice := Invoice{Amount: amount, TaxCalculator: IndiaTaxCalculator{}}
	fmt.Printf("Total (India) : %.2f\n", indiaInvoice.GetTotalAmount())

	usInvoice := Invoice{Amount: amount, TaxCalculator: USTaxCalculator{}}
	fmt.Printf("Total (US) : %.2f\n", usInvoice.GetTotalAmount())

	ukInvoice := Invoice{Amount: amount, TaxCalculator: UKTaxCalculator{}}
	fmt.Printf("Total (UK): £%.2f\n", ukInvoice.GetTotalAmount())
}
